package integration

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/vikkkko/safe-core-sdk-golang/api"
	"github.com/vikkkko/safe-core-sdk-golang/protocol"
	"github.com/vikkkko/safe-core-sdk-golang/types"
)

// Integration tests require real network connections and API keys
// These tests are skipped by default unless environment variables are set

func TestSafeClientIntegration(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test - set RUN_INTEGRATION_TESTS=true to run")
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("Skipping integration test - RPC_URL not set")
	}

	safeAddress := os.Getenv("SAFE_ADDRESS")
	if safeAddress == "" {
		t.Skip("Skipping integration test - SAFE_ADDRESS not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test Safe client creation
	t.Run("CreateSafeClient", func(t *testing.T) {
		client, err := protocol.NewSafe(protocol.SafeConfig{
			SafeAddress: safeAddress,
			RpcURL:      rpcURL,
			ChainID:     1, // Mainnet
		})

		if err != nil {
			t.Fatalf("Failed to create Safe client: %v", err)
		}

		if client.GetAddress().Hex() != safeAddress {
			t.Errorf("Expected Safe address %s, got %s", safeAddress, client.GetAddress().Hex())
		}
	})

	// Test Safe deployment check
	t.Run("CheckSafeDeployment", func(t *testing.T) {
		client, err := protocol.NewSafe(protocol.SafeConfig{
			SafeAddress: safeAddress,
			RpcURL:      rpcURL,
			ChainID:     1,
		})
		if err != nil {
			t.Fatalf("Failed to create Safe client: %v", err)
		}

		deployed, err := client.IsSafeDeployed(ctx)
		if err != nil {
			t.Errorf("Failed to check Safe deployment: %v", err)
		}

		t.Logf("Safe deployment status: %t", deployed)
	})

	// Test Safe info retrieval
	t.Run("GetSafeInfo", func(t *testing.T) {
		client, err := protocol.NewSafe(protocol.SafeConfig{
			SafeAddress: safeAddress,
			RpcURL:      rpcURL,
			ChainID:     1,
		})
		if err != nil {
			t.Fatalf("Failed to create Safe client: %v", err)
		}

		deployed, err := client.IsSafeDeployed(ctx)
		if err != nil {
			t.Fatalf("Failed to check Safe deployment: %v", err)
		}

		if !deployed {
			t.Skip("Safe not deployed, skipping info test")
		}

		info, err := client.GetSafeInfo(ctx)
		if err != nil {
			t.Errorf("Failed to get Safe info: %v", err)
			return
		}

		t.Logf("Safe info: Address=%s, Nonce=%d, Threshold=%d, Owners=%d",
			info.Address, info.Nonce, info.Threshold, len(info.Owners))

		if info.Address != safeAddress {
			t.Errorf("Expected address %s, got %s", safeAddress, info.Address)
		}
	})

	// Test transaction creation
	t.Run("CreateTransaction", func(t *testing.T) {
		client, err := protocol.NewSafe(protocol.SafeConfig{
			SafeAddress: safeAddress,
			RpcURL:      rpcURL,
			ChainID:     1,
		})
		if err != nil {
			t.Fatalf("Failed to create Safe client: %v", err)
		}

		txData := types.SafeTransactionDataPartial{
			To:    "0x1234567890123456789012345678901234567890",
			Value: "0", // No value for test
			Data:  "0x",
		}

		tx, err := client.CreateTransaction(ctx, txData)
		if err != nil {
			t.Errorf("Failed to create transaction: %v", err)
			return
		}

		if tx.Data.To != txData.To {
			t.Errorf("Expected To address %s, got %s", txData.To, tx.Data.To)
		}

		if tx.Data.Value != txData.Value {
			t.Errorf("Expected Value %s, got %s", txData.Value, tx.Data.Value)
		}

		t.Logf("Transaction created with nonce: %d", tx.Data.Nonce)
	})
}

func TestApiClientIntegration(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test - set RUN_INTEGRATION_TESTS=true to run")
	}

	apiKey := os.Getenv("SAFE_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test - SAFE_API_KEY not set")
	}

	safeAddress := os.Getenv("SAFE_ADDRESS")
	if safeAddress == "" {
		t.Skip("Skipping integration test - SAFE_ADDRESS not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test API client creation
	t.Run("CreateApiClient", func(t *testing.T) {
		client, err := api.NewSafeApiKit(api.SafeApiKitConfig{
			ChainID: 1,
			ApiKey:  apiKey,
		})

		if err != nil {
			t.Fatalf("Failed to create API client: %v", err)
		}

		if client == nil {
			t.Fatal("API client is nil")
		}
	})

	// Test Safe info from API
	t.Run("GetSafeInfoFromAPI", func(t *testing.T) {
		client, err := api.NewSafeApiKit(api.SafeApiKitConfig{
			ChainID: 1,
			ApiKey:  apiKey,
		})
		if err != nil {
			t.Fatalf("Failed to create API client: %v", err)
		}

		info, err := client.GetSafeInfo(ctx, safeAddress)
		if err != nil {
			t.Errorf("Failed to get Safe info from API: %v", err)
			return
		}

		t.Logf("API Safe info: Address=%s, Nonce=%d, Threshold=%d, Owners=%d",
			info.Address, info.Nonce, info.Threshold, len(info.Owners))

		if info.Address != safeAddress {
			t.Errorf("Expected address %s, got %s", safeAddress, info.Address)
		}
	})

	// Test transaction history
	t.Run("GetTransactionHistory", func(t *testing.T) {
		client, err := api.NewSafeApiKit(api.SafeApiKitConfig{
			ChainID: 1,
			ApiKey:  apiKey,
		})
		if err != nil {
			t.Fatalf("Failed to create API client: %v", err)
		}

		options := &api.GetMultisigTransactionsOptions{
			Limit: intPtr(5),
		}

		txs, err := client.GetMultisigTransactions(ctx, safeAddress, options)
		if err != nil {
			t.Errorf("Failed to get transaction history: %v", err)
			return
		}

		t.Logf("Found %d transactions in history (total: %d)", len(txs.Results), txs.Count)

		for i, tx := range txs.Results {
			if i >= 3 { // Only log first 3 for brevity
				break
			}
			t.Logf("  TX %d: To=%s, Value=%s, Executed=%t", i+1, tx.To, tx.Value, tx.IsExecuted)
		}
	})

	// Test pending transactions
	t.Run("GetPendingTransactions", func(t *testing.T) {
		client, err := api.NewSafeApiKit(api.SafeApiKitConfig{
			ChainID: 1,
			ApiKey:  apiKey,
		})
		if err != nil {
			t.Fatalf("Failed to create API client: %v", err)
		}

		options := &api.PendingTransactionsOptions{
			Limit: intPtr(10),
		}

		pendingTxs, err := client.GetPendingTransactions(ctx, safeAddress, options)
		if err != nil {
			t.Errorf("Failed to get pending transactions: %v", err)
			return
		}

		t.Logf("Found %d pending transactions", len(pendingTxs.Results))

		for i, tx := range pendingTxs.Results {
			if i >= 3 { // Only log first 3 for brevity
				break
			}
			t.Logf("  Pending TX %d: To=%s, Value=%s, Confirmations=%d/%d",
				i+1, tx.To, tx.Value, len(tx.Confirmations), tx.ConfirmationsRequired)
		}
	})
}

func TestSafeAddressPrediction(t *testing.T) {
	t.Run("PredictSafeAddress", func(t *testing.T) {
		config := types.SafeDeploymentConfig{
			SafeVersion: types.SafeVersion141,
			SafeSetupConfig: types.SafeSetupConfig{
				Owners: []string{
					"0x1111111111111111111111111111111111111111",
					"0x2222222222222222222222222222222222222222",
				},
				Threshold: 1,
			},
		}

		predictedAddress, err := protocol.PredictSafeAddress(config, 1) // Mainnet
		if err != nil {
			t.Errorf("Failed to predict Safe address: %v", err)
			return
		}

		t.Logf("Predicted Safe address: %s", predictedAddress)

		// Basic validation
		if len(predictedAddress) != 42 { // 0x + 40 hex chars
			t.Errorf("Expected address length 42, got %d", len(predictedAddress))
		}

		if predictedAddress[:2] != "0x" {
			t.Errorf("Expected address to start with 0x, got %s", predictedAddress[:2])
		}
	})
}

// Helper functions
func intPtr(i int) *int {
	return &i
}

// TestEnvironmentSetup helps verify that integration test environment is properly configured
func TestEnvironmentSetup(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping environment setup test - set RUN_INTEGRATION_TESTS=true to run")
	}

	requiredEnvVars := []string{
		"RPC_URL",
		"SAFE_ADDRESS",
		"SAFE_API_KEY",
	}

	for _, envVar := range requiredEnvVars {
		value := os.Getenv(envVar)
		if value == "" {
			t.Errorf("Required environment variable %s is not set", envVar)
		} else {
			// Only log first few characters for security
			displayValue := value
			if len(value) > 10 {
				displayValue = value[:10] + "..."
			}
			t.Logf("✓ %s is set: %s", envVar, displayValue)
		}
	}
}