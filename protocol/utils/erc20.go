package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// ERC20ABI contains the standard ERC20 token ABI
const ERC20ABI = `[
	{
		"constant": true,
		"inputs": [],
		"name": "name",
		"outputs": [{"name": "", "type": "string"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "symbol",
		"outputs": [{"name": "", "type": "string"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "decimals",
		"outputs": [{"name": "", "type": "uint8"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "totalSupply",
		"outputs": [{"name": "", "type": "uint256"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [{"name": "_owner", "type": "address"}],
		"name": "balanceOf",
		"outputs": [{"name": "balance", "type": "uint256"}],
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{"name": "_to", "type": "address"},
			{"name": "_value", "type": "uint256"}
		],
		"name": "transfer",
		"outputs": [{"name": "", "type": "bool"}],
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{"name": "_from", "type": "address"},
			{"name": "_to", "type": "address"},
			{"name": "_value", "type": "uint256"}
		],
		"name": "transferFrom",
		"outputs": [{"name": "", "type": "bool"}],
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{"name": "_spender", "type": "address"},
			{"name": "_value", "type": "uint256"}
		],
		"name": "approve",
		"outputs": [{"name": "", "type": "bool"}],
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{"name": "_owner", "type": "address"},
			{"name": "_spender", "type": "address"}
		],
		"name": "allowance",
		"outputs": [{"name": "remaining", "type": "uint256"}],
		"type": "function"
	}
]`

// CreateERC20TransferData creates the call data for ERC20 transfer function using ABI encoding
func CreateERC20TransferData(to string, amount *big.Int) ([]byte, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Encode the transfer function call
	toAddr := common.HexToAddress(to)
	data, err := parsedABI.Pack("transfer", toAddr, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to encode transfer call: %w", err)
	}

	return data, nil
}

// CreateERC20TransferFromData creates the call data for ERC20 transferFrom function using ABI encoding
func CreateERC20TransferFromData(from, to string, amount *big.Int) ([]byte, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Encode the transferFrom function call
	fromAddr := common.HexToAddress(from)
	toAddr := common.HexToAddress(to)
	data, err := parsedABI.Pack("transferFrom", fromAddr, toAddr, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to encode transferFrom call: %w", err)
	}

	return data, nil
}

// CreateERC20ApproveData creates the call data for ERC20 approve function using ABI encoding
func CreateERC20ApproveData(spender string, amount *big.Int) ([]byte, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Encode the approve function call
	spenderAddr := common.HexToAddress(spender)
	data, err := parsedABI.Pack("approve", spenderAddr, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to encode approve call: %w", err)
	}

	return data, nil
}

// CreateERC20BalanceOfData creates the call data for ERC20 balanceOf function using ABI encoding
func CreateERC20BalanceOfData(owner string) ([]byte, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Encode the balanceOf function call
	ownerAddr := common.HexToAddress(owner)
	data, err := parsedABI.Pack("balanceOf", ownerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to encode balanceOf call: %w", err)
	}

	return data, nil
}

// DecodeERC20TransferData decodes ERC20 transfer call data and returns the to address and amount
func DecodeERC20TransferData(data []byte) (to common.Address, amount *big.Int, err error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Decode the method and parameters
	method, err := parsedABI.MethodById(data[:4])
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to find method: %w", err)
	}

	if method.Name != "transfer" {
		return common.Address{}, nil, fmt.Errorf("not a transfer method: %s", method.Name)
	}

	// Unpack the parameters
	params, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to unpack parameters: %w", err)
	}

	if len(params) != 2 {
		return common.Address{}, nil, fmt.Errorf("invalid number of parameters: %d", len(params))
	}

	to = params[0].(common.Address)
	amount = params[1].(*big.Int)

	return to, amount, nil
}

// FormatTokenAmount formats a token amount with the specified decimals
func FormatTokenAmount(amount *big.Int, decimals uint8) string {
	if decimals == 0 {
		return amount.String()
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	quotient := new(big.Int).Div(amount, divisor)
	remainder := new(big.Int).Mod(amount, divisor)

	// Format the remainder with proper zero-padding
	remainderStr := remainder.String()
	for len(remainderStr) < int(decimals) {
		remainderStr = "0" + remainderStr
	}

	// Remove trailing zeros from remainder
	remainderStr = strings.TrimRight(remainderStr, "0")
	if remainderStr == "" {
		return quotient.String()
	}

	return quotient.String() + "." + remainderStr
}

// ParseTokenAmount parses a token amount string with the specified decimals
func ParseTokenAmount(amountStr string, decimals uint8) (*big.Int, error) {
	// Split by decimal point
	parts := strings.Split(amountStr, ".")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid amount format: %s", amountStr)
	}

	// Parse integer part
	integerPart, ok := new(big.Int).SetString(parts[0], 10)
	if !ok {
		return nil, fmt.Errorf("invalid integer part: %s", parts[0])
	}

	// Handle fractional part
	var fractionalPart *big.Int = big.NewInt(0)
	if len(parts) == 2 {
		fractionalStr := parts[1]
		if len(fractionalStr) > int(decimals) {
			return nil, fmt.Errorf("too many decimal places: %d, max: %d", len(fractionalStr), decimals)
		}

		// Pad with zeros to match decimals
		for len(fractionalStr) < int(decimals) {
			fractionalStr += "0"
		}

		var ok bool
		fractionalPart, ok = new(big.Int).SetString(fractionalStr, 10)
		if !ok {
			return nil, fmt.Errorf("invalid fractional part: %s", parts[1])
		}
	}

	// Calculate the final amount
	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	integerPartScaled := new(big.Int).Mul(integerPart, multiplier)
	result := new(big.Int).Add(integerPartScaled, fractionalPart)

	return result, nil
}
