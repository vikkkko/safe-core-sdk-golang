package types

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strings"
)

// SafeVersion represents the supported Safe contract versions
type SafeVersion string

const (
	SafeVersion141 SafeVersion = "1.4.1"
	SafeVersion130 SafeVersion = "1.3.0"
	SafeVersion120 SafeVersion = "1.2.0"
	SafeVersion111 SafeVersion = "1.1.1"
	SafeVersion100 SafeVersion = "1.0.0"
)

// OperationType represents the type of operation for a transaction
type OperationType uint8

const (
	Call         OperationType = 0 // Standard call
	DelegateCall OperationType = 1 // Delegate call
)

// SafeSetupConfig contains the configuration for setting up a new Safe
type SafeSetupConfig struct {
	Owners          []string `json:"owners"`                    // List of Safe owners
	Threshold       uint     `json:"threshold"`                 // Number of required confirmations
	To              string   `json:"to,omitempty"`              // Optional contract to call during setup
	Data            string   `json:"data,omitempty"`            // Optional data payload for setup call
	FallbackHandler string   `json:"fallbackHandler,omitempty"` // Optional fallback handler
	PaymentToken    string   `json:"paymentToken,omitempty"`    // Optional payment token
	Payment         string   `json:"payment,omitempty"`         // Optional payment amount
	PaymentReceiver string   `json:"paymentReceiver,omitempty"` // Optional payment receiver
}

// MetaTransactionData represents the essential data for a transaction
type MetaTransactionData struct {
	To        string         `json:"to"`                  // Target address
	Value     string         `json:"value"`               // Value in wei
	Data      string         `json:"data"`                // Transaction data
	Operation *OperationType `json:"operation,omitempty"` // Operation type (optional)
}

// SafeTransactionData represents complete transaction data for Safe
type SafeTransactionData struct {
	To             string        `json:"to"`             // Target address
	Value          string        `json:"value"`          // Value in wei
	Data           string        `json:"data"`           // Transaction data
	Operation      OperationType `json:"operation"`      // Operation type
	SafeTxGas      string        `json:"safeTxGas"`      // Gas for Safe transaction
	BaseGas        string        `json:"baseGas"`        // Base gas cost
	GasPrice       string        `json:"gasPrice"`       // Gas price
	GasToken       string        `json:"gasToken"`       // Gas token address
	RefundReceiver string        `json:"refundReceiver"` // Refund receiver address
	Nonce          uint64        `json:"nonce"`          // Transaction nonce
}

// SafeTransactionDataPartial represents partial transaction data (for creating transactions)
type SafeTransactionDataPartial struct {
	To             string         `json:"to"`                       // Target address
	Value          string         `json:"value"`                    // Value in wei
	Data           string         `json:"data"`                     // Transaction data
	Operation      *OperationType `json:"operation,omitempty"`      // Operation type (optional)
	SafeTxGas      *string        `json:"safeTxGas,omitempty"`      // Gas for Safe transaction (optional)
	BaseGas        *string        `json:"baseGas,omitempty"`        // Base gas cost (optional)
	GasPrice       *string        `json:"gasPrice,omitempty"`       // Gas price (optional)
	GasToken       *string        `json:"gasToken,omitempty"`       // Gas token address (optional)
	RefundReceiver *string        `json:"refundReceiver,omitempty"` // Refund receiver address (optional)
	Nonce          *uint64        `json:"nonce,omitempty"`          // Transaction nonce (optional)
}

// SafeSignature represents a signature for a Safe transaction or message
type SafeSignature struct {
	Signer              string `json:"signer"`              // Address of the signer
	Data                string `json:"data"`                // Signature data
	IsContractSignature bool   `json:"isContractSignature"` // Whether this is a contract signature
}

// StaticPart returns the static part of the signature
func (s *SafeSignature) StaticPart(dynamicOffset string) string {
	if s.IsContractSignature {
		if dynamicOffset == "" {
			return s.Data
		}
		return dynamicOffset
	}
	return s.Data
}

// DynamicPart returns the dynamic part of the signature
func (s *SafeSignature) DynamicPart() string {
	if s.IsContractSignature {
		return s.Data
	}
	return ""
}

// SafeTransaction represents a Safe transaction with signatures
type SafeTransaction struct {
	Data       SafeTransactionData      `json:"data"`       // Transaction data
	Signatures map[string]SafeSignature `json:"signatures"` // Map of signer address to signature
}

// GetSignature returns the signature for a specific signer
func (st *SafeTransaction) GetSignature(signer string) *SafeSignature {
	if sig, exists := st.Signatures[signer]; exists {
		return &sig
	}
	return nil
}

// AddSignature adds a signature to the transaction
func (st *SafeTransaction) AddSignature(signature SafeSignature) {
	if st.Signatures == nil {
		st.Signatures = make(map[string]SafeSignature)
	}
	st.Signatures[signature.Signer] = signature
}

// EncodedSignatures returns the encoded signatures for the transaction in hex format (0x...)
func (st *SafeTransaction) EncodedSignatures() string {
	bytes, err := st.EncodedSignaturesBytes()
	if err != nil {
		return ""
	}
	return "0x" + hex.EncodeToString(bytes)
}

// EncodedSignaturesBytes returns the raw signature bytes sorted by signer address.
func (st *SafeTransaction) EncodedSignaturesBytes() ([]byte, error) {
	if len(st.Signatures) == 0 {
		return nil, fmt.Errorf("no signatures present")
	}

	type signatureEntry struct {
		signer string
		sig    SafeSignature
	}

	entries := make([]signatureEntry, 0, len(st.Signatures))
	for signer, sig := range st.Signatures {
		s := signer
		if s == "" {
			s = sig.Signer
		}
		s = strings.ToLower(s)
		entries = append(entries, signatureEntry{signer: s, sig: sig})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].signer < entries[j].signer
	})

	result := make([]byte, 0, len(entries)*65)
	for _, entry := range entries {
		if entry.sig.IsContractSignature {
			return nil, fmt.Errorf("contract signatures are not supported in this encoding helper")
		}

		data := strings.TrimPrefix(entry.sig.Data, "0x")
		sigBytes, err := hex.DecodeString(data)
		if err != nil {
			return nil, fmt.Errorf("failed to decode signature for %s: %w", entry.sig.Signer, err)
		}
		if len(sigBytes) != 65 {
			return nil, fmt.Errorf("signature for %s has invalid length %d", entry.sig.Signer, len(sigBytes))
		}
		result = append(result, sigBytes...)
	}

	return result, nil
}

// SafeMessage represents a Safe message with signatures
type SafeMessage struct {
	Data       interface{}              `json:"data"`       // Message data (can be EIP712 typed data or string)
	Signatures map[string]SafeSignature `json:"signatures"` // Map of signer address to signature
}

// GetSignature returns the signature for a specific signer
func (sm *SafeMessage) GetSignature(signer string) *SafeSignature {
	if sig, exists := sm.Signatures[signer]; exists {
		return &sig
	}
	return nil
}

// AddSignature adds a signature to the message
func (sm *SafeMessage) AddSignature(signature SafeSignature) {
	if sm.Signatures == nil {
		sm.Signatures = make(map[string]SafeSignature)
	}
	sm.Signatures[signature.Signer] = signature
}

// EncodedSignatures returns the encoded signatures for the message
func (sm *SafeMessage) EncodedSignatures() string {
	// Implementation would encode all signatures according to Safe's format
	// This is a placeholder - actual implementation would depend on the signing format
	return ""
}

// TransactionBase represents basic transaction information
type TransactionBase struct {
	To    string `json:"to"`    // Target address
	Value string `json:"value"` // Value in wei
	Data  string `json:"data"`  // Transaction data
}

// TransactionOptions represents optional transaction parameters
type TransactionOptions struct {
	From                 *string  `json:"from,omitempty"`                 // From address
	GasLimit             *big.Int `json:"gasLimit,omitempty"`             // Gas limit
	GasPrice             *big.Int `json:"gasPrice,omitempty"`             // Gas price
	MaxFeePerGas         *big.Int `json:"maxFeePerGas,omitempty"`         // EIP-1559 max fee per gas
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas,omitempty"` // EIP-1559 max priority fee per gas
	Nonce                *uint64  `json:"nonce,omitempty"`                // Transaction nonce
}

// Transaction combines base transaction data with options
type Transaction struct {
	TransactionBase
	TransactionOptions
}

// BaseTransactionResult represents basic transaction result
type BaseTransactionResult struct {
	Hash string `json:"hash"` // Transaction hash
}

// TransactionResult represents complete transaction result
type TransactionResult struct {
	BaseTransactionResult
	TransactionResponse interface{}         `json:"transactionResponse"` // Provider-specific transaction response
	Options             *TransactionOptions `json:"options,omitempty"`   // Transaction options used
}

// EIP3770Address represents an address with chain prefix (EIP-3770)
type EIP3770Address struct {
	Prefix  string `json:"prefix"`  // Chain prefix (e.g., "eth")
	Address string `json:"address"` // Ethereum address
}

// String returns the EIP-3770 formatted address
func (addr *EIP3770Address) String() string {
	return addr.Prefix + ":" + addr.Address
}

// EIP712TypedData represents EIP-712 typed data structure
type EIP712TypedData struct {
	Types       map[string][]EIP712Type `json:"types"`       // Type definitions
	PrimaryType string                  `json:"primaryType"` // Primary type name
	Domain      EIP712Domain            `json:"domain"`      // Domain separator
	Message     map[string]interface{}  `json:"message"`     // Message data
}

// EIP712Type represents a single type definition in EIP-712
type EIP712Type struct {
	Name string `json:"name"` // Field name
	Type string `json:"type"` // Field type
}

// EIP712Domain represents the EIP-712 domain separator
type EIP712Domain struct {
	Name              *string  `json:"name,omitempty"`              // Domain name
	Version           *string  `json:"version,omitempty"`           // Domain version
	ChainId           *big.Int `json:"chainId,omitempty"`           // Chain ID
	VerifyingContract *string  `json:"verifyingContract,omitempty"` // Verifying contract address
	Salt              *string  `json:"salt,omitempty"`              // Salt for domain separation
}

// SigningMethod represents different methods of signing
type SigningMethod string

const (
	SigningMethodETHSign          SigningMethod = "eth_sign"
	SigningMethodETHSignTypedData SigningMethod = "eth_signTypedData"
	SigningMethodSafeSignMessage  SigningMethod = "safe_signMessage"
)

// SigningMethodType represents the type of signing method
type SigningMethodType struct {
	Method SigningMethod `json:"method"`
}
