// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IAccountManagementAccountInfo is an auto generated low-level Go binding around an user-defined struct.
type IAccountManagementAccountInfo struct {
	Account   common.Address
	CreatedAt *big.Int
	IsActive  bool
}

// IEnterpriseWalletTypesMethodConfig is an auto generated low-level Go binding around an user-defined struct.
type IEnterpriseWalletTypesMethodConfig struct {
	Controller common.Address
}

// IEnterpriseWalletTypesSafeSetupParams is an auto generated low-level Go binding around an user-defined struct.
type IEnterpriseWalletTypesSafeSetupParams struct {
	Owners          []common.Address
	Threshold       *big.Int
	To              common.Address
	Data            []byte
	FallbackHandler common.Address
	PaymentToken    common.Address
	Payment         *big.Int
	PaymentReceiver common.Address
	SaltNonce       *big.Int
}

// ISuperAdminManagementSuperAdminTransfer is an auto generated low-level Go binding around an user-defined struct.
type ISuperAdminManagementSuperAdminTransfer struct {
	CurrentSuperAdmin  common.Address
	ProposedSuperAdmin common.Address
	ProposedAt         *big.Int
	Timeout            *big.Int
	IsActive           bool
}

// EnterpriseWalletMetaData contains all meta data concerning the EnterpriseWallet contract.
var EnterpriseWalletMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"batchCollectFunds\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collectionAccounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchEmergencyFreeze\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"freeze\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelSuperAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collectFunds\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collectionAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmSuperAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createCollectionAccount\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"collectionTarget\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createPaymentAccount\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSafeAndCollectionAccount\",\"inputs\":[{\"name\":\"proxyFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"safeSingleton\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"safeParams\",\"type\":\"tuple\",\"internalType\":\"structIEnterpriseWalletTypes.SafeSetupParams\",\"components\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fallbackHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"saltNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"collectionTarget\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"safe\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collectionAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSafeAndPaymentAccount\",\"inputs\":[{\"name\":\"proxyFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"safeSingleton\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"safeParams\",\"type\":\"tuple\",\"internalType\":\"structIEnterpriseWalletTypes.SafeSetupParams\",\"components\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fallbackHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"saltNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"safe\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndUpdatePaymentAccountController\",\"inputs\":[{\"name\":\"paymentAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proxyFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"safeSingleton\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"safeParams\",\"type\":\"tuple\",\"internalType\":\"structIEnterpriseWalletTypes.SafeSetupParams\",\"components\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fallbackHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"saltNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emergencyFreeze\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"freeze\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emergencyPause\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCollectionAccountByIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAccountManagement.AccountInfo\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectionAccountNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectionAccounts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAccountManagement.AccountInfo[]\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectionAccountsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectionAccountsPaginated\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"accounts\",\"type\":\"tuple[]\",\"internalType\":\"structIAccountManagement.AccountInfo[]\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"total\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMethodConfig\",\"inputs\":[{\"name\":\"methodSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEnterpriseWalletTypes.MethodConfig\",\"components\":[{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentAccountByIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAccountManagement.AccountInfo\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentAccountNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentAccounts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAccountManagement.AccountInfo[]\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentAccountsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentAccountsPaginated\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"accounts\",\"type\":\"tuple[]\",\"internalType\":\"structIAccountManagement.AccountInfo[]\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"total\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSuperAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSuperAdminTransfer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISuperAdminManagement.SuperAdminTransfer\",\"components\":[{\"name\":\"currentSuperAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposedSuperAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"methods\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structIEnterpriseWalletTypes.MethodConfig[]\",\"components\":[{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"superAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isCollectionAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFrozen\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaymentAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSuperAdminTransfer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"predictCollectionAccountAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"predictPaymentAccountAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeSuperAdminTransfer\",\"inputs\":[{\"name\":\"newSuperAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rescueFunds\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCollectionTarget\",\"inputs\":[{\"name\":\"collectionAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMethodController\",\"inputs\":[{\"name\":\"methodSigs\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"},{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMethodController\",\"inputs\":[{\"name\":\"methodSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMethodControllers\",\"inputs\":[{\"name\":\"methodSigs\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"},{\"name\":\"controllers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePaymentAccountController\",\"inputs\":[{\"name\":\"paymentAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollectionAccountCreated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyFreeze\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPause\",\"inputs\":[{\"name\":\"paused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsCollected\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsRescued\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MethodControllerUpdated\",\"inputs\":[{\"name\":\"methodSig\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentAccountCreated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SafeAndCollectionAccountCreated\",\"inputs\":[{\"name\":\"safe\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collectionAccount\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SafeAndPaymentAccountCreated\",\"inputs\":[{\"name\":\"safe\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentAccount\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuperAdminTransferCancelled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuperAdminTransferProposed\",\"inputs\":[{\"name\":\"currentSuperAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposedSuperAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuperAdminTransferred\",\"inputs\":[{\"name\":\"oldSuperAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newSuperAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccountNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ActiveProposalExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFactoryAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMethodConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSafeAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSafeParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTimeout\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProposalExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProposalNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeDeploymentFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SameAddressAsCurrentAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TargetFrozen\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedProposal\",\"inputs\":[]}]",
}

// EnterpriseWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use EnterpriseWalletMetaData.ABI instead.
var EnterpriseWalletABI = EnterpriseWalletMetaData.ABI

// EnterpriseWallet is an auto generated Go binding around an Ethereum contract.
type EnterpriseWallet struct {
	EnterpriseWalletCaller     // Read-only binding to the contract
	EnterpriseWalletTransactor // Write-only binding to the contract
	EnterpriseWalletFilterer   // Log filterer for contract events
}

// EnterpriseWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnterpriseWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnterpriseWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnterpriseWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnterpriseWalletSession struct {
	Contract     *EnterpriseWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnterpriseWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnterpriseWalletCallerSession struct {
	Contract *EnterpriseWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EnterpriseWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnterpriseWalletTransactorSession struct {
	Contract     *EnterpriseWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EnterpriseWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnterpriseWalletRaw struct {
	Contract *EnterpriseWallet // Generic contract binding to access the raw methods on
}

// EnterpriseWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnterpriseWalletCallerRaw struct {
	Contract *EnterpriseWalletCaller // Generic read-only contract binding to access the raw methods on
}

// EnterpriseWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnterpriseWalletTransactorRaw struct {
	Contract *EnterpriseWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnterpriseWallet creates a new instance of EnterpriseWallet, bound to a specific deployed contract.
func NewEnterpriseWallet(address common.Address, backend bind.ContractBackend) (*EnterpriseWallet, error) {
	contract, err := bindEnterpriseWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWallet{EnterpriseWalletCaller: EnterpriseWalletCaller{contract: contract}, EnterpriseWalletTransactor: EnterpriseWalletTransactor{contract: contract}, EnterpriseWalletFilterer: EnterpriseWalletFilterer{contract: contract}}, nil
}

// NewEnterpriseWalletCaller creates a new read-only instance of EnterpriseWallet, bound to a specific deployed contract.
func NewEnterpriseWalletCaller(address common.Address, caller bind.ContractCaller) (*EnterpriseWalletCaller, error) {
	contract, err := bindEnterpriseWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletCaller{contract: contract}, nil
}

// NewEnterpriseWalletTransactor creates a new write-only instance of EnterpriseWallet, bound to a specific deployed contract.
func NewEnterpriseWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*EnterpriseWalletTransactor, error) {
	contract, err := bindEnterpriseWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletTransactor{contract: contract}, nil
}

// NewEnterpriseWalletFilterer creates a new log filterer instance of EnterpriseWallet, bound to a specific deployed contract.
func NewEnterpriseWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*EnterpriseWalletFilterer, error) {
	contract, err := bindEnterpriseWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFilterer{contract: contract}, nil
}

// bindEnterpriseWallet binds a generic wrapper to an already deployed contract.
func bindEnterpriseWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnterpriseWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseWallet *EnterpriseWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseWallet.Contract.EnterpriseWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseWallet *EnterpriseWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.EnterpriseWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseWallet *EnterpriseWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.EnterpriseWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseWallet *EnterpriseWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseWallet *EnterpriseWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseWallet *EnterpriseWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.contract.Transact(opts, method, params...)
}

// GetCollectionAccountByIndex is a free data retrieval call binding the contract method 0xed05263d.
//
// Solidity: function getCollectionAccountByIndex(uint256 index) view returns((address,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletCaller) GetCollectionAccountByIndex(opts *bind.CallOpts, index *big.Int) (IAccountManagementAccountInfo, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getCollectionAccountByIndex", index)

	if err != nil {
		return *new(IAccountManagementAccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IAccountManagementAccountInfo)).(*IAccountManagementAccountInfo)

	return out0, err

}

// GetCollectionAccountByIndex is a free data retrieval call binding the contract method 0xed05263d.
//
// Solidity: function getCollectionAccountByIndex(uint256 index) view returns((address,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletSession) GetCollectionAccountByIndex(index *big.Int) (IAccountManagementAccountInfo, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccountByIndex(&_EnterpriseWallet.CallOpts, index)
}

// GetCollectionAccountByIndex is a free data retrieval call binding the contract method 0xed05263d.
//
// Solidity: function getCollectionAccountByIndex(uint256 index) view returns((address,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetCollectionAccountByIndex(index *big.Int) (IAccountManagementAccountInfo, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccountByIndex(&_EnterpriseWallet.CallOpts, index)
}

// GetCollectionAccountNonce is a free data retrieval call binding the contract method 0x82be3dde.
//
// Solidity: function getCollectionAccountNonce() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletCaller) GetCollectionAccountNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getCollectionAccountNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollectionAccountNonce is a free data retrieval call binding the contract method 0x82be3dde.
//
// Solidity: function getCollectionAccountNonce() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletSession) GetCollectionAccountNonce() (*big.Int, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccountNonce(&_EnterpriseWallet.CallOpts)
}

// GetCollectionAccountNonce is a free data retrieval call binding the contract method 0x82be3dde.
//
// Solidity: function getCollectionAccountNonce() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetCollectionAccountNonce() (*big.Int, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccountNonce(&_EnterpriseWallet.CallOpts)
}

// GetCollectionAccounts is a free data retrieval call binding the contract method 0x0c6dcef7.
//
// Solidity: function getCollectionAccounts() view returns((address,uint256,bool)[])
func (_EnterpriseWallet *EnterpriseWalletCaller) GetCollectionAccounts(opts *bind.CallOpts) ([]IAccountManagementAccountInfo, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getCollectionAccounts")

	if err != nil {
		return *new([]IAccountManagementAccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountManagementAccountInfo)).(*[]IAccountManagementAccountInfo)

	return out0, err

}

// GetCollectionAccounts is a free data retrieval call binding the contract method 0x0c6dcef7.
//
// Solidity: function getCollectionAccounts() view returns((address,uint256,bool)[])
func (_EnterpriseWallet *EnterpriseWalletSession) GetCollectionAccounts() ([]IAccountManagementAccountInfo, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccounts(&_EnterpriseWallet.CallOpts)
}

// GetCollectionAccounts is a free data retrieval call binding the contract method 0x0c6dcef7.
//
// Solidity: function getCollectionAccounts() view returns((address,uint256,bool)[])
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetCollectionAccounts() ([]IAccountManagementAccountInfo, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccounts(&_EnterpriseWallet.CallOpts)
}

// GetCollectionAccountsCount is a free data retrieval call binding the contract method 0x0eb03246.
//
// Solidity: function getCollectionAccountsCount() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletCaller) GetCollectionAccountsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getCollectionAccountsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollectionAccountsCount is a free data retrieval call binding the contract method 0x0eb03246.
//
// Solidity: function getCollectionAccountsCount() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletSession) GetCollectionAccountsCount() (*big.Int, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccountsCount(&_EnterpriseWallet.CallOpts)
}

// GetCollectionAccountsCount is a free data retrieval call binding the contract method 0x0eb03246.
//
// Solidity: function getCollectionAccountsCount() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetCollectionAccountsCount() (*big.Int, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccountsCount(&_EnterpriseWallet.CallOpts)
}

// GetCollectionAccountsPaginated is a free data retrieval call binding the contract method 0x7b81d00d.
//
// Solidity: function getCollectionAccountsPaginated(uint256 offset, uint256 limit) view returns((address,uint256,bool)[] accounts, uint256 total)
func (_EnterpriseWallet *EnterpriseWalletCaller) GetCollectionAccountsPaginated(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	Accounts []IAccountManagementAccountInfo
	Total    *big.Int
}, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getCollectionAccountsPaginated", offset, limit)

	outstruct := new(struct {
		Accounts []IAccountManagementAccountInfo
		Total    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Accounts = *abi.ConvertType(out[0], new([]IAccountManagementAccountInfo)).(*[]IAccountManagementAccountInfo)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCollectionAccountsPaginated is a free data retrieval call binding the contract method 0x7b81d00d.
//
// Solidity: function getCollectionAccountsPaginated(uint256 offset, uint256 limit) view returns((address,uint256,bool)[] accounts, uint256 total)
func (_EnterpriseWallet *EnterpriseWalletSession) GetCollectionAccountsPaginated(offset *big.Int, limit *big.Int) (struct {
	Accounts []IAccountManagementAccountInfo
	Total    *big.Int
}, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccountsPaginated(&_EnterpriseWallet.CallOpts, offset, limit)
}

// GetCollectionAccountsPaginated is a free data retrieval call binding the contract method 0x7b81d00d.
//
// Solidity: function getCollectionAccountsPaginated(uint256 offset, uint256 limit) view returns((address,uint256,bool)[] accounts, uint256 total)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetCollectionAccountsPaginated(offset *big.Int, limit *big.Int) (struct {
	Accounts []IAccountManagementAccountInfo
	Total    *big.Int
}, error) {
	return _EnterpriseWallet.Contract.GetCollectionAccountsPaginated(&_EnterpriseWallet.CallOpts, offset, limit)
}

// GetMethodConfig is a free data retrieval call binding the contract method 0x3a6d19d5.
//
// Solidity: function getMethodConfig(bytes4 methodSig) view returns((address))
func (_EnterpriseWallet *EnterpriseWalletCaller) GetMethodConfig(opts *bind.CallOpts, methodSig [4]byte) (IEnterpriseWalletTypesMethodConfig, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getMethodConfig", methodSig)

	if err != nil {
		return *new(IEnterpriseWalletTypesMethodConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEnterpriseWalletTypesMethodConfig)).(*IEnterpriseWalletTypesMethodConfig)

	return out0, err

}

// GetMethodConfig is a free data retrieval call binding the contract method 0x3a6d19d5.
//
// Solidity: function getMethodConfig(bytes4 methodSig) view returns((address))
func (_EnterpriseWallet *EnterpriseWalletSession) GetMethodConfig(methodSig [4]byte) (IEnterpriseWalletTypesMethodConfig, error) {
	return _EnterpriseWallet.Contract.GetMethodConfig(&_EnterpriseWallet.CallOpts, methodSig)
}

// GetMethodConfig is a free data retrieval call binding the contract method 0x3a6d19d5.
//
// Solidity: function getMethodConfig(bytes4 methodSig) view returns((address))
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetMethodConfig(methodSig [4]byte) (IEnterpriseWalletTypesMethodConfig, error) {
	return _EnterpriseWallet.Contract.GetMethodConfig(&_EnterpriseWallet.CallOpts, methodSig)
}

// GetPaymentAccountByIndex is a free data retrieval call binding the contract method 0x8a22c6da.
//
// Solidity: function getPaymentAccountByIndex(uint256 index) view returns((address,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletCaller) GetPaymentAccountByIndex(opts *bind.CallOpts, index *big.Int) (IAccountManagementAccountInfo, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getPaymentAccountByIndex", index)

	if err != nil {
		return *new(IAccountManagementAccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IAccountManagementAccountInfo)).(*IAccountManagementAccountInfo)

	return out0, err

}

// GetPaymentAccountByIndex is a free data retrieval call binding the contract method 0x8a22c6da.
//
// Solidity: function getPaymentAccountByIndex(uint256 index) view returns((address,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletSession) GetPaymentAccountByIndex(index *big.Int) (IAccountManagementAccountInfo, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccountByIndex(&_EnterpriseWallet.CallOpts, index)
}

// GetPaymentAccountByIndex is a free data retrieval call binding the contract method 0x8a22c6da.
//
// Solidity: function getPaymentAccountByIndex(uint256 index) view returns((address,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetPaymentAccountByIndex(index *big.Int) (IAccountManagementAccountInfo, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccountByIndex(&_EnterpriseWallet.CallOpts, index)
}

// GetPaymentAccountNonce is a free data retrieval call binding the contract method 0x0282ee13.
//
// Solidity: function getPaymentAccountNonce() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletCaller) GetPaymentAccountNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getPaymentAccountNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentAccountNonce is a free data retrieval call binding the contract method 0x0282ee13.
//
// Solidity: function getPaymentAccountNonce() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletSession) GetPaymentAccountNonce() (*big.Int, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccountNonce(&_EnterpriseWallet.CallOpts)
}

// GetPaymentAccountNonce is a free data retrieval call binding the contract method 0x0282ee13.
//
// Solidity: function getPaymentAccountNonce() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetPaymentAccountNonce() (*big.Int, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccountNonce(&_EnterpriseWallet.CallOpts)
}

// GetPaymentAccounts is a free data retrieval call binding the contract method 0xcda1988f.
//
// Solidity: function getPaymentAccounts() view returns((address,uint256,bool)[])
func (_EnterpriseWallet *EnterpriseWalletCaller) GetPaymentAccounts(opts *bind.CallOpts) ([]IAccountManagementAccountInfo, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getPaymentAccounts")

	if err != nil {
		return *new([]IAccountManagementAccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAccountManagementAccountInfo)).(*[]IAccountManagementAccountInfo)

	return out0, err

}

// GetPaymentAccounts is a free data retrieval call binding the contract method 0xcda1988f.
//
// Solidity: function getPaymentAccounts() view returns((address,uint256,bool)[])
func (_EnterpriseWallet *EnterpriseWalletSession) GetPaymentAccounts() ([]IAccountManagementAccountInfo, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccounts(&_EnterpriseWallet.CallOpts)
}

// GetPaymentAccounts is a free data retrieval call binding the contract method 0xcda1988f.
//
// Solidity: function getPaymentAccounts() view returns((address,uint256,bool)[])
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetPaymentAccounts() ([]IAccountManagementAccountInfo, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccounts(&_EnterpriseWallet.CallOpts)
}

// GetPaymentAccountsCount is a free data retrieval call binding the contract method 0x31b35131.
//
// Solidity: function getPaymentAccountsCount() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletCaller) GetPaymentAccountsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getPaymentAccountsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentAccountsCount is a free data retrieval call binding the contract method 0x31b35131.
//
// Solidity: function getPaymentAccountsCount() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletSession) GetPaymentAccountsCount() (*big.Int, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccountsCount(&_EnterpriseWallet.CallOpts)
}

// GetPaymentAccountsCount is a free data retrieval call binding the contract method 0x31b35131.
//
// Solidity: function getPaymentAccountsCount() view returns(uint256)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetPaymentAccountsCount() (*big.Int, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccountsCount(&_EnterpriseWallet.CallOpts)
}

// GetPaymentAccountsPaginated is a free data retrieval call binding the contract method 0x4e50debf.
//
// Solidity: function getPaymentAccountsPaginated(uint256 offset, uint256 limit) view returns((address,uint256,bool)[] accounts, uint256 total)
func (_EnterpriseWallet *EnterpriseWalletCaller) GetPaymentAccountsPaginated(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	Accounts []IAccountManagementAccountInfo
	Total    *big.Int
}, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getPaymentAccountsPaginated", offset, limit)

	outstruct := new(struct {
		Accounts []IAccountManagementAccountInfo
		Total    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Accounts = *abi.ConvertType(out[0], new([]IAccountManagementAccountInfo)).(*[]IAccountManagementAccountInfo)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPaymentAccountsPaginated is a free data retrieval call binding the contract method 0x4e50debf.
//
// Solidity: function getPaymentAccountsPaginated(uint256 offset, uint256 limit) view returns((address,uint256,bool)[] accounts, uint256 total)
func (_EnterpriseWallet *EnterpriseWalletSession) GetPaymentAccountsPaginated(offset *big.Int, limit *big.Int) (struct {
	Accounts []IAccountManagementAccountInfo
	Total    *big.Int
}, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccountsPaginated(&_EnterpriseWallet.CallOpts, offset, limit)
}

// GetPaymentAccountsPaginated is a free data retrieval call binding the contract method 0x4e50debf.
//
// Solidity: function getPaymentAccountsPaginated(uint256 offset, uint256 limit) view returns((address,uint256,bool)[] accounts, uint256 total)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetPaymentAccountsPaginated(offset *big.Int, limit *big.Int) (struct {
	Accounts []IAccountManagementAccountInfo
	Total    *big.Int
}, error) {
	return _EnterpriseWallet.Contract.GetPaymentAccountsPaginated(&_EnterpriseWallet.CallOpts, offset, limit)
}

// GetSuperAdmin is a free data retrieval call binding the contract method 0x8204c326.
//
// Solidity: function getSuperAdmin() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletCaller) GetSuperAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getSuperAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSuperAdmin is a free data retrieval call binding the contract method 0x8204c326.
//
// Solidity: function getSuperAdmin() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletSession) GetSuperAdmin() (common.Address, error) {
	return _EnterpriseWallet.Contract.GetSuperAdmin(&_EnterpriseWallet.CallOpts)
}

// GetSuperAdmin is a free data retrieval call binding the contract method 0x8204c326.
//
// Solidity: function getSuperAdmin() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetSuperAdmin() (common.Address, error) {
	return _EnterpriseWallet.Contract.GetSuperAdmin(&_EnterpriseWallet.CallOpts)
}

// GetSuperAdminTransfer is a free data retrieval call binding the contract method 0xeccd9d29.
//
// Solidity: function getSuperAdminTransfer() view returns((address,address,uint256,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletCaller) GetSuperAdminTransfer(opts *bind.CallOpts) (ISuperAdminManagementSuperAdminTransfer, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "getSuperAdminTransfer")

	if err != nil {
		return *new(ISuperAdminManagementSuperAdminTransfer), err
	}

	out0 := *abi.ConvertType(out[0], new(ISuperAdminManagementSuperAdminTransfer)).(*ISuperAdminManagementSuperAdminTransfer)

	return out0, err

}

// GetSuperAdminTransfer is a free data retrieval call binding the contract method 0xeccd9d29.
//
// Solidity: function getSuperAdminTransfer() view returns((address,address,uint256,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletSession) GetSuperAdminTransfer() (ISuperAdminManagementSuperAdminTransfer, error) {
	return _EnterpriseWallet.Contract.GetSuperAdminTransfer(&_EnterpriseWallet.CallOpts)
}

// GetSuperAdminTransfer is a free data retrieval call binding the contract method 0xeccd9d29.
//
// Solidity: function getSuperAdminTransfer() view returns((address,address,uint256,uint256,bool))
func (_EnterpriseWallet *EnterpriseWalletCallerSession) GetSuperAdminTransfer() (ISuperAdminManagementSuperAdminTransfer, error) {
	return _EnterpriseWallet.Contract.GetSuperAdminTransfer(&_EnterpriseWallet.CallOpts)
}

// IsCollectionAccount is a free data retrieval call binding the contract method 0x0f10c8c8.
//
// Solidity: function isCollectionAccount(address account) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCaller) IsCollectionAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "isCollectionAccount", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollectionAccount is a free data retrieval call binding the contract method 0x0f10c8c8.
//
// Solidity: function isCollectionAccount(address account) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletSession) IsCollectionAccount(account common.Address) (bool, error) {
	return _EnterpriseWallet.Contract.IsCollectionAccount(&_EnterpriseWallet.CallOpts, account)
}

// IsCollectionAccount is a free data retrieval call binding the contract method 0x0f10c8c8.
//
// Solidity: function isCollectionAccount(address account) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) IsCollectionAccount(account common.Address) (bool, error) {
	return _EnterpriseWallet.Contract.IsCollectionAccount(&_EnterpriseWallet.CallOpts, account)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address target) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCaller) IsFrozen(opts *bind.CallOpts, target common.Address) (bool, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "isFrozen", target)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address target) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletSession) IsFrozen(target common.Address) (bool, error) {
	return _EnterpriseWallet.Contract.IsFrozen(&_EnterpriseWallet.CallOpts, target)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address target) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) IsFrozen(target common.Address) (bool, error) {
	return _EnterpriseWallet.Contract.IsFrozen(&_EnterpriseWallet.CallOpts, target)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletSession) IsPaused() (bool, error) {
	return _EnterpriseWallet.Contract.IsPaused(&_EnterpriseWallet.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) IsPaused() (bool, error) {
	return _EnterpriseWallet.Contract.IsPaused(&_EnterpriseWallet.CallOpts)
}

// IsPaymentAccount is a free data retrieval call binding the contract method 0x7f6a78cd.
//
// Solidity: function isPaymentAccount(address account) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCaller) IsPaymentAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "isPaymentAccount", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaymentAccount is a free data retrieval call binding the contract method 0x7f6a78cd.
//
// Solidity: function isPaymentAccount(address account) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletSession) IsPaymentAccount(account common.Address) (bool, error) {
	return _EnterpriseWallet.Contract.IsPaymentAccount(&_EnterpriseWallet.CallOpts, account)
}

// IsPaymentAccount is a free data retrieval call binding the contract method 0x7f6a78cd.
//
// Solidity: function isPaymentAccount(address account) view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) IsPaymentAccount(account common.Address) (bool, error) {
	return _EnterpriseWallet.Contract.IsPaymentAccount(&_EnterpriseWallet.CallOpts, account)
}

// IsValidSuperAdminTransfer is a free data retrieval call binding the contract method 0x61f1b5d8.
//
// Solidity: function isValidSuperAdminTransfer() view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCaller) IsValidSuperAdminTransfer(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "isValidSuperAdminTransfer")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidSuperAdminTransfer is a free data retrieval call binding the contract method 0x61f1b5d8.
//
// Solidity: function isValidSuperAdminTransfer() view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletSession) IsValidSuperAdminTransfer() (bool, error) {
	return _EnterpriseWallet.Contract.IsValidSuperAdminTransfer(&_EnterpriseWallet.CallOpts)
}

// IsValidSuperAdminTransfer is a free data retrieval call binding the contract method 0x61f1b5d8.
//
// Solidity: function isValidSuperAdminTransfer() view returns(bool)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) IsValidSuperAdminTransfer() (bool, error) {
	return _EnterpriseWallet.Contract.IsValidSuperAdminTransfer(&_EnterpriseWallet.CallOpts)
}

// PredictCollectionAccountAddress is a free data retrieval call binding the contract method 0x82d72e2f.
//
// Solidity: function predictCollectionAccountAddress() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletCaller) PredictCollectionAccountAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "predictCollectionAccountAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictCollectionAccountAddress is a free data retrieval call binding the contract method 0x82d72e2f.
//
// Solidity: function predictCollectionAccountAddress() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletSession) PredictCollectionAccountAddress() (common.Address, error) {
	return _EnterpriseWallet.Contract.PredictCollectionAccountAddress(&_EnterpriseWallet.CallOpts)
}

// PredictCollectionAccountAddress is a free data retrieval call binding the contract method 0x82d72e2f.
//
// Solidity: function predictCollectionAccountAddress() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) PredictCollectionAccountAddress() (common.Address, error) {
	return _EnterpriseWallet.Contract.PredictCollectionAccountAddress(&_EnterpriseWallet.CallOpts)
}

// PredictPaymentAccountAddress is a free data retrieval call binding the contract method 0x5ea854d5.
//
// Solidity: function predictPaymentAccountAddress() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletCaller) PredictPaymentAccountAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseWallet.contract.Call(opts, &out, "predictPaymentAccountAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictPaymentAccountAddress is a free data retrieval call binding the contract method 0x5ea854d5.
//
// Solidity: function predictPaymentAccountAddress() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletSession) PredictPaymentAccountAddress() (common.Address, error) {
	return _EnterpriseWallet.Contract.PredictPaymentAccountAddress(&_EnterpriseWallet.CallOpts)
}

// PredictPaymentAccountAddress is a free data retrieval call binding the contract method 0x5ea854d5.
//
// Solidity: function predictPaymentAccountAddress() view returns(address)
func (_EnterpriseWallet *EnterpriseWalletCallerSession) PredictPaymentAccountAddress() (common.Address, error) {
	return _EnterpriseWallet.Contract.PredictPaymentAccountAddress(&_EnterpriseWallet.CallOpts)
}

// BatchCollectFunds is a paid mutator transaction binding the contract method 0xa30796c6.
//
// Solidity: function batchCollectFunds(address token, address[] collectionAccounts) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) BatchCollectFunds(opts *bind.TransactOpts, token common.Address, collectionAccounts []common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "batchCollectFunds", token, collectionAccounts)
}

// BatchCollectFunds is a paid mutator transaction binding the contract method 0xa30796c6.
//
// Solidity: function batchCollectFunds(address token, address[] collectionAccounts) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) BatchCollectFunds(token common.Address, collectionAccounts []common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.BatchCollectFunds(&_EnterpriseWallet.TransactOpts, token, collectionAccounts)
}

// BatchCollectFunds is a paid mutator transaction binding the contract method 0xa30796c6.
//
// Solidity: function batchCollectFunds(address token, address[] collectionAccounts) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) BatchCollectFunds(token common.Address, collectionAccounts []common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.BatchCollectFunds(&_EnterpriseWallet.TransactOpts, token, collectionAccounts)
}

// BatchEmergencyFreeze is a paid mutator transaction binding the contract method 0x75682c13.
//
// Solidity: function batchEmergencyFreeze(address[] targets, bool freeze) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) BatchEmergencyFreeze(opts *bind.TransactOpts, targets []common.Address, freeze bool) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "batchEmergencyFreeze", targets, freeze)
}

// BatchEmergencyFreeze is a paid mutator transaction binding the contract method 0x75682c13.
//
// Solidity: function batchEmergencyFreeze(address[] targets, bool freeze) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) BatchEmergencyFreeze(targets []common.Address, freeze bool) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.BatchEmergencyFreeze(&_EnterpriseWallet.TransactOpts, targets, freeze)
}

// BatchEmergencyFreeze is a paid mutator transaction binding the contract method 0x75682c13.
//
// Solidity: function batchEmergencyFreeze(address[] targets, bool freeze) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) BatchEmergencyFreeze(targets []common.Address, freeze bool) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.BatchEmergencyFreeze(&_EnterpriseWallet.TransactOpts, targets, freeze)
}

// CancelSuperAdminTransfer is a paid mutator transaction binding the contract method 0xb4475f2d.
//
// Solidity: function cancelSuperAdminTransfer() returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) CancelSuperAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "cancelSuperAdminTransfer")
}

// CancelSuperAdminTransfer is a paid mutator transaction binding the contract method 0xb4475f2d.
//
// Solidity: function cancelSuperAdminTransfer() returns()
func (_EnterpriseWallet *EnterpriseWalletSession) CancelSuperAdminTransfer() (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CancelSuperAdminTransfer(&_EnterpriseWallet.TransactOpts)
}

// CancelSuperAdminTransfer is a paid mutator transaction binding the contract method 0xb4475f2d.
//
// Solidity: function cancelSuperAdminTransfer() returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) CancelSuperAdminTransfer() (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CancelSuperAdminTransfer(&_EnterpriseWallet.TransactOpts)
}

// CollectFunds is a paid mutator transaction binding the contract method 0xdd6890ef.
//
// Solidity: function collectFunds(address token, address collectionAccount) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) CollectFunds(opts *bind.TransactOpts, token common.Address, collectionAccount common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "collectFunds", token, collectionAccount)
}

// CollectFunds is a paid mutator transaction binding the contract method 0xdd6890ef.
//
// Solidity: function collectFunds(address token, address collectionAccount) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) CollectFunds(token common.Address, collectionAccount common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CollectFunds(&_EnterpriseWallet.TransactOpts, token, collectionAccount)
}

// CollectFunds is a paid mutator transaction binding the contract method 0xdd6890ef.
//
// Solidity: function collectFunds(address token, address collectionAccount) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) CollectFunds(token common.Address, collectionAccount common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CollectFunds(&_EnterpriseWallet.TransactOpts, token, collectionAccount)
}

// ConfirmSuperAdminTransfer is a paid mutator transaction binding the contract method 0x511bebf8.
//
// Solidity: function confirmSuperAdminTransfer() returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) ConfirmSuperAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "confirmSuperAdminTransfer")
}

// ConfirmSuperAdminTransfer is a paid mutator transaction binding the contract method 0x511bebf8.
//
// Solidity: function confirmSuperAdminTransfer() returns()
func (_EnterpriseWallet *EnterpriseWalletSession) ConfirmSuperAdminTransfer() (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.ConfirmSuperAdminTransfer(&_EnterpriseWallet.TransactOpts)
}

// ConfirmSuperAdminTransfer is a paid mutator transaction binding the contract method 0x511bebf8.
//
// Solidity: function confirmSuperAdminTransfer() returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) ConfirmSuperAdminTransfer() (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.ConfirmSuperAdminTransfer(&_EnterpriseWallet.TransactOpts)
}

// CreateAndUpdatePaymentAccountController is a paid mutator transaction binding the contract method 0x6f698fe2.
//
// Solidity: function createAndUpdatePaymentAccountController(address paymentAccount, address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) CreateAndUpdatePaymentAccountController(opts *bind.TransactOpts, paymentAccount common.Address, proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "createAndUpdatePaymentAccountController", paymentAccount, proxyFactory, safeSingleton, safeParams)
}

// CreateAndUpdatePaymentAccountController is a paid mutator transaction binding the contract method 0x6f698fe2.
//
// Solidity: function createAndUpdatePaymentAccountController(address paymentAccount, address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) CreateAndUpdatePaymentAccountController(paymentAccount common.Address, proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreateAndUpdatePaymentAccountController(&_EnterpriseWallet.TransactOpts, paymentAccount, proxyFactory, safeSingleton, safeParams)
}

// CreateAndUpdatePaymentAccountController is a paid mutator transaction binding the contract method 0x6f698fe2.
//
// Solidity: function createAndUpdatePaymentAccountController(address paymentAccount, address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) CreateAndUpdatePaymentAccountController(paymentAccount common.Address, proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreateAndUpdatePaymentAccountController(&_EnterpriseWallet.TransactOpts, paymentAccount, proxyFactory, safeSingleton, safeParams)
}

// CreateCollectionAccount is a paid mutator transaction binding the contract method 0xc8ac06ed.
//
// Solidity: function createCollectionAccount(string name, address collectionTarget) returns(address)
func (_EnterpriseWallet *EnterpriseWalletTransactor) CreateCollectionAccount(opts *bind.TransactOpts, name string, collectionTarget common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "createCollectionAccount", name, collectionTarget)
}

// CreateCollectionAccount is a paid mutator transaction binding the contract method 0xc8ac06ed.
//
// Solidity: function createCollectionAccount(string name, address collectionTarget) returns(address)
func (_EnterpriseWallet *EnterpriseWalletSession) CreateCollectionAccount(name string, collectionTarget common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreateCollectionAccount(&_EnterpriseWallet.TransactOpts, name, collectionTarget)
}

// CreateCollectionAccount is a paid mutator transaction binding the contract method 0xc8ac06ed.
//
// Solidity: function createCollectionAccount(string name, address collectionTarget) returns(address)
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) CreateCollectionAccount(name string, collectionTarget common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreateCollectionAccount(&_EnterpriseWallet.TransactOpts, name, collectionTarget)
}

// CreatePaymentAccount is a paid mutator transaction binding the contract method 0x08f25c4a.
//
// Solidity: function createPaymentAccount(string name, address controller) returns(address)
func (_EnterpriseWallet *EnterpriseWalletTransactor) CreatePaymentAccount(opts *bind.TransactOpts, name string, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "createPaymentAccount", name, controller)
}

// CreatePaymentAccount is a paid mutator transaction binding the contract method 0x08f25c4a.
//
// Solidity: function createPaymentAccount(string name, address controller) returns(address)
func (_EnterpriseWallet *EnterpriseWalletSession) CreatePaymentAccount(name string, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreatePaymentAccount(&_EnterpriseWallet.TransactOpts, name, controller)
}

// CreatePaymentAccount is a paid mutator transaction binding the contract method 0x08f25c4a.
//
// Solidity: function createPaymentAccount(string name, address controller) returns(address)
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) CreatePaymentAccount(name string, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreatePaymentAccount(&_EnterpriseWallet.TransactOpts, name, controller)
}

// CreateSafeAndCollectionAccount is a paid mutator transaction binding the contract method 0xbe13b69f.
//
// Solidity: function createSafeAndCollectionAccount(address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams, string name, address collectionTarget) returns(address safe, address collectionAccount)
func (_EnterpriseWallet *EnterpriseWalletTransactor) CreateSafeAndCollectionAccount(opts *bind.TransactOpts, proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams, name string, collectionTarget common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "createSafeAndCollectionAccount", proxyFactory, safeSingleton, safeParams, name, collectionTarget)
}

// CreateSafeAndCollectionAccount is a paid mutator transaction binding the contract method 0xbe13b69f.
//
// Solidity: function createSafeAndCollectionAccount(address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams, string name, address collectionTarget) returns(address safe, address collectionAccount)
func (_EnterpriseWallet *EnterpriseWalletSession) CreateSafeAndCollectionAccount(proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams, name string, collectionTarget common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreateSafeAndCollectionAccount(&_EnterpriseWallet.TransactOpts, proxyFactory, safeSingleton, safeParams, name, collectionTarget)
}

// CreateSafeAndCollectionAccount is a paid mutator transaction binding the contract method 0xbe13b69f.
//
// Solidity: function createSafeAndCollectionAccount(address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams, string name, address collectionTarget) returns(address safe, address collectionAccount)
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) CreateSafeAndCollectionAccount(proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams, name string, collectionTarget common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreateSafeAndCollectionAccount(&_EnterpriseWallet.TransactOpts, proxyFactory, safeSingleton, safeParams, name, collectionTarget)
}

// CreateSafeAndPaymentAccount is a paid mutator transaction binding the contract method 0xe089c6aa.
//
// Solidity: function createSafeAndPaymentAccount(address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams, string name) returns(address safe, address paymentAccount)
func (_EnterpriseWallet *EnterpriseWalletTransactor) CreateSafeAndPaymentAccount(opts *bind.TransactOpts, proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams, name string) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "createSafeAndPaymentAccount", proxyFactory, safeSingleton, safeParams, name)
}

// CreateSafeAndPaymentAccount is a paid mutator transaction binding the contract method 0xe089c6aa.
//
// Solidity: function createSafeAndPaymentAccount(address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams, string name) returns(address safe, address paymentAccount)
func (_EnterpriseWallet *EnterpriseWalletSession) CreateSafeAndPaymentAccount(proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams, name string) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreateSafeAndPaymentAccount(&_EnterpriseWallet.TransactOpts, proxyFactory, safeSingleton, safeParams, name)
}

// CreateSafeAndPaymentAccount is a paid mutator transaction binding the contract method 0xe089c6aa.
//
// Solidity: function createSafeAndPaymentAccount(address proxyFactory, address safeSingleton, (address[],uint256,address,bytes,address,address,uint256,address,uint256) safeParams, string name) returns(address safe, address paymentAccount)
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) CreateSafeAndPaymentAccount(proxyFactory common.Address, safeSingleton common.Address, safeParams IEnterpriseWalletTypesSafeSetupParams, name string) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.CreateSafeAndPaymentAccount(&_EnterpriseWallet.TransactOpts, proxyFactory, safeSingleton, safeParams, name)
}

// EmergencyFreeze is a paid mutator transaction binding the contract method 0x56e26b63.
//
// Solidity: function emergencyFreeze(address target, bool freeze) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) EmergencyFreeze(opts *bind.TransactOpts, target common.Address, freeze bool) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "emergencyFreeze", target, freeze)
}

// EmergencyFreeze is a paid mutator transaction binding the contract method 0x56e26b63.
//
// Solidity: function emergencyFreeze(address target, bool freeze) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) EmergencyFreeze(target common.Address, freeze bool) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.EmergencyFreeze(&_EnterpriseWallet.TransactOpts, target, freeze)
}

// EmergencyFreeze is a paid mutator transaction binding the contract method 0x56e26b63.
//
// Solidity: function emergencyFreeze(address target, bool freeze) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) EmergencyFreeze(target common.Address, freeze bool) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.EmergencyFreeze(&_EnterpriseWallet.TransactOpts, target, freeze)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0xe8f6940e.
//
// Solidity: function emergencyPause(bool pause) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) EmergencyPause(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "emergencyPause", pause)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0xe8f6940e.
//
// Solidity: function emergencyPause(bool pause) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) EmergencyPause(pause bool) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.EmergencyPause(&_EnterpriseWallet.TransactOpts, pause)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0xe8f6940e.
//
// Solidity: function emergencyPause(bool pause) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) EmergencyPause(pause bool) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.EmergencyPause(&_EnterpriseWallet.TransactOpts, pause)
}

// Initialize is a paid mutator transaction binding the contract method 0xc6a828af.
//
// Solidity: function initialize(bytes4[] methods, (address)[] configs, address superAdmin) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) Initialize(opts *bind.TransactOpts, methods [][4]byte, configs []IEnterpriseWalletTypesMethodConfig, superAdmin common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "initialize", methods, configs, superAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc6a828af.
//
// Solidity: function initialize(bytes4[] methods, (address)[] configs, address superAdmin) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) Initialize(methods [][4]byte, configs []IEnterpriseWalletTypesMethodConfig, superAdmin common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.Initialize(&_EnterpriseWallet.TransactOpts, methods, configs, superAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc6a828af.
//
// Solidity: function initialize(bytes4[] methods, (address)[] configs, address superAdmin) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) Initialize(methods [][4]byte, configs []IEnterpriseWalletTypesMethodConfig, superAdmin common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.Initialize(&_EnterpriseWallet.TransactOpts, methods, configs, superAdmin)
}

// ProposeSuperAdminTransfer is a paid mutator transaction binding the contract method 0x4c64d20e.
//
// Solidity: function proposeSuperAdminTransfer(address newSuperAdmin, uint256 timeout) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) ProposeSuperAdminTransfer(opts *bind.TransactOpts, newSuperAdmin common.Address, timeout *big.Int) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "proposeSuperAdminTransfer", newSuperAdmin, timeout)
}

// ProposeSuperAdminTransfer is a paid mutator transaction binding the contract method 0x4c64d20e.
//
// Solidity: function proposeSuperAdminTransfer(address newSuperAdmin, uint256 timeout) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) ProposeSuperAdminTransfer(newSuperAdmin common.Address, timeout *big.Int) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.ProposeSuperAdminTransfer(&_EnterpriseWallet.TransactOpts, newSuperAdmin, timeout)
}

// ProposeSuperAdminTransfer is a paid mutator transaction binding the contract method 0x4c64d20e.
//
// Solidity: function proposeSuperAdminTransfer(address newSuperAdmin, uint256 timeout) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) ProposeSuperAdminTransfer(newSuperAdmin common.Address, timeout *big.Int) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.ProposeSuperAdminTransfer(&_EnterpriseWallet.TransactOpts, newSuperAdmin, timeout)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.RescueFunds(&_EnterpriseWallet.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.RescueFunds(&_EnterpriseWallet.TransactOpts, token, amount)
}

// SetCollectionTarget is a paid mutator transaction binding the contract method 0x2bd959ee.
//
// Solidity: function setCollectionTarget(address collectionAccount, address target) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) SetCollectionTarget(opts *bind.TransactOpts, collectionAccount common.Address, target common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "setCollectionTarget", collectionAccount, target)
}

// SetCollectionTarget is a paid mutator transaction binding the contract method 0x2bd959ee.
//
// Solidity: function setCollectionTarget(address collectionAccount, address target) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) SetCollectionTarget(collectionAccount common.Address, target common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.SetCollectionTarget(&_EnterpriseWallet.TransactOpts, collectionAccount, target)
}

// SetCollectionTarget is a paid mutator transaction binding the contract method 0x2bd959ee.
//
// Solidity: function setCollectionTarget(address collectionAccount, address target) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) SetCollectionTarget(collectionAccount common.Address, target common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.SetCollectionTarget(&_EnterpriseWallet.TransactOpts, collectionAccount, target)
}

// SetMethodController is a paid mutator transaction binding the contract method 0x5aeee838.
//
// Solidity: function setMethodController(bytes4[] methodSigs, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) SetMethodController(opts *bind.TransactOpts, methodSigs [][4]byte, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "setMethodController", methodSigs, controller)
}

// SetMethodController is a paid mutator transaction binding the contract method 0x5aeee838.
//
// Solidity: function setMethodController(bytes4[] methodSigs, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) SetMethodController(methodSigs [][4]byte, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.SetMethodController(&_EnterpriseWallet.TransactOpts, methodSigs, controller)
}

// SetMethodController is a paid mutator transaction binding the contract method 0x5aeee838.
//
// Solidity: function setMethodController(bytes4[] methodSigs, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) SetMethodController(methodSigs [][4]byte, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.SetMethodController(&_EnterpriseWallet.TransactOpts, methodSigs, controller)
}

// UpdateMethodController is a paid mutator transaction binding the contract method 0x4358ad24.
//
// Solidity: function updateMethodController(bytes4 methodSig, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) UpdateMethodController(opts *bind.TransactOpts, methodSig [4]byte, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "updateMethodController", methodSig, controller)
}

// UpdateMethodController is a paid mutator transaction binding the contract method 0x4358ad24.
//
// Solidity: function updateMethodController(bytes4 methodSig, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) UpdateMethodController(methodSig [4]byte, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.UpdateMethodController(&_EnterpriseWallet.TransactOpts, methodSig, controller)
}

// UpdateMethodController is a paid mutator transaction binding the contract method 0x4358ad24.
//
// Solidity: function updateMethodController(bytes4 methodSig, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) UpdateMethodController(methodSig [4]byte, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.UpdateMethodController(&_EnterpriseWallet.TransactOpts, methodSig, controller)
}

// UpdateMethodControllers is a paid mutator transaction binding the contract method 0x51ab59a2.
//
// Solidity: function updateMethodControllers(bytes4[] methodSigs, address[] controllers) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) UpdateMethodControllers(opts *bind.TransactOpts, methodSigs [][4]byte, controllers []common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "updateMethodControllers", methodSigs, controllers)
}

// UpdateMethodControllers is a paid mutator transaction binding the contract method 0x51ab59a2.
//
// Solidity: function updateMethodControllers(bytes4[] methodSigs, address[] controllers) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) UpdateMethodControllers(methodSigs [][4]byte, controllers []common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.UpdateMethodControllers(&_EnterpriseWallet.TransactOpts, methodSigs, controllers)
}

// UpdateMethodControllers is a paid mutator transaction binding the contract method 0x51ab59a2.
//
// Solidity: function updateMethodControllers(bytes4[] methodSigs, address[] controllers) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) UpdateMethodControllers(methodSigs [][4]byte, controllers []common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.UpdateMethodControllers(&_EnterpriseWallet.TransactOpts, methodSigs, controllers)
}

// UpdatePaymentAccountController is a paid mutator transaction binding the contract method 0xf0f3fb5b.
//
// Solidity: function updatePaymentAccountController(address paymentAccount, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) UpdatePaymentAccountController(opts *bind.TransactOpts, paymentAccount common.Address, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.Transact(opts, "updatePaymentAccountController", paymentAccount, controller)
}

// UpdatePaymentAccountController is a paid mutator transaction binding the contract method 0xf0f3fb5b.
//
// Solidity: function updatePaymentAccountController(address paymentAccount, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletSession) UpdatePaymentAccountController(paymentAccount common.Address, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.UpdatePaymentAccountController(&_EnterpriseWallet.TransactOpts, paymentAccount, controller)
}

// UpdatePaymentAccountController is a paid mutator transaction binding the contract method 0xf0f3fb5b.
//
// Solidity: function updatePaymentAccountController(address paymentAccount, address controller) returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) UpdatePaymentAccountController(paymentAccount common.Address, controller common.Address) (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.UpdatePaymentAccountController(&_EnterpriseWallet.TransactOpts, paymentAccount, controller)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EnterpriseWallet *EnterpriseWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EnterpriseWallet *EnterpriseWalletSession) Receive() (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.Receive(&_EnterpriseWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EnterpriseWallet *EnterpriseWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _EnterpriseWallet.Contract.Receive(&_EnterpriseWallet.TransactOpts)
}

// EnterpriseWalletCollectionAccountCreatedIterator is returned from FilterCollectionAccountCreated and is used to iterate over the raw logs and unpacked data for CollectionAccountCreated events raised by the EnterpriseWallet contract.
type EnterpriseWalletCollectionAccountCreatedIterator struct {
	Event *EnterpriseWalletCollectionAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletCollectionAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletCollectionAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletCollectionAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletCollectionAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletCollectionAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletCollectionAccountCreated represents a CollectionAccountCreated event raised by the EnterpriseWallet contract.
type EnterpriseWalletCollectionAccountCreated struct {
	Account common.Address
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCollectionAccountCreated is a free log retrieval operation binding the contract event 0x6550fac320580d8d6c59493378f2d1dd10b7f7e0bec627d067be994280dea31e.
//
// Solidity: event CollectionAccountCreated(address indexed account, address indexed creator)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterCollectionAccountCreated(opts *bind.FilterOpts, account []common.Address, creator []common.Address) (*EnterpriseWalletCollectionAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "CollectionAccountCreated", accountRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletCollectionAccountCreatedIterator{contract: _EnterpriseWallet.contract, event: "CollectionAccountCreated", logs: logs, sub: sub}, nil
}

// WatchCollectionAccountCreated is a free log subscription operation binding the contract event 0x6550fac320580d8d6c59493378f2d1dd10b7f7e0bec627d067be994280dea31e.
//
// Solidity: event CollectionAccountCreated(address indexed account, address indexed creator)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchCollectionAccountCreated(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletCollectionAccountCreated, account []common.Address, creator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "CollectionAccountCreated", accountRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletCollectionAccountCreated)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "CollectionAccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectionAccountCreated is a log parse operation binding the contract event 0x6550fac320580d8d6c59493378f2d1dd10b7f7e0bec627d067be994280dea31e.
//
// Solidity: event CollectionAccountCreated(address indexed account, address indexed creator)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseCollectionAccountCreated(log types.Log) (*EnterpriseWalletCollectionAccountCreated, error) {
	event := new(EnterpriseWalletCollectionAccountCreated)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "CollectionAccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletEmergencyFreezeIterator is returned from FilterEmergencyFreeze and is used to iterate over the raw logs and unpacked data for EmergencyFreeze events raised by the EnterpriseWallet contract.
type EnterpriseWalletEmergencyFreezeIterator struct {
	Event *EnterpriseWalletEmergencyFreeze // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletEmergencyFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletEmergencyFreeze)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletEmergencyFreeze)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletEmergencyFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletEmergencyFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletEmergencyFreeze represents a EmergencyFreeze event raised by the EnterpriseWallet contract.
type EnterpriseWalletEmergencyFreeze struct {
	Target common.Address
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyFreeze is a free log retrieval operation binding the contract event 0xeb91f931ecd76eb4e66768d05e7d6aa0b8754bc7084f417566027b6e19a3919a.
//
// Solidity: event EmergencyFreeze(address indexed target, bool frozen)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterEmergencyFreeze(opts *bind.FilterOpts, target []common.Address) (*EnterpriseWalletEmergencyFreezeIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "EmergencyFreeze", targetRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletEmergencyFreezeIterator{contract: _EnterpriseWallet.contract, event: "EmergencyFreeze", logs: logs, sub: sub}, nil
}

// WatchEmergencyFreeze is a free log subscription operation binding the contract event 0xeb91f931ecd76eb4e66768d05e7d6aa0b8754bc7084f417566027b6e19a3919a.
//
// Solidity: event EmergencyFreeze(address indexed target, bool frozen)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchEmergencyFreeze(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletEmergencyFreeze, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "EmergencyFreeze", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletEmergencyFreeze)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "EmergencyFreeze", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyFreeze is a log parse operation binding the contract event 0xeb91f931ecd76eb4e66768d05e7d6aa0b8754bc7084f417566027b6e19a3919a.
//
// Solidity: event EmergencyFreeze(address indexed target, bool frozen)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseEmergencyFreeze(log types.Log) (*EnterpriseWalletEmergencyFreeze, error) {
	event := new(EnterpriseWalletEmergencyFreeze)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "EmergencyFreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletEmergencyPauseIterator is returned from FilterEmergencyPause and is used to iterate over the raw logs and unpacked data for EmergencyPause events raised by the EnterpriseWallet contract.
type EnterpriseWalletEmergencyPauseIterator struct {
	Event *EnterpriseWalletEmergencyPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletEmergencyPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletEmergencyPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletEmergencyPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletEmergencyPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletEmergencyPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletEmergencyPause represents a EmergencyPause event raised by the EnterpriseWallet contract.
type EnterpriseWalletEmergencyPause struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyPause is a free log retrieval operation binding the contract event 0xb80d1ae5628c2af91424c54c87acfd15016be4bebeef0425b328238702f34831.
//
// Solidity: event EmergencyPause(bool paused)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterEmergencyPause(opts *bind.FilterOpts) (*EnterpriseWalletEmergencyPauseIterator, error) {

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "EmergencyPause")
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletEmergencyPauseIterator{contract: _EnterpriseWallet.contract, event: "EmergencyPause", logs: logs, sub: sub}, nil
}

// WatchEmergencyPause is a free log subscription operation binding the contract event 0xb80d1ae5628c2af91424c54c87acfd15016be4bebeef0425b328238702f34831.
//
// Solidity: event EmergencyPause(bool paused)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchEmergencyPause(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletEmergencyPause) (event.Subscription, error) {

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "EmergencyPause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletEmergencyPause)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "EmergencyPause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyPause is a log parse operation binding the contract event 0xb80d1ae5628c2af91424c54c87acfd15016be4bebeef0425b328238702f34831.
//
// Solidity: event EmergencyPause(bool paused)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseEmergencyPause(log types.Log) (*EnterpriseWalletEmergencyPause, error) {
	event := new(EnterpriseWalletEmergencyPause)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "EmergencyPause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletFundsCollectedIterator is returned from FilterFundsCollected and is used to iterate over the raw logs and unpacked data for FundsCollected events raised by the EnterpriseWallet contract.
type EnterpriseWalletFundsCollectedIterator struct {
	Event *EnterpriseWalletFundsCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletFundsCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletFundsCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletFundsCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletFundsCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletFundsCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletFundsCollected represents a FundsCollected event raised by the EnterpriseWallet contract.
type EnterpriseWalletFundsCollected struct {
	From   common.Address
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsCollected is a free log retrieval operation binding the contract event 0x6db57dba961bb022856264c685785a66657081b9612505524d881c2abc8d375e.
//
// Solidity: event FundsCollected(address indexed from, address indexed to, address indexed token, uint256 amount)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterFundsCollected(opts *bind.FilterOpts, from []common.Address, to []common.Address, token []common.Address) (*EnterpriseWalletFundsCollectedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "FundsCollected", fromRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFundsCollectedIterator{contract: _EnterpriseWallet.contract, event: "FundsCollected", logs: logs, sub: sub}, nil
}

// WatchFundsCollected is a free log subscription operation binding the contract event 0x6db57dba961bb022856264c685785a66657081b9612505524d881c2abc8d375e.
//
// Solidity: event FundsCollected(address indexed from, address indexed to, address indexed token, uint256 amount)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchFundsCollected(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletFundsCollected, from []common.Address, to []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "FundsCollected", fromRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletFundsCollected)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "FundsCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsCollected is a log parse operation binding the contract event 0x6db57dba961bb022856264c685785a66657081b9612505524d881c2abc8d375e.
//
// Solidity: event FundsCollected(address indexed from, address indexed to, address indexed token, uint256 amount)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseFundsCollected(log types.Log) (*EnterpriseWalletFundsCollected, error) {
	event := new(EnterpriseWalletFundsCollected)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "FundsCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletFundsRescuedIterator is returned from FilterFundsRescued and is used to iterate over the raw logs and unpacked data for FundsRescued events raised by the EnterpriseWallet contract.
type EnterpriseWalletFundsRescuedIterator struct {
	Event *EnterpriseWalletFundsRescued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletFundsRescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletFundsRescued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletFundsRescued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletFundsRescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletFundsRescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletFundsRescued represents a FundsRescued event raised by the EnterpriseWallet contract.
type EnterpriseWalletFundsRescued struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsRescued is a free log retrieval operation binding the contract event 0xed2837b80b3489773c5f1ed30dd2884b4c90dbf7b87428ea03c49b24ef59a805.
//
// Solidity: event FundsRescued(address indexed to, address indexed token, uint256 amount)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterFundsRescued(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*EnterpriseWalletFundsRescuedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "FundsRescued", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFundsRescuedIterator{contract: _EnterpriseWallet.contract, event: "FundsRescued", logs: logs, sub: sub}, nil
}

// WatchFundsRescued is a free log subscription operation binding the contract event 0xed2837b80b3489773c5f1ed30dd2884b4c90dbf7b87428ea03c49b24ef59a805.
//
// Solidity: event FundsRescued(address indexed to, address indexed token, uint256 amount)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchFundsRescued(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletFundsRescued, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "FundsRescued", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletFundsRescued)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "FundsRescued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsRescued is a log parse operation binding the contract event 0xed2837b80b3489773c5f1ed30dd2884b4c90dbf7b87428ea03c49b24ef59a805.
//
// Solidity: event FundsRescued(address indexed to, address indexed token, uint256 amount)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseFundsRescued(log types.Log) (*EnterpriseWalletFundsRescued, error) {
	event := new(EnterpriseWalletFundsRescued)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "FundsRescued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EnterpriseWallet contract.
type EnterpriseWalletInitializedIterator struct {
	Event *EnterpriseWalletInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletInitialized represents a Initialized event raised by the EnterpriseWallet contract.
type EnterpriseWalletInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterInitialized(opts *bind.FilterOpts) (*EnterpriseWalletInitializedIterator, error) {

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletInitializedIterator{contract: _EnterpriseWallet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletInitialized) (event.Subscription, error) {

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletInitialized)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseInitialized(log types.Log) (*EnterpriseWalletInitialized, error) {
	event := new(EnterpriseWalletInitialized)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletMethodControllerUpdatedIterator is returned from FilterMethodControllerUpdated and is used to iterate over the raw logs and unpacked data for MethodControllerUpdated events raised by the EnterpriseWallet contract.
type EnterpriseWalletMethodControllerUpdatedIterator struct {
	Event *EnterpriseWalletMethodControllerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletMethodControllerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletMethodControllerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletMethodControllerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletMethodControllerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletMethodControllerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletMethodControllerUpdated represents a MethodControllerUpdated event raised by the EnterpriseWallet contract.
type EnterpriseWalletMethodControllerUpdated struct {
	MethodSig  [4]byte
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMethodControllerUpdated is a free log retrieval operation binding the contract event 0x9b1635e988a263ce51376fde006cfa712ff764a45652ec4c2d42e534d6514091.
//
// Solidity: event MethodControllerUpdated(bytes4 indexed methodSig, address indexed controller)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterMethodControllerUpdated(opts *bind.FilterOpts, methodSig [][4]byte, controller []common.Address) (*EnterpriseWalletMethodControllerUpdatedIterator, error) {

	var methodSigRule []interface{}
	for _, methodSigItem := range methodSig {
		methodSigRule = append(methodSigRule, methodSigItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "MethodControllerUpdated", methodSigRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletMethodControllerUpdatedIterator{contract: _EnterpriseWallet.contract, event: "MethodControllerUpdated", logs: logs, sub: sub}, nil
}

// WatchMethodControllerUpdated is a free log subscription operation binding the contract event 0x9b1635e988a263ce51376fde006cfa712ff764a45652ec4c2d42e534d6514091.
//
// Solidity: event MethodControllerUpdated(bytes4 indexed methodSig, address indexed controller)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchMethodControllerUpdated(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletMethodControllerUpdated, methodSig [][4]byte, controller []common.Address) (event.Subscription, error) {

	var methodSigRule []interface{}
	for _, methodSigItem := range methodSig {
		methodSigRule = append(methodSigRule, methodSigItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "MethodControllerUpdated", methodSigRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletMethodControllerUpdated)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "MethodControllerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMethodControllerUpdated is a log parse operation binding the contract event 0x9b1635e988a263ce51376fde006cfa712ff764a45652ec4c2d42e534d6514091.
//
// Solidity: event MethodControllerUpdated(bytes4 indexed methodSig, address indexed controller)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseMethodControllerUpdated(log types.Log) (*EnterpriseWalletMethodControllerUpdated, error) {
	event := new(EnterpriseWalletMethodControllerUpdated)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "MethodControllerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletPaymentAccountCreatedIterator is returned from FilterPaymentAccountCreated and is used to iterate over the raw logs and unpacked data for PaymentAccountCreated events raised by the EnterpriseWallet contract.
type EnterpriseWalletPaymentAccountCreatedIterator struct {
	Event *EnterpriseWalletPaymentAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletPaymentAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletPaymentAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletPaymentAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletPaymentAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletPaymentAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletPaymentAccountCreated represents a PaymentAccountCreated event raised by the EnterpriseWallet contract.
type EnterpriseWalletPaymentAccountCreated struct {
	Account    common.Address
	Creator    common.Address
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaymentAccountCreated is a free log retrieval operation binding the contract event 0x3f2d64265a1d8f3b6d5a211bc2b2b7f6e43f4757acfd574154902e01dd71f912.
//
// Solidity: event PaymentAccountCreated(address indexed account, address indexed creator, address indexed controller)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterPaymentAccountCreated(opts *bind.FilterOpts, account []common.Address, creator []common.Address, controller []common.Address) (*EnterpriseWalletPaymentAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "PaymentAccountCreated", accountRule, creatorRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletPaymentAccountCreatedIterator{contract: _EnterpriseWallet.contract, event: "PaymentAccountCreated", logs: logs, sub: sub}, nil
}

// WatchPaymentAccountCreated is a free log subscription operation binding the contract event 0x3f2d64265a1d8f3b6d5a211bc2b2b7f6e43f4757acfd574154902e01dd71f912.
//
// Solidity: event PaymentAccountCreated(address indexed account, address indexed creator, address indexed controller)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchPaymentAccountCreated(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletPaymentAccountCreated, account []common.Address, creator []common.Address, controller []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "PaymentAccountCreated", accountRule, creatorRule, controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletPaymentAccountCreated)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "PaymentAccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentAccountCreated is a log parse operation binding the contract event 0x3f2d64265a1d8f3b6d5a211bc2b2b7f6e43f4757acfd574154902e01dd71f912.
//
// Solidity: event PaymentAccountCreated(address indexed account, address indexed creator, address indexed controller)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParsePaymentAccountCreated(log types.Log) (*EnterpriseWalletPaymentAccountCreated, error) {
	event := new(EnterpriseWalletPaymentAccountCreated)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "PaymentAccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletSafeAndCollectionAccountCreatedIterator is returned from FilterSafeAndCollectionAccountCreated and is used to iterate over the raw logs and unpacked data for SafeAndCollectionAccountCreated events raised by the EnterpriseWallet contract.
type EnterpriseWalletSafeAndCollectionAccountCreatedIterator struct {
	Event *EnterpriseWalletSafeAndCollectionAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletSafeAndCollectionAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletSafeAndCollectionAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletSafeAndCollectionAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletSafeAndCollectionAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletSafeAndCollectionAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletSafeAndCollectionAccountCreated represents a SafeAndCollectionAccountCreated event raised by the EnterpriseWallet contract.
type EnterpriseWalletSafeAndCollectionAccountCreated struct {
	Safe              common.Address
	CollectionAccount common.Address
	Name              string
	Target            common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSafeAndCollectionAccountCreated is a free log retrieval operation binding the contract event 0x86c8aea7aab7fc4ededf3146e3e4bc4136ef1a6df1fe018c530ec1d87f0ee285.
//
// Solidity: event SafeAndCollectionAccountCreated(address indexed safe, address indexed collectionAccount, string name, address target)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterSafeAndCollectionAccountCreated(opts *bind.FilterOpts, safe []common.Address, collectionAccount []common.Address) (*EnterpriseWalletSafeAndCollectionAccountCreatedIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}
	var collectionAccountRule []interface{}
	for _, collectionAccountItem := range collectionAccount {
		collectionAccountRule = append(collectionAccountRule, collectionAccountItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "SafeAndCollectionAccountCreated", safeRule, collectionAccountRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletSafeAndCollectionAccountCreatedIterator{contract: _EnterpriseWallet.contract, event: "SafeAndCollectionAccountCreated", logs: logs, sub: sub}, nil
}

// WatchSafeAndCollectionAccountCreated is a free log subscription operation binding the contract event 0x86c8aea7aab7fc4ededf3146e3e4bc4136ef1a6df1fe018c530ec1d87f0ee285.
//
// Solidity: event SafeAndCollectionAccountCreated(address indexed safe, address indexed collectionAccount, string name, address target)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchSafeAndCollectionAccountCreated(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletSafeAndCollectionAccountCreated, safe []common.Address, collectionAccount []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}
	var collectionAccountRule []interface{}
	for _, collectionAccountItem := range collectionAccount {
		collectionAccountRule = append(collectionAccountRule, collectionAccountItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "SafeAndCollectionAccountCreated", safeRule, collectionAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletSafeAndCollectionAccountCreated)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "SafeAndCollectionAccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSafeAndCollectionAccountCreated is a log parse operation binding the contract event 0x86c8aea7aab7fc4ededf3146e3e4bc4136ef1a6df1fe018c530ec1d87f0ee285.
//
// Solidity: event SafeAndCollectionAccountCreated(address indexed safe, address indexed collectionAccount, string name, address target)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseSafeAndCollectionAccountCreated(log types.Log) (*EnterpriseWalletSafeAndCollectionAccountCreated, error) {
	event := new(EnterpriseWalletSafeAndCollectionAccountCreated)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "SafeAndCollectionAccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletSafeAndPaymentAccountCreatedIterator is returned from FilterSafeAndPaymentAccountCreated and is used to iterate over the raw logs and unpacked data for SafeAndPaymentAccountCreated events raised by the EnterpriseWallet contract.
type EnterpriseWalletSafeAndPaymentAccountCreatedIterator struct {
	Event *EnterpriseWalletSafeAndPaymentAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletSafeAndPaymentAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletSafeAndPaymentAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletSafeAndPaymentAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletSafeAndPaymentAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletSafeAndPaymentAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletSafeAndPaymentAccountCreated represents a SafeAndPaymentAccountCreated event raised by the EnterpriseWallet contract.
type EnterpriseWalletSafeAndPaymentAccountCreated struct {
	Safe           common.Address
	PaymentAccount common.Address
	Name           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSafeAndPaymentAccountCreated is a free log retrieval operation binding the contract event 0x460aff70a8d100b1173216345ee13e8a5fbcc0bdfae2737268c5db6791a37828.
//
// Solidity: event SafeAndPaymentAccountCreated(address indexed safe, address indexed paymentAccount, string name)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterSafeAndPaymentAccountCreated(opts *bind.FilterOpts, safe []common.Address, paymentAccount []common.Address) (*EnterpriseWalletSafeAndPaymentAccountCreatedIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}
	var paymentAccountRule []interface{}
	for _, paymentAccountItem := range paymentAccount {
		paymentAccountRule = append(paymentAccountRule, paymentAccountItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "SafeAndPaymentAccountCreated", safeRule, paymentAccountRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletSafeAndPaymentAccountCreatedIterator{contract: _EnterpriseWallet.contract, event: "SafeAndPaymentAccountCreated", logs: logs, sub: sub}, nil
}

// WatchSafeAndPaymentAccountCreated is a free log subscription operation binding the contract event 0x460aff70a8d100b1173216345ee13e8a5fbcc0bdfae2737268c5db6791a37828.
//
// Solidity: event SafeAndPaymentAccountCreated(address indexed safe, address indexed paymentAccount, string name)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchSafeAndPaymentAccountCreated(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletSafeAndPaymentAccountCreated, safe []common.Address, paymentAccount []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}
	var paymentAccountRule []interface{}
	for _, paymentAccountItem := range paymentAccount {
		paymentAccountRule = append(paymentAccountRule, paymentAccountItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "SafeAndPaymentAccountCreated", safeRule, paymentAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletSafeAndPaymentAccountCreated)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "SafeAndPaymentAccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSafeAndPaymentAccountCreated is a log parse operation binding the contract event 0x460aff70a8d100b1173216345ee13e8a5fbcc0bdfae2737268c5db6791a37828.
//
// Solidity: event SafeAndPaymentAccountCreated(address indexed safe, address indexed paymentAccount, string name)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseSafeAndPaymentAccountCreated(log types.Log) (*EnterpriseWalletSafeAndPaymentAccountCreated, error) {
	event := new(EnterpriseWalletSafeAndPaymentAccountCreated)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "SafeAndPaymentAccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletSuperAdminTransferCancelledIterator is returned from FilterSuperAdminTransferCancelled and is used to iterate over the raw logs and unpacked data for SuperAdminTransferCancelled events raised by the EnterpriseWallet contract.
type EnterpriseWalletSuperAdminTransferCancelledIterator struct {
	Event *EnterpriseWalletSuperAdminTransferCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletSuperAdminTransferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletSuperAdminTransferCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletSuperAdminTransferCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletSuperAdminTransferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletSuperAdminTransferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletSuperAdminTransferCancelled represents a SuperAdminTransferCancelled event raised by the EnterpriseWallet contract.
type EnterpriseWalletSuperAdminTransferCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSuperAdminTransferCancelled is a free log retrieval operation binding the contract event 0xbc84c2603adde04ce6300533f34797b9a57392341700d79615f8647d6942065c.
//
// Solidity: event SuperAdminTransferCancelled()
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterSuperAdminTransferCancelled(opts *bind.FilterOpts) (*EnterpriseWalletSuperAdminTransferCancelledIterator, error) {

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "SuperAdminTransferCancelled")
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletSuperAdminTransferCancelledIterator{contract: _EnterpriseWallet.contract, event: "SuperAdminTransferCancelled", logs: logs, sub: sub}, nil
}

// WatchSuperAdminTransferCancelled is a free log subscription operation binding the contract event 0xbc84c2603adde04ce6300533f34797b9a57392341700d79615f8647d6942065c.
//
// Solidity: event SuperAdminTransferCancelled()
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchSuperAdminTransferCancelled(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletSuperAdminTransferCancelled) (event.Subscription, error) {

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "SuperAdminTransferCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletSuperAdminTransferCancelled)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "SuperAdminTransferCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSuperAdminTransferCancelled is a log parse operation binding the contract event 0xbc84c2603adde04ce6300533f34797b9a57392341700d79615f8647d6942065c.
//
// Solidity: event SuperAdminTransferCancelled()
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseSuperAdminTransferCancelled(log types.Log) (*EnterpriseWalletSuperAdminTransferCancelled, error) {
	event := new(EnterpriseWalletSuperAdminTransferCancelled)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "SuperAdminTransferCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletSuperAdminTransferProposedIterator is returned from FilterSuperAdminTransferProposed and is used to iterate over the raw logs and unpacked data for SuperAdminTransferProposed events raised by the EnterpriseWallet contract.
type EnterpriseWalletSuperAdminTransferProposedIterator struct {
	Event *EnterpriseWalletSuperAdminTransferProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletSuperAdminTransferProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletSuperAdminTransferProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletSuperAdminTransferProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletSuperAdminTransferProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletSuperAdminTransferProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletSuperAdminTransferProposed represents a SuperAdminTransferProposed event raised by the EnterpriseWallet contract.
type EnterpriseWalletSuperAdminTransferProposed struct {
	CurrentSuperAdmin  common.Address
	ProposedSuperAdmin common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSuperAdminTransferProposed is a free log retrieval operation binding the contract event 0xef5ca9734ccba0eda923940b1b8c0313ac32f516a2ca06c28fd77c5d01e78208.
//
// Solidity: event SuperAdminTransferProposed(address indexed currentSuperAdmin, address indexed proposedSuperAdmin)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterSuperAdminTransferProposed(opts *bind.FilterOpts, currentSuperAdmin []common.Address, proposedSuperAdmin []common.Address) (*EnterpriseWalletSuperAdminTransferProposedIterator, error) {

	var currentSuperAdminRule []interface{}
	for _, currentSuperAdminItem := range currentSuperAdmin {
		currentSuperAdminRule = append(currentSuperAdminRule, currentSuperAdminItem)
	}
	var proposedSuperAdminRule []interface{}
	for _, proposedSuperAdminItem := range proposedSuperAdmin {
		proposedSuperAdminRule = append(proposedSuperAdminRule, proposedSuperAdminItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "SuperAdminTransferProposed", currentSuperAdminRule, proposedSuperAdminRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletSuperAdminTransferProposedIterator{contract: _EnterpriseWallet.contract, event: "SuperAdminTransferProposed", logs: logs, sub: sub}, nil
}

// WatchSuperAdminTransferProposed is a free log subscription operation binding the contract event 0xef5ca9734ccba0eda923940b1b8c0313ac32f516a2ca06c28fd77c5d01e78208.
//
// Solidity: event SuperAdminTransferProposed(address indexed currentSuperAdmin, address indexed proposedSuperAdmin)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchSuperAdminTransferProposed(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletSuperAdminTransferProposed, currentSuperAdmin []common.Address, proposedSuperAdmin []common.Address) (event.Subscription, error) {

	var currentSuperAdminRule []interface{}
	for _, currentSuperAdminItem := range currentSuperAdmin {
		currentSuperAdminRule = append(currentSuperAdminRule, currentSuperAdminItem)
	}
	var proposedSuperAdminRule []interface{}
	for _, proposedSuperAdminItem := range proposedSuperAdmin {
		proposedSuperAdminRule = append(proposedSuperAdminRule, proposedSuperAdminItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "SuperAdminTransferProposed", currentSuperAdminRule, proposedSuperAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletSuperAdminTransferProposed)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "SuperAdminTransferProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSuperAdminTransferProposed is a log parse operation binding the contract event 0xef5ca9734ccba0eda923940b1b8c0313ac32f516a2ca06c28fd77c5d01e78208.
//
// Solidity: event SuperAdminTransferProposed(address indexed currentSuperAdmin, address indexed proposedSuperAdmin)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseSuperAdminTransferProposed(log types.Log) (*EnterpriseWalletSuperAdminTransferProposed, error) {
	event := new(EnterpriseWalletSuperAdminTransferProposed)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "SuperAdminTransferProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletSuperAdminTransferredIterator is returned from FilterSuperAdminTransferred and is used to iterate over the raw logs and unpacked data for SuperAdminTransferred events raised by the EnterpriseWallet contract.
type EnterpriseWalletSuperAdminTransferredIterator struct {
	Event *EnterpriseWalletSuperAdminTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnterpriseWalletSuperAdminTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletSuperAdminTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnterpriseWalletSuperAdminTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnterpriseWalletSuperAdminTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletSuperAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletSuperAdminTransferred represents a SuperAdminTransferred event raised by the EnterpriseWallet contract.
type EnterpriseWalletSuperAdminTransferred struct {
	OldSuperAdmin common.Address
	NewSuperAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSuperAdminTransferred is a free log retrieval operation binding the contract event 0x0f62530a074f4e1e883a8c916fa7f8639d52598edb7f9b5aa3148d991db5610d.
//
// Solidity: event SuperAdminTransferred(address indexed oldSuperAdmin, address indexed newSuperAdmin)
func (_EnterpriseWallet *EnterpriseWalletFilterer) FilterSuperAdminTransferred(opts *bind.FilterOpts, oldSuperAdmin []common.Address, newSuperAdmin []common.Address) (*EnterpriseWalletSuperAdminTransferredIterator, error) {

	var oldSuperAdminRule []interface{}
	for _, oldSuperAdminItem := range oldSuperAdmin {
		oldSuperAdminRule = append(oldSuperAdminRule, oldSuperAdminItem)
	}
	var newSuperAdminRule []interface{}
	for _, newSuperAdminItem := range newSuperAdmin {
		newSuperAdminRule = append(newSuperAdminRule, newSuperAdminItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.FilterLogs(opts, "SuperAdminTransferred", oldSuperAdminRule, newSuperAdminRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletSuperAdminTransferredIterator{contract: _EnterpriseWallet.contract, event: "SuperAdminTransferred", logs: logs, sub: sub}, nil
}

// WatchSuperAdminTransferred is a free log subscription operation binding the contract event 0x0f62530a074f4e1e883a8c916fa7f8639d52598edb7f9b5aa3148d991db5610d.
//
// Solidity: event SuperAdminTransferred(address indexed oldSuperAdmin, address indexed newSuperAdmin)
func (_EnterpriseWallet *EnterpriseWalletFilterer) WatchSuperAdminTransferred(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletSuperAdminTransferred, oldSuperAdmin []common.Address, newSuperAdmin []common.Address) (event.Subscription, error) {

	var oldSuperAdminRule []interface{}
	for _, oldSuperAdminItem := range oldSuperAdmin {
		oldSuperAdminRule = append(oldSuperAdminRule, oldSuperAdminItem)
	}
	var newSuperAdminRule []interface{}
	for _, newSuperAdminItem := range newSuperAdmin {
		newSuperAdminRule = append(newSuperAdminRule, newSuperAdminItem)
	}

	logs, sub, err := _EnterpriseWallet.contract.WatchLogs(opts, "SuperAdminTransferred", oldSuperAdminRule, newSuperAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletSuperAdminTransferred)
				if err := _EnterpriseWallet.contract.UnpackLog(event, "SuperAdminTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSuperAdminTransferred is a log parse operation binding the contract event 0x0f62530a074f4e1e883a8c916fa7f8639d52598edb7f9b5aa3148d991db5610d.
//
// Solidity: event SuperAdminTransferred(address indexed oldSuperAdmin, address indexed newSuperAdmin)
func (_EnterpriseWallet *EnterpriseWalletFilterer) ParseSuperAdminTransferred(log types.Log) (*EnterpriseWalletSuperAdminTransferred, error) {
	event := new(EnterpriseWalletSuperAdminTransferred)
	if err := _EnterpriseWallet.contract.UnpackLog(event, "SuperAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
