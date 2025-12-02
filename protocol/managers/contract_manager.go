package managers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/contracts"
	"github.com/vikkkko/safe-core-sdk-golang/types"
)

// ContractManager manages Safe-related contracts
type ContractManager struct {
	client  *ethclient.Client
	chainID *big.Int
}

// NewContractManager creates a new contract manager
func NewContractManager(client *ethclient.Client, chainID *big.Int) (*ContractManager, error) {
	return &ContractManager{
		client:  client,
		chainID: chainID,
	}, nil
}

// GetSafeContract returns a Safe contract instance
func (cm *ContractManager) GetSafeContract(address common.Address) (*contracts.SafeContract, error) {
	return contracts.NewSafeContract(address, cm.client)
}

// GetSafeProxyFactoryContract returns a Safe proxy factory contract instance
func (cm *ContractManager) GetSafeProxyFactoryContract(version types.SafeVersion) (*contracts.SafeProxyFactoryContract, error) {
	address, err := cm.getSafeProxyFactoryAddress(version)
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe proxy factory address: %w", err)
	}

	return contracts.NewSafeProxyFactoryContract(address, cm.client)
}

// GetMultiSendContract returns a MultiSend contract instance
func (cm *ContractManager) GetMultiSendContract(version types.SafeVersion) (*contracts.MultiSendContract, error) {
	address, err := cm.getMultiSendAddress(version)
	if err != nil {
		return nil, fmt.Errorf("failed to get MultiSend address: %w", err)
	}

	return contracts.NewMultiSendContract(address, cm.client)
}

// GetMultiSendCallOnlyContract returns a MultiSendCallOnly contract instance
func (cm *ContractManager) GetMultiSendCallOnlyContract(version types.SafeVersion) (*contracts.MultiSendCallOnlyContract, error) {
	address, err := cm.getMultiSendCallOnlyAddress(version)
	if err != nil {
		return nil, fmt.Errorf("failed to get MultiSendCallOnly address: %w", err)
	}

	return contracts.NewMultiSendCallOnlyContract(address, cm.client)
}

// getSafeProxyFactoryAddress returns the Safe proxy factory address for the given version and chain
func (cm *ContractManager) getSafeProxyFactoryAddress(version types.SafeVersion) (common.Address, error) {
	// This is a placeholder implementation
	// In a real implementation, this would return the actual deployed addresses
	// based on the chain ID and Safe version

	chainID := cm.chainID.Int64()

	// Example addresses for Ethereum mainnet (chain ID 1)
	if chainID == 1 {
		switch version {
		case types.SafeVersion141:
			return common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"), nil
		case types.SafeVersion130:
			return common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	// Sepolia multichannel deployments
	if chainID == 11155111 {
		switch version {
		case types.SafeVersion150Multichannel:
			return common.HexToAddress("0x72D89c510AFBeC255b81482C8DCC720FC8743175"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	// For other chains, return zero address as placeholder
	return common.Address{}, fmt.Errorf("unsupported chain ID: %d", chainID)
}

// getMultiSendAddress returns the MultiSend address for the given version and chain
func (cm *ContractManager) getMultiSendAddress(version types.SafeVersion) (common.Address, error) {
	chainID := cm.chainID.Int64()

	// Example addresses for Ethereum mainnet (chain ID 1)
	if chainID == 1 {
		switch version {
		case types.SafeVersion141:
			return common.HexToAddress("0xA238CBeb142c10Ef7Ad8442C6D1f9E89e07e7761"), nil
		case types.SafeVersion130:
			return common.HexToAddress("0xA238CBeb142c10Ef7Ad8442C6D1f9E89e07e7761"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	if chainID == 11155111 {
		switch version {
		case types.SafeVersion150Multichannel:
			return common.HexToAddress("0xAf4bc1f38ab171b69edD46fD11B4b66580488c74"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	return common.Address{}, fmt.Errorf("unsupported chain ID: %d", chainID)
}

// getMultiSendCallOnlyAddress returns the MultiSendCallOnly address for the given version and chain
func (cm *ContractManager) getMultiSendCallOnlyAddress(version types.SafeVersion) (common.Address, error) {
	chainID := cm.chainID.Int64()

	// Example addresses for Ethereum mainnet (chain ID 1)
	if chainID == 1 {
		switch version {
		case types.SafeVersion141:
			return common.HexToAddress("0x40A2aCCbd92BCA938b02010E17A5b8929b49130D"), nil
		case types.SafeVersion130:
			return common.HexToAddress("0x40A2aCCbd92BCA938b02010E17A5b8929b49130D"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	if chainID == 11155111 {
		switch version {
		case types.SafeVersion150Multichannel:
			return common.HexToAddress("0x4A8E76A4eCaD4Fb0B3e8f51ec43A6979D384B0bF"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	return common.Address{}, fmt.Errorf("unsupported chain ID: %d", chainID)
}

// GetSafeMasterCopyAddress returns the Safe master copy address for the given version and chain
func (cm *ContractManager) GetSafeMasterCopyAddress(version types.SafeVersion) (common.Address, error) {
	chainID := cm.chainID.Int64()

	// Example addresses for Ethereum mainnet (chain ID 1)
	if chainID == 1 {
		switch version {
		case types.SafeVersion141:
			return common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552"), nil
		case types.SafeVersion130:
			return common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	if chainID == 11155111 {
		switch version {
		case types.SafeVersion150Multichannel:
			return common.HexToAddress("0x7E4aFC215CBdeB92151379692602faa37B40Edd7"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	return common.Address{}, fmt.Errorf("unsupported chain ID: %d", chainID)
}

// GetCompatibilityFallbackHandlerAddress returns the compatibility fallback handler address
func (cm *ContractManager) GetCompatibilityFallbackHandlerAddress(version types.SafeVersion) (common.Address, error) {
	chainID := cm.chainID.Int64()

	// Example addresses for Ethereum mainnet (chain ID 1)
	if chainID == 1 {
		switch version {
		case types.SafeVersion141:
			return common.HexToAddress("0xf48f2B2d2a534e402487b3ee7C18c33Aec0Fe5e4"), nil
		case types.SafeVersion130:
			return common.HexToAddress("0xf48f2B2d2a534e402487b3ee7C18c33Aec0Fe5e4"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	if chainID == 11155111 {
		switch version {
		case types.SafeVersion150Multichannel:
			return common.HexToAddress("0x51Eb07162CDd89e0575d059e848888e283155710"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	return common.Address{}, fmt.Errorf("unsupported chain ID: %d", chainID)
}

// GetCreateCallAddress returns the CreateCall contract address
func (cm *ContractManager) GetCreateCallAddress(version types.SafeVersion) (common.Address, error) {
	chainID := cm.chainID.Int64()

	// Example addresses for Ethereum mainnet (chain ID 1)
	if chainID == 1 {
		switch version {
		case types.SafeVersion141:
			return common.HexToAddress("0x7cbB62EaA69F79e6873cD1ecB2392971036cFAa4"), nil
		case types.SafeVersion130:
			return common.HexToAddress("0x7cbB62EaA69F79e6873cD1ecB2392971036cFAa4"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	if chainID == 11155111 {
		switch version {
		case types.SafeVersion150Multichannel:
			return common.HexToAddress("0xe324d7DFe4B1cf38Deac8397cFE98fBBFE46D349"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	return common.Address{}, fmt.Errorf("unsupported chain ID: %d", chainID)
}

// GetSignMessageLibAddress returns the SignMessageLib contract address
func (cm *ContractManager) GetSignMessageLibAddress(version types.SafeVersion) (common.Address, error) {
	chainID := cm.chainID.Int64()

	// Example addresses for Ethereum mainnet (chain ID 1)
	if chainID == 1 {
		switch version {
		case types.SafeVersion141:
			return common.HexToAddress("0xA65387F16B013cf2Af4605Ad8aA5ec25a2cbA3a2"), nil
		case types.SafeVersion130:
			return common.HexToAddress("0xA65387F16B013cf2Af4605Ad8aA5ec25a2cbA3a2"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	if chainID == 11155111 {
		switch version {
		case types.SafeVersion150Multichannel:
			return common.HexToAddress("0x662a9BAB37AF013D5A9a3e3B3b50E5133Ebfb6d5"), nil
		default:
			return common.Address{}, fmt.Errorf("unsupported Safe version: %s", version)
		}
	}

	return common.Address{}, fmt.Errorf("unsupported chain ID: %d", chainID)
}
