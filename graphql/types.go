package graphql

import (
	"encoding/json"
	"fmt"
	"math/big"
)

// BigInt is a custom type for handling big integers from GraphQL
// GraphQL returns numbers as strings, so we need custom unmarshaling
type BigInt struct {
	*big.Int
}

// UnmarshalJSON implements json.Unmarshaler for BigInt
func (b *BigInt) UnmarshalJSON(data []byte) error {
	// Remove quotes from the string
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	// Parse the string as a big integer
	bi := new(big.Int)
	if _, ok := bi.SetString(s, 10); !ok {
		return fmt.Errorf("invalid big integer: %s", s)
	}

	b.Int = bi
	return nil
}

// MarshalJSON implements json.Marshaler for BigInt
func (b BigInt) MarshalJSON() ([]byte, error) {
	if b.Int == nil {
		return []byte("\"0\""), nil
	}
	return json.Marshal(b.String())
}

// PaymentAllowance represents an allowance record from the graph
type PaymentAllowance struct {
	Token  string `json:"token"`  // Token contract address
	Owner  string `json:"owner"`  // Address that granted the allowance
	Amount BigInt `json:"amount"` // Allowance amount
}

// PaymentAllowancesResponse represents the GraphQL response for paymentAllowances query
type PaymentAllowancesResponse struct {
	Data struct {
		PaymentAllowances []PaymentAllowance `json:"paymentAllowances"`
	} `json:"data"`
}

// PaymentApproval represents an approval event for a payment account
type PaymentApproval struct {
	Token     string `json:"token"`     // Token contract address
	Spender   string `json:"spender"`   // Spender address
	Amount    BigInt `json:"amount"`    // Approved amount
	Timestamp string `json:"timestamp"` // Timestamp of the approval
	TxHash    string `json:"txHash"`    // Transaction hash
}

// PaymentApprovalsResponse represents GraphQL response for payment approvals
type PaymentApprovalsResponse struct {
	Data struct {
		PaymentApprovals []PaymentApproval `json:"paymentApprovals"`
	} `json:"data"`
}

// TransactionInfo represents detailed info of a transaction returned by the graph
type TransactionInfo struct {
	ID             string `json:"id"`             // Transaction hash
	BlockNumber    string `json:"blockNumber"`    // Block number containing the transaction
	GasLimit       string `json:"gasLimit"`       // Gas limit set for the transaction
	GasPrice       string `json:"gasPrice"`       // Gas price used for the transaction
	GasUsed        string `json:"gasUsed"`        // Gas actually used
	Status         string `json:"status"`         // Transaction execution status
	Timestamp      string `json:"timestamp"`      // Block timestamp
	TransactionFee string `json:"transactionFee"` // Total transaction fee
}

// TransactionInfoResponse represents the GraphQL response for transactionInfo query
type TransactionInfoResponse struct {
	Data struct {
		TransactionInfo *TransactionInfo `json:"transactionInfo"`
	} `json:"data"`
}

// TokenMetadata represents ERC20 token metadata
type TokenMetadata struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int32  `json:"decimals,string"`
}

// TokenBalance represents a token balance entry for payment/collection account
type TokenBalance struct {
	Balance   BigInt        `json:"balance"`
	UpdatedAt string        `json:"updatedAt"`
	Token     TokenMetadata `json:"token"`
}

// PaymentAccount represents data returned for payment account token balances
type PaymentAccount struct {
	ID            string         `json:"id"`
	TokenBalances []TokenBalance `json:"tokenBalances"`
}

// CollectionAccount represents data returned for collection account token balances
type CollectionAccount struct {
	ID            string         `json:"id"`
	TokenBalances []TokenBalance `json:"tokenBalances"`
}

// PaymentAccountsResponse represents GraphQL response for paymentAccounts query
type PaymentAccountsResponse struct {
	Data struct {
		PaymentAccounts []PaymentAccount `json:"paymentAccounts"`
	} `json:"data"`
}

// CollectionAccountsResponse represents GraphQL response for collectionAccounts query
type CollectionAccountsResponse struct {
	Data struct {
		CollectionAccounts []CollectionAccount `json:"collectionAccounts"`
	} `json:"data"`
}

// PaymentAuthorizations aggregates allowance and approval data
type PaymentAuthorizations struct {
	Allowances []PaymentAllowance `json:"allowances"`
	Approvals  []PaymentApproval  `json:"approvals"`
}

// GraphQLRequest represents a GraphQL request body
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQLError represents a GraphQL error
type GraphQLError struct {
	Message string `json:"message"`
}

// GraphQLResponse represents a generic GraphQL response
type GraphQLResponse struct {
	Data   interface{}    `json:"data"`
	Errors []GraphQLError `json:"errors,omitempty"`
}
