package contracts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SafeContract represents a Safe smart contract
type SafeContract struct {
	address common.Address
	client  *ethclient.Client
}

// NewSafeContract creates a new Safe contract instance
func NewSafeContract(address common.Address, client *ethclient.Client) (*SafeContract, error) {
	return &SafeContract{
		address: address,
		client:  client,
	}, nil
}

// Address returns the contract address
func (sc *SafeContract) Address() common.Address {
	return sc.address
}

// GetNonce returns the current nonce of the Safe
func (sc *SafeContract) GetNonce(ctx context.Context) (*big.Int, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	nonce, err := binding.Nonce(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	return nonce, nil
}

// GetThreshold returns the current threshold of the Safe
func (sc *SafeContract) GetThreshold(ctx context.Context) (*big.Int, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	threshold, err := binding.GetThreshold(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get threshold: %w", err)
	}

	return threshold, nil
}

// GetOwners returns the list of Safe owners
func (sc *SafeContract) GetOwners(ctx context.Context) ([]common.Address, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	owners, err := binding.GetOwners(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get owners: %w", err)
	}

	return owners, nil
}

// IsOwner checks if an address is a Safe owner
func (sc *SafeContract) IsOwner(ctx context.Context, address common.Address) (bool, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return false, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	isOwner, err := binding.IsOwner(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return false, fmt.Errorf("failed to check if owner: %w", err)
	}

	return isOwner, nil
}

// GetModules returns the list of enabled modules
func (sc *SafeContract) GetModules(ctx context.Context) ([]common.Address, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	// Use SENTINEL_MODULES address as start and large page size to get all modules
	sentinel := common.HexToAddress("0x0000000000000000000000000000000000000001")
	result, err := binding.GetModulesPaginated(&bind.CallOpts{Context: ctx}, sentinel, big.NewInt(100))
	if err != nil {
		return nil, fmt.Errorf("failed to get modules: %w", err)
	}

	return result.Array, nil
}

// IsModuleEnabled checks if a module is enabled
func (sc *SafeContract) IsModuleEnabled(ctx context.Context, moduleAddress common.Address) (bool, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return false, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	isEnabled, err := binding.IsModuleEnabled(&bind.CallOpts{Context: ctx}, moduleAddress)
	if err != nil {
		return false, fmt.Errorf("failed to check if module is enabled: %w", err)
	}

	return isEnabled, nil
}

// GetGuard returns the current transaction guard address
func (sc *SafeContract) GetGuard(ctx context.Context) (common.Address, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	guard, err := binding.GetGuard(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get guard: %w", err)
	}

	return guard, nil
}

// GetFallbackHandler returns the current fallback handler address
func (sc *SafeContract) GetFallbackHandler(ctx context.Context) (common.Address, error) {
	// Fallback handler is stored in a specific storage slot
	// FALLBACK_HANDLER_STORAGE_SLOT = keccak256("fallback_manager.handler.address")
	storageSlot := common.HexToHash("0x6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d5")

	storageValue, err := sc.client.StorageAt(ctx, sc.address, storageSlot, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to read fallback handler from storage: %w", err)
	}

	// Convert bytes to address (last 20 bytes)
	return common.BytesToAddress(storageValue), nil
}

// ExecTransaction executes a Safe transaction
func (sc *SafeContract) ExecTransaction(
	ctx context.Context,
	to common.Address,
	value *big.Int,
	data []byte,
	operation uint8,
	safeTxGas *big.Int,
	baseGas *big.Int,
	gasPrice *big.Int,
	gasToken common.Address,
	refundReceiver common.Address,
	signatures []byte,
) (common.Hash, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Prepare the transaction data
	// 2. Send the transaction to the blockchain
	// 3. Return the transaction hash

	return common.Hash{}, fmt.Errorf("execTransaction not implemented")
}

// GetTransactionHash calculates the transaction hash for signing
func (sc *SafeContract) GetTransactionHash(
	to common.Address,
	value *big.Int,
	data []byte,
	operation uint8,
	safeTxGas *big.Int,
	baseGas *big.Int,
	gasPrice *big.Int,
	gasToken common.Address,
	refundReceiver common.Address,
	nonce *big.Int,
) ([32]byte, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	txHash, err := binding.GetTransactionHash(
		&bind.CallOpts{},
		to,
		value,
		data,
		operation,
		safeTxGas,
		baseGas,
		gasPrice,
		gasToken,
		refundReceiver,
		nonce,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to get transaction hash: %w", err)
	}

	return txHash, nil
}

// GetMessageHash calculates the message hash for signing
func (sc *SafeContract) GetMessageHash(message []byte) ([32]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would calculate the actual message hash
	// according to the Safe's getMessageHash method
	return [32]byte{}, nil
}

// GetChainId returns the chain ID from the Safe contract
func (sc *SafeContract) GetChainId(ctx context.Context) (*big.Int, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	chainId, err := binding.GetChainId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	return chainId, nil
}

// DomainSeparator returns the domain separator from the Safe contract
func (sc *SafeContract) DomainSeparator(ctx context.Context) ([32]byte, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create Safe binding: %w", err)
	}

	domainSeparator, err := binding.DomainSeparator(&bind.CallOpts{Context: ctx})
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to get domain separator: %w", err)
	}

	return domainSeparator, nil
}

// VERSION returns the Safe version
func (sc *SafeContract) VERSION(ctx context.Context) (string, error) {
	binding, err := NewSafeBinding(sc.address, sc.client)
	if err != nil {
		return "", fmt.Errorf("failed to create Safe binding: %w", err)
	}

	version, err := binding.VERSION(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", fmt.Errorf("failed to get version: %w", err)
	}

	return version, nil
}
