package utils

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vikkkko/safe-core-sdk-golang/types"
)

// StandardizeSafeTransactionData fills in missing fields in transaction data with defaults
func StandardizeSafeTransactionData(txData types.SafeTransactionDataPartial, currentNonce uint64) *types.SafeTransactionData {
	// Set default operation type if not specified
	operation := types.Call
	if txData.Operation != nil {
		operation = *txData.Operation
	}

	// Use provided nonce or current nonce
	nonce := currentNonce
	if txData.Nonce != nil {
		nonce = *txData.Nonce
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
	}
}

// EncodeMultiSendData encodes multiple transactions for MultiSend contract
func EncodeMultiSendData(transactions []types.MetaTransactionData) ([]byte, error) {
	var encoded []byte

	for _, tx := range transactions {
		// Default operation to Call if not specified
		operation := types.Call
		if tx.Operation != nil {
			operation = *tx.Operation
		}

		// Encode operation (1 byte)
		encoded = append(encoded, byte(operation))

		// Encode to address (20 bytes)
		toAddr := common.HexToAddress(tx.To)
		encoded = append(encoded, toAddr.Bytes()...)

		// Encode value (32 bytes)
		value, ok := new(big.Int).SetString(tx.Value, 10)
		if !ok {
			return nil, fmt.Errorf("invalid value: %s", tx.Value)
		}
		valueBytes := make([]byte, 32)
		value.FillBytes(valueBytes)
		encoded = append(encoded, valueBytes...)

		// Encode data length (32 bytes)
		data := common.FromHex(tx.Data)
		dataLength := make([]byte, 32)
		binary.BigEndian.PutUint64(dataLength[24:], uint64(len(data)))
		encoded = append(encoded, dataLength...)

		// Encode data
		encoded = append(encoded, data...)
	}

	return encoded, nil
}

// EstimateTxGas estimates gas for a transaction
func EstimateTxGas(txData types.SafeTransactionData) (*big.Int, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Calculate gas based on transaction complexity
	// 2. Consider the data size, operation type, etc.
	// 3. Return an estimated gas amount

	// Simple estimation based on data size
	dataSize := len(common.FromHex(txData.Data))
	baseGas := big.NewInt(21000) // Base transaction gas
	dataGas := big.NewInt(int64(dataSize * 68)) // Rough estimation for data gas

	return new(big.Int).Add(baseGas, dataGas), nil
}

// EstimateSafeTxGas estimates Safe transaction gas
func EstimateSafeTxGas(txData types.SafeTransactionData) (*big.Int, error) {
	// This is a placeholder implementation
	// In a real implementation, this would estimate the gas needed
	// for the Safe-specific operations (signature verification, etc.)

	return big.NewInt(0), nil
}

// EstimateBaseGas estimates base gas for Safe operations
func EstimateBaseGas() (*big.Int, error) {
	// This is a placeholder implementation
	// In a real implementation, this would estimate the base gas
	// needed for Safe contract operations

	return big.NewInt(0), nil
}

// ValidateTransactionData validates Safe transaction data
func ValidateTransactionData(txData types.SafeTransactionData) error {
	// Validate to address
	if !common.IsHexAddress(txData.To) {
		return fmt.Errorf("invalid to address: %s", txData.To)
	}

	// Validate value
	if _, ok := new(big.Int).SetString(txData.Value, 10); !ok {
		return fmt.Errorf("invalid value: %s", txData.Value)
	}

	// Validate gas values
	if _, ok := new(big.Int).SetString(txData.SafeTxGas, 10); !ok {
		return fmt.Errorf("invalid safeTxGas: %s", txData.SafeTxGas)
	}

	if _, ok := new(big.Int).SetString(txData.BaseGas, 10); !ok {
		return fmt.Errorf("invalid baseGas: %s", txData.BaseGas)
	}

	if _, ok := new(big.Int).SetString(txData.GasPrice, 10); !ok {
		return fmt.Errorf("invalid gasPrice: %s", txData.GasPrice)
	}

	// Validate gas token address
	if !common.IsHexAddress(txData.GasToken) {
		return fmt.Errorf("invalid gasToken address: %s", txData.GasToken)
	}

	// Validate refund receiver address
	if !common.IsHexAddress(txData.RefundReceiver) {
		return fmt.Errorf("invalid refundReceiver address: %s", txData.RefundReceiver)
	}

	return nil
}


// ParseTransactionValue parses a string value to big.Int
func ParseTransactionValue(value string) (*big.Int, error) {
	// Try to parse as decimal
	if val, ok := new(big.Int).SetString(value, 10); ok {
		return val, nil
	}

	// Try to parse as hex
	if val, ok := new(big.Int).SetString(value, 0); ok {
		return val, nil
	}

	return nil, fmt.Errorf("invalid transaction value: %s", value)
}

// FormatTransactionValue formats a big.Int value as string
func FormatTransactionValue(value *big.Int) string {
	return value.String()
}

// CalculateTransactionSize calculates the size of a transaction in bytes
func CalculateTransactionSize(txData types.SafeTransactionData) int {
	size := 0

	// Address fields (20 bytes each)
	size += 20 // to
	size += 20 // gasToken
	size += 20 // refundReceiver

	// Value and gas fields (32 bytes each when encoded)
	size += 32 // value
	size += 32 // safeTxGas
	size += 32 // baseGas
	size += 32 // gasPrice

	// Operation (1 byte)
	size += 1

	// Nonce (8 bytes when encoded as uint64)
	size += 8

	// Data length
	dataSize := len(common.FromHex(txData.Data))
	size += dataSize

	return size
}

// IsEmptyData checks if transaction data is empty
func IsEmptyData(data string) bool {
	return data == "" || data == "0x" || data == "0x0"
}

// HasDelegateCalls checks if any transaction in a batch uses delegate call
func HasDelegateCalls(transactions []types.MetaTransactionData) bool {
	for _, tx := range transactions {
		if tx.Operation != nil && *tx.Operation == types.DelegateCall {
			return true
		}
	}
	return false
}

// ConvertWeiToEther converts wei to ether (as string)
func ConvertWeiToEther(wei *big.Int) string {
	ether := new(big.Rat).SetFrac(wei, big.NewInt(1e18))
	return ether.FloatString(18)
}

// ConvertEtherToWei converts ether to wei
func ConvertEtherToWei(ether string) (*big.Int, error) {
	etherFloat, err := strconv.ParseFloat(ether, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid ether value: %s", ether)
	}

	// Convert to wei (multiply by 10^18)
	wei := new(big.Float).Mul(big.NewFloat(etherFloat), big.NewFloat(1e18))
	weiInt, _ := wei.Int(nil)

	return weiInt, nil
}