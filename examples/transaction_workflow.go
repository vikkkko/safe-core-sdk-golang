//go:build ignore
// +build ignore

package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/vikkkko/safe-core-sdk-golang/api"
	"github.com/vikkkko/safe-core-sdk-golang/protocol"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
	"github.com/vikkkko/safe-core-sdk-golang/types"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	fmt.Println("🔐 Safe交易工作流程演示")
	fmt.Println("======================")

	// 从环境变量读取配置
	safeAddress := os.Getenv("SAFE_ADDRESS")
	if safeAddress == "" {
		log.Fatal("SAFE_ADDRESS not set in .env")
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("RPC_URL not set in .env")
	}

	chainIDStr := os.Getenv("CHAIN_ID")
	chainID, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil || chainID == 0 {
		chainID = 11155111 // 默认 Sepolia
		log.Printf("Using default CHAIN_ID: %d", chainID)
	}

	apiKey := os.Getenv("SAFE_API_KEY")

	testPrivateKey := os.Getenv("OWNER_PRIVATE_KEY")
	if testPrivateKey == "" {
		log.Fatal("OWNER_PRIVATE_KEY not set in .env")
	}

	ctx := context.Background()

	// 第一步：创建客户端
	fmt.Printf("🔧 创建客户端...")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress,
		RpcURL:      rpcURL,
		ChainID:     chainID,
	})
	if err != nil {
		log.Fatalf("创建Safe客户端失败: %v", err)
	}

	apiClient, err := api.NewSafeApiKit(api.SafeApiKitConfig{
		ChainID: chainID,
		ApiKey:  apiKey,
	})
	if err != nil {
		log.Fatalf("创建API客户端失败: %v", err)
	}
	fmt.Printf(" ✅\n")

	// 第二步：获取Safe信息
	fmt.Printf("📊 获取Safe信息...")
	safeInfo, err := apiClient.GetSafeInfo(ctx, safeAddress)
	if err != nil {
		log.Fatalf("获取Safe信息失败: %v", err)
	}

	currentNonce, err := strconv.ParseUint(safeInfo.Nonce, 10, 64)
	if err != nil {
		log.Fatalf("解析随机数失败: %v", err)
	}
	fmt.Printf(" ✅ (阈值: %d/%d, 随机数: %d)\n", safeInfo.Threshold, len(safeInfo.Owners), currentNonce)

	// 第三步：创建ERC20转账交易
	fmt.Printf("📝 创建USDT转账交易...")
	usdtAddress := "0xAD2B0439ed98F50eDEB0e04F064d492bAFDAd73B"      // Sepolia USDC
	recipientAddress := "0x9C126aa4Eb6D110D646139969774F2c5b64dD279" // 接收地址
	transferAmount := big.NewInt(1000000)                            // 1 USDC (6位小数)

	// 创建ERC20转账数据: transfer(address to, uint256 amount)
	transferData, err := utils.CreateERC20TransferData(recipientAddress, transferAmount)
	if err != nil {
		log.Fatalf("创建ERC20转账数据失败: %v", err)
	}
	fmt.Printf(" ✅\n")

	txData := types.SafeTransactionDataPartial{
		To:    usdtAddress,                             // USDC合约地址
		Value: "0",                                     // ERC20转账无需ETH
		Data:  "0x" + hex.EncodeToString(transferData), // ERC20转账调用数据
		Nonce: &currentNonce,                           // 使用当前随机数
	}

	transaction, err := safeClient.CreateTransaction(ctx, txData)
	if err != nil {
		log.Fatalf("创建交易失败: %v", err)
	}

	// 第四步：从Safe合约获取交易哈希（确保与链上一致）
	fmt.Printf("🔐 从Safe合约获取交易哈希...")

	// 解析交易数据字段
	value := new(big.Int)
	value.SetString(transaction.Data.Value, 10)

	safeTxGas := new(big.Int)
	safeTxGas.SetString(transaction.Data.SafeTxGas, 10)

	baseGas := new(big.Int)
	baseGas.SetString(transaction.Data.BaseGas, 10)

	gasPrice := new(big.Int)
	gasPrice.SetString(transaction.Data.GasPrice, 10)

	txHashBytes, err := safeClient.GetTransactionHash(
		ctx,
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
		log.Fatalf("获取交易哈希失败: %v", err)
	}
	txHash := txHashBytes[:]
	fmt.Printf(" ✅\n")

	// 第五步：签名交易
	fmt.Printf("✍️  签名交易...")
	privateKey, err := crypto.HexToECDSA(testPrivateKey)
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}

	signerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 检查签名者是否为所有者
	isOwner := false
	for _, owner := range safeInfo.Owners {
		if common.HexToAddress(owner) == signerAddress {
			isOwner = true
			break
		}
	}

	signature, err := utils.SignMessage(txHash, privateKey)
	if err != nil {
		log.Fatalf("签名交易失败: %v", err)
	}

	if !isOwner {
		fmt.Printf(" ⚠️  (签名者非所有者)\n")
	} else {
		fmt.Printf(" ✅\n")
	}

	// 第六步：提交到Safe服务
	fmt.Printf("📤 提交交易到Safe服务...")

	// 计算Safe交易哈希用于服务
	safeTxHash := hex.EncodeToString(txHash)

	// 准备提案
	proposal := api.ProposeTransactionProps{
		SafeAddress:             safeAddress,
		SafeTxHash:              "0x" + safeTxHash,
		To:                      transaction.Data.To,
		Value:                   transaction.Data.Value,
		Data:                    transaction.Data.Data,
		Operation:               int(transaction.Data.Operation),
		GasToken:                transaction.Data.GasToken,
		SafeTxGas:               0, // 将自动估算
		BaseGas:                 0, // 将自动估算
		GasPrice:                transaction.Data.GasPrice,
		RefundReceiver:          transaction.Data.RefundReceiver,
		Nonce:                   int64(transaction.Data.Nonce),
		Sender:                  signerAddress.Hex(),
		Signature:               "0x" + hex.EncodeToString(signature),
		ContractTransactionHash: "0x" + safeTxHash,
	}

	if isOwner {
		response, err := apiClient.ProposeTransaction(ctx, proposal)
		if err != nil {
			fmt.Printf(" ❌\n❌ 提交失败: %v\n", err)
		} else {
			fmt.Printf(" ✅\n✅ 交易提案已提交!\n")
			// 使用响应中的哈希，如果为空则使用我们计算的哈希
			displayHash := response.SafeTxHash
			if displayHash == "" {
				displayHash = "0x" + safeTxHash
			}
			fmt.Printf("🔗 SAFE Transaction Hash: %s\n", displayHash)
			fmt.Printf("📊 需要确认: %d/%d\n", len(response.Confirmations), safeInfo.Threshold)
		}
	} else {
		fmt.Printf(" ⏭️  (跳过提交，签名者非所有者)\n")
	}

	fmt.Println("\n✅ Safe交易工作流程演示完成!")
}
