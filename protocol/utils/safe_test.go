package utils_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
)

// Example demonstrates basic Safe deployment data preparation
func ExamplePrepareSafeDeployment() {
	// Define owners
	owners := []common.Address{
		common.HexToAddress("0x1234567890123456789012345678901234567890"),
		common.HexToAddress("0x2345678901234567890123456789012345678901"),
		common.HexToAddress("0x3456789012345678901234567890123456789012"),
	}

	// Prepare deployment configuration
	config := utils.DeploySafeConfig{
		Owners:           owners,
		Threshold:        2, // 2 out of 3 multisig
		FactoryAddress:   common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
		SingletonAddress: common.HexToAddress("0x29fcB43b46531BcA003ddC8FCB67FFE91900C762"),
		SaltNonce:        big.NewInt(0),
	}

	// Get factory call data
	callData, err := utils.PrepareSafeDeployment(config)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Factory call data length: %d bytes\n", len(callData))
	// Output: Factory call data length: 580 bytes
}

// Example demonstrates parsing owner addresses from strings
func ExampleParseOwnersFromStrings() {
	ownerStrings := []string{
		"0x1234567890123456789012345678901234567890",
		"0x2345678901234567890123456789012345678901",
	}

	owners, err := utils.ParseOwnersFromStrings(ownerStrings)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Parsed %d owners\n", len(owners))
	// Output: Parsed 2 owners
}

// Example demonstrates creating Safe init data with simple parameters
func ExampleCreateSafeInitDataSimple() {
	owners := []common.Address{
		common.HexToAddress("0x1234567890123456789012345678901234567890"),
		common.HexToAddress("0x2345678901234567890123456789012345678901"),
	}

	initData, err := utils.CreateSafeInitDataSimple(owners, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Init data length: %d bytes\n", len(initData))
	// Output: Init data length: 388 bytes
}

// TestPrepareSafeDeployment tests the complete deployment data preparation
func TestPrepareSafeDeployment(t *testing.T) {
	owners := []common.Address{
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
	}

	config := utils.DeploySafeConfig{
		Owners:           owners,
		Threshold:        2,
		FactoryAddress:   common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
		SingletonAddress: common.HexToAddress("0x29fcB43b46531BcA003ddC8FCB67FFE91900C762"),
		SaltNonce:        big.NewInt(0),
	}

	callData, err := utils.PrepareSafeDeployment(config)
	if err != nil {
		t.Fatalf("Failed to prepare deployment: %v", err)
	}

	if len(callData) == 0 {
		t.Error("Expected non-empty call data")
	}

	// Call data should start with function selector for createProxyWithNonce
	// Function selector is first 4 bytes
	if len(callData) < 4 {
		t.Error("Call data too short to contain function selector")
	}
}

// TestParseOwnersFromStrings tests owner address parsing
func TestParseOwnersFromStrings(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expectError bool
	}{
		{
			name: "valid addresses",
			input: []string{
				"0x1234567890123456789012345678901234567890",
				"0x2345678901234567890123456789012345678901",
			},
			expectError: false,
		},
		{
			name:        "empty input",
			input:       []string{},
			expectError: true,
		},
		{
			name: "invalid address",
			input: []string{
				"0x1234567890123456789012345678901234567890",
				"not-an-address",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			owners, err := utils.ParseOwnersFromStrings(tt.input)
			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if len(owners) != len(tt.input) {
					t.Errorf("Expected %d owners, got %d", len(tt.input), len(owners))
				}
			}
		})
	}
}

// TestCreateSafeInitData tests Safe initialization data creation
func TestCreateSafeInitData(t *testing.T) {
	owners := []common.Address{
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
	}

	tests := []struct {
		name        string
		config      utils.SafeSetupConfig
		expectError bool
	}{
		{
			name: "valid config",
			config: utils.SafeSetupConfig{
				Owners:          owners,
				Threshold:       big.NewInt(1),
				To:              common.Address{},
				Data:            []byte{},
				FallbackHandler: common.HexToAddress("0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99"),
				PaymentToken:    common.Address{},
				Payment:         big.NewInt(0),
				PaymentReceiver: common.Address{},
			},
			expectError: false,
		},
		{
			name: "no owners",
			config: utils.SafeSetupConfig{
				Owners:    []common.Address{},
				Threshold: big.NewInt(1),
			},
			expectError: true,
		},
		{
			name: "threshold too high",
			config: utils.SafeSetupConfig{
				Owners:    owners,
				Threshold: big.NewInt(2), // More than number of owners
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := utils.CreateSafeInitData(tt.config)
			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if len(data) == 0 {
					t.Error("Expected non-empty init data")
				}
			}
		})
	}
}

// TestDefaultSafeSetupConfig tests default configuration creation
func TestDefaultSafeSetupConfig(t *testing.T) {
	owners := []common.Address{
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
	}

	config := utils.DefaultSafeSetupConfig(owners, 2)

	if len(config.Owners) != 2 {
		t.Errorf("Expected 2 owners, got %d", len(config.Owners))
	}

	if config.Threshold.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("Expected threshold 2, got %s", config.Threshold.String())
	}

	// Should have default fallback handler
	if config.FallbackHandler == (common.Address{}) {
		t.Error("Expected default fallback handler to be set")
	}
}
