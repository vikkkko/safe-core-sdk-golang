package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SafeSetupConfig represents the configuration for Safe setup
type SafeSetupConfig struct {
	Owners            []common.Address // List of Safe owners
	Threshold         *big.Int         // Number of required confirmations
	To                common.Address   // Contract address for optional delegate call during setup
	Data              []byte           // Data payload for optional delegate call
	FallbackHandler   common.Address   // Handler for fallback calls to this contract
	PaymentToken      common.Address   // Token that should be used for the payment (0 is ETH)
	Payment           *big.Int         // Value that should be paid
	PaymentReceiver   common.Address   // Address that should receive the payment (or 0 if tx.origin)
}

// DefaultSafeSetupConfig creates a Safe setup configuration with sensible defaults
// Only owners and threshold are required
func DefaultSafeSetupConfig(owners []common.Address, threshold uint) SafeSetupConfig {
	return SafeSetupConfig{
		Owners:    owners,
		Threshold: big.NewInt(int64(threshold)),
		To:        common.Address{}, // No delegate call
		Data:      []byte{},          // No data
		// Default fallback handler for Sepolia
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{}, // No payment token
		Payment:         big.NewInt(0),    // No payment
		PaymentReceiver: common.Address{}, // No payment receiver
	}
}

// CreateSafeInitData creates the initialization data for Safe setup
// This encodes the setup() function call with all necessary parameters
// Uses the generated contract bindings for accurate ABI encoding
func CreateSafeInitData(config SafeSetupConfig) ([]byte, error) {
	// Validate configuration
	if len(config.Owners) == 0 {
		return nil, fmt.Errorf("at least one owner is required")
	}
	if config.Threshold == nil || config.Threshold.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("threshold must be greater than 0")
	}
	if config.Threshold.Cmp(big.NewInt(int64(len(config.Owners)))) > 0 {
		return nil, fmt.Errorf("threshold (%s) cannot be greater than number of owners (%d)",
			config.Threshold.String(), len(config.Owners))
	}

	// Set defaults if not provided
	payment := config.Payment
	if payment == nil {
		payment = big.NewInt(0)
	}

	data := config.Data
	if data == nil {
		data = []byte{}
	}

	// Use the generated ABI to pack the setup function call
	abi, err := SafeContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	// Encode setup function call
	encodedData, err := abi.Pack(
		"setup",
		config.Owners,
		config.Threshold,
		config.To,
		data,
		config.FallbackHandler,
		config.PaymentToken,
		payment,
		config.PaymentReceiver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode setup call: %w", err)
	}

	return encodedData, nil
}

// CreateSafeInitDataSimple is a simplified version that only requires owners and threshold
// Uses default values for all other parameters
func CreateSafeInitDataSimple(owners []common.Address, threshold uint) ([]byte, error) {
	config := DefaultSafeSetupConfig(owners, threshold)
	return CreateSafeInitData(config)
}

// ParseOwnersFromStrings converts string addresses to common.Address slice
func ParseOwnersFromStrings(ownerStrings []string) ([]common.Address, error) {
	if len(ownerStrings) == 0 {
		return nil, fmt.Errorf("at least one owner is required")
	}

	owners := make([]common.Address, len(ownerStrings))
	for i, ownerStr := range ownerStrings {
		if !common.IsHexAddress(ownerStr) {
			return nil, fmt.Errorf("invalid owner address at index %d: %s", i, ownerStr)
		}
		owners[i] = common.HexToAddress(ownerStr)
	}

	return owners, nil
}

// NewSafeContractWrapper creates a new Safe contract instance
// This wraps the generated binding for easier use
func NewSafeContractWrapper(address common.Address, client *ethclient.Client) (*SafeContract, error) {
	return NewSafeContract(address, client)
}

// CallSafeMethod is a helper to call read-only Safe methods
// This provides a generic way to call any Safe view/pure function
func CallSafeMethod(client *ethclient.Client, safeAddress common.Address, method string, args ...interface{}) ([]interface{}, error) {
	abi, err := SafeContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	callData, err := abi.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack method %s: %w", method, err)
	}

	msg := ethereum.CallMsg{
		To:   &safeAddress,
		Data: callData,
	}
	result, err := client.CallContract(nil, msg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call method %s: %w", method, err)
	}

	outputs, err := abi.Unpack(method, result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return outputs, nil
}

// SafeAddOwnerWithThresholdData creates calldata for addOwnerWithThreshold function
// Parameters:
//   - owner: New owner address to add
//   - threshold: New threshold (use 0 to keep current threshold)
func SafeAddOwnerWithThresholdData(owner common.Address, threshold *big.Int) ([]byte, error) {
	abi, err := SafeContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	data, err := abi.Pack("addOwnerWithThreshold", owner, threshold)
	if err != nil {
		return nil, fmt.Errorf("failed to encode addOwnerWithThreshold call: %w", err)
	}

	return data, nil
}

// SafeRemoveOwnerData creates calldata for removeOwner function
// Parameters:
//   - prevOwner: Previous owner in the linked list (use sentinel 0x1 if removing first owner)
//   - owner: Owner address to remove
//   - threshold: New threshold
func SafeRemoveOwnerData(prevOwner, owner common.Address, threshold *big.Int) ([]byte, error) {
	abi, err := SafeContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	data, err := abi.Pack("removeOwner", prevOwner, owner, threshold)
	if err != nil {
		return nil, fmt.Errorf("failed to encode removeOwner call: %w", err)
	}

	return data, nil
}

// SafeSwapOwnerData creates calldata for swapOwner function
// Parameters:
//   - prevOwner: Previous owner in the linked list (use sentinel 0x1 if swapping first owner)
//   - oldOwner: Owner address to replace
//   - newOwner: New owner address
func SafeSwapOwnerData(prevOwner, oldOwner, newOwner common.Address) ([]byte, error) {
	abi, err := SafeContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	data, err := abi.Pack("swapOwner", prevOwner, oldOwner, newOwner)
	if err != nil {
		return nil, fmt.Errorf("failed to encode swapOwner call: %w", err)
	}

	return data, nil
}

// SafeChangeThresholdData creates calldata for changeThreshold function
// Parameters:
//   - threshold: New threshold
func SafeChangeThresholdData(threshold *big.Int) ([]byte, error) {
	abi, err := SafeContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	data, err := abi.Pack("changeThreshold", threshold)
	if err != nil {
		return nil, fmt.Errorf("failed to encode changeThreshold call: %w", err)
	}

	return data, nil
}

// SafeSetGuardData creates calldata for setGuard function
// Parameters:
//   - guard: Guard contract address (use zero address to disable guard)
func SafeSetGuardData(guard common.Address) ([]byte, error) {
	abi, err := SafeContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	data, err := abi.Pack("setGuard", guard)
	if err != nil {
		return nil, fmt.Errorf("failed to encode setGuard call: %w", err)
	}

	return data, nil
}
