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
// This matches SafeProxyFactory.createProxyWithNonce logic
func CalculateProxyAddress(
	factory common.Address,
	singleton common.Address,
	initializer []byte,
	saltNonce *big.Int,
) (common.Address, error) {
	// Step 1: Calculate salt = keccak256(abi.encodePacked(keccak256(initializer), saltNonce))
	// This matches the Solidity code: bytes32 salt = keccak256(abi.encodePacked(keccak256(initializer), saltNonce));
	initializerHash := crypto.Keccak256(initializer)

	// Encode saltNonce as bytes32 (32 bytes, big-endian)
	saltNonceBytes := make([]byte, 32)
	saltNonce.FillBytes(saltNonceBytes)

	// Pack initializerHash and saltNonce
	saltData := append(initializerHash, saltNonceBytes...)
	salt := crypto.Keccak256(saltData)

	// Step 2: Calculate init code hash
	// This matches: abi.encodePacked(type(SafeProxy).creationCode, uint256(uint160(_singleton)))
	// SafeProxy creation code (from the factory contract's proxyCreationCode() method)
	proxyCreationCode := common.FromHex("0x608060405234801561001057600080fd5b506040516101e63803806101e68339818101604052602081101561003357600080fd5b8101908080519060200190929190505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156100ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806101c46022913960400191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060ab806101196000396000f3fe608060405273ffffffffffffffffffffffffffffffffffffffff600054167fa619486e0000000000000000000000000000000000000000000000000000000060003514156050578060005260206000f35b3660008037600080366000845af43d6000803e60008114156070573d6000fd5b3d6000f3fea2646970667358221220d1429297349653a4918076d650332de1a1068c5f3e07c5c82360c277770b955264736f6c63430007060033496e76616c69642073696e676c65746f6e20616464726573732070726f7669646564")

	// Encode singleton address as uint256 (32 bytes, left-padded)
	singletonUint256 := make([]byte, 32)
	copy(singletonUint256[12:], singleton.Bytes()) // address is 20 bytes, so pad with 12 zeros on left

	// Combine creation code with singleton address
	initCode := append(proxyCreationCode, singletonUint256...)
	initCodeHash := crypto.Keccak256(initCode)

	// Step 3: Calculate CREATE2 address
	// Formula: keccak256(0xff ++ factory ++ salt ++ keccak256(initCode))[12:]
	data := []byte{0xff}
	data = append(data, factory.Bytes()...)
	data = append(data, salt...)
	data = append(data, initCodeHash...)

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