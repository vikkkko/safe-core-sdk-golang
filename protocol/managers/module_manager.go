package managers

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yinwei/safe-core-sdk-golang/protocol/contracts"
)

// ModuleManager manages Safe modules
type ModuleManager struct {
	client      *ethclient.Client
	safeAddress common.Address
}

// NewModuleManager creates a new module manager
func NewModuleManager(client *ethclient.Client, safeAddress common.Address) *ModuleManager {
	return &ModuleManager{
		client:      client,
		safeAddress: safeAddress,
	}
}

// GetModules returns the list of enabled modules
func (mm *ModuleManager) GetModules(ctx context.Context) ([]common.Address, error) {
	safeContract, err := contracts.NewSafeContract(mm.safeAddress, mm.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Safe contract instance: %w", err)
	}

	return safeContract.GetModules(ctx)
}

// IsModuleEnabled checks if a module is enabled
func (mm *ModuleManager) IsModuleEnabled(ctx context.Context, moduleAddress common.Address) (bool, error) {
	safeContract, err := contracts.NewSafeContract(mm.safeAddress, mm.client)
	if err != nil {
		return false, fmt.Errorf("failed to create Safe contract instance: %w", err)
	}

	return safeContract.IsModuleEnabled(ctx, moduleAddress)
}

// GetModulesCount returns the number of enabled modules
func (mm *ModuleManager) GetModulesCount(ctx context.Context) (uint, error) {
	modules, err := mm.GetModules(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get modules: %w", err)
	}

	return uint(len(modules)), nil
}

// EnableModuleTxParams represents parameters for enabling a module
type EnableModuleTxParams struct {
	ModuleAddress string `json:"moduleAddress"` // Address of the module to enable
}

// CreateEnableModuleTx creates a transaction to enable a module
func (mm *ModuleManager) CreateEnableModuleTx(ctx context.Context, params EnableModuleTxParams) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate the module address
	// 2. Check if the module is not already enabled
	// 3. Create the transaction data for enabling the module
	// 4. Return the encoded transaction data

	if !common.IsHexAddress(params.ModuleAddress) {
		return nil, fmt.Errorf("invalid module address: %s", params.ModuleAddress)
	}

	// Check if module is already enabled
	isEnabled, err := mm.IsModuleEnabled(ctx, common.HexToAddress(params.ModuleAddress))
	if err != nil {
		return nil, fmt.Errorf("failed to check if module is enabled: %w", err)
	}

	if isEnabled {
		return nil, fmt.Errorf("module %s is already enabled", params.ModuleAddress)
	}

	// TODO: Implement actual transaction data creation
	return []byte{}, nil
}

// DisableModuleTxParams represents parameters for disabling a module
type DisableModuleTxParams struct {
	ModuleAddress string `json:"moduleAddress"` // Address of the module to disable
}

// CreateDisableModuleTx creates a transaction to disable a module
func (mm *ModuleManager) CreateDisableModuleTx(ctx context.Context, params DisableModuleTxParams) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate the module address
	// 2. Check if the module is currently enabled
	// 3. Create the transaction data for disabling the module
	// 4. Return the encoded transaction data

	if !common.IsHexAddress(params.ModuleAddress) {
		return nil, fmt.Errorf("invalid module address: %s", params.ModuleAddress)
	}

	// Check if module is enabled
	isEnabled, err := mm.IsModuleEnabled(ctx, common.HexToAddress(params.ModuleAddress))
	if err != nil {
		return nil, fmt.Errorf("failed to check if module is enabled: %w", err)
	}

	if !isEnabled {
		return nil, fmt.Errorf("module %s is not enabled", params.ModuleAddress)
	}

	// TODO: Implement actual transaction data creation
	return []byte{}, nil
}