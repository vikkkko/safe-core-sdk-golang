package managers

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yinwei/safe-core-sdk-golang/protocol/contracts"
)

// FallbackHandlerManager manages Safe fallback handlers
type FallbackHandlerManager struct {
	client      *ethclient.Client
	safeAddress common.Address
}

// NewFallbackHandlerManager creates a new fallback handler manager
func NewFallbackHandlerManager(client *ethclient.Client, safeAddress common.Address) *FallbackHandlerManager {
	return &FallbackHandlerManager{
		client:      client,
		safeAddress: safeAddress,
	}
}

// GetFallbackHandler returns the current fallback handler address
func (fhm *FallbackHandlerManager) GetFallbackHandler(ctx context.Context) (common.Address, error) {
	safeContract, err := contracts.NewSafeContract(fhm.safeAddress, fhm.client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create Safe contract instance: %w", err)
	}

	return safeContract.GetFallbackHandler(ctx)
}

// SetFallbackHandlerTxParams represents parameters for setting a fallback handler
type SetFallbackHandlerTxParams struct {
	FallbackHandlerAddress string `json:"fallbackHandlerAddress"` // Address of the fallback handler contract
}

// CreateSetFallbackHandlerTx creates a transaction to set a fallback handler
func (fhm *FallbackHandlerManager) CreateSetFallbackHandlerTx(ctx context.Context, params SetFallbackHandlerTxParams) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate the fallback handler address
	// 2. Create the transaction data for setting the fallback handler
	// 3. Return the encoded transaction data

	if !common.IsHexAddress(params.FallbackHandlerAddress) {
		return nil, fmt.Errorf("invalid fallback handler address: %s", params.FallbackHandlerAddress)
	}

	// TODO: Implement actual transaction data creation
	return []byte{}, nil
}