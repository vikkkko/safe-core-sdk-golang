package protocol

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yinwei/safe-core-sdk-golang/protocol/managers"
	"github.com/yinwei/safe-core-sdk-golang/protocol/utils"
	"github.com/yinwei/safe-core-sdk-golang/types"
)

// SafeConfig represents the configuration for connecting to a Safe
type SafeConfig struct {
	SafeAddress string `json:"safeAddress"`          // Address of the existing Safe
	RpcURL      string `json:"rpcUrl"`               // RPC URL for blockchain connection
	ChainID     int64  `json:"chainId"`              // Chain ID
	PrivateKey  string `json:"privateKey,omitempty"` // Private key for signing (optional)
}

// SafeConfigWithPredicted represents configuration with predicted Safe properties
type SafeConfigWithPredicted struct {
	Predicted  types.PredictedSafeProps `json:"predicted"`            // Predicted Safe properties
	RpcURL     string                   `json:"rpcUrl"`               // RPC URL for blockchain connection
	ChainID    int64                    `json:"chainId"`              // Chain ID
	PrivateKey string                   `json:"privateKey,omitempty"` // Private key for signing (optional)
}

// Safe represents a Safe Smart Account client
type Safe struct {
	config          SafeConfig
	client          *ethclient.Client
	predictedSafe   *types.PredictedSafeProps
	contractManager *managers.ContractManager
	ownerManager    *managers.OwnerManager
	moduleManager   *managers.ModuleManager
	guardManager    *managers.GuardManager
	fallbackManager *managers.FallbackHandlerManager
}

// NewSafe creates a new Safe client for an existing Safe
func NewSafe(config SafeConfig) (*Safe, error) {
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	// Validate the Safe address
	if !common.IsHexAddress(config.SafeAddress) {
		return nil, fmt.Errorf("invalid Safe address: %s", config.SafeAddress)
	}

	contractManager, err := managers.NewContractManager(client, big.NewInt(config.ChainID))
	if err != nil {
		return nil, fmt.Errorf("failed to create contract manager: %w", err)
	}

	safe := &Safe{
		config:          config,
		client:          client,
		contractManager: contractManager,
	}

	// Initialize managers
	safe.ownerManager = managers.NewOwnerManager(safe.client, common.HexToAddress(config.SafeAddress))
	safe.moduleManager = managers.NewModuleManager(safe.client, common.HexToAddress(config.SafeAddress))
	safe.guardManager = managers.NewGuardManager(safe.client, common.HexToAddress(config.SafeAddress))
	safe.fallbackManager = managers.NewFallbackHandlerManager(safe.client, common.HexToAddress(config.SafeAddress))

	return safe, nil
}

// NewSafeWithPredicted creates a new Safe client for a predicted (not yet deployed) Safe
func NewSafeWithPredicted(config SafeConfigWithPredicted) (*Safe, error) {
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	contractManager, err := managers.NewContractManager(client, big.NewInt(config.ChainID))
	if err != nil {
		return nil, fmt.Errorf("failed to create contract manager: %w", err)
	}

	safe := &Safe{
		config: SafeConfig{
			SafeAddress: config.Predicted.SafeAddress,
			RpcURL:      config.RpcURL,
			ChainID:     config.ChainID,
			PrivateKey:  config.PrivateKey,
		},
		client:          client,
		predictedSafe:   &config.Predicted,
		contractManager: contractManager,
	}

	// Initialize managers
	safe.ownerManager = managers.NewOwnerManager(safe.client, common.HexToAddress(config.Predicted.SafeAddress))
	safe.moduleManager = managers.NewModuleManager(safe.client, common.HexToAddress(config.Predicted.SafeAddress))
	safe.guardManager = managers.NewGuardManager(safe.client, common.HexToAddress(config.Predicted.SafeAddress))
	safe.fallbackManager = managers.NewFallbackHandlerManager(safe.client, common.HexToAddress(config.Predicted.SafeAddress))

	return safe, nil
}

// GetAddress returns the Safe address
func (s *Safe) GetAddress() common.Address {
	return common.HexToAddress(s.config.SafeAddress)
}

// GetChainID returns the chain ID
func (s *Safe) GetChainID() int64 {
	return s.config.ChainID
}

// IsSafeDeployed checks if the Safe is deployed on the blockchain
func (s *Safe) IsSafeDeployed(ctx context.Context) (bool, error) {
	address := s.GetAddress()
	code, err := s.client.CodeAt(ctx, address, nil)
	if err != nil {
		return false, fmt.Errorf("failed to get code at address %s: %w", address.Hex(), err)
	}
	return len(code) > 0, nil
}

// GetNonce returns the current nonce of the Safe
func (s *Safe) GetNonce(ctx context.Context) (uint64, error) {
	safeContract, err := s.contractManager.GetSafeContract(s.GetAddress())
	if err != nil {
		return 0, fmt.Errorf("failed to get Safe contract: %w", err)
	}

	nonce, err := safeContract.GetNonce(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get nonce: %w", err)
	}

	return nonce.Uint64(), nil
}

// GetThreshold returns the current threshold of the Safe
func (s *Safe) GetThreshold(ctx context.Context) (uint, error) {
	safeContract, err := s.contractManager.GetSafeContract(s.GetAddress())
	if err != nil {
		return 0, fmt.Errorf("failed to get Safe contract: %w", err)
	}

	threshold, err := safeContract.GetThreshold(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get threshold: %w", err)
	}

	return uint(threshold.Uint64()), nil
}

// GetOwners returns the list of Safe owners
func (s *Safe) GetOwners(ctx context.Context) ([]common.Address, error) {
	return s.ownerManager.GetOwners(ctx)
}

// IsOwner checks if an address is a Safe owner
func (s *Safe) IsOwner(ctx context.Context, address common.Address) (bool, error) {
	return s.ownerManager.IsOwner(ctx, address)
}

// GetSafeInfo returns complete information about the Safe
func (s *Safe) GetSafeInfo(ctx context.Context) (*types.SafeInfo, error) {
	address := s.GetAddress()

	nonce, err := s.GetNonce(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	threshold, err := s.GetThreshold(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get threshold: %w", err)
	}

	owners, err := s.GetOwners(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get owners: %w", err)
	}

	// Convert addresses to strings
	ownerStrings := make([]string, len(owners))
	for i, owner := range owners {
		ownerStrings[i] = owner.Hex()
	}

	// Get modules
	modules, err := s.moduleManager.GetModules(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get modules: %w", err)
	}

	moduleStrings := make([]string, len(modules))
	for i, module := range modules {
		moduleStrings[i] = module.Hex()
	}

	// Get fallback handler
	fallbackHandler, err := s.fallbackManager.GetFallbackHandler(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get fallback handler: %w", err)
	}

	// Get guard
	guard, err := s.guardManager.GetGuard(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get guard: %w", err)
	}

	// Get version (placeholder for now)
	version := string(types.DefaultSafeVersion)

	return &types.SafeInfo{
		Address:         address.Hex(),
		Nonce:           nonce,
		Threshold:       threshold,
		Owners:          ownerStrings,
		MasterCopy:      "", // TODO: Implement getting master copy
		Modules:         moduleStrings,
		FallbackHandler: fallbackHandler.Hex(),
		Guard:           guard.Hex(),
		Version:         version,
	}, nil
}

// CreateTransaction creates a new Safe transaction
func (s *Safe) CreateTransaction(ctx context.Context, txData types.SafeTransactionDataPartial) (*types.SafeTransaction, error) {
	// Fill in missing transaction data with defaults
	fullTxData, err := s.standardizeSafeTransactionData(ctx, txData)
	if err != nil {
		return nil, fmt.Errorf("failed to standardize transaction data: %w", err)
	}

	return &types.SafeTransaction{
		Data:       *fullTxData,
		Signatures: make(map[string]types.SafeSignature),
	}, nil
}

// standardizeSafeTransactionData fills in missing fields in transaction data
func (s *Safe) standardizeSafeTransactionData(ctx context.Context, txData types.SafeTransactionDataPartial) (*types.SafeTransactionData, error) {
	// Set default operation type if not specified
	operation := types.Call
	if txData.Operation != nil {
		operation = *txData.Operation
	}

	// Get current nonce if not specified
	nonce := uint64(0)
	if txData.Nonce != nil {
		nonce = *txData.Nonce
	} else {
		currentNonce, err := s.GetNonce(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get current nonce: %w", err)
		}
		nonce = currentNonce
	}

	// Set default gas values if not specified
	safeTxGas := "0"
	if txData.SafeTxGas != nil {
		safeTxGas = *txData.SafeTxGas
	}

	baseGas := "0"
	if txData.BaseGas != nil {
		baseGas = *txData.BaseGas
	}

	gasPrice := "0"
	if txData.GasPrice != nil {
		gasPrice = *txData.GasPrice
	}

	gasToken := common.HexToAddress("0x0").Hex()
	if txData.GasToken != nil {
		gasToken = *txData.GasToken
	}

	refundReceiver := common.HexToAddress("0x0").Hex()
	if txData.RefundReceiver != nil {
		refundReceiver = *txData.RefundReceiver
	}

	return &types.SafeTransactionData{
		To:             txData.To,
		Value:          txData.Value,
		Data:           txData.Data,
		Operation:      operation,
		SafeTxGas:      safeTxGas,
		BaseGas:        baseGas,
		GasPrice:       gasPrice,
		GasToken:       gasToken,
		RefundReceiver: refundReceiver,
		Nonce:          nonce,
	}, nil
}

// SignTransaction signs a Safe transaction
func (s *Safe) SignTransaction(ctx context.Context, transaction *types.SafeTransaction, signerAddress common.Address) error {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Calculate the transaction hash
	// 2. Sign the hash using the appropriate signing method
	// 3. Add the signature to the transaction

	// For now, create a placeholder signature
	signature := types.SafeSignature{
		Signer:              signerAddress.Hex(),
		Data:                "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", // Placeholder
		IsContractSignature: false,
	}

	transaction.AddSignature(signature)
	return nil
}

// ExecuteTransaction executes a Safe transaction
func (s *Safe) ExecuteTransaction(ctx context.Context, transaction *types.SafeTransaction) (*types.TransactionResult, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Validate the transaction and signatures
	// 2. Submit the transaction to the blockchain
	// 3. Return the transaction result

	return &types.TransactionResult{
		BaseTransactionResult: types.BaseTransactionResult{
			Hash: "0x0000000000000000000000000000000000000000000000000000000000000000", // Placeholder
		},
		TransactionResponse: nil,
		Options:             nil,
	}, nil
}

// PredictSafeAddress predicts the address of a Safe before deployment
func PredictSafeAddress(config types.SafeDeploymentConfig, chainID *big.Int) (string, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Calculate the CREATE2 address based on the Safe configuration
	// 2. Return the predicted address

	return utils.PredictSafeAddress(config, chainID)
}

// DeploySafe deploys a new Safe with the given configuration
func (s *Safe) DeploySafe(ctx context.Context, config types.SafeDeploymentConfig) (*types.TransactionResult, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Create the Safe deployment transaction
	// 2. Submit it to the blockchain
	// 3. Return the deployment result

	return &types.TransactionResult{
		BaseTransactionResult: types.BaseTransactionResult{
			Hash: "0x0000000000000000000000000000000000000000000000000000000000000000", // Placeholder
		},
		TransactionResponse: nil,
		Options:             nil,
	}, nil
}
