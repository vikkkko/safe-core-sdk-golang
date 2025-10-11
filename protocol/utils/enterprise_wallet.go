package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// EnterpriseWalletFactoryABI contains the ABI for the EnterpriseWalletFactory contract
const EnterpriseWalletFactoryABI = `[
	{
		"type": "function",
		"name": "createWallet",
		"inputs": [
			{"name": "implementation", "type": "address"},
			{"name": "salt", "type": "bytes32"},
			{
				"name": "params",
				"type": "tuple",
				"components": [
					{"name": "methods", "type": "bytes4[]"},
					{
						"name": "configs",
						"type": "tuple[]",
						"components": [
							{"name": "controller", "type": "address"}
						]
					},
					{"name": "superAdmin", "type": "address"}
				]
			}
		],
		"outputs": [{"name": "", "type": "address"}]
	},
	{
		"type": "function",
		"name": "predictWalletAddress",
		"inputs": [
			{"name": "implementation", "type": "address"},
			{"name": "salt", "type": "bytes32"},
			{"name": "deployer", "type": "address"}
		],
		"outputs": [{"name": "", "type": "address"}]
	},
	{
		"type": "function",
		"name": "isImplementationWhitelisted",
		"inputs": [{"name": "implementation", "type": "address"}],
		"outputs": [{"name": "", "type": "bool"}]
	}
]`

// EnterpriseWalletABI contains the key functions from the EnterpriseWallet contract
const EnterpriseWalletABI = `[
	{
		"type": "function",
		"name": "createPaymentAccount",
		"inputs": [
			{"name": "name", "type": "string"},
			{"name": "controller", "type": "address"}
		],
		"outputs": [{"name": "", "type": "address"}]
	},
	{
		"type": "function",
		"name": "createCollectionAccount",
		"inputs": [
			{"name": "name", "type": "string"},
			{"name": "collectionTarget", "type": "address"}
		],
		"outputs": [{"name": "", "type": "address"}]
	},
	{
		"type": "function",
		"name": "approveTokenForPayment",
		"inputs": [
			{"name": "token", "type": "address"},
			{"name": "paymentAccount", "type": "address"},
			{"name": "amount", "type": "uint256"}
		],
		"outputs": []
	},
	{
		"type": "function",
		"name": "transferETHToPayment",
		"inputs": [
			{"name": "paymentAccount", "type": "address"},
			{"name": "amount", "type": "uint256"}
		],
		"outputs": []
	},
	{
		"type": "function",
		"name": "collectFunds",
		"inputs": [
			{"name": "token", "type": "address"},
			{"name": "collectionAccount", "type": "address"}
		],
		"outputs": []
	},
	{
		"type": "function",
		"name": "getPaymentAccounts",
		"inputs": [],
		"outputs": [
			{
				"name": "",
				"type": "tuple[]",
				"components": [
					{"name": "account", "type": "address"},
					{"name": "createdAt", "type": "uint256"},
					{"name": "isActive", "type": "bool"}
				]
			}
		]
	},
	{
		"type": "function",
		"name": "getCollectionAccounts",
		"inputs": [],
		"outputs": [
			{
				"name": "",
				"type": "tuple[]",
				"components": [
					{"name": "account", "type": "address"},
					{"name": "createdAt", "type": "uint256"},
					{"name": "isActive", "type": "bool"}
				]
			}
		]
	},
	{
		"type": "function",
		"name": "getSuperAdmin",
		"inputs": [],
		"outputs": [{"name": "", "type": "address"}]
	},
	{
		"type": "function",
		"name": "predictPaymentAccountAddress",
		"inputs": [],
		"outputs": [{"name": "", "type": "address"}]
	},
	{
		"type": "function",
		"name": "predictCollectionAccountAddress",
		"inputs": [],
		"outputs": [{"name": "", "type": "address"}]
	},
	{
		"type": "function",
		"name": "updateMethodControllers",
		"inputs": [
			{"name": "methodSigs", "type": "bytes4[]"},
			{"name": "controllers", "type": "address[]"}
		],
		"outputs": []
	},
	{
		"type": "function",
		"name": "setMethodController",
		"inputs": [
			{"name": "methodSigs", "type": "bytes4[]"},
			{"name": "controller", "type": "address"}
		],
		"outputs": []
	},
	{
		"type": "function",
		"name": "emergencyFreeze",
		"inputs": [
			{"name": "target", "type": "address"},
			{"name": "freeze", "type": "bool"}
		],
		"outputs": []
	}
]`

// MethodConfig represents the configuration for a method in the enterprise wallet
type MethodConfig struct {
	Controller common.Address
}

// InitParams represents the initialization parameters for an enterprise wallet
type InitParams struct {
	Methods    [][4]byte
	Configs    []MethodConfig
	SuperAdmin common.Address
}

// CreateEnterpriseWalletData creates the call data for deploying an enterprise wallet through the factory
func CreateEnterpriseWalletData(implementation common.Address, salt [32]byte, params InitParams) ([]byte, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletFactoryABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWalletFactory ABI: %w", err)
	}

	// Convert InitParams to the format expected by the ABI
	type abiMethodConfig struct {
		Controller common.Address
	}
	type abiInitParams struct {
		Methods    [][4]byte
		Configs    []abiMethodConfig
		SuperAdmin common.Address
	}

	abiConfigs := make([]abiMethodConfig, len(params.Configs))
	for i, cfg := range params.Configs {
		abiConfigs[i] = abiMethodConfig{Controller: cfg.Controller}
	}

	abiParams := abiInitParams{
		Methods:    params.Methods,
		Configs:    abiConfigs,
		SuperAdmin: params.SuperAdmin,
	}

	// Encode the createWallet function call
	data, err := parsedABI.Pack("createWallet", implementation, salt, abiParams)
	if err != nil {
		return nil, fmt.Errorf("failed to encode createWallet call: %w", err)
	}

	return data, nil
}

// PredictEnterpriseWalletAddress predicts the address of an enterprise wallet before deployment
func PredictEnterpriseWalletAddress(implementation common.Address, salt [32]byte, deployer common.Address, factoryAddress common.Address) (common.Address, error) {
	// The factory uses: keccak256(abi.encodePacked(deployer, salt))
	fullSalt := crypto.Keccak256Hash(
		append(deployer.Bytes(), salt[:]...),
	)

	// Calculate the clones CREATE2 address
	// Clones uses: keccak256(abi.encodePacked(hex"ff", address, salt, keccak256(code)))
	// The code for minimal proxy is: 0x3d602d80600a3d3981f3363d3d373d3d3d363d73<implementation>5af43d82803e903d91602b57fd5bf3
	code := append(
		[]byte{0x3d, 0x60, 0x2d, 0x80, 0x60, 0x0a, 0x3d, 0x39, 0x81, 0xf3, 0x36, 0x3d, 0x3d, 0x37, 0x3d, 0x3d, 0x3d, 0x36, 0x3d, 0x73},
		append(implementation.Bytes(), []byte{0x5a, 0xf4, 0x3d, 0x82, 0x80, 0x3e, 0x90, 0x3d, 0x91, 0x60, 0x2b, 0x57, 0xfd, 0x5b, 0xf3}...)...,
	)

	codeHash := crypto.Keccak256Hash(code)

	// Calculate CREATE2 address: keccak256(0xff ++ factory ++ fullSalt ++ codeHash)[12:]
	data := append([]byte{0xff}, factoryAddress.Bytes()...)
	data = append(data, fullSalt.Bytes()...)
	data = append(data, codeHash.Bytes()...)

	hash := crypto.Keccak256Hash(data)
	return common.BytesToAddress(hash[12:]), nil
}

// CreatePaymentAccountData creates the call data for creating a payment account
func CreatePaymentAccountData(name string, controller common.Address) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWallet ABI: %w", err)
	}

	data, err := parsedABI.Pack("createPaymentAccount", name, controller)
	if err != nil {
		return nil, fmt.Errorf("failed to encode createPaymentAccount call: %w", err)
	}

	return data, nil
}

// CreateCollectionAccountData creates the call data for creating a collection account
func CreateCollectionAccountData(name string, collectionTarget common.Address) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWallet ABI: %w", err)
	}

	data, err := parsedABI.Pack("createCollectionAccount", name, collectionTarget)
	if err != nil {
		return nil, fmt.Errorf("failed to encode createCollectionAccount call: %w", err)
	}

	return data, nil
}

// ApproveTokenForPaymentData creates the call data for approving tokens for a payment account
func ApproveTokenForPaymentData(token common.Address, paymentAccount common.Address, amount *big.Int) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWallet ABI: %w", err)
	}

	data, err := parsedABI.Pack("approveTokenForPayment", token, paymentAccount, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to encode approveTokenForPayment call: %w", err)
	}

	return data, nil
}

// TransferETHToPaymentData creates the call data for transferring ETH to a payment account
func TransferETHToPaymentData(paymentAccount common.Address, amount *big.Int) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWallet ABI: %w", err)
	}

	data, err := parsedABI.Pack("transferETHToPayment", paymentAccount, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to encode transferETHToPayment call: %w", err)
	}

	return data, nil
}

// CollectFundsData creates the call data for collecting funds from a collection account
func CollectFundsData(token common.Address, collectionAccount common.Address) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWallet ABI: %w", err)
	}

	data, err := parsedABI.Pack("collectFunds", token, collectionAccount)
	if err != nil {
		return nil, fmt.Errorf("failed to encode collectFunds call: %w", err)
	}

	return data, nil
}

// GetMethodSelector returns the 4-byte method selector for a function signature
func GetMethodSelector(signature string) [4]byte {
	hash := crypto.Keccak256([]byte(signature))
	var selector [4]byte
	copy(selector[:], hash[:4])
	return selector
}

// Common method selectors for enterprise wallet
var (
	CreatePaymentAccountSelector    = GetMethodSelector("createPaymentAccount(string,address)")
	CreateCollectionAccountSelector = GetMethodSelector("createCollectionAccount(string,address)")
	ApproveTokenForPaymentSelector  = GetMethodSelector("approveTokenForPayment(address,address,uint256)")
	TransferETHToPaymentSelector    = GetMethodSelector("transferETHToPayment(address,uint256)")
	CollectFundsSelector            = GetMethodSelector("collectFunds(address,address)")

	// SuperAdmin transfer selectors
	ProposeSuperAdminTransferSelector = GetMethodSelector("proposeSuperAdminTransfer(address,uint256)")
	ConfirmSuperAdminTransferSelector = GetMethodSelector("confirmSuperAdminTransfer(uint256)")
	CancelSuperAdminTransferSelector  = GetMethodSelector("cancelSuperAdminTransfer(uint256)")
)

// ProposeSuperAdminTransferData creates the call data for proposing a super admin transfer
func ProposeSuperAdminTransferData(newSuperAdmin common.Address, timeout *big.Int) ([]byte, error) {
	// Manually encode the function call
	selector := ProposeSuperAdminTransferSelector

	// Encode parameters (address, uint256)
	data := make([]byte, 4+32+32)
	copy(data[0:4], selector[:])
	copy(data[4+12:4+32], newSuperAdmin.Bytes())

	// Encode timeout as uint256
	timeoutBytes := timeout.Bytes()
	copy(data[36+32-len(timeoutBytes):36+32], timeoutBytes)

	return data, nil
}

// ConfirmSuperAdminTransferData creates the call data for confirming a super admin transfer
func ConfirmSuperAdminTransferData(proposalId *big.Int) ([]byte, error) {
	selector := ConfirmSuperAdminTransferSelector

	// Encode parameter (uint256)
	data := make([]byte, 4+32)
	copy(data[0:4], selector[:])

	// Encode proposalId as uint256
	proposalIdBytes := proposalId.Bytes()
	copy(data[4+32-len(proposalIdBytes):4+32], proposalIdBytes)

	return data, nil
}

// CancelSuperAdminTransferData creates the call data for cancelling a super admin transfer
func CancelSuperAdminTransferData(proposalId *big.Int) ([]byte, error) {
	selector := CancelSuperAdminTransferSelector

	// Encode parameter (uint256)
	data := make([]byte, 4+32)
	copy(data[0:4], selector[:])

	// Encode proposalId as uint256
	proposalIdBytes := proposalId.Bytes()
	copy(data[4+32-len(proposalIdBytes):4+32], proposalIdBytes)

	return data, nil
}

// UpdateMethodControllersData creates the call data for batch updating method controllers with different controllers
// This allows updating multiple methods with different controllers in a single transaction
func UpdateMethodControllersData(methodSigs [][4]byte, controllers []common.Address) ([]byte, error) {
	if len(methodSigs) != len(controllers) {
		return nil, fmt.Errorf("methodSigs and controllers must have the same length")
	}
	if len(methodSigs) == 0 {
		return nil, fmt.Errorf("methodSigs cannot be empty")
	}

	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWallet ABI: %w", err)
	}

	data, err := parsedABI.Pack("updateMethodControllers", methodSigs, controllers)
	if err != nil {
		return nil, fmt.Errorf("failed to encode updateMethodControllers call: %w", err)
	}

	return data, nil
}

// SetMethodControllerData creates the call data for batch setting the same controller for multiple methods
// This allows setting the same controller for multiple methods in a single transaction
func SetMethodControllerData(methodSigs [][4]byte, controller common.Address) ([]byte, error) {
	if len(methodSigs) == 0 {
		return nil, fmt.Errorf("methodSigs cannot be empty")
	}

	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWallet ABI: %w", err)
	}

	data, err := parsedABI.Pack("setMethodController", methodSigs, controller)
	if err != nil {
		return nil, fmt.Errorf("failed to encode setMethodController call: %w", err)
	}

	return data, nil
}

// EmergencyFreezeData creates the call data for emergency freezing/unfreezing an account
// Note: This now requires method controller permission (changed from superAdmin)
func EmergencyFreezeData(target common.Address, freeze bool) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnterpriseWallet ABI: %w", err)
	}

	data, err := parsedABI.Pack("emergencyFreeze", target, freeze)
	if err != nil {
		return nil, fmt.Errorf("failed to encode emergencyFreeze call: %w", err)
	}

	return data, nil
}
