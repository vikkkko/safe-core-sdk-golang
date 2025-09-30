package utils

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vikkkko/safe-core-sdk-golang/types"
)

// PredictSafeAddress predicts the address of a Safe before deployment
func PredictSafeAddress(config types.SafeDeploymentConfig, chainID *big.Int) (string, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Calculate the initializer data for the Safe setup
	// 2. Calculate the CREATE2 address using the factory, salt, and initializer
	// 3. Return the predicted address

	// For now, return a placeholder address
	return "0x0000000000000000000000000000000000000000", nil
}

// ValidateEthereumAddress validates if a string is a valid Ethereum address
func ValidateEthereumAddress(address string) bool {
	return common.IsHexAddress(address)
}

// ValidateEIP3770Address validates if an address follows the EIP-3770 format
func ValidateEIP3770Address(address string) bool {
	// EIP-3770 format: <shortName>:<address>
	parts := strings.Split(address, ":")
	if len(parts) != 2 {
		return false
	}

	// Validate the address part
	return ValidateEthereumAddress(parts[1])
}

// ParseEIP3770Address parses an EIP-3770 address
func ParseEIP3770Address(address string) (*types.EIP3770Address, error) {
	parts := strings.Split(address, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid EIP-3770 address format: %s", address)
	}

	if !ValidateEthereumAddress(parts[1]) {
		return nil, fmt.Errorf("invalid Ethereum address in EIP-3770 format: %s", parts[1])
	}

	return &types.EIP3770Address{
		Prefix:  parts[0],
		Address: parts[1],
	}, nil
}

// ChecksumAddress returns the checksummed version of an address
func ChecksumAddress(address string) string {
	return common.HexToAddress(address).Hex()
}

// IsSameAddress compares two addresses (case-insensitive)
func IsSameAddress(addr1, addr2 string) bool {
	return strings.EqualFold(addr1, addr2)
}

// CalculateProxyAddress calculates the CREATE2 address for a proxy
func CalculateProxyAddress(
	factory common.Address,
	singleton common.Address,
	initializer []byte,
	saltNonce *big.Int,
) (common.Address, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Get the proxy creation code from the factory
	// 2. Calculate the salt from saltNonce and keccak256(initializer)
	// 3. Calculate the CREATE2 address

	// Placeholder calculation
	salt := crypto.Keccak256Hash(saltNonce.Bytes(), crypto.Keccak256(initializer))

	// This is a simplified version - real implementation would use the actual proxy bytecode
	initCodeHash := crypto.Keccak256Hash([]byte("placeholder"))

	// CREATE2 address calculation: keccak256(0xff ++ factory ++ salt ++ keccak256(initCode))[12:]
	data := append([]byte{0xff}, factory.Bytes()...)
	data = append(data, salt.Bytes()...)
	data = append(data, initCodeHash.Bytes()...)

	hash := crypto.Keccak256(data)
	return common.BytesToAddress(hash[12:]), nil
}

// EncodePackedData encodes data in packed format (similar to abi.encodePacked)
func EncodePackedData(data ...[]byte) []byte {
	var result []byte
	for _, d := range data {
		result = append(result, d...)
	}
	return result
}

// Keccak256 calculates the Keccak256 hash of data
func Keccak256(data []byte) []byte {
	return crypto.Keccak256(data)
}

// Sha256 calculates the SHA256 hash of data
func Sha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}