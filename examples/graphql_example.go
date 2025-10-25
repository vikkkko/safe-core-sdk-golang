package main

import (
	"bufio"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/vikkkko/safe-core-sdk-golang/graphql"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Get GraphQL endpoint from environment or use default
	endpoint := os.Getenv("GRAPH_ENDPOINT")
	if endpoint == "" {
		endpoint = "https://api.studio.thegraph.com/query/103887/mvp/version/latest"
	}

	fmt.Println("=== GraphQL Query Interactive CLI ===")
	fmt.Printf("GraphQL Endpoint: %s\n", endpoint)
	fmt.Println()

	// Create GraphQL client
	config := graphql.Config{
		Endpoint: endpoint,
		RPCURL:   os.Getenv("RPC_URL"),
	}
	client := graphql.NewClient(config)
	defer client.Close()

	// Show menu and run queries
	for {
		showMenu()
		choice := getUserInput()

		if choice == "0" || choice == "exit" || choice == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		runQuery(client, choice)
		fmt.Println("\nPress Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func showMenu() {
	fmt.Println("\n===============================================")
	fmt.Println("   GraphQL Query Interactive CLI")
	fmt.Println("===============================================")
	fmt.Println("\nAvailable Queries:")
	fmt.Println("  1. Query payment allowances (查询付款账户授权记录)")
	fmt.Println("  2. Query transaction info by hash (通过交易哈希获取详情)")
	fmt.Println("  3. Query payment account token balances (查询付款账户代币余额)")
	fmt.Println("  4. Query collection account token balances (查询收款账户代币余额)")
	fmt.Println("  0. Exit")
	fmt.Println("===============================================")
	fmt.Print("\nEnter your choice: ")
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func prompt(label string) string {
	fmt.Printf("%s: ", label)
	return getUserInput()
}

func runQuery(client *graphql.Client, choice string) {
	fmt.Println()
	switch choice {
	case "1":
		queryPaymentAllowances(client)
	case "2":
		queryTransactionInfo(client)
	case "3":
		queryPaymentAccountBalances(client)
	case "4":
		queryCollectionAccountBalances(client)
	default:
		fmt.Println("Invalid choice.")
	}
}

// ============= Query Functions =============

func queryPaymentAllowances(client *graphql.Client) {
	fmt.Println("=== Query Payment Allowances (查询授权记录) ===")
	fmt.Println("查询哪些地址授权了指定的付款账户")
	fmt.Println()

	// Prompt for payment account address
	paymentAccountInput := prompt("Payment account address (付款账户地址)")
	if !common.IsHexAddress(paymentAccountInput) {
		log.Printf("Error: Invalid payment account address")
		return
	}

	paymentAccount := common.HexToAddress(paymentAccountInput)

	// Query allowances
	fmt.Printf("\nQuerying allowances for payment account: %s\n", paymentAccount.Hex())
	ctx := context.Background()

	// GraphQL requires lowercase addresses
	paymentAccountLower := strings.ToLower(paymentAccount.Hex())

	allowances, err := client.GetPaymentAllowances(ctx, paymentAccountLower)
	if err != nil {
		log.Printf("Failed to query payment allowances: %v", err)
		return
	}

	// Display results
	if len(allowances) == 0 {
		fmt.Printf("\n❌ No allowances found for payment account: %s\n", paymentAccount.Hex())
		fmt.Println("This means no addresses have granted approvals to this payment account.")
	} else {
		fmt.Printf("\n✅ Found %d allowance(s) for payment account: %s\n\n", len(allowances), paymentAccount.Hex())

		// Create a table header
		fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
		fmt.Printf("%-5s %-42s %-42s %s\n", "No.", "Owner (授权者)", "Token (代币)", "Amount (金额)")
		fmt.Println("═══════════════════════════════════════════════════════════════════════════════")

		for i, allowance := range allowances {
			fmt.Printf("%-5d %-42s %-42s %s\n",
				i+1,
				allowance.Owner,
				allowance.Token,
				allowance.Amount.String(),
			)
		}

		fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
		fmt.Printf("\n说明:\n")
		fmt.Printf("- Owner: 授权者的地址（拥有代币的地址）\n")
		fmt.Printf("- Token: ERC20 代币合约地址\n")
		fmt.Printf("- Amount: 授权的金额（以最小单位计，例如 wei）\n")
		fmt.Printf("\n这些记录表示哪些地址（Owner）授权了付款账户可以使用他们的代币。\n")
	}
}

func queryPaymentAccountBalances(client *graphql.Client) {
	fmt.Println("=== Query Payment Account Token Balances (付款账户代币余额) ===")
	fmt.Println("查询指定付款账户持有的代币及其余额")
	fmt.Println()

	addresses := promptAddressList("Payment account addresses (多个地址使用逗号分隔)")
	if len(addresses) == 0 {
		log.Printf("Error: please provide at least one valid payment account address")
		return
	}

	ctx := context.Background()
	accounts, err := client.GetPaymentAccounts(ctx, addresses)
	if err != nil {
		log.Printf("Failed to query payment accounts: %v", err)
		return
	}

	printPaymentAccountBalancesResult(accounts)
}

func queryCollectionAccountBalances(client *graphql.Client) {
	fmt.Println("=== Query Collection Account Token Balances (收款账户代币余额) ===")
	fmt.Println("查询指定收款账户持有的代币及其余额")
	fmt.Println()

	addresses := promptAddressList("Collection account addresses (多个地址使用逗号分隔)")
	if len(addresses) == 0 {
		log.Printf("Error: please provide at least one valid collection account address")
		return
	}

	ctx := context.Background()
	accounts, err := client.GetCollectionAccounts(ctx, addresses)
	if err != nil {
		log.Printf("Failed to query collection accounts: %v", err)
		return
	}

	printCollectionAccountBalancesResult(accounts)
}

func queryTransactionInfo(client *graphql.Client) {
	fmt.Println("=== Query Transaction Info (交易详情查询) ===")
	fmt.Println("根据交易哈希获取区块高度、Gas 使用等信息")
	fmt.Println()

	txHashInput := prompt("Transaction hash (交易哈希)")
	txHash := strings.TrimSpace(txHashInput)
	if !strings.HasPrefix(txHash, "0x") {
		txHash = "0x" + txHash
	}

	if !isValidTransactionHash(txHash) {
		log.Printf("Error: Invalid transaction hash")
		return
	}

	txHashLower := strings.ToLower(txHash)

	fmt.Printf("\nQuerying transaction info for hash: %s\n", txHashLower)
	ctx := context.Background()

	info, err := client.GetTransactionInfo(ctx, txHashLower)
	if err != nil {
		log.Printf("Failed to query transaction info: %v", err)
		return
	}

	fmt.Println("\n=== Transaction Details ===")
	fmt.Printf("Hash:           %s\n", info.ID)
	fmt.Printf("Block Number:   %s\n", info.BlockNumber)
	fmt.Printf("Status:         %s\n", formatStatus(info.Status))
	fmt.Printf("Gas Limit:      %s\n", fallbackValue(info.GasLimit))
	fmt.Printf("Gas Used:       %s\n", fallbackValue(info.GasUsed))
	fmt.Printf("Gas Price:      %s\n", info.GasPrice)
	fmt.Printf("Transaction Fee:%s\n", fallbackValue(info.TransactionFee))

	if ts, err := strconv.ParseInt(info.Timestamp, 10, 64); err == nil {
		fmt.Printf("Timestamp:      %s (%s)\n", info.Timestamp, time.Unix(ts, 0).UTC().Format("2006-01-02 15:04:05 MST"))
	} else {
		fmt.Printf("Timestamp:      %s\n", info.Timestamp)
	}
}

func isValidTransactionHash(hash string) bool {
	if !strings.HasPrefix(hash, "0x") {
		return false
	}
	trimmed := strings.TrimPrefix(hash, "0x")
	if len(trimmed) != 64 {
		return false
	}
	_, err := hex.DecodeString(trimmed)
	return err == nil
}

func formatStatus(status string) string {
	switch status {
	case "1":
		return "1 (Success)"
	case "0":
		return "0 (Failed)"
	case "":
		return "N/A"
	default:
		return status
	}
}

func fallbackValue(value string) string {
	if value == "" {
		return "N/A"
	}
	return value
}

func promptAddressList(label string) []string {
	input := prompt(label)
	if input == "" {
		return nil
	}

	parts := strings.Split(input, ",")
	addresses := make([]string, 0, len(parts))
	for _, part := range parts {
		addr := strings.TrimSpace(part)
		if addr == "" {
			continue
		}
		if !strings.HasPrefix(addr, "0x") {
			addr = "0x" + addr
		}
		if !common.IsHexAddress(addr) {
			log.Printf("Warning: skipping invalid address %s", addr)
			continue
		}
		addresses = append(addresses, addr)
	}
	return addresses
}

func printPaymentAccountBalancesResult(accounts []graphql.PaymentAccount) {
	if len(accounts) == 0 {
		fmt.Println("\n❌ No accounts found for the provided addresses.")
		return
	}

	for _, account := range accounts {
		displayAccountTokenBalances("Payment Account", account.ID, account.TokenBalances)
	}
	fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
}

func printCollectionAccountBalancesResult(accounts []graphql.CollectionAccount) {
	if len(accounts) == 0 {
		fmt.Println("\n❌ No accounts found for the provided addresses.")
		return
	}

	for _, account := range accounts {
		displayAccountTokenBalances("Collection Account", account.ID, account.TokenBalances)
	}
	fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
}

func displayAccountTokenBalances(accountType, accountID string, balances []graphql.TokenBalance) {
	fmt.Println("\n═══════════════════════════════════════════════════════════════════════════════")
	fmt.Printf("%s: %s\n", accountType, accountID)

	if len(balances) == 0 {
		fmt.Println("No token balances recorded.")
		return
	}

	fmt.Println("\nToken Balances:")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s %-20s %-10s %-32s %s\n", "Symbol", "Name", "Decimals", "Balance", "Updated At (UTC)")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")

	for _, balance := range balances {
		fmt.Printf("%-15s %-20s %-10d %-32s %s\n",
			balance.Token.Symbol,
			balance.Token.Name,
			balance.Token.Decimals,
			balance.Balance.String(),
			formatUnixTimestamp(balance.UpdatedAt),
		)
	}
}

func formatUnixTimestamp(ts string) string {
	if ts == "" {
		return "N/A"
	}
	value, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return ts
	}
	return time.Unix(value, 0).UTC().Format("2006-01-02 15:04:05 MST")
}
