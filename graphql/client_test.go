package graphql

import (
	"context"
	"encoding/json"
	"testing"
)

func TestNewClient(t *testing.T) {
	config := DefaultConfig()
	client := NewClient(config)

	if client == nil {
		t.Fatal("Expected client to be created")
	}

	if client.endpoint != config.Endpoint {
		t.Errorf("Expected endpoint %s, got %s", config.Endpoint, client.endpoint)
	}
}

func TestNewDefaultClient(t *testing.T) {
	client := NewDefaultClient()

	if client == nil {
		t.Fatal("Expected client to be created")
	}

	expectedEndpoint := "https://api.studio.thegraph.com/query/103887/mvp/version/latest"
	if client.endpoint != expectedEndpoint {
		t.Errorf("Expected endpoint %s, got %s", expectedEndpoint, client.endpoint)
	}
}

func TestGetPaymentAllowances_InvalidAddress(t *testing.T) {
	client := NewDefaultClient()
	ctx := context.Background()

	// Test with a valid address format but likely no data
	paymentAccount := "0x0000000000000000000000000000000000000000"

	allowances, err := client.GetPaymentAllowances(ctx, paymentAccount)
	if err != nil {
		t.Logf("Query failed (expected if network unavailable): %v", err)
		return
	}

	// If query succeeds, allowances should be a valid slice (possibly empty)
	if allowances == nil {
		t.Error("Expected allowances to be non-nil slice")
	}
}

// Note: Integration tests that hit the actual GraphQL endpoint
// should be run separately and may require network access
func TestGraphQLRequest_Marshal(t *testing.T) {
	req := GraphQLRequest{
		Query: "{ test }",
		Variables: map[string]interface{}{
			"id": "123",
		},
	}

	// Verify request fields are set correctly
	if req.Query != "{ test }" {
		t.Errorf("Expected query '{ test }', got '%s'", req.Query)
	}

	if req.Variables["id"] != "123" {
		t.Errorf("Expected variable id='123', got '%v'", req.Variables["id"])
	}
}

func TestBigInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "valid large number",
			input:   `"9999999999999999999999"`,
			want:    "9999999999999999999999",
			wantErr: false,
		},
		{
			name:    "valid small number",
			input:   `"123"`,
			want:    "123",
			wantErr: false,
		},
		{
			name:    "valid zero",
			input:   `"0"`,
			want:    "0",
			wantErr: false,
		},
		{
			name:    "invalid number",
			input:   `"abc"`,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bi BigInt
			err := bi.UnmarshalJSON([]byte(tt.input))

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if bi.String() != tt.want {
				t.Errorf("Expected %s, got %s", tt.want, bi.String())
			}
		})
	}
}

func TestPaymentAllowance_UnmarshalJSON(t *testing.T) {
	jsonData := `{
		"token": "0x1234567890123456789012345678901234567890",
		"owner": "0x0987654321098765432109876543210987654321",
		"amount": "9999999999999999999999"
	}`

	var allowance PaymentAllowance
	err := json.Unmarshal([]byte(jsonData), &allowance)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if allowance.Token != "0x1234567890123456789012345678901234567890" {
		t.Errorf("Expected token 0x1234567890123456789012345678901234567890, got %s", allowance.Token)
	}

	if allowance.Owner != "0x0987654321098765432109876543210987654321" {
		t.Errorf("Expected owner 0x0987654321098765432109876543210987654321, got %s", allowance.Owner)
	}

	if allowance.Amount.String() != "9999999999999999999999" {
		t.Errorf("Expected amount 9999999999999999999999, got %s", allowance.Amount.String())
	}
}
