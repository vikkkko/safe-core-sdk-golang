package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client represents a GraphQL client for querying subgraph data
type Client struct {
	endpoint   string
	httpClient *http.Client
	rpcURL     string

	ethClient   *ethclient.Client
	ethClientMu sync.Mutex
}

// Config represents configuration for GraphQL client
type Config struct {
	Endpoint string        // GraphQL endpoint URL
	RPCURL   string        // Optional RPC endpoint for receipt enrichment
	Timeout  time.Duration // HTTP timeout (default: 30s)
}

// DefaultConfig returns default GraphQL configuration
func DefaultConfig() Config {
	return Config{
		Endpoint: "https://api.studio.thegraph.com/query/103887/mvp/version/latest",
		Timeout:  30 * time.Second,
	}
}

// NewClient creates a new GraphQL client
func NewClient(config Config) *Client {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return &Client{
		endpoint: config.Endpoint,
		rpcURL:   config.RPCURL,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// NewDefaultClient creates a new GraphQL client with default configuration
func NewDefaultClient() *Client {
	return NewClient(DefaultConfig())
}

// Query executes a GraphQL query and returns the raw response
func (c *Client) Query(ctx context.Context, query string, variables map[string]interface{}) ([]byte, error) {
	reqBody := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for GraphQL errors
	var graphqlResp GraphQLResponse
	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(graphqlResp.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", graphqlResp.Errors[0].Message)
	}

	return body, nil
}

// GetPaymentAllowances queries allowances granted to a specific payment account
// Parameters:
//   - paymentAccount: The address of the payment account to query allowances for
//
// Returns:
//   - List of allowances showing which addresses have granted approvals to this payment account
func (c *Client) GetPaymentAllowances(ctx context.Context, paymentAccount string) ([]PaymentAllowance, error) {
	accountID := normalizeAddress(paymentAccount)
	if accountID == "" {
		return nil, fmt.Errorf("invalid payment account address: %s", paymentAccount)
	}

	query := `
		query GetPaymentAllowances($paymentAccount: String!) {
			paymentAllowances(where: {paymentAccount: $paymentAccount}) {
				token
				owner
				amount
			}
		}
	`

	variables := map[string]interface{}{
		"paymentAccount": accountID,
	}

	respBody, err := c.Query(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to query payment allowances: %w", err)
	}

	var response PaymentAllowancesResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Data.PaymentAllowances, nil
}

// GetPaymentApprovals queries approvals granted by a payment account
func (c *Client) GetPaymentApprovals(ctx context.Context, account string) ([]PaymentApproval, error) {
	accountID := normalizeAddress(account)
	if accountID == "" {
		return nil, fmt.Errorf("invalid payment account address: %s", account)
	}

	query := `
		query GetPaymentApprovals($account: String!) {
			paymentApprovals(
				where: {account: $account}
				orderBy: timestamp
				orderDirection: desc
			) {
				token
				spender
				amount
				timestamp
				txHash
			}
		}
	`

	variables := map[string]interface{}{
		"account": accountID,
	}

	respBody, err := c.Query(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to query payment approvals: %w", err)
	}

	var response PaymentApprovalsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payment approvals response: %w", err)
	}

	return response.Data.PaymentApprovals, nil
}

// GetTransactionInfo queries a transaction by its hash
func (c *Client) GetTransactionInfo(ctx context.Context, txHash string) (*TransactionInfo, error) {
	query := `
		query GetTransactionInfo($id: ID!) {
			transactionInfo(id: $id) {
				id
				blockNumber
				gasLimit
				gasPrice
				gasUsed
				status
				timestamp
				transactionFee
			}
		}
	`

	variables := map[string]interface{}{
		"id": txHash,
	}

	respBody, err := c.Query(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to query transaction info: %w", err)
	}

	var response TransactionInfoResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal transaction info response: %w", err)
	}

	if response.Data.TransactionInfo == nil {
		return nil, fmt.Errorf("transaction %s not found", txHash)
	}

	if c.rpcURL != "" {
		if err := c.enrichTransactionInfo(ctx, response.Data.TransactionInfo); err != nil {
			// Do not fail the call if receipt lookup fails – return partial graph data
			// while still surfacing the primary result.
		}
	}

	return response.Data.TransactionInfo, nil
}

// GetPaymentAccounts retrieves payment accounts with their token balances.
// accountIDs accepts one or more account addresses (case-insensitive).
func (c *Client) GetPaymentAccounts(ctx context.Context, accountIDs []string) ([]PaymentAccount, error) {
	if len(accountIDs) == 0 {
		return []PaymentAccount{}, nil
	}

	query := `
        query GetPaymentAccounts($ids: [ID!]) {
            paymentAccounts(where: {id_in: $ids}) {
                id
                tokenBalances {
                    balance
                    updatedAt
                    token {
                        name
                        symbol
                        decimals
                    }
                }
            }
        }
    `

	variables := map[string]interface{}{
		"ids": normalizeAddresses(accountIDs),
	}

	respBody, err := c.Query(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to query payment accounts: %w", err)
	}

	var response PaymentAccountsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payment accounts response: %w", err)
	}

	return response.Data.PaymentAccounts, nil
}

// GetCollectionAccounts retrieves collection accounts with their token balances.
func (c *Client) GetCollectionAccounts(ctx context.Context, accountIDs []string) ([]CollectionAccount, error) {
	if len(accountIDs) == 0 {
		return []CollectionAccount{}, nil
	}

	query := `
        query GetCollectionAccounts($ids: [ID!]) {
            collectionAccounts(where: {id_in: $ids}) {
                id
                tokenBalances {
                    balance
                    updatedAt
                    token {
                        name
                        symbol
                        decimals
                    }
                }
            }
        }
    `

	variables := map[string]interface{}{
		"ids": normalizeAddresses(accountIDs),
	}

	respBody, err := c.Query(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to query collection accounts: %w", err)
	}

	var response CollectionAccountsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal collection accounts response: %w", err)
	}

	return response.Data.CollectionAccounts, nil
}

// GetPaymentAuthorizations retrieves both allowances and approvals for a payment account.
func (c *Client) GetPaymentAuthorizations(ctx context.Context, account string) (*PaymentAuthorizations, error) {
	allowances, err := c.GetPaymentAllowances(ctx, account)
	if err != nil {
		return nil, err
	}

	approvals, err := c.GetPaymentApprovals(ctx, account)
	if err != nil {
		return nil, err
	}

	return &PaymentAuthorizations{
		Allowances: allowances,
		Approvals:  approvals,
	}, nil
}

// Close releases any resources held by the client.
func (c *Client) Close() {
	c.ethClientMu.Lock()
	defer c.ethClientMu.Unlock()
	if c.ethClient != nil {
		c.ethClient.Close()
		c.ethClient = nil
	}
}

func (c *Client) enrichTransactionInfo(ctx context.Context, info *TransactionInfo) error {
	ethClient, err := c.getEthClient(ctx)
	if err != nil {
		return err
	}

	receipt, err := ethClient.TransactionReceipt(ctx, common.HexToHash(info.ID))
	if err != nil {
		return err
	}

	info.GasUsed = strconv.FormatUint(receipt.GasUsed, 10)
	info.Status = strconv.FormatUint(uint64(receipt.Status), 10)

	if info.GasLimit == "" {
		info.GasLimit = strconv.FormatUint(receipt.GasUsed, 10)
	}

	if info.TransactionFee == "" {
		if gasPrice, ok := new(big.Int).SetString(info.GasPrice, 10); ok {
			fee := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(receipt.GasUsed))
			info.TransactionFee = fee.String()
		}
	}

	return nil
}

func (c *Client) getEthClient(ctx context.Context) (*ethclient.Client, error) {
	c.ethClientMu.Lock()
	defer c.ethClientMu.Unlock()

	if c.rpcURL == "" {
		return nil, fmt.Errorf("rpc url not configured")
	}

	if c.ethClient != nil {
		return c.ethClient, nil
	}

	client, err := ethclient.DialContext(ctx, c.rpcURL)
	if err != nil {
		return nil, err
	}

	c.ethClient = client
	return c.ethClient, nil
}

func normalizeAddresses(addresses []string) []string {
	normalized := make([]string, 0, len(addresses))
	for _, addr := range addresses {
		if addr == "" {
			continue
		}
		lower := strings.ToLower(addr)
		if !strings.HasPrefix(lower, "0x") {
			lower = "0x" + lower
		}
		normalized = append(normalized, lower)
	}
	return normalized
}

func normalizeAddress(address string) string {
	normalized := normalizeAddresses([]string{address})
	if len(normalized) == 0 {
		return ""
	}
	return normalized[0]
}
