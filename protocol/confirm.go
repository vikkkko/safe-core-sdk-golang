package protocol

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vikkkko/safe-core-sdk-golang/api"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
	"github.com/vikkkko/safe-core-sdk-golang/types"
)

// ConfirmTransactionConfig contains configuration for confirming a Safe transaction
type ConfirmTransactionConfig struct {
	SafeTxHash  string          // Safe transaction hash to confirm
	APIClient   *api.SafeApiKit // API client for Safe Transaction Service
	AutoExecute bool            // Whether to automatically execute when threshold is met (default: true)
}

// ConfirmTransactionResult contains the result of confirming a transaction
type ConfirmTransactionResult struct {
	AlreadySigned       bool                     // Whether the current signer already signed
	SignatureSubmitted  bool                     // Whether a new signature was submitted
	ThresholdMet        bool                     // Whether signature threshold is met
	TransactionExecuted bool                     // Whether transaction was executed
	ExecutionResult     *types.TransactionResult // Execution result if transaction was executed
	CurrentSignatures   int                      // Current number of signatures
	RequiredSignatures  int                      // Required number of signatures
}

// ConfirmTransaction confirms an existing Safe transaction by safeTxHash
// This is a high-level SDK method that simplifies the multisig workflow:
// 1. Fetches transaction details from Safe Transaction Service
// 2. Checks if current signer already signed
// 3. Signs and submits signature if not already signed
// 4. Optionally auto-executes if threshold is met
func (s *Safe) ConfirmTransaction(ctx context.Context, config ConfirmTransactionConfig) (*ConfirmTransactionResult, error) {
	if config.SafeTxHash == "" {
		return nil, fmt.Errorf("safeTxHash is required")
	}

	if config.APIClient == nil {
		return nil, fmt.Errorf("APIClient is required")
	}

	if s.config.PrivateKey == "" {
		return nil, fmt.Errorf("private key is required to confirm transaction")
	}

	// Ensure safeTxHash has 0x prefix
	safeTxHash := config.SafeTxHash
	if !strings.HasPrefix(safeTxHash, "0x") && !strings.HasPrefix(safeTxHash, "0X") {
		safeTxHash = "0x" + safeTxHash
	}

	result := &ConfirmTransactionResult{}

	autoExecute := config.AutoExecute

	// 1. Fetch transaction details from Safe Transaction Service
	txDetails, err := config.APIClient.GetMultisigTransaction(ctx, safeTxHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction details: %w", err)
	}

	result.CurrentSignatures = len(txDetails.Confirmations)
	result.RequiredSignatures = txDetails.ConfirmationsRequired

	// Check if already executed
	if txDetails.IsExecuted {
		return nil, fmt.Errorf("transaction already executed")
	}

	// 2. Check if current signer already signed
	privateKeyHex := strings.TrimPrefix(s.config.PrivateKey, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	signerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	currentSignerAddr := strings.ToLower(signerAddress.Hex())

	alreadySigned := false
	for _, confirmation := range txDetails.Confirmations {
		if strings.EqualFold(confirmation.Owner, currentSignerAddr) {
			alreadySigned = true
			break
		}
	}

	result.AlreadySigned = alreadySigned

	// 3. Sign and submit if not already signed and more signatures are required
	if !alreadySigned && result.CurrentSignatures < result.RequiredSignatures {
		// Convert safeTxHash string to bytes for signing
		// safeTxHash from API is already the correct transaction hash
		txHash := common.FromHex(safeTxHash)
		if len(txHash) != 32 {
			return nil, fmt.Errorf("invalid safeTxHash length: expected 32 bytes, got %d", len(txHash))
		}

		// Sign the transaction
		signature, err := utils.SignMessage(txHash, privateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to sign transaction: %w", err)
		}

		// Submit signature to Safe Transaction Service
		_, err = config.APIClient.ConfirmTransaction(ctx, safeTxHash, "0x"+hex.EncodeToString(signature))
		if err != nil {
			return nil, fmt.Errorf("failed to submit signature: %w", err)
		}

		result.SignatureSubmitted = true

		// Update signature count
		result.CurrentSignatures++
	}

	// 4. Check if threshold is met
	result.ThresholdMet = result.CurrentSignatures >= result.RequiredSignatures

	fmt.Printf("autoExecute: %v\n", autoExecute)
	// 5. Auto-execute if threshold met and auto-execute enabled
	if result.ThresholdMet && autoExecute {
		// Fetch updated transaction details
		txDetails, err = config.APIClient.GetMultisigTransaction(ctx, safeTxHash)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch updated transaction details: %w", err)
		}

		// Rebuild transaction with all signatures
		safeTx, err := buildSafeTransactionFromAPIResponse(txDetails)
		if err != nil {
			return nil, fmt.Errorf("failed to rebuild transaction for execution: %w", err)
		}

		// Execute the transaction
		execResult, err := s.ExecuteTransaction(ctx, safeTx)
		if err != nil {
			return nil, fmt.Errorf("failed to execute transaction: %w", err)
		}

		result.TransactionExecuted = true
		result.ExecutionResult = execResult
	}

	return result, nil
}

// buildSafeTransactionFromAPIResponse rebuilds a SafeTransaction from API response
func buildSafeTransactionFromAPIResponse(txDetails *api.SafeMultisigTransactionResponse) (*types.SafeTransaction, error) {
	if txDetails == nil {
		return nil, fmt.Errorf("transaction details cannot be nil")
	}

	// Parse operation type
	operation := types.OperationType(txDetails.Operation)

	// Build transaction data
	txData := types.SafeTransactionData{
		To:             txDetails.To,
		Value:          txDetails.Value,
		Data:           txDetails.Data,
		Operation:      operation,
		SafeTxGas:      fmt.Sprintf("%d", txDetails.SafeTxGas),
		BaseGas:        fmt.Sprintf("%d", txDetails.BaseGas),
		GasPrice:       txDetails.GasPrice,
		GasToken:       txDetails.GasToken,
		RefundReceiver: txDetails.RefundReceiver,
		Nonce:          uint64(txDetails.Nonce),
	}

	// Build signatures map
	signatures := make(map[string]types.SafeSignature)
	for _, confirmation := range txDetails.Confirmations {
		sig := types.SafeSignature{
			Signer: confirmation.Owner,
			Data:   confirmation.Signature,
		}
		signatures[strings.ToLower(confirmation.Owner)] = sig
	}

	return &types.SafeTransaction{
		Data:       txData,
		Signatures: signatures,
	}, nil
}
