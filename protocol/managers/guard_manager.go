package managers

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yinwei/safe-core-sdk-golang/protocol/contracts"
)

// GuardManager manages Safe transaction guards
type GuardManager struct {
	client      *ethclient.Client
	safeAddress common.Address
}

// NewGuardManager creates a new guard manager
func NewGuardManager(client *ethclient.Client, safeAddress common.Address) *GuardManager {
	return &GuardManager{
		client:      client,
		safeAddress: safeAddress,
	}
}

// GetGuard returns the current transaction guard address
func (gm *GuardManager) GetGuard(ctx context.Context) (common.Address, error) {
	safeContract, err := contracts.NewSafeContract(gm.safeAddress, gm.client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create Safe contract instance: %w", err)
	}

	return safeContract.GetGuard(ctx)
}

// SetGuardTxParams represents parameters for setting a guard
type SetGuardTxParams struct {
	GuardAddress string `json:"guardAddress"` // Address of the guard contract (use zero address to disable)
}

// CreateSetGuardTx creates a transaction to set a transaction guard
func (gm *GuardManager) CreateSetGuardTx(ctx context.Context, params SetGuardTxParams) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate the guard address (can be zero address to disable guard)
	// 2. Create the transaction data for setting the guard
	// 3. Return the encoded transaction data

	if params.GuardAddress != "" && !common.IsHexAddress(params.GuardAddress) {
		return nil, fmt.Errorf("invalid guard address: %s", params.GuardAddress)
	}

	// TODO: Implement actual transaction data creation
	return []byte{}, nil
}

// CreateDisableGuardTx creates a transaction to disable the current guard
func (gm *GuardManager) CreateDisableGuardTx(ctx context.Context) ([]byte, error) {
	// This creates a transaction to set the guard to zero address (disabled)
	return gm.CreateSetGuardTx(ctx, SetGuardTxParams{
		GuardAddress: common.HexToAddress("0x0").Hex(),
	})
}