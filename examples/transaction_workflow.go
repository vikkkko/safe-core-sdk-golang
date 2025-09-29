package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/yinwei/safe-core-sdk-golang/api"
	"github.com/yinwei/safe-core-sdk-golang/protocol"
	"github.com/yinwei/safe-core-sdk-golang/protocol/utils"
	"github.com/yinwei/safe-core-sdk-golang/types"
)

func main() {
	fmt.Println("🔐 Safe交易工作流程演示")
	fmt.Println("======================")

	// 配置参数
	safeAddress := "0x447d4227d88D6A7fB1486879be24Be00418A5fB7"
	rpcURL := ""
	chainID := int64(11155111) // Sepolia测试网
	apiKey := ""
	// 警告：生产环境中绝不要使用真实私钥！
	// 这里仅用于演示目的
	testPrivateKey := "" // 测试私钥

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
	fmt.Printf("📝 创建USDC转账交易...")
	usdcAddress := "0xEDC9b422dC055939F63e9Dc808ACEc05B515C28e"      // Sepolia USDC
	recipientAddress := "0x9C126aa4Eb6D110D646139969774F2c5b64dD279" // 接收地址
	transferAmount := big.NewInt(1000000)                            // 1 USDC (6位小数)

	// 创建ERC20转账数据: transfer(address to, uint256 amount)
	transferData, err := utils.CreateERC20TransferData(recipientAddress, transferAmount)
	if err != nil {
		log.Fatalf("创建ERC20转账数据失败: %v", err)
	}
	fmt.Printf(" ✅\n")

	txData := types.SafeTransactionDataPartial{
		To:    usdcAddress,                             // USDC合约地址
		Value: "0",                                     // ERC20转账无需ETH
		Data:  "0x" + hex.EncodeToString(transferData), // ERC20转账调用数据
		Nonce: &currentNonce,                           // 使用当前随机数
	}

	transaction, err := safeClient.CreateTransaction(ctx, txData)
	if err != nil {
		log.Fatalf("创建交易失败: %v", err)
	}

	// 第四步：计算交易哈希
	fmt.Printf("🔐 计算交易哈希...")
	txHash, err := calculateSafeTransactionHash(transaction.Data, safeAddress, chainID)
	if err != nil {
		log.Fatalf("计算交易哈希失败: %v", err)
	}
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
			fmt.Printf("🔗 交易哈希: %s\n", response.SafeTxHash)
			fmt.Printf("📊 需要确认: %d/%d\n", len(response.Confirmations), safeInfo.Threshold)
		}
	} else {
		fmt.Printf(" ⏭️  (跳过提交，签名者非所有者)\n")
	}

	fmt.Println("\n✅ Safe交易工作流程演示完成!")
}

// calculateSafeTransactionHash 计算需要签名的交易哈希
func calculateSafeTransactionHash(txData types.SafeTransactionData, safeAddress string, chainID int64) ([]byte, error) {
	// 这是一个简化实现
	// 实际实现中应该使用Safe合约的getTransactionHash方法

	to := common.HexToAddress(txData.To)
	value, _ := new(big.Int).SetString(txData.Value, 10)
	data := common.FromHex(txData.Data)
	operation := uint8(txData.Operation)
	safeTxGas, _ := new(big.Int).SetString(txData.SafeTxGas, 10)
	baseGas, _ := new(big.Int).SetString(txData.BaseGas, 10)
	gasPrice, _ := new(big.Int).SetString(txData.GasPrice, 10)
	gasToken := common.HexToAddress(txData.GasToken)
	refundReceiver := common.HexToAddress(txData.RefundReceiver)
	nonce := new(big.Int).SetUint64(txData.Nonce)
	safeAddr := common.HexToAddress(safeAddress)
	chainIDBig := big.NewInt(chainID)

	return utils.CalculateTransactionHash(
		safeAddr,
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
		chainIDBig,
	)
}
