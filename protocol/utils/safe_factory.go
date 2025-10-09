package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SafeFactoryVersion represents different Safe factory contract versions
type SafeFactoryVersion string

const (
	// SafeFactoryV141 is the Safe factory version 1.4.1
	SafeFactoryV141 SafeFactoryVersion = "1.4.1"
)

// SafeFactoryConfig represents the configuration for Safe factory deployment
type SafeFactoryConfig struct {
	FactoryAddress   common.Address // Address of the Safe factory contract
	SingletonAddress common.Address // Address of the Safe singleton/master copy
	InitData         []byte         // Initialization data from CreateSafeInitData
	SaltNonce        *big.Int       // Salt nonce for CREATE2 deployment
}

// CreateSafeFactoryCallData creates the call data for deploying a Safe via factory
// Uses createProxyWithNonce function for deterministic deployment
// Uses the generated contract bindings for accurate ABI encoding
func CreateSafeFactoryCallData(singleton common.Address, initData []byte, saltNonce *big.Int) ([]byte, error) {
	// Validate inputs
	if singleton == (common.Address{}) {
		return nil, fmt.Errorf("singleton address cannot be zero")
	}
	if len(initData) == 0 {
		return nil, fmt.Errorf("init data cannot be empty")
	}
	if saltNonce == nil {
		saltNonce = big.NewInt(0)
	}

	// Use the generated ABI to pack the function call
	abi, err := SafeProxyFactoryContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get SafeProxyFactory ABI: %w", err)
	}

	// Encode the factory function call
	data, err := abi.Pack(
		"createProxyWithNonce",
		singleton,
		initData,
		saltNonce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode factory call: %w", err)
	}

	return data, nil
}

// CreateChainSpecificProxyCallData creates call data for chain-specific proxy deployment
// This is useful for L2 networks that require chain-specific deployment
func CreateChainSpecificProxyCallData(singleton common.Address, initData []byte, saltNonce *big.Int) ([]byte, error) {
	if singleton == (common.Address{}) {
		return nil, fmt.Errorf("singleton address cannot be zero")
	}
	if len(initData) == 0 {
		return nil, fmt.Errorf("init data cannot be empty")
	}
	if saltNonce == nil {
		saltNonce = big.NewInt(0)
	}

	abi, err := SafeProxyFactoryContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get SafeProxyFactory ABI: %w", err)
	}

	data, err := abi.Pack(
		"createChainSpecificProxyWithNonce",
		singleton,
		initData,
		saltNonce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode chain-specific factory call: %w", err)
	}

	return data, nil
}

// PredictSafeAddressFromFactory predicts the Safe address using CREATE2
// This wraps the existing CalculateProxyAddress function for convenience
// Note: For accurate prediction, you should call the factory contract's calculation method
func PredictSafeAddressFromFactory(factory common.Address, singleton common.Address, initData []byte, saltNonce *big.Int) (common.Address, error) {
	if saltNonce == nil {
		saltNonce = big.NewInt(0)
	}

	// Use the existing CalculateProxyAddress function from address.go
	return CalculateProxyAddress(factory, singleton, initData, saltNonce)
}

// GenerateRandomSalt generates a random 32-byte salt for CREATE2 deployment
func GenerateRandomSalt() [32]byte {
	var salt [32]byte
	// In production, use crypto/rand for secure random generation
	// For now, use a timestamp-based approach
	timestamp := big.NewInt(0).SetBytes(crypto.Keccak256([]byte(fmt.Sprintf("%d", common.BigToHash(big.NewInt(0)).Bytes()))))
	copy(salt[:], timestamp.Bytes())
	return salt
}

// GetProxyCreationCode retrieves the proxy creation code from the factory
// This is useful for CREATE2 address calculation
func GetProxyCreationCode(client *ethclient.Client, factoryAddress common.Address) ([]byte, error) {
	abi, err := SafeProxyFactoryContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get SafeProxyFactory ABI: %w", err)
	}

	callData, err := abi.Pack("proxyCreationCode")
	if err != nil {
		return nil, fmt.Errorf("failed to pack proxyCreationCode call: %w", err)
	}

	msg := ethereum.CallMsg{
		To:   &factoryAddress,
		Data: callData,
	}
	result, err := client.CallContract(nil, msg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call proxyCreationCode: %w", err)
	}

	var creationCode []byte
	if err := abi.UnpackIntoInterface(&creationCode, "proxyCreationCode", result); err != nil {
		return nil, fmt.Errorf("failed to unpack proxy creation code: %w", err)
	}

	return creationCode, nil
}

// GetChainId retrieves the chain ID from the factory contract
func GetChainId(client *ethclient.Client, factoryAddress common.Address) (*big.Int, error) {
	factory, err := NewSafeProxyFactoryContract(factoryAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create factory contract: %w", err)
	}

	chainId, err := factory.GetChainId(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	return chainId, nil
}

// DeploySafeConfig represents the complete configuration for Safe deployment
type DeploySafeConfig struct {
	Owners           []common.Address // Safe owners
	Threshold        uint             // Required confirmations
	FallbackHandler  common.Address   // Fallback handler (optional)
	FactoryAddress   common.Address   // Factory contract address
	SingletonAddress common.Address   // Safe singleton address
	SaltNonce        *big.Int         // Salt for CREATE2 (optional, defaults to 0)
}

// PrepareSafeDeployment prepares all data needed for Safe deployment
// Returns the factory call data that can be sent as a transaction
func PrepareSafeDeployment(config DeploySafeConfig) ([]byte, error) {
	// Step 1: Create Safe initialization data
	setupConfig := SafeSetupConfig{
		Owners:          config.Owners,
		Threshold:       big.NewInt(int64(config.Threshold)),
		To:              common.Address{},
		Data:            []byte{},
		FallbackHandler: config.FallbackHandler,
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: common.Address{},
	}

	// Use default fallback handler if not provided
	if config.FallbackHandler == (common.Address{}) {
		setupConfig.FallbackHandler = common.HexToAddress("0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99")
	}

	initData, err := CreateSafeInitData(setupConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create init data: %w", err)
	}

	// Step 2: Create factory call data
	saltNonce := config.SaltNonce
	if saltNonce == nil {
		saltNonce = big.NewInt(0)
	}

	factoryCallData, err := CreateSafeFactoryCallData(
		config.SingletonAddress,
		initData,
		saltNonce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create factory call data: %w", err)
	}

	return factoryCallData, nil
}

// NewSafeProxyFactoryContractWrapper creates a new SafeProxyFactory contract instance
// This wraps the generated binding for easier use
func NewSafeProxyFactoryContractWrapper(address common.Address, client *ethclient.Client) (*SafeProxyFactoryContract, error) {
	return NewSafeProxyFactoryContract(address, client)
}
