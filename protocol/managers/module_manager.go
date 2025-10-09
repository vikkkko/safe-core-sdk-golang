package managers

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/contracts"
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

	// Get Safe ABI to encode function call
	abi, err := contracts.SafeBindingMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	// Pack the enableModule function call
	data, err := abi.Pack("enableModule", common.HexToAddress(params.ModuleAddress))
	if err != nil {
		return nil, fmt.Errorf("failed to pack enableModule call: %w", err)
	}

	return data, nil
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

	// Get all modules to find previous module in linked list
	modules, err := mm.GetModules(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get modules: %w", err)
	}

	// Find previous module in the linked list
	sentinel := common.HexToAddress("0x0000000000000000000000000000000000000001")
	moduleToDisable := common.HexToAddress(params.ModuleAddress)
	var prevModule common.Address = sentinel

	for i, module := range modules {
		if module == moduleToDisable {
			if i == 0 {
				prevModule = sentinel
			} else {
				prevModule = modules[i-1]
			}
			break
		}
	}

	// Get Safe ABI to encode function call
	abi, err := contracts.SafeBindingMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	// Pack the disableModule function call
	data, err := abi.Pack("disableModule", prevModule, moduleToDisable)
	if err != nil {
		return nil, fmt.Errorf("failed to pack disableModule call: %w", err)
	}

	return data, nil
}
