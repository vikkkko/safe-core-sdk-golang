package unit

import (
	"testing"

	"github.com/yinwei/safe-core-sdk-golang/api"
)

func TestSafeApiKitConfig(t *testing.T) {
	tests := []struct {
		name        string
		config      api.SafeApiKitConfig
		expectError bool
		errorMsg    string
	}{
		{
			name: "ValidConfigWithCustomURL",
			config: api.SafeApiKitConfig{
				ChainID:      1,
				TxServiceURL: "https://custom-api.example.com",
				ApiKey:       "test-api-key",
			},
			expectError: false,
		},
		{
			name: "ValidConfigWithDefaultURL",
			config: api.SafeApiKitConfig{
				ChainID: 1,
				ApiKey:  "test-api-key",
			},
			expectError: false,
		},
		{
			name: "InvalidConfigSafeGlobalNoApiKey",
			config: api.SafeApiKitConfig{
				ChainID:      1,
				TxServiceURL: "https://api.safe.global",
			},
			expectError: true,
			errorMsg:    "apiKey is mandatory when using api.safe.global or api.5afe.dev domains",
		},
		{
			name: "InvalidConfig5afeDevNoApiKey",
			config: api.SafeApiKitConfig{
				ChainID:      1,
				TxServiceURL: "https://api.5afe.dev",
			},
			expectError: true,
			errorMsg:    "apiKey is mandatory when using api.safe.global or api.5afe.dev domains",
		},
		{
			name: "InvalidConfigNoURLNoApiKey",
			config: api.SafeApiKitConfig{
				ChainID: 1,
			},
			expectError: true,
			errorMsg:    "apiKey is mandatory when txServiceUrl is not defined",
		},
		{
			name: "ValidConfigSafeGlobalWithApiKey",
			config: api.SafeApiKitConfig{
				ChainID:      1,
				TxServiceURL: "https://api.safe.global",
				ApiKey:       "test-api-key",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := api.NewSafeApiKit(tt.config)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got nil")
				}
				if tt.errorMsg != "" && err != nil {
					// Check if error message contains expected text
					if !containsString(err.Error(), tt.errorMsg) {
						t.Errorf("Expected error to contain '%s', got '%s'", tt.errorMsg, err.Error())
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if client == nil {
					t.Error("Expected client to be created but got nil")
				}
			}
		})
	}
}

func TestGetMultisigTransactionsOptions(t *testing.T) {
	options := &api.GetMultisigTransactionsOptions{
		Executed:    boolPtr(true),
		TrustedOnly: boolPtr(false),
		Limit:       intPtr(10),
		Offset:      intPtr(20),
	}

	if options.Executed == nil || *options.Executed != true {
		t.Error("Expected Executed to be true")
	}

	if options.TrustedOnly == nil || *options.TrustedOnly != false {
		t.Error("Expected TrustedOnly to be false")
	}

	if options.Limit == nil || *options.Limit != 10 {
		t.Error("Expected Limit to be 10")
	}

	if options.Offset == nil || *options.Offset != 20 {
		t.Error("Expected Offset to be 20")
	}
}

func TestProposeTransactionProps(t *testing.T) {
	props := api.ProposeTransactionProps{
		SafeAddress:    "0x1234567890123456789012345678901234567890",
		SafeTxHash:     "0xabcdef1234567890",
		To:             "0x9876543210987654321098765432109876543210",
		Value:          "1000000000000000000",
		Data:           "0x",
		Operation:      0,
		GasToken:       "0x0000000000000000000000000000000000000000",
		SafeTxGas:      21000,
		BaseGas:        5000,
		GasPrice:       "20000000000",
		RefundReceiver: "0x0000000000000000000000000000000000000000",
		Nonce:          1,
		Sender:         "0x1111111111111111111111111111111111111111",
		Signature:      "0xsignature",
	}

	if props.SafeAddress != "0x1234567890123456789012345678901234567890" {
		t.Errorf("Expected SafeAddress to be set correctly")
	}

	if props.Value != "1000000000000000000" {
		t.Errorf("Expected Value to be 1000000000000000000, got %s", props.Value)
	}

	if props.Operation != 0 {
		t.Errorf("Expected Operation to be 0, got %d", props.Operation)
	}

	if props.SafeTxGas != 21000 {
		t.Errorf("Expected SafeTxGas to be 21000, got %d", props.SafeTxGas)
	}
}

func TestSafeInfoResponse(t *testing.T) {
	response := api.SafeInfoResponse{
		Address:         "0x1234567890123456789012345678901234567890",
		Nonce:           5,
		Threshold:       2,
		Owners:          []string{"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222"},
		MasterCopy:      "0x3333333333333333333333333333333333333333",
		Modules:         []string{"0x4444444444444444444444444444444444444444"},
		FallbackHandler: "0x5555555555555555555555555555555555555555",
		Guard:           "0x6666666666666666666666666666666666666666",
		Version:         "1.4.1",
	}

	if response.Address != "0x1234567890123456789012345678901234567890" {
		t.Error("Address not set correctly")
	}

	if response.Nonce != 5 {
		t.Errorf("Expected Nonce to be 5, got %d", response.Nonce)
	}

	if response.Threshold != 2 {
		t.Errorf("Expected Threshold to be 2, got %d", response.Threshold)
	}

	if len(response.Owners) != 2 {
		t.Errorf("Expected 2 owners, got %d", len(response.Owners))
	}

	if response.Version != "1.4.1" {
		t.Errorf("Expected Version to be 1.4.1, got %s", response.Version)
	}
}

func TestSafeMultisigTransactionEstimate(t *testing.T) {
	estimate := api.SafeMultisigTransactionEstimate{
		To:        "0x1234567890123456789012345678901234567890",
		Value:     "1000000000000000000",
		Data:      "0xa9059cbb",
		Operation: 0,
	}

	if estimate.To != "0x1234567890123456789012345678901234567890" {
		t.Error("To address not set correctly")
	}

	if estimate.Value != "1000000000000000000" {
		t.Errorf("Expected Value to be 1000000000000000000, got %s", estimate.Value)
	}

	if estimate.Data != "0xa9059cbb" {
		t.Errorf("Expected Data to be 0xa9059cbb, got %s", estimate.Data)
	}

	if estimate.Operation != 0 {
		t.Errorf("Expected Operation to be 0, got %d", estimate.Operation)
	}
}

// Helper functions
func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr ||
		   len(s) > len(substr) && s[len(s)-len(substr):] == substr ||
		   (len(s) > len(substr) && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}