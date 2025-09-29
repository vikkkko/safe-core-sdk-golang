package api

import (
	"time"
)

// SafeInfoResponse represents the response from the Safe info endpoint
type SafeInfoResponse struct {
	Address         string   `json:"address"`
	Nonce           string   `json:"nonce"`           // Changed from int64 to string to match API
	Threshold       int      `json:"threshold"`
	Owners          []string `json:"owners"`
	MasterCopy      string   `json:"masterCopy"`
	Modules         []string `json:"modules"`
	FallbackHandler string   `json:"fallbackHandler"`
	Guard           string   `json:"guard"`
	Version         string   `json:"version"`
}

// SafeMultisigTransactionResponse represents a multisig transaction from the API
type SafeMultisigTransactionResponse struct {
	Safe                    string                                    `json:"safe"`
	To                      string                                    `json:"to"`
	Value                   string                                    `json:"value"`
	Data                    string                                    `json:"data"`
	Operation               int                                       `json:"operation"`
	GasToken                string                                    `json:"gasToken"`
	SafeTxGas               int64                                     `json:"safeTxGas"`
	BaseGas                 int64                                     `json:"baseGas"`
	GasPrice                string                                    `json:"gasPrice"`
	RefundReceiver          string                                    `json:"refundReceiver"`
	Nonce                   int64                                     `json:"nonce"`
	ExecutionDate           *time.Time                                `json:"executionDate"`
	SubmissionDate          time.Time                                 `json:"submissionDate"`
	Modified                time.Time                                 `json:"modified"`
	BlockNumber             *int64                                    `json:"blockNumber"`
	TransactionHash         *string                                   `json:"transactionHash"`
	SafeTxHash              string                                    `json:"safeTxHash"`
	Executor                *string                                   `json:"executor"`
	IsExecuted              bool                                      `json:"isExecuted"`
	IsSuccessful            *bool                                     `json:"isSuccessful"`
	EthGasPrice             *string                                   `json:"ethGasPrice"`
	MaxFeePerGas            *string                                   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas    *string                                   `json:"maxPriorityFeePerGas"`
	GasUsed                 *int64                                    `json:"gasUsed"`
	Fee                     *string                                   `json:"fee"`
	Origin                  *string                                   `json:"origin"`
	DataDecoded             *DataDecoded                              `json:"dataDecoded"`
	ConfirmationsRequired   int                                       `json:"confirmationsRequired"`
	Confirmations           []SafeMultisigConfirmationResponse        `json:"confirmations"`
	Trusted                 bool                                      `json:"trusted"`
	Signatures              *string                                   `json:"signatures"`
}

// SafeMultisigConfirmationResponse represents a transaction confirmation
type SafeMultisigConfirmationResponse struct {
	Owner           string     `json:"owner"`
	SubmissionDate  time.Time  `json:"submissionDate"`
	TransactionHash *string    `json:"transactionHash"`
	Signature       string     `json:"signature"`
	SignatureType   string     `json:"signatureType"`
}

// SafeMultisigTransactionListResponse represents a list of multisig transactions
type SafeMultisigTransactionListResponse struct {
	Count    int                               `json:"count"`
	Next     *string                           `json:"next"`
	Previous *string                           `json:"previous"`
	Results  []SafeMultisigTransactionResponse `json:"results"`
}

// DataDecoded represents decoded transaction data
type DataDecoded struct {
	Method     string                   `json:"method"`
	Parameters []DataDecodedParameter   `json:"parameters"`
}

// DataDecodedParameter represents a decoded parameter
type DataDecodedParameter struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Value        interface{} `json:"value"`
	ValueDecoded interface{} `json:"valueDecoded,omitempty"`
}

// ProposeTransactionProps represents parameters for proposing a transaction
type ProposeTransactionProps struct {
	SafeAddress             string  `json:"safeAddress"`
	SafeTxHash              string  `json:"safeTxHash"`
	To                      string  `json:"to"`
	Value                   string  `json:"value"`
	Data                    string  `json:"data"`
	Operation               int     `json:"operation"`
	GasToken                string  `json:"gasToken"`
	SafeTxGas               int64   `json:"safeTxGas"`
	BaseGas                 int64   `json:"baseGas"`
	GasPrice                string  `json:"gasPrice"`
	RefundReceiver          string  `json:"refundReceiver"`
	Nonce                   int64   `json:"nonce"`
	Sender                  string  `json:"sender"`
	Signature               string  `json:"signature"`
	ContractTransactionHash string  `json:"contractTransactionHash"`
	Origin                  *string `json:"origin,omitempty"`
}

// SignatureResponse represents a signature response
type SignatureResponse struct {
	Signature string `json:"signature"`
}

// GetMultisigTransactionsOptions represents options for getting multisig transactions
type GetMultisigTransactionsOptions struct {
	Executed    *bool `json:"executed,omitempty"`
	TrustedOnly *bool `json:"trusted,omitempty"`
	Limit       *int  `json:"limit,omitempty"`
	Offset      *int  `json:"offset,omitempty"`
}

// PendingTransactionsOptions represents options for getting pending transactions
type PendingTransactionsOptions struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// GetIncomingTransactionsOptions represents options for getting incoming transactions
type GetIncomingTransactionsOptions struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// TransferListResponse represents a list of transfers
type TransferListResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Transfer `json:"results"`
}

// Transfer represents a token transfer
type Transfer struct {
	Type             string     `json:"type"`
	ExecutionDate    time.Time  `json:"executionDate"`
	BlockNumber      int64      `json:"blockNumber"`
	TransactionHash  string     `json:"transactionHash"`
	To               string     `json:"to"`
	Value            string     `json:"value"`
	TokenId          *string    `json:"tokenId,omitempty"`
	TokenAddress     *string    `json:"tokenAddress,omitempty"`
	From             string     `json:"from"`
}

// SafesByOwnerResponse represents Safes owned by an address
type SafesByOwnerResponse struct {
	Safes []string `json:"safes"`
}

// SafeMultisigTransactionEstimate represents transaction estimation parameters
type SafeMultisigTransactionEstimate struct {
	To        string `json:"to"`
	Value     string `json:"value"`
	Data      string `json:"data"`
	Operation int    `json:"operation"`
}

// SafeMultisigTransactionEstimateResponse represents transaction estimation response
type SafeMultisigTransactionEstimateResponse struct {
	SafeTxGas string `json:"safeTxGas"`
}

// SafeCreationInfoResponse represents Safe creation information
type SafeCreationInfoResponse struct {
	Created           time.Time `json:"created"`
	Creator           string    `json:"creator"`
	TransactionHash   string    `json:"transactionHash"`
	FactoryAddress    string    `json:"factoryAddress"`
	MasterCopy        string    `json:"masterCopy"`
	SetupData         string    `json:"setupData"`
	DataDecoded       *DataDecoded `json:"dataDecoded,omitempty"`
}

// SafeServiceInfoResponse represents Safe service information
type SafeServiceInfoResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	BuildNumber string `json:"buildNumber"`
}

// OwnerResponse represents owner information
type OwnerResponse struct {
	SafeAddress string `json:"safeAddress"`
	Owners      []string `json:"owners"`
}

// ModulesResponse represents enabled modules
type ModulesResponse struct {
	SafeAddress string   `json:"safeAddress"`
	Modules     []string `json:"modules"`
}

// TokenInfoResponse represents token information
type TokenInfoResponse struct {
	Type     string `json:"type"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	LogoUri  string `json:"logoUri,omitempty"`
}

// TokenInfoListResponse represents a list of tokens
type TokenInfoListResponse struct {
	Count    int                 `json:"count"`
	Next     *string             `json:"next"`
	Previous *string             `json:"previous"`
	Results  []TokenInfoResponse `json:"results"`
}

// TokenInfoListOptions represents options for getting token info
type TokenInfoListOptions struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// SafeMessage represents a Safe message
type SafeMessage struct {
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	Safe         string    `json:"safe"`
	MessageHash  string    `json:"messageHash"`
	Message      interface{} `json:"message"`
	ProposedBy   string    `json:"proposedBy"`
	SafeAppId    *int64    `json:"safeAppId,omitempty"`
	Confirmations []SafeMessageConfirmation `json:"confirmations"`
	PreparedSignature *string `json:"preparedSignature,omitempty"`
}

// SafeMessageConfirmation represents a message confirmation
type SafeMessageConfirmation struct {
	Created        time.Time `json:"created"`
	Modified       time.Time `json:"modified"`
	Owner          string    `json:"owner"`
	Signature      string    `json:"signature"`
	SignatureType  string    `json:"signatureType"`
}

// SafeMessageListResponse represents a list of Safe messages
type SafeMessageListResponse struct {
	Count    int           `json:"count"`
	Next     *string       `json:"next"`
	Previous *string       `json:"previous"`
	Results  []SafeMessage `json:"results"`
}

// GetSafeMessageListOptions represents options for getting Safe messages
type GetSafeMessageListOptions struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// AddMessageOptions represents options for adding a message
type AddMessageOptions struct {
	Message        interface{} `json:"message"`
	SafeAppId      *int64      `json:"safeAppId,omitempty"`
	Signature      string      `json:"signature"`
}

// AllTransactionsListResponse represents all transactions (multisig + module + incoming)
type AllTransactionsListResponse struct {
	Count    int                    `json:"count"`
	Next     *string                `json:"next"`
	Previous *string                `json:"previous"`
	Results  []AllTransactionResult `json:"results"`
}

// AllTransactionResult represents any type of transaction result
type AllTransactionResult struct {
	Type                    string                            `json:"type"` // "MULTISIG_TRANSACTION", "MODULE_TRANSACTION", "ETHEREUM_TRANSACTION"
	MultisigTransaction     *SafeMultisigTransactionResponse  `json:"multisigTransaction,omitempty"`
	ModuleTransaction       *SafeModuleTransactionResponse    `json:"moduleTransaction,omitempty"`
	EthereumTransaction     *SafeEthereumTransactionResponse  `json:"ethereumTransaction,omitempty"`
}

// SafeModuleTransactionResponse represents a module transaction
type SafeModuleTransactionResponse struct {
	Created          time.Time    `json:"created"`
	ExecutionDate    time.Time    `json:"executionDate"`
	BlockNumber      int64        `json:"blockNumber"`
	IsSuccessful     bool         `json:"isSuccessful"`
	TransactionHash  string       `json:"transactionHash"`
	Safe             string       `json:"safe"`
	Module           string       `json:"module"`
	To               string       `json:"to"`
	Value            string       `json:"value"`
	Data             string       `json:"data"`
	Operation        int          `json:"operation"`
	DataDecoded      *DataDecoded `json:"dataDecoded,omitempty"`
}

// SafeEthereumTransactionResponse represents an Ethereum transaction
type SafeEthereumTransactionResponse struct {
	ExecutionDate   time.Time `json:"executionDate"`
	To              string    `json:"to"`
	Data            string    `json:"data"`
	TxHash          string    `json:"txHash"`
	BlockNumber     int64     `json:"blockNumber"`
	Transfers       []Transfer `json:"transfers,omitempty"`
	From            string    `json:"from"`
}

// AllTransactionsOptions represents options for getting all transactions
type AllTransactionsOptions struct {
	Executed    *bool   `json:"executed,omitempty"`
	Queued      *bool   `json:"queued,omitempty"`
	Trusted     *bool   `json:"trusted,omitempty"`
	Limit       *int    `json:"limit,omitempty"`
	Offset      *int    `json:"offset,omitempty"`
}