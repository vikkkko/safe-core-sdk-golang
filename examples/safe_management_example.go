package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/vikkkko/safe-core-sdk-golang/api"
	"github.com/vikkkko/safe-core-sdk-golang/protocol"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
	safetypes "github.com/vikkkko/safe-core-sdk-golang/types"
)

// SafeManagementContext holds necessary data for Safe management operations
type SafeManagementContext struct {
	Client        *ethclient.Client
	RPCURL        string
	PrivateKey    *ecdsa.PrivateKey
	PrivateKeyHex string
	FromAddress   common.Address
	ChainID       *big.Int
	SafeAPIKey    string
	SafeAPIURL    string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	ctx := initializeContext()
	if ctx == nil {
		log.Fatal("Failed to initialize context")
	}

	for {
		showMenu()
		choice := prompt("Enter your choice")

		switch choice {
		case "1":
			addOwnerWithThreshold(ctx)
		case "2":
			removeOwnerWithThreshold(ctx)
		case "3":
			swapOwner(ctx)
		case "4":
			changeThreshold(ctx)
		case "5":
			querySafeInfo(ctx)
		case "6":
			confirmSafeTransaction(ctx)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func initializeContext() *SafeManagementContext {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("RPC_URL not set in .env")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	chainIDStr := os.Getenv("CHAIN_ID")
	chainID, ok := new(big.Int).SetString(chainIDStr, 10)
	if !ok {
		log.Fatal("Invalid CHAIN_ID in .env")
	}

	privateKeyHex := os.Getenv("DEPLOYER_PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatal("DEPLOYER_PRIVATE_KEY not set in .env")
	}

	cleanKey := strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(cleanKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	safeAPIKey := os.Getenv("SAFE_API_KEY")
	safeAPIURL := os.Getenv("SAFE_API_BASE_URL")

	fmt.Printf("Connected to %s (Chain ID: %s)\n", rpcURL, chainID.String())
	fmt.Printf("Default signer address: %s\n", fromAddress.Hex())

	return &SafeManagementContext{
		Client:        client,
		RPCURL:        rpcURL,
		PrivateKey:    privateKey,
		PrivateKeyHex: cleanKey,
		FromAddress:   fromAddress,
		ChainID:       chainID,
		SafeAPIKey:    safeAPIKey,
		SafeAPIURL:    safeAPIURL,
	}
}

func showMenu() {
	fmt.Println("\n=== Safe Multisig Management ===")
	fmt.Println("1. Add Owner (with threshold)")
	fmt.Println("2. Remove Owner (with threshold)")
	fmt.Println("3. Swap Owner (replace one owner with another)")
	fmt.Println("4. Change Threshold")
	fmt.Println("5. Query Safe Info (owners, threshold)")
	fmt.Println("6. Confirm Safe Transaction")
	fmt.Println("0. Exit")
	fmt.Println()
}

func prompt(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", message)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func promptAddress(message, defaultValue string) common.Address {
	input := prompt(fmt.Sprintf("%s [%s]", message, defaultValue))
	if input == "" {
		input = defaultValue
	}
	if !common.IsHexAddress(input) {
		log.Printf("Warning: Invalid address format: %s, using zero address", input)
		return common.Address{}
	}
	return common.HexToAddress(input)
}

func confirmSend() bool {
	response := strings.ToLower(prompt("Confirm? (yes/no) [no]"))
	return response == "yes" || response == "y"
}

// addOwnerWithThreshold adds a new owner to the Safe with an optional threshold change
func addOwnerWithThreshold(ctx *SafeManagementContext) {
	fmt.Println("\n=== Add Owner with Threshold ===")

	safeAddress := promptAddress("Safe address", "")
	if safeAddress == (common.Address{}) {
		log.Printf("Error: Valid Safe address is required")
		return
	}

	newOwner := promptAddress("New owner address", "")
	if newOwner == (common.Address{}) {
		log.Printf("Error: Valid owner address is required")
		return
	}

	thresholdStr := prompt("New threshold [default:0]")
	var threshold *big.Int
	if thresholdStr == "" {
		threshold = big.NewInt(0) // 0 means keep current threshold
	} else {
		parsed, err := strconv.ParseInt(thresholdStr, 10, 64)
		if err != nil || parsed < 0 {
			log.Printf("Error: Invalid threshold")
			return
		}
		threshold = big.NewInt(parsed)
	}

	// Generate calldata for addOwnerWithThreshold
	calldata, err := utils.SafeAddOwnerWithThresholdData(newOwner, threshold)
	if err != nil {
		log.Printf("Error generating calldata: %v", err)
		return
	}

	fmt.Printf("\n=== Transaction Details ===\n")
	fmt.Printf("Safe: %s\n", safeAddress.Hex())
	fmt.Printf("New Owner: %s\n", newOwner.Hex())
	fmt.Printf("Threshold: %s (0 = keep current)\n", threshold.String())
	fmt.Printf("Calldata: 0x%x\n", calldata)
	fmt.Printf("Calldata length: %d bytes\n\n", len(calldata))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Create Safe transaction through multisig
	proposeSafeTransaction(ctx, safeAddress, safeAddress, calldata)
}

// removeOwnerWithThreshold removes an owner from the Safe
func removeOwnerWithThreshold(ctx *SafeManagementContext) {
	fmt.Println("\n=== Remove Owner with Threshold ===")

	safeAddress := promptAddress("Safe address", "")
	if safeAddress == (common.Address{}) {
		log.Printf("Error: Valid Safe address is required")
		return
	}

	// First, get current owners to help user identify prevOwner
	fmt.Println("\n📋 Fetching current owners...")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress.Hex(),
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  ctx.PrivateKeyHex,
	})
	if err != nil {
		log.Printf("Error creating Safe client: %v", err)
		return
	}

	owners, err := safeClient.GetOwners(context.Background())
	if err != nil {
		log.Printf("Error getting owners: %v", err)
		return
	}

	fmt.Printf("\nCurrent owners (%d):\n", len(owners))
	for i, owner := range owners {
		fmt.Printf("  %d. %s\n", i+1, owner.Hex())
	}

	ownerToRemove := promptAddress("\nOwner to remove", "")
	if ownerToRemove == (common.Address{}) {
		log.Printf("Error: Valid owner address is required")
		return
	}

	// Find prevOwner in the linked list
	var prevOwner common.Address
	sentinel := common.HexToAddress("0x0000000000000000000000000000000000000001")

	found := false
	for i, owner := range owners {
		if strings.EqualFold(owner.Hex(), ownerToRemove.Hex()) {
			if i == 0 {
				prevOwner = sentinel
			} else {
				prevOwner = owners[i-1]
			}
			found = true
			break
		}
	}

	if !found {
		log.Printf("Error: Address %s is not an owner of this Safe", ownerToRemove.Hex())
		return
	}

	fmt.Printf("Previous owner in list: %s\n", prevOwner.Hex())

	thresholdStr := prompt("New threshold")
	threshold, err := strconv.ParseInt(thresholdStr, 10, 64)
	if err != nil || threshold < 1 {
		log.Printf("Error: Invalid threshold (must be >= 1)")
		return
	}

	// Generate calldata for removeOwner
	calldata, err := utils.SafeRemoveOwnerData(prevOwner, ownerToRemove, big.NewInt(threshold))
	if err != nil {
		log.Printf("Error generating calldata: %v", err)
		return
	}

	fmt.Printf("\n=== Transaction Details ===\n")
	fmt.Printf("Safe: %s\n", safeAddress.Hex())
	fmt.Printf("Previous Owner: %s\n", prevOwner.Hex())
	fmt.Printf("Owner to Remove: %s\n", ownerToRemove.Hex())
	fmt.Printf("New Threshold: %d\n", threshold)
	fmt.Printf("Calldata: 0x%x\n", calldata)
	fmt.Printf("Calldata length: %d bytes\n\n", len(calldata))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Create Safe transaction through multisig
	proposeSafeTransaction(ctx, safeAddress, safeAddress, calldata)
}

// swapOwner replaces one owner with another
func swapOwner(ctx *SafeManagementContext) {
	fmt.Println("\n=== Swap Owner ===")

	safeAddress := promptAddress("Safe address", "")
	if safeAddress == (common.Address{}) {
		log.Printf("Error: Valid Safe address is required")
		return
	}

	// Get current owners
	fmt.Println("\n📋 Fetching current owners...")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress.Hex(),
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  ctx.PrivateKeyHex,
	})
	if err != nil {
		log.Printf("Error creating Safe client: %v", err)
		return
	}

	owners, err := safeClient.GetOwners(context.Background())
	if err != nil {
		log.Printf("Error getting owners: %v", err)
		return
	}

	fmt.Printf("\nCurrent owners (%d):\n", len(owners))
	for i, owner := range owners {
		fmt.Printf("  %d. %s\n", i+1, owner.Hex())
	}

	oldOwner := promptAddress("\nOld owner (to replace)", "")
	if oldOwner == (common.Address{}) {
		log.Printf("Error: Valid owner address is required")
		return
	}

	newOwner := promptAddress("New owner (replacement)", "")
	if newOwner == (common.Address{}) {
		log.Printf("Error: Valid owner address is required")
		return
	}

	// Find prevOwner in the linked list
	var prevOwner common.Address
	sentinel := common.HexToAddress("0x0000000000000000000000000000000000000001")

	found := false
	for i, owner := range owners {
		if strings.EqualFold(owner.Hex(), oldOwner.Hex()) {
			if i == 0 {
				prevOwner = sentinel
			} else {
				prevOwner = owners[i-1]
			}
			found = true
			break
		}
	}

	if !found {
		log.Printf("Error: Address %s is not an owner of this Safe", oldOwner.Hex())
		return
	}

	fmt.Printf("Previous owner in list: %s\n", prevOwner.Hex())

	// Generate calldata for swapOwner
	calldata, err := utils.SafeSwapOwnerData(prevOwner, oldOwner, newOwner)
	if err != nil {
		log.Printf("Error generating calldata: %v", err)
		return
	}

	fmt.Printf("\n=== Transaction Details ===\n")
	fmt.Printf("Safe: %s\n", safeAddress.Hex())
	fmt.Printf("Previous Owner: %s\n", prevOwner.Hex())
	fmt.Printf("Old Owner: %s\n", oldOwner.Hex())
	fmt.Printf("New Owner: %s\n", newOwner.Hex())
	fmt.Printf("Calldata: 0x%x\n", calldata)
	fmt.Printf("Calldata length: %d bytes\n\n", len(calldata))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Create Safe transaction through multisig
	proposeSafeTransaction(ctx, safeAddress, safeAddress, calldata)
}

// changeThreshold changes the threshold of the Safe
func changeThreshold(ctx *SafeManagementContext) {
	fmt.Println("\n=== Change Threshold ===")

	safeAddress := promptAddress("Safe address", "")
	if safeAddress == (common.Address{}) {
		log.Printf("Error: Valid Safe address is required")
		return
	}

	// Get current threshold
	fmt.Println("\n📋 Fetching current Safe info...")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress.Hex(),
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  ctx.PrivateKeyHex,
	})
	if err != nil {
		log.Printf("Error creating Safe client: %v", err)
		return
	}

	currentThreshold, err := safeClient.GetThreshold(context.Background())
	if err != nil {
		log.Printf("Error getting threshold: %v", err)
		return
	}

	owners, err := safeClient.GetOwners(context.Background())
	if err != nil {
		log.Printf("Error getting owners: %v", err)
		return
	}

	fmt.Printf("\nCurrent threshold: %d\n", currentThreshold)
	fmt.Printf("Number of owners: %d\n", len(owners))

	thresholdStr := prompt("New threshold")
	threshold, err := strconv.ParseInt(thresholdStr, 10, 64)
	if err != nil || threshold < 1 || threshold > int64(len(owners)) {
		log.Printf("Error: Invalid threshold (must be between 1 and %d)", len(owners))
		return
	}

	// Generate calldata for changeThreshold
	calldata, err := utils.SafeChangeThresholdData(big.NewInt(threshold))
	if err != nil {
		log.Printf("Error generating calldata: %v", err)
		return
	}

	fmt.Printf("\n=== Transaction Details ===\n")
	fmt.Printf("Safe: %s\n", safeAddress.Hex())
	fmt.Printf("Current Threshold: %d\n", currentThreshold)
	fmt.Printf("New Threshold: %d\n", threshold)
	fmt.Printf("Calldata: 0x%x\n", calldata)
	fmt.Printf("Calldata length: %d bytes\n\n", len(calldata))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Create Safe transaction through multisig
	proposeSafeTransaction(ctx, safeAddress, safeAddress, calldata)
}

// querySafeInfo displays information about a Safe
func querySafeInfo(ctx *SafeManagementContext) {
	fmt.Println("\n=== Query Safe Info ===")

	safeAddress := promptAddress("Safe address", "")
	if safeAddress == (common.Address{}) {
		log.Printf("Error: Valid Safe address is required")
		return
	}

	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress.Hex(),
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  ctx.PrivateKeyHex,
	})
	if err != nil {
		log.Printf("Error creating Safe client: %v", err)
		return
	}

	// Get owners
	owners, err := safeClient.GetOwners(context.Background())
	if err != nil {
		log.Printf("Error getting owners: %v", err)
		return
	}

	// Get threshold
	threshold, err := safeClient.GetThreshold(context.Background())
	if err != nil {
		log.Printf("Error getting threshold: %v", err)
		return
	}

	// Get nonce
	nonce, err := safeClient.GetNonce(context.Background())
	if err != nil {
		log.Printf("Error getting nonce: %v", err)
		return
	}

	fmt.Printf("\n=== Safe Information ===\n")
	fmt.Printf("Address: %s\n", safeAddress.Hex())
	fmt.Printf("Threshold: %d\n", threshold)
	fmt.Printf("Nonce: %d\n", nonce)
	fmt.Printf("Owners (%d):\n", len(owners))
	for i, owner := range owners {
		fmt.Printf("  %d. %s\n", i+1, owner.Hex())
	}
}

// proposeSafeTransaction creates and proposes a Safe transaction
func proposeSafeTransaction(ctx *SafeManagementContext, safeAddress, targetAddress common.Address, calldata []byte) {
	// Create Safe client
	fmt.Printf("\n🔧 创建Safe客户端...")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress.Hex(),
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  ctx.PrivateKeyHex,
	})
	if err != nil {
		log.Printf("创建Safe客户端失败: %v", err)
		return
	}

	// Create API client
	apiConfig := api.SafeApiKitConfig{
		ChainID: ctx.ChainID.Int64(),
		ApiKey:  ctx.SafeAPIKey,
	}
	if ctx.SafeAPIURL != "" {
		apiConfig.TxServiceURL = ctx.SafeAPIURL
	}
	apiClient, err := api.NewSafeApiKit(apiConfig)
	if err != nil {
		log.Printf("创建API客户端失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// Get Safe info
	fmt.Printf("📊 获取Safe信息...")
	safeInfo, err := apiClient.GetSafeInfo(context.Background(), safeAddress.Hex())
	if err != nil {
		log.Printf("获取Safe信息失败: %v", err)
		return
	}

	currentNonce, err := strconv.ParseUint(safeInfo.Nonce, 10, 64)
	if err != nil {
		log.Printf("解析随机数失败: %v", err)
		return
	}
	fmt.Printf(" ✅ (阈值: %d/%d, 随机数: %d)\n", safeInfo.Threshold, len(safeInfo.Owners), currentNonce)

	// Create Safe transaction
	fmt.Printf("📋 创建Safe交易...")
	channel := uint64(0)
	txData := safetypes.SafeTransactionDataPartial{
		To:      targetAddress.Hex(),
		Value:   "0",
		Data:    "0x" + hex.EncodeToString(calldata),
		Nonce:   &currentNonce,
		Channel: &channel,
	}

	transaction, err := safeClient.CreateTransaction(context.Background(), txData)
	if err != nil {
		log.Printf("创建交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// Get transaction hash
	fmt.Printf("🔐 获取Safe交易哈希...")
	value := new(big.Int)
	value.SetString(transaction.Data.Value, 10)

	safeTxGas := new(big.Int)
	safeTxGas.SetString(transaction.Data.SafeTxGas, 10)

	baseGas := new(big.Int)
	baseGas.SetString(transaction.Data.BaseGas, 10)

	gasPrice := new(big.Int)
	gasPrice.SetString(transaction.Data.GasPrice, 10)

	txHashBytes, err := safeClient.GetTransactionHash(
		context.Background(),
		transaction.Data.Channel,
		common.HexToAddress(transaction.Data.To),
		value,
		common.FromHex(transaction.Data.Data),
		uint8(transaction.Data.Operation),
		safeTxGas,
		baseGas,
		gasPrice,
		common.HexToAddress(transaction.Data.GasToken),
		common.HexToAddress(transaction.Data.RefundReceiver),
		new(big.Int).SetUint64(transaction.Data.Nonce),
	)
	if err != nil {
		log.Printf("获取交易哈希失败: %v", err)
		return
	}
	txHash := txHashBytes[:]
	safeTxHash := hex.EncodeToString(txHash)
	fmt.Printf(" ✅\n   交易哈希: 0x%s\n", safeTxHash)

	// Get first owner's private key for signing
	ownerKeyHex := os.Getenv("OWNER_PRIVATE_KEY")
	if ownerKeyHex == "" {
		log.Printf("Error: OWNER_PRIVATE_KEY not set in .env")
		return
	}

	ownerPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(ownerKeyHex, "0x"))
	if err != nil {
		log.Printf("解析 OWNER_PRIVATE_KEY 失败: %v", err)
		return
	}
	ownerAddress := crypto.PubkeyToAddress(ownerPrivateKey.PublicKey)

	// Sign transaction
	fmt.Printf("\n✍️  签名交易...")
	signature, err := utils.SignMessage(txHash, ownerPrivateKey)
	if err != nil {
		log.Printf("签名交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n   签名者: %s\n", ownerAddress.Hex())

	// Submit to Safe service
	fmt.Printf("\n📤 提交交易到Safe服务...")
	proposal := api.ProposeTransactionProps{
		SafeAddress:             safeAddress.Hex(),
		SafeTxHash:              "0x" + safeTxHash,
		To:                      transaction.Data.To,
		Value:                   transaction.Data.Value,
		Data:                    transaction.Data.Data,
		Operation:               int(transaction.Data.Operation),
		GasToken:                transaction.Data.GasToken,
		SafeTxGas:               0,
		BaseGas:                 0,
		GasPrice:                transaction.Data.GasPrice,
		RefundReceiver:          transaction.Data.RefundReceiver,
		Nonce:                   int64(transaction.Data.Nonce),
		Sender:                  ownerAddress.Hex(),
		Signature:               "0x" + hex.EncodeToString(signature),
		ContractTransactionHash: "0x" + safeTxHash,
	}

	_, err = apiClient.ProposeTransaction(context.Background(), proposal)
	if err != nil {
		log.Printf("提交失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")
	fmt.Printf("\n✅ 交易已提交到Safe服务\n")
	fmt.Printf("交易哈希: 0x%s\n", safeTxHash)

	// Ask if user wants to auto-confirm and execute
	fmt.Println("\n=== 自动执行选项 ===")
	autoConfirmChoice := prompt("是否使用 ConfirmTransaction SDK 方法自动收集签名并执行? (yes/no) [no]")

	if strings.ToLower(autoConfirmChoice) == "yes" || strings.ToLower(autoConfirmChoice) == "y" {
		fmt.Printf("\n🔐 使用SDK的ConfirmTransaction方法...\n")

		// Use the SDK's ConfirmTransaction method with auto-execute enabled
		result, err := safeClient.ConfirmTransaction(context.Background(), protocol.ConfirmTransactionConfig{
			SafeTxHash:  "0x" + safeTxHash,
			APIClient:   apiClient,
			AutoExecute: true, // 自动执行
		})
		if err != nil {
			log.Printf("确认失败: %v", err)
			fmt.Printf("请手动使用选项6或Safe UI来收集其他签名并执行交易\n")
			return
		}

		// Display result
		fmt.Printf("\n=== 确认结果 ===\n")
		fmt.Printf("已签名: %v\n", result.AlreadySigned)
		fmt.Printf("提交了新签名: %v\n", result.SignatureSubmitted)
		fmt.Printf("签名数量: %d/%d\n", result.CurrentSignatures, result.RequiredSignatures)
		fmt.Printf("达到阈值: %v\n", result.ThresholdMet)
		fmt.Printf("已执行: %v\n", result.TransactionExecuted)

		if result.TransactionExecuted && result.ExecutionResult != nil {
			fmt.Printf("\n✅ 交易已成功执行并上链!\n")
			fmt.Printf("链上交易哈希: %s\n", result.ExecutionResult.Hash)
		} else if result.ThresholdMet {
			fmt.Printf("\n⚠️  已达到签名阈值，但未自动执行\n")
			fmt.Printf("请检查Safe配置或手动执行\n")
		} else {
			fmt.Printf("\n⏳ 等待更多签名 (还需要 %d 个)\n", result.RequiredSignatures-result.CurrentSignatures)
			fmt.Printf("请使用选项6或Safe UI来收集其他签名并执行交易\n")
		}
	} else {
		fmt.Printf("请使用选项6或Safe UI来收集其他签名并执行交易\n")
	}
}

// confirmSafeTransaction confirms an existing Safe transaction
func confirmSafeTransaction(ctx *SafeManagementContext) {
	fmt.Println("\n=== Confirm Safe Transaction ===")

	safeTxHash := prompt("Safe transaction hash (0x...)")
	if safeTxHash == "" {
		log.Printf("Error: Safe transaction hash is required")
		return
	}

	// Ensure 0x prefix
	if !strings.HasPrefix(safeTxHash, "0x") && !strings.HasPrefix(safeTxHash, "0X") {
		safeTxHash = "0x" + safeTxHash
	}

	// Select which private key to use for signing
	fmt.Println("\n=== 选择签名私钥 ===")
	fmt.Println("1. DEPLOYER_PRIVATE_KEY")
	fmt.Println("2. OWNER_PRIVATE_KEY")
	fmt.Println("3. OWNER2_PRIVATE_KEY")
	fmt.Println("4. OWNER3_PRIVATE_KEY")
	keyChoice := prompt("选择私钥 [2]")
	if keyChoice == "" {
		keyChoice = "2"
	}

	var selectedPrivateKey string
	var keyLabel string
	switch keyChoice {
	case "1":
		selectedPrivateKey = os.Getenv("DEPLOYER_PRIVATE_KEY")
		keyLabel = "DEPLOYER_PRIVATE_KEY"
	case "2":
		selectedPrivateKey = os.Getenv("OWNER_PRIVATE_KEY")
		keyLabel = "OWNER_PRIVATE_KEY"
	case "3":
		selectedPrivateKey = os.Getenv("OWNER2_PRIVATE_KEY")
		keyLabel = "OWNER2_PRIVATE_KEY"
	case "4":
		selectedPrivateKey = os.Getenv("OWNER3_PRIVATE_KEY")
		keyLabel = "OWNER3_PRIVATE_KEY"
	default:
		selectedPrivateKey = os.Getenv("OWNER_PRIVATE_KEY")
		keyLabel = "OWNER_PRIVATE_KEY"
	}

	if selectedPrivateKey == "" {
		log.Printf("Error: %s not set in .env", keyLabel)
		return
	}

	fmt.Printf("使用私钥: %s\n", keyLabel)

	// Create API client
	apiConfig := api.SafeApiKitConfig{
		ChainID: ctx.ChainID.Int64(),
		ApiKey:  ctx.SafeAPIKey,
	}
	if ctx.SafeAPIURL != "" {
		apiConfig.TxServiceURL = ctx.SafeAPIURL
	}
	apiClient, err := api.NewSafeApiKit(apiConfig)
	if err != nil {
		log.Printf("创建API客户端失败: %v", err)
		return
	}

	// Fetch transaction details first to get Safe address
	fmt.Printf("\n📋 获取交易详情...")
	txDetails, err := apiClient.GetMultisigTransaction(context.Background(), safeTxHash)
	if err != nil {
		log.Printf("获取失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	safeAddress := txDetails.Safe

	// Display transaction info
	fmt.Printf("\n=== 交易信息 ===\n")
	fmt.Printf("Safe地址: %s\n", safeAddress)
	fmt.Printf("目标地址: %s\n", txDetails.To)
	fmt.Printf("数据: %s\n", txDetails.Data)
	fmt.Printf("当前签名: %d/%d\n", len(txDetails.Confirmations), txDetails.ConfirmationsRequired)
	fmt.Printf("已执行: %v\n\n", txDetails.IsExecuted)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Create Safe client with selected private key
	cleanKey := strings.TrimPrefix(selectedPrivateKey, "0x")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress,
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  cleanKey,
	})
	if err != nil {
		log.Printf("创建Safe客户端失败: %v", err)
		return
	}

	// Use SDK's ConfirmTransaction method
	fmt.Printf("\n🔐 确认交易...\n")
	autoExecute := true
	result, err := safeClient.ConfirmTransaction(context.Background(), protocol.ConfirmTransactionConfig{
		SafeTxHash:  safeTxHash,
		APIClient:   apiClient,
		AutoExecute: autoExecute,
	})
	if err != nil {
		log.Printf("确认失败: %v", err)
		return
	}

	// Display result
	fmt.Printf("\n=== 确认结果 ===\n")
	fmt.Printf("已签名: %v\n", result.AlreadySigned)
	fmt.Printf("提交了新签名: %v\n", result.SignatureSubmitted)
	fmt.Printf("签名数量: %d/%d\n", result.CurrentSignatures, result.RequiredSignatures)
	fmt.Printf("达到阈值: %v\n", result.ThresholdMet)
	fmt.Printf("已执行: %v\n", result.TransactionExecuted)

	if result.TransactionExecuted && result.ExecutionResult != nil {
		fmt.Printf("\n✅ 交易已执行!\n")
		fmt.Printf("交易哈希: %s\n", result.ExecutionResult.Hash)
	} else if result.ThresholdMet {
		fmt.Printf("\n⚠️  已达到签名阈值，但未自动执行\n")
	} else {
		fmt.Printf("\n⏳ 等待更多签名 (还需要 %d 个)\n", result.RequiredSignatures-result.CurrentSignatures)
	}
}
