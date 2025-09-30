package unit

import (
	"testing"

	"github.com/vikkkko/safe-core-sdk-golang/types"
)

func TestSafeVersion(t *testing.T) {
	tests := []struct {
		name    string
		version types.SafeVersion
		want    string
	}{
		{
			name:    "SafeVersion141",
			version: types.SafeVersion141,
			want:    "1.4.1",
		},
		{
			name:    "SafeVersion130",
			version: types.SafeVersion130,
			want:    "1.3.0",
		},
		{
			name:    "SafeVersion120",
			version: types.SafeVersion120,
			want:    "1.2.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.version) != tt.want {
				t.Errorf("SafeVersion = %v, want %v", tt.version, tt.want)
			}
		})
	}
}

func TestOperationType(t *testing.T) {
	tests := []struct {
		name      string
		operation types.OperationType
		want      uint8
	}{
		{
			name:      "Call",
			operation: types.Call,
			want:      0,
		},
		{
			name:      "DelegateCall",
			operation: types.DelegateCall,
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint8(tt.operation) != tt.want {
				t.Errorf("OperationType = %v, want %v", tt.operation, tt.want)
			}
		})
	}
}

func TestSafeSignature(t *testing.T) {
	sig := &types.SafeSignature{
		Signer:              "0x1234567890123456789012345678901234567890",
		Data:                "0xabcdef",
		IsContractSignature: false,
	}

	t.Run("StaticPart", func(t *testing.T) {
		staticPart := sig.StaticPart("")
		if staticPart != "0xabcdef" {
			t.Errorf("StaticPart() = %v, want %v", staticPart, "0xabcdef")
		}
	})

	t.Run("DynamicPart", func(t *testing.T) {
		dynamicPart := sig.DynamicPart()
		if dynamicPart != "" {
			t.Errorf("DynamicPart() = %v, want empty string", dynamicPart)
		}
	})
}

func TestSafeTransaction(t *testing.T) {
	tx := &types.SafeTransaction{
		Data: types.SafeTransactionData{
			To:             "0x1234567890123456789012345678901234567890",
			Value:          "1000000000000000000",
			Data:           "0x",
			Operation:      types.Call,
			SafeTxGas:      "0",
			BaseGas:        "0",
			GasPrice:       "0",
			GasToken:       "0x0000000000000000000000000000000000000000",
			RefundReceiver: "0x0000000000000000000000000000000000000000",
			Nonce:          1,
		},
		Signatures: make(map[string]types.SafeSignature),
	}

	sig := types.SafeSignature{
		Signer:              "0x1111111111111111111111111111111111111111",
		Data:                "0xsignature",
		IsContractSignature: false,
	}

	t.Run("AddSignature", func(t *testing.T) {
		tx.AddSignature(sig)
		if len(tx.Signatures) != 1 {
			t.Errorf("Expected 1 signature, got %d", len(tx.Signatures))
		}
	})

	t.Run("GetSignature", func(t *testing.T) {
		retrievedSig := tx.GetSignature(sig.Signer)
		if retrievedSig == nil {
			t.Error("Expected to find signature, got nil")
		}
		if retrievedSig.Data != sig.Data {
			t.Errorf("Expected signature data %v, got %v", sig.Data, retrievedSig.Data)
		}
	})

	t.Run("GetNonexistentSignature", func(t *testing.T) {
		retrievedSig := tx.GetSignature("0x9999999999999999999999999999999999999999")
		if retrievedSig != nil {
			t.Error("Expected nil for nonexistent signature, got signature")
		}
	})
}

func TestEIP3770Address(t *testing.T) {
	addr := &types.EIP3770Address{
		Prefix:  "eth",
		Address: "0x1234567890123456789012345678901234567890",
	}

	t.Run("String", func(t *testing.T) {
		expected := "eth:0x1234567890123456789012345678901234567890"
		result := addr.String()
		if result != expected {
			t.Errorf("String() = %v, want %v", result, expected)
		}
	})
}

func TestSafeSetupConfig(t *testing.T) {
	config := types.SafeSetupConfig{
		Owners:    []string{"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222"},
		Threshold: 2,
	}

	t.Run("ValidConfig", func(t *testing.T) {
		if len(config.Owners) != 2 {
			t.Errorf("Expected 2 owners, got %d", len(config.Owners))
		}
		if config.Threshold != 2 {
			t.Errorf("Expected threshold 2, got %d", config.Threshold)
		}
	})
}

func TestDefaultSafeVersion(t *testing.T) {
	if types.DefaultSafeVersion != types.SafeVersion141 {
		t.Errorf("Expected default version %v, got %v", types.SafeVersion141, types.DefaultSafeVersion)
	}
}

func TestSigningMethod(t *testing.T) {
	tests := []struct {
		name   string
		method types.SigningMethod
		want   string
	}{
		{
			name:   "ETH_Sign",
			method: types.SigningMethodETHSign,
			want:   "eth_sign",
		},
		{
			name:   "ETH_SignTypedData",
			method: types.SigningMethodETHSignTypedData,
			want:   "eth_signTypedData",
		},
		{
			name:   "Safe_SignMessage",
			method: types.SigningMethodSafeSignMessage,
			want:   "safe_signMessage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.method) != tt.want {
				t.Errorf("SigningMethod = %v, want %v", tt.method, tt.want)
			}
		})
	}
}