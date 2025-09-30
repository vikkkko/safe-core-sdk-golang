package managers

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/contracts"
)

// OwnerManager manages Safe owners
type OwnerManager struct {
	client      *ethclient.Client
	safeAddress common.Address
}

// NewOwnerManager creates a new owner manager
func NewOwnerManager(client *ethclient.Client, safeAddress common.Address) *OwnerManager {
	return &OwnerManager{
		client:      client,
		safeAddress: safeAddress,
	}
}

// GetOwners returns the list of Safe owners
func (om *OwnerManager) GetOwners(ctx context.Context) ([]common.Address, error) {
	safeContract, err := contracts.NewSafeContract(om.safeAddress, om.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Safe contract instance: %w", err)
	}

	return safeContract.GetOwners(ctx)
}

// IsOwner checks if an address is a Safe owner
func (om *OwnerManager) IsOwner(ctx context.Context, address common.Address) (bool, error) {
	safeContract, err := contracts.NewSafeContract(om.safeAddress, om.client)
	if err != nil {
		return false, fmt.Errorf("failed to create Safe contract instance: %w", err)
	}

	return safeContract.IsOwner(ctx, address)
}

// GetOwnersCount returns the number of Safe owners
func (om *OwnerManager) GetOwnersCount(ctx context.Context) (uint, error) {
	owners, err := om.GetOwners(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get owners: %w", err)
	}

	return uint(len(owners)), nil
}

// AddOwnerTxParams represents parameters for adding an owner
type AddOwnerTxParams struct {
	OwnerAddress string `json:"ownerAddress"` // Address of the new owner
	Threshold    *uint  `json:"threshold,omitempty"` // New threshold (optional, defaults to current + 1)
}

// CreateAddOwnerTx creates a transaction to add a new owner
func (om *OwnerManager) CreateAddOwnerTx(ctx context.Context, params AddOwnerTxParams) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate the new owner address
	// 2. Check if the address is already an owner
	// 3. Create the transaction data for adding the owner
	// 4. Return the encoded transaction data

	if !common.IsHexAddress(params.OwnerAddress) {
		return nil, fmt.Errorf("invalid owner address: %s", params.OwnerAddress)
	}

	// Check if already an owner
	isOwner, err := om.IsOwner(ctx, common.HexToAddress(params.OwnerAddress))
	if err != nil {
		return nil, fmt.Errorf("failed to check if address is owner: %w", err)
	}

	if isOwner {
		return nil, fmt.Errorf("address %s is already an owner", params.OwnerAddress)
	}

	// TODO: Implement actual transaction data creation
	return []byte{}, nil
}

// RemoveOwnerTxParams represents parameters for removing an owner
type RemoveOwnerTxParams struct {
	OwnerAddress string `json:"ownerAddress"` // Address of the owner to remove
	Threshold    *uint  `json:"threshold,omitempty"` // New threshold (optional, defaults to current - 1)
}

// CreateRemoveOwnerTx creates a transaction to remove an owner
func (om *OwnerManager) CreateRemoveOwnerTx(ctx context.Context, params RemoveOwnerTxParams) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate the owner address
	// 2. Check if the address is actually an owner
	// 3. Ensure removing the owner won't make threshold impossible to reach
	// 4. Create the transaction data for removing the owner
	// 5. Return the encoded transaction data

	if !common.IsHexAddress(params.OwnerAddress) {
		return nil, fmt.Errorf("invalid owner address: %s", params.OwnerAddress)
	}

	// Check if actually an owner
	isOwner, err := om.IsOwner(ctx, common.HexToAddress(params.OwnerAddress))
	if err != nil {
		return nil, fmt.Errorf("failed to check if address is owner: %w", err)
	}

	if !isOwner {
		return nil, fmt.Errorf("address %s is not an owner", params.OwnerAddress)
	}

	// TODO: Implement actual transaction data creation
	return []byte{}, nil
}

// SwapOwnerTxParams represents parameters for swapping an owner
type SwapOwnerTxParams struct {
	OldOwnerAddress string `json:"oldOwnerAddress"` // Address of the owner to replace
	NewOwnerAddress string `json:"newOwnerAddress"` // Address of the new owner
}

// CreateSwapOwnerTx creates a transaction to swap an owner
func (om *OwnerManager) CreateSwapOwnerTx(ctx context.Context, params SwapOwnerTxParams) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate both addresses
	// 2. Check if old address is an owner and new address is not
	// 3. Create the transaction data for swapping the owner
	// 4. Return the encoded transaction data

	if !common.IsHexAddress(params.OldOwnerAddress) {
		return nil, fmt.Errorf("invalid old owner address: %s", params.OldOwnerAddress)
	}

	if !common.IsHexAddress(params.NewOwnerAddress) {
		return nil, fmt.Errorf("invalid new owner address: %s", params.NewOwnerAddress)
	}

	// Check if old address is an owner
	isOldOwner, err := om.IsOwner(ctx, common.HexToAddress(params.OldOwnerAddress))
	if err != nil {
		return nil, fmt.Errorf("failed to check if old address is owner: %w", err)
	}

	if !isOldOwner {
		return nil, fmt.Errorf("address %s is not an owner", params.OldOwnerAddress)
	}

	// Check if new address is not already an owner
	isNewOwner, err := om.IsOwner(ctx, common.HexToAddress(params.NewOwnerAddress))
	if err != nil {
		return nil, fmt.Errorf("failed to check if new address is owner: %w", err)
	}

	if isNewOwner {
		return nil, fmt.Errorf("address %s is already an owner", params.NewOwnerAddress)
	}

	// TODO: Implement actual transaction data creation
	return []byte{}, nil
}

// ChangeThresholdTxParams represents parameters for changing the threshold
type ChangeThresholdTxParams struct {
	Threshold uint `json:"threshold"` // New threshold value
}

// CreateChangeThresholdTx creates a transaction to change the threshold
func (om *OwnerManager) CreateChangeThresholdTx(ctx context.Context, params ChangeThresholdTxParams) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate the new threshold
	// 2. Check if it's achievable with current number of owners
	// 3. Create the transaction data for changing the threshold
	// 4. Return the encoded transaction data

	if params.Threshold == 0 {
		return nil, fmt.Errorf("threshold must be greater than 0")
	}

	ownersCount, err := om.GetOwnersCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get owners count: %w", err)
	}

	if params.Threshold > ownersCount {
		return nil, fmt.Errorf("threshold (%d) cannot be greater than number of owners (%d)", params.Threshold, ownersCount)
	}

	// TODO: Implement actual transaction data creation
	return []byte{}, nil
}