package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// SafeApiKitConfig represents the configuration for the Safe API client
type SafeApiKitConfig struct {
	ChainID      int64  `json:"chainId"`                // Chain ID
	TxServiceURL string `json:"txServiceUrl,omitempty"` // Custom transaction service URL
	ApiKey       string `json:"apiKey,omitempty"`       // API key for Safe Transaction Service
}

// SafeApiKit provides methods to interact with Safe Transaction Service API
type SafeApiKit struct {
	chainID       int64
	apiKey        string
	baseURL       string
	httpClient    *http.Client
}

// NewSafeApiKit creates a new Safe API client
func NewSafeApiKit(config SafeApiKitConfig) (*SafeApiKit, error) {
	var baseURL string

	if config.TxServiceURL != "" {
		// If custom URL contains safe.global or 5afe.dev, API key is mandatory
		if (strings.Contains(config.TxServiceURL, "api.safe.global") ||
			strings.Contains(config.TxServiceURL, "api.5afe.dev")) &&
			config.ApiKey == "" {
			return nil, fmt.Errorf("apiKey is mandatory when using api.safe.global or api.5afe.dev domains")
		}
		baseURL = config.TxServiceURL
	} else {
		// If no custom URL, API key is mandatory and we use default URL
		if config.ApiKey == "" {
			return nil, fmt.Errorf("apiKey is mandatory when txServiceUrl is not defined")
		}

		defaultURL, err := getTransactionServiceURL(config.ChainID)
		if err != nil {
			return nil, fmt.Errorf("failed to get transaction service URL: %w", err)
		}
		baseURL = defaultURL
	}

	return &SafeApiKit{
		chainID:    config.ChainID,
		apiKey:     config.ApiKey,
		baseURL:    strings.TrimSuffix(baseURL, "/"),
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}, nil
}

// GetSafeInfo retrieves information about a Safe
func (api *SafeApiKit) GetSafeInfo(ctx context.Context, safeAddress string) (*SafeInfoResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/safes/%s/", safeAddress)

	var response SafeInfoResponse
	err := api.makeRequest(ctx, "GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe info: %w", err)
	}

	return &response, nil
}

// GetMultisigTransactions retrieves multisig transactions for a Safe
func (api *SafeApiKit) GetMultisigTransactions(ctx context.Context, safeAddress string, options *GetMultisigTransactionsOptions) (*SafeMultisigTransactionListResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/safes/%s/multisig-transactions/", safeAddress)

	// Add query parameters
	if options != nil {
		params := url.Values{}
		if options.Executed != nil {
			params.Add("executed", strconv.FormatBool(*options.Executed))
		}
		if options.TrustedOnly != nil {
			params.Add("trusted", strconv.FormatBool(*options.TrustedOnly))
		}
		if options.Limit != nil {
			params.Add("limit", strconv.Itoa(*options.Limit))
		}
		if options.Offset != nil {
			params.Add("offset", strconv.Itoa(*options.Offset))
		}
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var response SafeMultisigTransactionListResponse
	err := api.makeRequest(ctx, "GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get multisig transactions: %w", err)
	}

	return &response, nil
}

// GetMultisigTransaction retrieves a specific multisig transaction
func (api *SafeApiKit) GetMultisigTransaction(ctx context.Context, safeTxHash string) (*SafeMultisigTransactionResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/multisig-transactions/%s/", safeTxHash)

	var response SafeMultisigTransactionResponse
	err := api.makeRequest(ctx, "GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get multisig transaction: %w", err)
	}

	return &response, nil
}

// ProposeTransaction proposes a new transaction to the Safe
func (api *SafeApiKit) ProposeTransaction(ctx context.Context, params ProposeTransactionProps) (*SafeMultisigTransactionResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/safes/%s/multisig-transactions/", params.SafeAddress)

	var response SafeMultisigTransactionResponse
	err := api.makeRequest(ctx, "POST", endpoint, params, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to propose transaction: %w", err)
	}

	return &response, nil
}

// ConfirmTransaction confirms a proposed transaction
func (api *SafeApiKit) ConfirmTransaction(ctx context.Context, safeTxHash string, signature string) (*SignatureResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/multisig-transactions/%s/confirmations/", safeTxHash)

	requestBody := map[string]string{
		"signature": signature,
	}

	var response SignatureResponse
	err := api.makeRequest(ctx, "POST", endpoint, requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to confirm transaction: %w", err)
	}

	return &response, nil
}

// GetPendingTransactions retrieves pending transactions for a Safe
func (api *SafeApiKit) GetPendingTransactions(ctx context.Context, safeAddress string, options *PendingTransactionsOptions) (*SafeMultisigTransactionListResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/safes/%s/multisig-transactions/", safeAddress)

	params := url.Values{}
	params.Add("executed", "false")

	if options != nil {
		if options.Limit != nil {
			params.Add("limit", strconv.Itoa(*options.Limit))
		}
		if options.Offset != nil {
			params.Add("offset", strconv.Itoa(*options.Offset))
		}
	}

	endpoint += "?" + params.Encode()

	var response SafeMultisigTransactionListResponse
	err := api.makeRequest(ctx, "GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending transactions: %w", err)
	}

	return &response, nil
}

// GetIncomingTransactions retrieves incoming transactions for a Safe
func (api *SafeApiKit) GetIncomingTransactions(ctx context.Context, safeAddress string, options *GetIncomingTransactionsOptions) (*TransferListResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/safes/%s/incoming-transfers/", safeAddress)

	if options != nil {
		params := url.Values{}
		if options.Limit != nil {
			params.Add("limit", strconv.Itoa(*options.Limit))
		}
		if options.Offset != nil {
			params.Add("offset", strconv.Itoa(*options.Offset))
		}
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var response TransferListResponse
	err := api.makeRequest(ctx, "GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get incoming transactions: %w", err)
	}

	return &response, nil
}

// GetSafesByOwner retrieves Safes owned by a specific address
func (api *SafeApiKit) GetSafesByOwner(ctx context.Context, ownerAddress string) (*SafesByOwnerResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/owners/%s/safes/", ownerAddress)

	var response SafesByOwnerResponse
	err := api.makeRequest(ctx, "GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get Safes by owner: %w", err)
	}

	return &response, nil
}

// EstimateTransaction estimates gas for a Safe transaction
func (api *SafeApiKit) EstimateTransaction(ctx context.Context, safeAddress string, txData SafeMultisigTransactionEstimate) (*SafeMultisigTransactionEstimateResponse, error) {
	endpoint := fmt.Sprintf("/api/v1/safes/%s/multisig-transactions/estimations/", safeAddress)

	var response SafeMultisigTransactionEstimateResponse
	err := api.makeRequest(ctx, "POST", endpoint, txData, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate transaction: %w", err)
	}

	return &response, nil
}

// makeRequest makes an HTTP request to the Safe Transaction Service API
func (api *SafeApiKit) makeRequest(ctx context.Context, method, endpoint string, body interface{}, result interface{}) error {
	// Prepare request body
	var requestBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		requestBody = bytes.NewReader(jsonBody)
	}

	// Create request
	url := api.baseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, method, url, requestBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if api.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+api.apiKey)
	}

	// Make request
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(responseBody))
	}

	// Parse response - handle empty responses for successful operations
	if result != nil && len(responseBody) > 0 {
		err = json.Unmarshal(responseBody, result)
		if err != nil {
			// Add debug info for JSON parsing errors
			return fmt.Errorf("failed to unmarshal response: %w (response body: %s)", err, string(responseBody))
		}
	}

	return nil
}

// getTransactionServiceURL returns the default transaction service URL for a chain
func getTransactionServiceURL(chainID int64) (string, error) {
	// Default URLs for common chains with correct endpoints
	switch chainID {
	case 1: // Ethereum Mainnet
		return "https://api.safe.global/tx-service/eth", nil
	case 5: // Goerli
		return "https://api.safe.global/tx-service/gor", nil
	case 11155111: // Sepolia
		return "https://api.safe.global/tx-service/sep", nil
	case 137: // Polygon
		return "https://api.safe.global/tx-service/matic", nil
	case 56: // BSC
		return "https://api.safe.global/tx-service/bsc", nil
	case 42161: // Arbitrum
		return "https://api.safe.global/tx-service/arb1", nil
	case 10: // Optimism
		return "https://api.safe.global/tx-service/oeth", nil
	default:
		return "", fmt.Errorf("unsupported chain ID: %d", chainID)
	}
}