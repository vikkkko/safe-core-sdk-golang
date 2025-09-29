package utils

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/yinwei/safe-core-sdk-golang/types"
)

// SignMessage signs a message with a private key using Safe's expected signature format
func SignMessage(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	// The message is already a hash, sign it directly
	signature, err := crypto.Sign(message, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}

	// Adjust v value for Safe compatibility
	// Safe expects v to be 27 or 28 for ECDSA signatures
	if signature[64] < 27 {
		signature[64] += 27
	}

	// For Safe, we need to check if the signature has the EIP-191 prefix
	// If recovery doesn't match the expected signer, it means the message was signed with prefix
	signerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Test signature recovery
	testSig := make([]byte, 65)
	copy(testSig, signature)
	if testSig[64] >= 27 {
		testSig[64] -= 27
	}

	recoveredPubKey, err := crypto.SigToPub(message, testSig)
	if err != nil {
		return nil, fmt.Errorf("failed to verify signature: %w", err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey)

	// If addresses don't match, it means we need to add 4 to v (EIP-191 prefix)
	if recoveredAddress != signerAddress {
		signature[64] += 4
	}

	return signature, nil
}

// RecoverSigner recovers the signer address from a signature
func RecoverSigner(message []byte, signature []byte) (common.Address, error) {
	// The message is already a hash, use it directly
	if len(signature) != 65 {
		return common.Address{}, fmt.Errorf("invalid signature length: %d", len(signature))
	}

	// Create a copy to avoid modifying the original
	sig := make([]byte, 65)
	copy(sig, signature)

	// Adjust v value for recovery (crypto.SigToPub expects v = 0 or 1)
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	// Recover the public key
	pubkey, err := crypto.SigToPub(message, sig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to recover public key: %w", err)
	}

	// Get the address from the public key
	address := crypto.PubkeyToAddress(*pubkey)
	return address, nil
}

// BuildSignatureBytes builds the signature bytes for a Safe transaction
func BuildSignatureBytes(signatures []types.SafeSignature) []byte {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Sort signatures by signer address
	// 2. Concatenate all signature data according to Safe's format
	// 3. Return the final signature bytes

	var result []byte
	for _, sig := range signatures {
		// Convert hex signature to bytes
		sigBytes := common.FromHex(sig.Data)
		result = append(result, sigBytes...)
	}

	return result
}

// GenerateSignature generates a signature for a Safe transaction or message
func GenerateSignature(hash []byte, signerAddress common.Address, privateKey *ecdsa.PrivateKey) (*types.SafeSignature, error) {
	// Sign the hash
	signature, err := SignMessage(hash, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign hash: %w", err)
	}

	return &types.SafeSignature{
		Signer:              signerAddress.Hex(),
		Data:                common.Bytes2Hex(signature),
		IsContractSignature: false,
	}, nil
}

// GeneratePreValidatedSignature generates a pre-validated signature
func GeneratePreValidatedSignature(signerAddress common.Address) *types.SafeSignature {
	// Pre-validated signatures are used when the signer is a Safe owner
	// and the signature validation is done on-chain
	return &types.SafeSignature{
		Signer:              signerAddress.Hex(),
		Data:                "0x000000000000000000000000" + signerAddress.Hex()[2:] + "0000000000000000000000000000000000000000000000000000000000000000" + "01",
		IsContractSignature: false,
	}
}

// GenerateContractSignature generates a contract signature
func GenerateContractSignature(contractAddress common.Address, signatureData []byte) *types.SafeSignature {
	return &types.SafeSignature{
		Signer:              contractAddress.Hex(),
		Data:                common.Bytes2Hex(signatureData),
		IsContractSignature: true,
	}
}

// CalculateTransactionHash calculates the hash of a Safe transaction for signing using EIP-712
func CalculateTransactionHash(
	safeAddress common.Address,
	to common.Address,
	value *big.Int,
	data []byte,
	operation uint8,
	safeTxGas *big.Int,
	baseGas *big.Int,
	gasPrice *big.Int,
	gasToken common.Address,
	refundReceiver common.Address,
	nonce *big.Int,
	chainID *big.Int,
) ([]byte, error) {
	// EIP-712 domain separator
	domainSeparator := calculateDomainSeparator(safeAddress, chainID)

	// SafeTx type hash (keccak256 of the type string)
	safeTxTypeHash := crypto.Keccak256([]byte("SafeTx(address to,uint256 value,bytes data,uint8 operation,uint256 safeTxGas,uint256 baseGas,uint256 gasPrice,address gasToken,address refundReceiver,uint256 nonce)"))

	// Encode transaction data according to EIP-712
	encodedTxData := encodeSafeTransactionData(
		safeTxTypeHash,
		to,
		value,
		data,
		operation,
		safeTxGas,
		baseGas,
		gasPrice,
		gasToken,
		refundReceiver,
		nonce,
	)

	// Final EIP-712 hash: keccak256("\x19\x01" + domainSeparator + structHash)
	finalHash := crypto.Keccak256(
		[]byte("\x19\x01"),
		domainSeparator,
		encodedTxData,
	)

	return finalHash, nil
}

// calculateDomainSeparator calculates the EIP-712 domain separator for Safe v1.4.1+
func calculateDomainSeparator(safeAddress common.Address, chainID *big.Int) []byte {
	// Domain type hash for Safe v1.3.0+ (includes chainId)
	domainTypeHash := crypto.Keccak256([]byte("EIP712Domain(uint256 chainId,address verifyingContract)"))

	// Encode domain separator
	domainSeparator := crypto.Keccak256(
		domainTypeHash,
		common.LeftPadBytes(chainID.Bytes(), 32),
		common.LeftPadBytes(safeAddress.Bytes(), 32),
	)

	return domainSeparator
}

// encodeSafeTransactionData encodes the Safe transaction data according to EIP-712
func encodeSafeTransactionData(
	typehash []byte,
	to common.Address,
	value *big.Int,
	data []byte,
	operation uint8,
	safeTxGas *big.Int,
	baseGas *big.Int,
	gasPrice *big.Int,
	gasToken common.Address,
	refundReceiver common.Address,
	nonce *big.Int,
) []byte {
	// Hash of the data field (keccak256 of bytes data)
	dataHash := crypto.Keccak256(data)

	// Encode all fields as 32-byte values
	return crypto.Keccak256(
		typehash,                                       // SafeTx typehash
		common.LeftPadBytes(to.Bytes(), 32),           // to address
		common.LeftPadBytes(value.Bytes(), 32),        // value
		dataHash,                                      // keccak256(data)
		common.LeftPadBytes([]byte{operation}, 32),    // operation
		common.LeftPadBytes(safeTxGas.Bytes(), 32),    // safeTxGas
		common.LeftPadBytes(baseGas.Bytes(), 32),      // baseGas
		common.LeftPadBytes(gasPrice.Bytes(), 32),     // gasPrice
		common.LeftPadBytes(gasToken.Bytes(), 32),     // gasToken
		common.LeftPadBytes(refundReceiver.Bytes(), 32), // refundReceiver
		common.LeftPadBytes(nonce.Bytes(), 32),        // nonce
	)
}

// CalculateMessageHash calculates the hash of a Safe message for signing
func CalculateMessageHash(
	safeAddress common.Address,
	message []byte,
	chainID *big.Int,
) ([]byte, error) {
	// This is a placeholder implementation
	// In a real implementation, this would:
	// 1. Encode the message according to Safe's format
	// 2. Calculate the hash using the Safe's domain separator
	// 3. Return the final hash for signing

	// Placeholder hash calculation
	data := append(safeAddress.Bytes(), message...)
	data = append(data, chainID.Bytes()...)

	hash := crypto.Keccak256(data)
	return hash, nil
}

// AdjustVInSignature adjusts the v value in a signature for Safe compatibility
func AdjustVInSignature(signature []byte) []byte {
	if len(signature) != 65 {
		return signature
	}

	// Adjust v value if needed
	if signature[64] < 27 {
		signature[64] += 27
	}

	return signature
}