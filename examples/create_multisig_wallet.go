package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/yinwei/safe-core-sdk-golang/api"
)

func main() {
	fmt.Println("🏗️  Safe Multisig Wallet Creation")
	fmt.Println("================================")

	// 多签钱包配置
	owners := []string{
		"0x9C126aa4Eb6D110D646139969774F2c5b64dD279", // 所有者1
		"0xeB7E951F2D1A38188762dF12E0703aE16F76ab73", // 所有者2
		"0x74f4EFFb0B538BAec703346b03B6d9292f53A4CD", // 所有者3
	}
	threshold := 2 // 需要2个签名才能执行交易

	// 网络配置
	rpcURL := ""               //"https://ethereum-sepolia-rpc.publicnode.com"
	chainID := int64(11155111) // Sepolia测试网
	apiKey := ""

	// 部署者私钥 (用于支付gas费用)
	deployerPrivateKey := ""

	// 运行模式：true = 实际部署, false = 演示模式
	deployMode := false // 设置为true进行实际部署

	ctx := context.Background()

	fmt.Printf("Config: %d/%d multisig, %d owners\n", threshold, len(owners), len(owners))

	// 验证所有者地址格式
	for i, owner := range owners {
		if !common.IsHexAddress(owner) {
			log.Fatalf("Invalid owner address at index %d: %s", i, owner)
		}
	}

	// 验证阈值
	if threshold <= 0 || threshold > len(owners) {
		log.Fatalf("Invalid threshold: %d (must be between 1 and %d)", threshold, len(owners))
	}

	fmt.Printf("✅ Config validated\n")

	// =====================================================
	// STEP 2: 连接到以太坊网络 (仅在部署模式下)
	// =====================================================
	var client *ethclient.Client
	var networkID *big.Int

	if deployMode {
		fmt.Printf("🌐 Connecting to Sepolia...")
		var err error
		client, err = ethclient.Dial(rpcURL)
		if err != nil {
			log.Fatalf("Connection failed: %v", err)
		}
		defer client.Close()

		networkID, err = client.NetworkID(ctx)
		if err != nil {
			log.Fatalf("Network error: %v", err)
		}

		if networkID.Int64() != chainID {
			log.Fatalf("Wrong network: expected %d, got %d", chainID, networkID.Int64())
		}
		fmt.Printf(" ✅\n")
	} else {
		fmt.Printf("🎯 Demo mode (offline)\n")
		networkID = big.NewInt(chainID)
	}

	// 获取部署者账户信息
	privateKey, err := crypto.HexToECDSA(deployerPrivateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	deployerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Printf("💰 Deployer: %s\n", deployerAddress.Hex())

	if deployMode {
		balance, err := client.BalanceAt(ctx, deployerAddress, nil)
		if err != nil {
			log.Fatalf("Balance check failed: %v", err)
		}

		fmt.Printf("💰 Balance: %s ETH\n", formatEther(balance))
		if balance.Cmp(big.NewInt(0)) == 0 {
			fmt.Printf("⚠️  Warning: No ETH for gas\n")
		}
	}

	fmt.Printf("🔮 Predicting Safe address...")

	// Safe工厂合约地址 (Sepolia)
	safeFactoryAddress := "0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"
	safeSingletonAddress := "0x29fcB43b46531BcA003ddC8FCB67FFE91900C762" // Safe v1.4.1 L2版本

	fmt.Printf("📍 Safe Factory: %s\n", safeFactoryAddress)
	fmt.Printf("📍 Safe Singleton: %s\n", safeSingletonAddress)

	// 生成随机salt用于CREATE2
	salt := generateRandomSalt()
	fmt.Printf("📍 Salt: 0x%s\n", hex.EncodeToString(salt[:]))

	// 预测Safe地址
	predictedAddress, err := predictSafeAddress(
		safeFactoryAddress,
		safeSingletonAddress,
		owners,
		threshold,
		salt,
	)
	if err != nil {
		log.Fatalf("Failed to predict Safe address: %v", err)
	}

	fmt.Printf(" ✅\n🎯 Safe Address: %s\n", predictedAddress.Hex())

	fmt.Printf("📝 Preparing transaction...")

	// 创建Safe初始化数据
	initData, err := createSafeInitData(owners, threshold)
	if err != nil {
		log.Fatalf("Failed to create Safe init data: %v", err)
	}

	// 创建工厂调用数据
	factoryCallData, err := createSafeFactoryCallData(
		safeSingletonAddress,
		initData,
		salt,
	)
	if err != nil {
		log.Fatalf("Failed to create factory call data: %v", err)
	}

	fmt.Printf(" ✅ (%d bytes)\n", len(factoryCallData))

	// =====================================================
	// STEP 5: 部署Safe钱包
	// =====================================================
	if deployMode {
		fmt.Printf("🚀 Deploying Safe...")

		// 实际发送部署交易
		actualAddress, err := deploySafeWallet(ctx, client, privateKey, safeFactoryAddress, factoryCallData)
		if err != nil {
			fmt.Printf(" ❌\n❌ Deployment failed: %v\n", err)
			return
		}
		fmt.Printf(" ✅\n")

		// 显示实际地址与预测地址的对比
		if actualAddress != predictedAddress {
			fmt.Printf("⚠️  Actual address differs from predicted:\n")
			fmt.Printf("   Predicted: %s\n", predictedAddress.Hex())
			fmt.Printf("   Actual:    %s\n", actualAddress.Hex())
		}

		fmt.Printf("🔍 Verifying deployment...")
		// 使用实际地址进行验证
		err = verifySafeDeployment(ctx, client, actualAddress, owners, threshold, chainID, apiKey)
		if err != nil {
			fmt.Printf(" ❌\n❌ Verification failed: %v\n", err)
			return
		}
		fmt.Printf(" ✅\n")

		// 更新最终消息中的地址
		predictedAddress = actualAddress
	} else {
		fmt.Printf("💡 Demo mode - Ready to deploy (set deployMode=true)\n")
	}

	if deployMode {
		fmt.Printf("\n✅ Safe multisig wallet deployed successfully!\n")
		fmt.Printf("🔗 View at: https://app.safe.global/sep:%s\n", predictedAddress.Hex())
	}
}

// generateRandomSalt 生成用于CREATE2的32字节盐值
func generateRandomSalt() [32]byte {
	var salt [32]byte
	// 生产环境中应使用crypto/rand生成安全随机数
	// 演示中使用可预测的值
	copy(salt[:], []byte("SafeMultisigDemo2025"))
	return salt
}

// predictSafeAddress 使用CREATE2预测Safe地址
func predictSafeAddress(
	factoryAddress string,
	singletonAddress string,
	owners []string,
	threshold int,
	salt [32]byte,
) (common.Address, error) {
	// 这是一个简化的预测实现
	// 实际实现中应该：
	// 1. 编码Safe初始化数据
	// 2. 使用工厂合约计算CREATE2地址
	// 3. 返回预测地址

	// 演示用途，生成模拟地址
	hash := crypto.Keccak256(
		common.HexToAddress(factoryAddress).Bytes(),
		common.HexToAddress(singletonAddress).Bytes(),
		salt[:],
	)

	var addr common.Address
	copy(addr[:], hash[12:])
	return addr, nil
}

// createSafeInitData 创建Safe初始化数据
func createSafeInitData(owners []string, threshold int) ([]byte, error) {
	// Safe设置函数签名: setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver)
	safeSetupABI := `[{
		"inputs": [
			{"name": "_owners", "type": "address[]"},
			{"name": "_threshold", "type": "uint256"},
			{"name": "to", "type": "address"},
			{"name": "data", "type": "bytes"},
			{"name": "fallbackHandler", "type": "address"},
			{"name": "paymentToken", "type": "address"},
			{"name": "payment", "type": "uint256"},
			{"name": "paymentReceiver", "type": "address"}
		],
		"name": "setup",
		"outputs": [],
		"type": "function"
	}]`

	parsedABI, err := abi.JSON(strings.NewReader(safeSetupABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse Safe setup ABI: %w", err)
	}

	// 将字符串地址转换为common.Address
	ownerAddresses := make([]common.Address, len(owners))
	for i, owner := range owners {
		ownerAddresses[i] = common.HexToAddress(owner)
	}

	// 编码设置函数调用
	data, err := parsedABI.Pack(
		"setup",
		ownerAddresses,               // _owners
		big.NewInt(int64(threshold)), // _threshold
		common.Address{},             // to (no delegate call)
		[]byte{},                     // data (empty)
		common.HexToAddress("0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99"), // fallbackHandler
		common.Address{}, // paymentToken (no payment)
		big.NewInt(0),    // payment (0)
		common.Address{}, // paymentReceiver (none)
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode setup call: %w", err)
	}

	return data, nil
}

// createSafeFactoryCallData creates the call data for Safe factory
func createSafeFactoryCallData(singletonAddress string, initData []byte, salt [32]byte) ([]byte, error) {
	// Safe factory createProxyWithNonce function
	factoryABI := `[{
		"inputs": [
			{"name": "singleton", "type": "address"},
			{"name": "data", "type": "bytes"},
			{"name": "saltNonce", "type": "uint256"}
		],
		"name": "createProxyWithNonce",
		"outputs": [{"name": "proxy", "type": "address"}],
		"type": "function"
	}]`

	parsedABI, err := abi.JSON(strings.NewReader(factoryABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse factory ABI: %w", err)
	}

	// Convert salt to big.Int
	saltBig := new(big.Int).SetBytes(salt[:])

	// Encode the factory function call
	data, err := parsedABI.Pack(
		"createProxyWithNonce",
		common.HexToAddress(singletonAddress), // singleton
		initData,                              // data
		saltBig,                               // saltNonce
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode factory call: %w", err)
	}

	return data, nil
}

// deploySafeWallet deploys the Safe wallet to the blockchain and returns the created address
func deploySafeWallet(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey, factoryAddress string, callData []byte) (common.Address, error) {

	// Get deployer address
	deployerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Get current nonce
	nonce, err := client.PendingNonceAt(ctx, deployerAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get gas price: %w", err)
	}

	// Ensure minimum gas price for Sepolia
	minGasPrice := big.NewInt(1000000000) // 1 Gwei
	if gasPrice.Cmp(minGasPrice) < 0 {
		gasPrice = minGasPrice
	}

	// Get chain ID
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Create transaction
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(factoryAddress),
		big.NewInt(0), // value = 0
		500000,        // gas limit
		gasPrice,
		callData,
	)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	fmt.Printf("✅ Transaction sent successfully!\n")
	fmt.Printf("   TX Hash: %s\n", signedTx.Hash().Hex())
	fmt.Printf("   Factory: %s\n", factoryAddress)
	fmt.Printf("   Gas Used: %d\n", signedTx.Gas())
	fmt.Printf("   Gas Price: %s Gwei\n", formatGwei(gasPrice))

	// Wait for transaction receipt
	fmt.Printf("⏳ Waiting for transaction confirmation...\n")
	receipt, err := waitForTransaction(ctx, client, signedTx.Hash())
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to wait for transaction: %w", err)
	}

	if receipt.Status != 1 {
		return common.Address{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	fmt.Printf("🎉 Transaction confirmed successfully!\n")
	fmt.Printf("   Block Number: %d\n", receipt.BlockNumber.Uint64())
	fmt.Printf("   Gas Used: %d\n", receipt.GasUsed)
	fmt.Printf("   Transaction Index: %d\n", receipt.TransactionIndex)

	// Extract the created Safe address from logs
	var createdAddress common.Address
	for _, log := range receipt.Logs {
		// Look for ProxyCreation event from SafeProxyFactory
		if len(log.Topics) > 0 && len(log.Data) >= 32 {
			// The first topic should be the ProxyCreation event hash
			// The created proxy address is typically in the data or topics
			if log.Address.Hex() == factoryAddress {
				// For SafeProxyFactory, the created proxy address is in the first 32 bytes of data
				if len(log.Data) >= 32 {
					createdAddress = common.BytesToAddress(log.Data[0:32])
					break
				}
			}
		}
	}

	if createdAddress == (common.Address{}) {
		return common.Address{}, fmt.Errorf("could not find created Safe address in transaction logs")
	}

	return createdAddress, nil
}

// verifySafeDeployment verifies that the Safe was deployed correctly
func verifySafeDeployment(ctx context.Context, client *ethclient.Client, safeAddress common.Address, expectedOwners []string, expectedThreshold int, chainID int64, apiKey string) error {

	// Check if contract exists
	code, err := client.CodeAt(ctx, safeAddress, nil)
	if err != nil {
		return fmt.Errorf("failed to get contract code: %w", err)
	}

	if len(code) == 0 {
		return fmt.Errorf("no contract found at address %s", safeAddress.Hex())
	}

	fmt.Printf("✅ Contract deployed at address\n")
	fmt.Printf("   Contract Size: %d bytes\n", len(code))

	// Try to verify with Safe API (give it a moment to index)
	fmt.Printf("⏳ Waiting for Safe API indexing...\n")
	time.Sleep(10 * time.Second)

	// Create API client
	apiClient, err := api.NewSafeApiKit(api.SafeApiKitConfig{
		ChainID: chainID,
		ApiKey:  apiKey,
	})
	if err != nil {
		fmt.Printf("⚠️  Could not create API client: %v\n", err)
		fmt.Printf("   Safe deployed but API verification skipped\n")
		return nil
	}

	// Get Safe info from API
	safeInfo, err := apiClient.GetSafeInfo(ctx, safeAddress.Hex())
	if err != nil {
		fmt.Printf("⚠️  Safe not yet indexed by API service: %v\n", err)
		fmt.Printf("   (This is normal for new deployments)\n")
		return nil
	}

	// Verify configuration
	fmt.Printf("✅ Safe successfully indexed by API!\n")
	fmt.Printf("   Address: %s\n", safeInfo.Address)
	fmt.Printf("   Threshold: %d (expected: %d)\n", safeInfo.Threshold, expectedThreshold)
	fmt.Printf("   Owners: %d (expected: %d)\n", len(safeInfo.Owners), len(expectedOwners))

	// Verify threshold
	if safeInfo.Threshold != expectedThreshold {
		return fmt.Errorf("threshold mismatch: got %d, expected %d", safeInfo.Threshold, expectedThreshold)
	}

	// Verify owners
	if len(safeInfo.Owners) != len(expectedOwners) {
		return fmt.Errorf("owner count mismatch: got %d, expected %d", len(safeInfo.Owners), len(expectedOwners))
	}

	// Check each owner
	for i, expectedOwner := range expectedOwners {
		found := false
		for _, actualOwner := range safeInfo.Owners {
			if strings.EqualFold(actualOwner, expectedOwner) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("owner %d not found: %s", i+1, expectedOwner)
		}
	}

	fmt.Printf("🎯 All configurations verified successfully!\n")
	fmt.Printf("   Safe Version: %s\n", safeInfo.Version)
	fmt.Printf("   Master Copy: %s\n", safeInfo.MasterCopy)

	return nil
}

// waitForTransaction waits for a transaction to be mined
func waitForTransaction(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for i := 0; i < 60; i++ { // Wait up to 60 seconds
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		// Sleep and retry
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("transaction not mined after 60 seconds")
}

// formatGwei formats wei to gwei
func formatGwei(wei *big.Int) string {
	gwei := new(big.Float).SetInt(wei)
	gwei.Quo(gwei, big.NewFloat(1e9))
	return gwei.Text('f', 2)
}

// formatEther formats wei to ether with proper decimal places
func formatEther(wei *big.Int) string {
	eth := new(big.Float).SetInt(wei)
	eth.Quo(eth, big.NewFloat(1e18))
	return eth.Text('f', 6)
}
