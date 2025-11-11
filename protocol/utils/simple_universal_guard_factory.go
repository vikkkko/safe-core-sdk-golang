package utils

import (
	"fmt"
	"strings"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	projectabi "github.com/vikkkko/safe-core-sdk-golang/abi"
)

// SimpleUniversalGuardFactoryABI is the ABI string for SimpleUniversalGuardFactory contract
var SimpleUniversalGuardFactoryABI = string(projectabi.SimpleUniversalGuardFactory)

// SimpleUniversalGuardFactoryContract provides typed helpers for the SimpleUniversalGuardFactory contract
type SimpleUniversalGuardFactoryContract struct {
	address  common.Address
	abi      ethabi.ABI
	contract *bind.BoundContract
}

// NewSimpleUniversalGuardFactoryContract creates a SimpleUniversalGuardFactory helper
func NewSimpleUniversalGuardFactoryContract(address common.Address, backend bind.ContractBackend) (*SimpleUniversalGuardFactoryContract, error) {
	// Parse ABI from embedded abi package
	parsed, err := ethabi.JSON(strings.NewReader(SimpleUniversalGuardFactoryABI))
	if err != nil {
		return nil, fmt.Errorf("parse SimpleUniversalGuardFactory ABI: %w", err)
	}

	bound := bind.NewBoundContract(address, parsed, backend, backend, backend)

	return &SimpleUniversalGuardFactoryContract{
		address:  address,
		abi:      parsed,
		contract: bound,
	}, nil
}

// GetImplementation returns the implementation contract address
func (s *SimpleUniversalGuardFactoryContract) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	if err := s.contract.Call(opts, &out, "getImplementation"); err != nil {
		return common.Address{}, err
	}
	return *ethabi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// GetGuard returns the guard address for a required signer
func (s *SimpleUniversalGuardFactoryContract) GetGuard(opts *bind.CallOpts, requiredSigner common.Address) (common.Address, error) {
	var out []interface{}
	if err := s.contract.Call(opts, &out, "getGuard", requiredSigner); err != nil {
		return common.Address{}, err
	}
	return *ethabi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// PredictGuardAddress predicts the guard address for a required signer
func (s *SimpleUniversalGuardFactoryContract) PredictGuardAddress(opts *bind.CallOpts, requiredSigner common.Address) (common.Address, error) {
	var out []interface{}
	if err := s.contract.Call(opts, &out, "predictGuardAddress", requiredSigner); err != nil {
		return common.Address{}, err
	}
	return *ethabi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// CreateGuard creates a new guard for the required signer
func (s *SimpleUniversalGuardFactoryContract) CreateGuard(auth *bind.TransactOpts, requiredSigner common.Address) (*gethtypes.Transaction, error) {
	return s.contract.Transact(auth, "createGuard", requiredSigner)
}
