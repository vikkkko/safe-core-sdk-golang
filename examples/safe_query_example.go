package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/vikkkko/safe-core-sdk-golang/api"
	"github.com/vikkkko/safe-core-sdk-golang/protocol"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	fmt.Println("=== Safe Query Examples ===\n")

	for {
		fmt.Println("\n请选择查询方式:")
		fmt.Println("1. 通过Safe地址查询")
		fmt.Println("2. 通过SafeTxHash查询")
		fmt.Println("0. 退出")
		fmt.Print("\n请输入选项: ")

		choice := prompt("")

		switch choice {
		case "1":
			querySafeByAddress()
		case "2":
			queryBySafeTxHash()
		case "0":
			fmt.Println("退出程序")
			return
		default:
			fmt.Println("无效选项，请重新选择")
		}
	}
}

func prompt(message string) string {
	reader := bufio.NewReader(os.Stdin)
	if message != "" {
		fmt.Printf("%s: ", message)
	}
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// querySafeByAddress 通过Safe地址查询Safe信息
func querySafeByAddress() {
	fmt.Println("\n=== 通过Safe地址查询 ===")

	// 输入Safe地址
	safeAddrStr := prompt("请输入Safe地址")
	if !common.IsHexAddress(safeAddrStr) {
		log.Printf("Error: 无效的地址格式: %s", safeAddrStr)
		return
	}
	safeAddress := common.HexToAddress(safeAddrStr)

	// 获取RPC URL
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("Error: RPC_URL not found in environment")
	}

	// 获取ChainID
	chainIDStr := os.Getenv("CHAIN_ID")
	if chainIDStr == "" {
		log.Fatal("Error: CHAIN_ID not found in environment")
	}
	chainIDInt, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Error: Invalid CHAIN_ID: %v", err)
	}

	// 创建API客户端
	apiConfig := api.SafeApiKitConfig{
		ChainID: chainIDInt,
		ApiKey:  os.Getenv("SAFE_API_KEY"),
	}
	if customURL := os.Getenv("SAFE_API_BASE_URL"); customURL != "" {
		apiConfig.TxServiceURL = customURL
	}

	apiClient, err := api.NewSafeApiKit(apiConfig)
	if err != nil {
		log.Fatalf("Failed to create API client: %v", err)
	}

	// 创建Safe客户端
	privateKey := os.Getenv("DEPLOYER_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("Error: DEPLOYER_PRIVATE_KEY not found in environment")
	}

	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress.Hex(),
		RpcURL:      rpcURL,
		ChainID:     chainIDInt,
		PrivateKey:  strings.TrimPrefix(privateKey, "0x"),
	})
	if err != nil {
		log.Fatalf("Failed to create Safe client: %v", err)
	}

	fmt.Println("\n正在查询Safe信息...")

	// 1. 查询Safe基本信息
	fmt.Println("\n--- Safe基本信息 ---")
	safeInfo, err := apiClient.GetSafeInfo(context.Background(), safeAddress.Hex())
	if err != nil {
		log.Printf("获取Safe信息失败: %v", err)
	} else {
		fmt.Printf("Safe地址: %s\n", safeInfo.Address)
		fmt.Printf("Nonce: %s\n", safeInfo.Nonce)
		fmt.Printf("Threshold: %d\n", safeInfo.Threshold)
		fmt.Printf("Owners数量: %d\n", len(safeInfo.Owners))
		fmt.Println("Owners列表:")
		for i, owner := range safeInfo.Owners {
			fmt.Printf("  %d. %s\n", i+1, owner)
		}
		fmt.Printf("Master Copy: %s\n", safeInfo.MasterCopy)
		fmt.Printf("Version: %s\n", safeInfo.Version)
		fmt.Printf("Fallback Handler: %s\n", safeInfo.FallbackHandler)
		fmt.Printf("Guard: %s\n", safeInfo.Guard)
	}

	// 2. 查询链上信息（需要连接RPC）
	fmt.Println("\n--- 链上信息 ---")
	threshold, err := safeClient.GetThreshold(context.Background())
	if err != nil {
		log.Printf("获取Threshold失败: %v", err)
	} else {
		fmt.Printf("链上Threshold: %d\n", threshold)
	}

	nonce, err := safeClient.GetNonce(context.Background())
	if err != nil {
		log.Printf("获取Nonce失败: %v", err)
	} else {
		fmt.Printf("链上Nonce: %d\n", nonce)
	}

	owners, err := safeClient.GetOwners(context.Background())
	if err != nil {
		log.Printf("获取Owners失败: %v", err)
	} else {
		fmt.Printf("链上Owners数量: %d\n", len(owners))
		fmt.Println("链上Owners列表:")
		for i, owner := range owners {
			fmt.Printf("  %d. %s\n", i+1, owner.Hex())
		}
	}

	// 3. 查询待处理交易
	fmt.Println("\n--- 待处理交易 ---")
	pendingTxs, err := apiClient.GetPendingTransactions(context.Background(), safeAddress.Hex(), nil)
	if err != nil {
		log.Printf("获取待处理交易失败: %v", err)
	} else {
		fmt.Printf("待处理交易数量: %d\n", pendingTxs.Count)
		if pendingTxs.Count > 0 {
			fmt.Println("最近的待处理交易:")
			for i, tx := range pendingTxs.Results {
				if i >= 5 { // 只显示前5个
					break
				}
				fmt.Printf("\n  交易 %d:\n", i+1)
				fmt.Printf("    SafeTxHash: %s\n", tx.SafeTxHash)
				fmt.Printf("    Nonce: %d\n", tx.Nonce)
				fmt.Printf("    To: %s\n", tx.To)
				fmt.Printf("    Value: %s\n", tx.Value)
				fmt.Printf("    确认数: %d/%d\n", len(tx.Confirmations), safeInfo.Threshold)
				fmt.Printf("    已执行: %v\n", tx.IsExecuted)
			}
		}
	}

	// 4. 查询所有交易
	fmt.Println("\n--- 所有交易 ---")
	allTxs, err := apiClient.GetMultisigTransactions(context.Background(), safeAddress.Hex(), nil)
	if err != nil {
		log.Printf("获取所有交易失败: %v", err)
	} else {
		fmt.Printf("总交易数量: %d\n", allTxs.Count)
		if allTxs.Count > 0 {
			fmt.Println("最近的交易:")
			for i, tx := range allTxs.Results {
				if i >= 5 { // 只显示前5个
					break
				}
				fmt.Printf("\n  交易 %d:\n", i+1)
				fmt.Printf("    SafeTxHash: %s\n", tx.SafeTxHash)
				fmt.Printf("    Nonce: %d\n", tx.Nonce)
				fmt.Printf("    已执行: %v\n", tx.IsExecuted)
				if tx.ExecutionDate != nil {
					fmt.Printf("    执行时间: %s\n", tx.ExecutionDate.Format("2006-01-02 15:04:05"))
				}
			}
		}
	}
}

// queryBySafeTxHash 通过SafeTxHash查询交易详情
func queryBySafeTxHash() {
	fmt.Println("\n=== 通过SafeTxHash查询 ===")

	// 输入SafeTxHash
	safeTxHash := prompt("请输入SafeTxHash (带或不带0x前缀)")
	safeTxHash = strings.TrimPrefix(safeTxHash, "0x")

	if len(safeTxHash) != 64 {
		log.Printf("Error: 无效的SafeTxHash格式，应为64位十六进制字符串")
		return
	}

	// 获取ChainID
	chainIDStr := os.Getenv("CHAIN_ID")
	if chainIDStr == "" {
		log.Fatal("Error: CHAIN_ID not found in environment")
	}
	chainIDInt, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Error: Invalid CHAIN_ID: %v", err)
	}

	// 创建API客户端
	apiConfig := api.SafeApiKitConfig{
		ChainID: chainIDInt,
		ApiKey:  os.Getenv("SAFE_API_KEY"),
	}
	if customURL := os.Getenv("SAFE_API_BASE_URL"); customURL != "" {
		apiConfig.TxServiceURL = customURL
	}

	apiClient, err := api.NewSafeApiKit(apiConfig)
	if err != nil {
		log.Fatalf("Failed to create API client: %v", err)
	}

	fmt.Println("\n正在查询交易详情...")

	// 查询交易详情
	tx, err := apiClient.GetMultisigTransaction(context.Background(), "0x"+safeTxHash)
	if err != nil {
		log.Fatalf("获取交易详情失败: %v", err)
	}

	// 显示交易详情
	fmt.Println("\n=== 交易详情 ===")
	fmt.Printf("SafeTxHash: %s\n", tx.SafeTxHash)
	fmt.Printf("Safe地址: %s\n", tx.Safe)
	fmt.Printf("Nonce: %d\n", tx.Nonce)
	fmt.Printf("目标地址: %s\n", tx.To)
	fmt.Printf("Value: %s Wei\n", tx.Value)
	fmt.Printf("Data: %s\n", tx.Data)
	fmt.Printf("Operation: %d\n", tx.Operation)

	fmt.Println("\n--- Gas参数 ---")
	fmt.Printf("SafeTxGas: %s\n", tx.SafeTxGas)
	fmt.Printf("BaseGas: %s\n", tx.BaseGas)
	fmt.Printf("GasPrice: %s\n", tx.GasPrice)
	fmt.Printf("GasToken: %s\n", tx.GasToken)
	fmt.Printf("RefundReceiver: %s\n", tx.RefundReceiver)

	fmt.Println("\n--- 签名信息 ---")
	fmt.Printf("确认数: %d\n", len(tx.Confirmations))
	if len(tx.Confirmations) > 0 {
		fmt.Println("签名列表:")
		for i, conf := range tx.Confirmations {
			fmt.Printf("  %d. Owner: %s\n", i+1, conf.Owner)
			fmt.Printf("     Signature: %s\n", conf.Signature)
			fmt.Printf("     签名时间: %s\n", conf.SubmissionDate.Format("2006-01-02 15:04:05"))
		}
	}

	fmt.Println("\n--- 执行状态 ---")
	fmt.Printf("已提交: %v\n", tx.IsSuccessful != nil)
	fmt.Printf("已执行: %v\n", tx.IsExecuted)
	if tx.IsExecuted {
		if tx.IsSuccessful != nil {
			fmt.Printf("执行成功: %v\n", *tx.IsSuccessful)
		}
		if tx.TransactionHash != nil && *tx.TransactionHash != "" {
			fmt.Printf("链上交易Hash: %s\n", *tx.TransactionHash)
		}
		if tx.ExecutionDate != nil {
			fmt.Printf("执行时间: %s\n", tx.ExecutionDate.Format("2006-01-02 15:04:05"))
		}
		if tx.BlockNumber != nil {
			fmt.Printf("区块号: %d\n", *tx.BlockNumber)
		}
	}

	fmt.Println("\n--- 提交信息 ---")
	if !tx.SubmissionDate.IsZero() {
		fmt.Printf("提交时间: %s\n", tx.SubmissionDate.Format("2006-01-02 15:04:05"))
	}
	if !tx.Modified.IsZero() {
		fmt.Printf("修改时间: %s\n", tx.Modified.Format("2006-01-02 15:04:05"))
	}

	// 如果交易未执行，显示还需要多少签名
	if !tx.IsExecuted {
		fmt.Println("\n--- 待处理 ---")
		// 获取Safe信息以获取threshold
		safeInfo, err := apiClient.GetSafeInfo(context.Background(), tx.Safe)
		if err == nil {
			remaining := int(safeInfo.Threshold) - len(tx.Confirmations)
			if remaining > 0 {
				fmt.Printf("⚠️  还需要 %d 个签名才能执行\n", remaining)
			} else {
				fmt.Println("✅ 已收集足够签名，可以执行")
			}
		}
	}
}
