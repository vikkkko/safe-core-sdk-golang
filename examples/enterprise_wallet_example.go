package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/contracts"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
)

const (
	// Enterprise Wallet deployed contracts on your network
	FactoryAddress     = "0x19cd09AA77a74f92fC12D4D2f5D63ea61193E157"
	ImplementationAddr = "0x3d6850a4A9790c3aD3924A5d66b4fEEC8cd25bE2"
)

// Context holds all necessary data for examples
type ExampleContext struct {
	Client          *ethclient.Client
	PrivateKey      *ecdsa.PrivateKey
	FromAddress     common.Address
	ChainID         *big.Int
	Auth            *bind.TransactOpts
	FactoryContract *contracts.EnterpriseWalletFactory
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Initialize context
	ctx, err := initializeContext()
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}
	defer ctx.Client.Close()

	fmt.Printf("Using account: %s\n", ctx.FromAddress.Hex())
	fmt.Printf("Factory address: %s\n", FactoryAddress)
	fmt.Printf("Implementation address: %s\n\n", ImplementationAddr)

	// Show menu and run examples
	for {
		showMenu()
		choice := getUserInput()

		if choice == "0" || choice == "exit" || choice == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		runExample(ctx, choice)
		fmt.Println("\nPress Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func initializeContext() (*ExampleContext, error) {
	// Get configuration from environment
	rpcURL := os.Getenv("RPC_URL")
	privateKeyHex := os.Getenv("DEPLOYER_PRIVATE_KEY")
	chainIDStr := os.Getenv("CHAIN_ID")

	if rpcURL == "" || privateKeyHex == "" || chainIDStr == "" {
		return nil, fmt.Errorf("please set RPC_URL, DEPLOYER_PRIVATE_KEY, and CHAIN_ID in .env file")
	}

	// Parse chain ID
	chainID := new(big.Int)
	chainID.SetString(chainIDStr, 10)

	// Connect to Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		client.Close()
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	// Create factory contract instance
	factoryContract, err := contracts.NewEnterpriseWalletFactory(common.HexToAddress(FactoryAddress), client)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to instantiate factory contract: %w", err)
	}

	return &ExampleContext{
		Client:          client,
		PrivateKey:      privateKey,
		FromAddress:     fromAddress,
		ChainID:         chainID,
		Auth:            auth,
		FactoryContract: factoryContract,
	}, nil
}

func showMenu() {
	fmt.Println("\n===============================================")
	fmt.Println("   Enterprise Wallet SDK Examples")
	fmt.Println("===============================================")
	fmt.Println("\nFactory Contract Examples:")
	fmt.Println("  1.  Check if implementation is whitelisted")
	fmt.Println("  2.  Get all whitelisted implementations")
	fmt.Println("  3.  Predict wallet address")
	fmt.Println("  4.  Prepare wallet deployment data")
	fmt.Println("  5.  Deploy enterprise wallet (TX)")
	fmt.Println("\nEnterprise Wallet - Prepare Data:")
	fmt.Println("  6.  Show method selectors")
	fmt.Println("  7.  Prepare createPaymentAccount data")
	fmt.Println("  8.  Prepare createCollectionAccount data")
	fmt.Println("  9.  Prepare approveTokenForPayment data")
	fmt.Println("  10. Prepare transferETHToPayment data")
	fmt.Println("  11. Prepare setCollectionTarget data")
	fmt.Println("  12. Prepare collectFunds data")
	fmt.Println("  13. Prepare updateMethodController data")
	fmt.Println("  14. Prepare updatePaymentAccountController data")
	fmt.Println("\nEnterprise Wallet - Execute on Chain (TX):")
	fmt.Println("  22. Execute createPaymentAccount")
	fmt.Println("  23. Execute createCollectionAccount")
	fmt.Println("  24. Execute approveTokenForPayment")
	fmt.Println("  25. Execute transferETHToPayment")
	fmt.Println("  26. Execute setCollectionTarget")
	fmt.Println("  27. Execute collectFunds")
	fmt.Println("  28. Execute updateMethodController")
	fmt.Println("  29. Execute updatePaymentAccountController")
	fmt.Println("\nQuery Functions (View):")
	fmt.Println("  15. Query payment accounts")
	fmt.Println("  16. Query collection accounts")
	fmt.Println("  17. Query method config")
	fmt.Println("  18. Check if address is payment account")
	fmt.Println("  19. Check if address is collection account")
	fmt.Println("  20. Get token allowance")
	fmt.Println("\nUtility:")
	fmt.Println("  21. Run all read-only examples")
	fmt.Println("  0.  Exit")
	fmt.Println("===============================================")
	fmt.Print("\nEnter your choice: ")
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func runExample(ctx *ExampleContext, choice string) {
	fmt.Println()
	switch choice {
	case "1":
		exampleCheckImplementationWhitelisted(ctx)
	case "2":
		exampleGetWhitelistedImplementations(ctx)
	case "3":
		examplePredictWalletAddress(ctx)
	case "4":
		examplePrepareDeploymentData(ctx)
	case "5":
		exampleDeployWallet(ctx)
	case "6":
		exampleShowMethodSelectors(ctx)
	case "7":
		examplePrepareCreatePaymentAccount(ctx)
	case "8":
		examplePrepareCreateCollectionAccount(ctx)
	case "9":
		examplePrepareApproveToken(ctx)
	case "10":
		examplePrepareTransferETH(ctx)
	case "11":
		examplePrepareSetCollectionTarget(ctx)
	case "12":
		examplePrepareCollectFunds(ctx)
	case "13":
		examplePrepareUpdateMethodController(ctx)
	case "14":
		examplePrepareUpdatePaymentAccountController(ctx)
	case "15":
		exampleQueryPaymentAccounts(ctx)
	case "16":
		exampleQueryCollectionAccounts(ctx)
	case "17":
		exampleQueryMethodConfig(ctx)
	case "18":
		exampleCheckIsPaymentAccount(ctx)
	case "19":
		exampleCheckIsCollectionAccount(ctx)
	case "20":
		exampleGetAllowance(ctx)
	case "21":
		runAllReadOnlyExamples(ctx)
	case "22":
		exampleExecuteCreatePaymentAccount(ctx)
	case "23":
		exampleExecuteCreateCollectionAccount(ctx)
	case "24":
		exampleExecuteApproveToken(ctx)
	case "25":
		exampleExecuteTransferETH(ctx)
	case "26":
		exampleExecuteSetCollectionTarget(ctx)
	case "27":
		exampleExecuteCollectFunds(ctx)
	case "28":
		exampleExecuteUpdateMethodController(ctx)
	case "29":
		exampleExecuteUpdatePaymentAccountController(ctx)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

// ============= Factory Contract Examples =============

func exampleCheckImplementationWhitelisted(ctx *ExampleContext) {
	fmt.Println("=== Check Implementation Whitelisted ===")
	isWhitelisted, err := ctx.FactoryContract.IsImplementationWhitelisted(
		&bind.CallOpts{},
		common.HexToAddress(ImplementationAddr),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Implementation %s is whitelisted: %v\n", ImplementationAddr, isWhitelisted)
}

func exampleGetWhitelistedImplementations(ctx *ExampleContext) {
	fmt.Println("=== Get Whitelisted Implementations ===")
	implementations, err := ctx.FactoryContract.GetWhitelistedImplementations(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Found %d whitelisted implementation(s):\n", len(implementations))
	for i, impl := range implementations {
		fmt.Printf("  %d. %s\n", i+1, impl.Hex())
	}
}

func examplePredictWalletAddress(ctx *ExampleContext) {
	fmt.Println("=== Predict Wallet Address ===")
	var salt [32]byte
	copy(salt[:], []byte("my-enterprise-wallet-v1"))

	predictedAddr, err := ctx.FactoryContract.PredictWalletAddress(
		&bind.CallOpts{},
		common.HexToAddress(ImplementationAddr),
		salt,
		ctx.FromAddress,
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Salt: 0x%x\n", salt)
	fmt.Printf("Deployer: %s\n", ctx.FromAddress.Hex())
	fmt.Printf("Predicted wallet address: %s\n", predictedAddr.Hex())
}

func examplePrepareDeploymentData(ctx *ExampleContext) {
	fmt.Println("=== Prepare Deployment Data ===")
	var salt [32]byte
	copy(salt[:], []byte("my-enterprise-wallet-v1"))

	methodSelectors := [][4]byte{
		utils.CreatePaymentAccountSelector,
		utils.CreateCollectionAccountSelector,
		utils.ApproveTokenForPaymentSelector,
	}

	configs := make([]utils.MethodConfig, len(methodSelectors))
	for i := range methodSelectors {
		configs[i] = utils.MethodConfig{
			Controller: ctx.FromAddress,
		}
	}

	initParams := utils.InitParams{
		Methods:    methodSelectors,
		Configs:    configs,
		SuperAdmin: ctx.FromAddress,
	}

	deployData, err := utils.CreateEnterpriseWalletData(
		common.HexToAddress(ImplementationAddr),
		salt,
		initParams,
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Deployment calldata prepared:\n")
	fmt.Printf("  Length: %d bytes\n", len(deployData))
	fmt.Printf("  First 10 bytes: 0x%x\n", deployData[:10])
	fmt.Printf("  Full data: 0x%x\n", deployData)
}

func exampleDeployWallet(ctx *ExampleContext) {
	fmt.Println("=== Deploy Enterprise Wallet ===")
	fmt.Println("WARNING: This will send an actual transaction!")
	fmt.Print("Type 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Deployment cancelled.")
		return
	}

	var salt [32]byte
	copy(salt[:], []byte("my-enterprise-wallet-v1"))

	methodSelectors := [][4]byte{
		utils.CreatePaymentAccountSelector,
		utils.CreateCollectionAccountSelector,
		utils.ApproveTokenForPaymentSelector,
	}

	// Create init params
	configs := make([]contracts.IEnterpriseWalletMethodConfig, len(methodSelectors))
	for i := range methodSelectors {
		configs[i] = contracts.IEnterpriseWalletMethodConfig{Controller: ctx.FromAddress}
	}

	contractInitParams := contracts.IEnterpriseWalletFactoryInitParams{
		Methods:    methodSelectors,
		Configs:    configs,
		SuperAdmin: ctx.FromAddress,
	}

	tx, err := ctx.FactoryContract.CreateWallet(
		ctx.Auth,
		common.HexToAddress(ImplementationAddr),
		salt,
		contractInitParams,
	)
	if err != nil {
		log.Printf("Error deploying wallet: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		predictedAddr, _ := ctx.FactoryContract.PredictWalletAddress(
			&bind.CallOpts{},
			common.HexToAddress(ImplementationAddr),
			salt,
			ctx.FromAddress,
		)
		fmt.Printf("✓ Enterprise wallet deployed successfully!\n")
		fmt.Printf("  Address: %s\n", predictedAddr.Hex())
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Deployment transaction failed")
	}
}

// ============= Enterprise Wallet Examples =============

func exampleShowMethodSelectors(ctx *ExampleContext) {
	fmt.Println("=== Method Selectors ===")
	fmt.Printf("createPaymentAccount:             0x%x\n", utils.CreatePaymentAccountSelector)
	fmt.Printf("createCollectionAccount:          0x%x\n", utils.CreateCollectionAccountSelector)
	fmt.Printf("approveTokenForPayment:           0x%x\n", utils.ApproveTokenForPaymentSelector)
	fmt.Printf("transferETHToPayment:             0x%x\n", utils.TransferETHToPaymentSelector)
	fmt.Printf("collectFunds:                     0x%x\n", utils.CollectFundsSelector)

	// Show additional selectors
	setCollectionTargetSelector := utils.GetMethodSelector("setCollectionTarget(address,address)")
	updateMethodControllerSelector := utils.GetMethodSelector("updateMethodController(bytes4,address)")
	updatePaymentControllerSelector := utils.GetMethodSelector("updatePaymentAccountController(address,address)")
	emergencyFreezeSelector := utils.GetMethodSelector("emergencyFreeze(address,bool)")
	emergencyPauseSelector := utils.GetMethodSelector("emergencyPause(bool)")

	fmt.Printf("setCollectionTarget:              0x%x\n", setCollectionTargetSelector)
	fmt.Printf("updateMethodController:           0x%x\n", updateMethodControllerSelector)
	fmt.Printf("updatePaymentAccountController:   0x%x\n", updatePaymentControllerSelector)
	fmt.Printf("emergencyFreeze:                  0x%x\n", emergencyFreezeSelector)
	fmt.Printf("emergencyPause:                   0x%x\n", emergencyPauseSelector)
}

func examplePrepareCreatePaymentAccount(ctx *ExampleContext) {
	fmt.Println("=== Prepare Create Payment Account ===")
	accountName := "Treasury Payment Account"
	controller := ctx.FromAddress

	data, err := utils.CreatePaymentAccountData(accountName, controller)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Account name: %s\n", accountName)
	fmt.Printf("Controller: %s\n", controller.Hex())
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareCreateCollectionAccount(ctx *ExampleContext) {
	fmt.Println("=== Prepare Create Collection Account ===")
	accountName := "Revenue Collection Account"
	collectionTarget := common.Address{} // address(0) means enterprise wallet itself

	data, err := utils.CreateCollectionAccountData(accountName, collectionTarget)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Account name: %s\n", accountName)
	fmt.Printf("Collection target: %s (address(0) = enterprise wallet)\n", collectionTarget.Hex())
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareApproveToken(ctx *ExampleContext) {
	fmt.Println("=== Prepare Approve Token for Payment ===")
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // USDC
	paymentAccount := common.HexToAddress("0x1111111111111111111111111111111111111111")
	amount := new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e6)) // 1000 USDC

	data, err := utils.ApproveTokenForPaymentData(usdcAddress, paymentAccount, amount)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Token: %s (USDC)\n", usdcAddress.Hex())
	fmt.Printf("Payment account: %s\n", paymentAccount.Hex())
	fmt.Printf("Amount: 1000 USDC\n")
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareTransferETH(ctx *ExampleContext) {
	fmt.Println("=== Prepare Transfer ETH to Payment ===")
	paymentAccount := common.HexToAddress("0x1111111111111111111111111111111111111111")
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18)) // 1 ETH

	data, err := utils.TransferETHToPaymentData(paymentAccount, amount)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Payment account: %s\n", paymentAccount.Hex())
	fmt.Printf("Amount: 1 ETH\n")
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareSetCollectionTarget(ctx *ExampleContext) {
	fmt.Println("=== Prepare Set Collection Target ===")

	// Create ABI-encoded data for setCollectionTarget
	collectionAccount := common.HexToAddress("0x2222222222222222222222222222222222222222")
	newTarget := common.HexToAddress("0x3333333333333333333333333333333333333333")

	selector := utils.GetMethodSelector("setCollectionTarget(address,address)")

	// Manually encode parameters (address, address)
	data := make([]byte, 4+32+32)
	copy(data[0:4], selector[:])
	copy(data[4+12:4+32], collectionAccount.Bytes())
	copy(data[36+12:36+32], newTarget.Bytes())

	fmt.Printf("Collection account: %s\n", collectionAccount.Hex())
	fmt.Printf("New target: %s\n", newTarget.Hex())
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareCollectFunds(ctx *ExampleContext) {
	fmt.Println("=== Prepare Collect Funds ===")
	collectionAccount := common.HexToAddress("0x2222222222222222222222222222222222222222")

	// Example 1: Collect ETH
	ethData, err := utils.CollectFundsData(common.Address{}, collectionAccount)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("Collect ETH:")
	fmt.Printf("  Token: address(0) (ETH)\n")
	fmt.Printf("  Collection account: %s\n", collectionAccount.Hex())
	fmt.Printf("  Calldata: 0x%x\n\n", ethData)

	// Example 2: Collect USDC
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	usdcData, err := utils.CollectFundsData(usdcAddress, collectionAccount)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("Collect USDC:")
	fmt.Printf("  Token: %s (USDC)\n", usdcAddress.Hex())
	fmt.Printf("  Collection account: %s\n", collectionAccount.Hex())
	fmt.Printf("  Calldata: 0x%x\n", usdcData)
}

func examplePrepareUpdateMethodController(ctx *ExampleContext) {
	fmt.Println("=== Prepare Update Method Controller ===")

	methodSig := utils.CreatePaymentAccountSelector
	newController := common.HexToAddress("0x4444444444444444444444444444444444444444")

	selector := utils.GetMethodSelector("updateMethodController(bytes4,address)")

	// Manually encode parameters (bytes4, address)
	data := make([]byte, 4+32+32)
	copy(data[0:4], selector[:])
	copy(data[4:4+4], methodSig[:])
	copy(data[36+12:36+32], newController.Bytes())

	fmt.Printf("Method selector: 0x%x (createPaymentAccount)\n", methodSig)
	fmt.Printf("New controller: %s\n", newController.Hex())
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareUpdatePaymentAccountController(ctx *ExampleContext) {
	fmt.Println("=== Prepare Update Payment Account Controller ===")

	paymentAccount := common.HexToAddress("0x1111111111111111111111111111111111111111")
	newController := common.HexToAddress("0x5555555555555555555555555555555555555555")

	selector := utils.GetMethodSelector("updatePaymentAccountController(address,address)")

	// Manually encode parameters (address, address)
	data := make([]byte, 4+32+32)
	copy(data[0:4], selector[:])
	copy(data[4+12:4+32], paymentAccount.Bytes())
	copy(data[36+12:36+32], newController.Bytes())

	fmt.Printf("Payment account: %s\n", paymentAccount.Hex())
	fmt.Printf("New controller: %s\n", newController.Hex())
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

// ============= View Functions Examples =============

func exampleQueryPaymentAccounts(ctx *ExampleContext) {
	fmt.Println("=== Query Payment Accounts ===")
	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()

	if walletAddrStr == "" {
		fmt.Println("Using example address...")
		walletAddrStr = "0xCD6c4962346F5680C765127ED29A8F5cc53a6B66"
	}

	walletAddr := common.HexToAddress(walletAddrStr)
	wallet, err := contracts.NewEnterpriseWallet(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	accounts, err := wallet.GetPaymentAccounts(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error querying payment accounts: %v", err)
		return
	}

	fmt.Printf("Found %d payment account(s):\n", len(accounts))
	for i, acc := range accounts {
		fmt.Printf("  %d. Address: %s\n", i+1, acc.Account.Hex())
		fmt.Printf("     Created: %s\n", acc.CreatedAt.String())
		fmt.Printf("     Active: %v\n", acc.IsActive)
	}
}

func exampleQueryCollectionAccounts(ctx *ExampleContext) {
	fmt.Println("=== Query Collection Accounts ===")
	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()

	if walletAddrStr == "" {
		fmt.Println("Using example address...")
		walletAddrStr = "0xCD6c4962346F5680C765127ED29A8F5cc53a6B66"
	}

	walletAddr := common.HexToAddress(walletAddrStr)
	wallet, err := contracts.NewEnterpriseWallet(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	accounts, err := wallet.GetCollectionAccounts(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error querying collection accounts: %v", err)
		return
	}

	fmt.Printf("Found %d collection account(s):\n", len(accounts))
	for i, acc := range accounts {
		fmt.Printf("  %d. Address: %s\n", i+1, acc.Account.Hex())
		fmt.Printf("     Created: %s\n", acc.CreatedAt.String())
		fmt.Printf("     Active: %v\n", acc.IsActive)
	}
}

func exampleQueryMethodConfig(ctx *ExampleContext) {
	fmt.Println("=== Query Method Config ===")
	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()

	if walletAddrStr == "" {
		fmt.Println("Using example address...")
		walletAddrStr = "0xCD6c4962346F5680C765127ED29A8F5cc53a6B66"
	}

	walletAddr := common.HexToAddress(walletAddrStr)
	wallet, err := contracts.NewEnterpriseWallet(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	methodSig := utils.CreatePaymentAccountSelector
	config, err := wallet.GetMethodConfig(&bind.CallOpts{}, methodSig)
	if err != nil {
		log.Printf("Error querying method config: %v", err)
		return
	}

	fmt.Printf("Method: createPaymentAccount (0x%x)\n", methodSig)
	fmt.Printf("Controller: %s\n", config.Controller.Hex())
}

func exampleCheckIsPaymentAccount(ctx *ExampleContext) {
	fmt.Println("=== Check Is Payment Account ===")
	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()

	if walletAddrStr == "" {
		fmt.Println("Using example address...")
		walletAddrStr = "0xCD6c4962346F5680C765127ED29A8F5cc53a6B66"
	}

	fmt.Print("Enter account address to check: ")
	accountAddrStr := getUserInput()

	if accountAddrStr == "" {
		accountAddrStr = "0x1111111111111111111111111111111111111111"
	}

	walletAddr := common.HexToAddress(walletAddrStr)
	accountAddr := common.HexToAddress(accountAddrStr)

	wallet, err := contracts.NewEnterpriseWallet(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	isPaymentAccount, err := wallet.IsPaymentAccount(&bind.CallOpts{}, accountAddr)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Address %s is payment account: %v\n", accountAddr.Hex(), isPaymentAccount)
}

func exampleCheckIsCollectionAccount(ctx *ExampleContext) {
	fmt.Println("=== Check Is Collection Account ===")
	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()

	if walletAddrStr == "" {
		fmt.Println("Using example address...")
		walletAddrStr = "0xCD6c4962346F5680C765127ED29A8F5cc53a6B66"
	}

	fmt.Print("Enter account address to check: ")
	accountAddrStr := getUserInput()

	if accountAddrStr == "" {
		accountAddrStr = "0x2222222222222222222222222222222222222222"
	}

	walletAddr := common.HexToAddress(walletAddrStr)
	accountAddr := common.HexToAddress(accountAddrStr)

	wallet, err := contracts.NewEnterpriseWallet(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	isCollectionAccount, err := wallet.IsCollectionAccount(&bind.CallOpts{}, accountAddr)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Address %s is collection account: %v\n", accountAddr.Hex(), isCollectionAccount)
}

func exampleGetAllowance(ctx *ExampleContext) {
	fmt.Println("=== Get Token Allowance ===")
	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()

	if walletAddrStr == "" {
		fmt.Println("Using example address...")
		walletAddrStr = "0xCD6c4962346F5680C765127ED29A8F5cc53a6B66"
	}

	walletAddr := common.HexToAddress(walletAddrStr)
	tokenAddr := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // USDC
	paymentAccount := common.HexToAddress("0x1111111111111111111111111111111111111111")

	wallet, err := contracts.NewEnterpriseWallet(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	allowance, err := wallet.GetAllowance(&bind.CallOpts{}, tokenAddr, paymentAccount)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Token: %s (USDC)\n", tokenAddr.Hex())
	fmt.Printf("Payment account: %s\n", paymentAccount.Hex())
	fmt.Printf("Allowance: %s (raw)\n", allowance.String())

	// Convert to human-readable USDC (6 decimals)
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	readable := new(big.Int).Div(allowance, divisor)
	fmt.Printf("Allowance: %s USDC\n", readable.String())
}

// ============= Utility Functions =============

func runAllReadOnlyExamples(ctx *ExampleContext) {
	fmt.Println("=== Running All Read-Only Examples ===\n")

	examples := []struct {
		name string
		fn   func(*ExampleContext)
	}{
		{"Check Implementation Whitelisted", exampleCheckImplementationWhitelisted},
		{"Get Whitelisted Implementations", exampleGetWhitelistedImplementations},
		{"Predict Wallet Address", examplePredictWalletAddress},
		{"Show Method Selectors", exampleShowMethodSelectors},
		{"Prepare Create Payment Account", examplePrepareCreatePaymentAccount},
		{"Prepare Create Collection Account", examplePrepareCreateCollectionAccount},
		{"Prepare Approve Token", examplePrepareApproveToken},
		{"Prepare Transfer ETH", examplePrepareTransferETH},
		{"Prepare Set Collection Target", examplePrepareSetCollectionTarget},
		{"Prepare Collect Funds", examplePrepareCollectFunds},
		{"Prepare Update Method Controller", examplePrepareUpdateMethodController},
		{"Prepare Update Payment Account Controller", examplePrepareUpdatePaymentAccountController},
	}

	for i, ex := range examples {
		fmt.Printf("\n[%d/%d] %s\n", i+1, len(examples), ex.name)
		fmt.Println(strings.Repeat("-", 50))
		ex.fn(ctx)
	}

	fmt.Println("\n=== All Read-Only Examples Completed ===")
}

// ============= Execute on Chain Functions =============

// Helper function to get fresh auth for transactions
func getFreshAuth(ctx *ExampleContext) (*bind.TransactOpts, error) {
	nonce, err := ctx.Client.PendingNonceAt(context.Background(), ctx.FromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := ctx.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ctx.PrivateKey, ctx.ChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth, nil
}

func exampleExecuteCreatePaymentAccount(ctx *ExampleContext) {
	fmt.Println("=== Execute Create Payment Account ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter account name: ")
	accountName := getUserInput()
	if accountName == "" {
		accountName = "Payment Account " + fmt.Sprint(time.Now().Unix())
	}

	fmt.Print("Enter controller address (press Enter for your address): ")
	controllerStr := getUserInput()
	var controller common.Address
	if controllerStr == "" {
		controller = ctx.FromAddress
	} else {
		controller = common.HexToAddress(controllerStr)
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Account name: %s\n", accountName)
	fmt.Printf("  Controller: %s\n", controller.Hex())
	fmt.Print("\nType 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Transaction cancelled.")
		return
	}

	// Get fresh auth
	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Connect to wallet contract
	wallet, err := contracts.NewEnterpriseWallet(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Execute transaction
	tx, err := wallet.CreatePaymentAccount(auth, accountName, controller)
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ Payment account created successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)

		// Try to get the created account address from events
		fmt.Println("\nNote: Check transaction logs for the created payment account address")
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteCreateCollectionAccount(ctx *ExampleContext) {
	fmt.Println("=== Execute Create Collection Account ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter account name: ")
	accountName := getUserInput()
	if accountName == "" {
		accountName = "Collection Account " + fmt.Sprint(time.Now().Unix())
	}

	fmt.Print("Enter collection target address (press Enter for wallet itself): ")
	targetStr := getUserInput()
	var target common.Address
	if targetStr == "" {
		target = common.Address{} // address(0) defaults to wallet
	} else {
		target = common.HexToAddress(targetStr)
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Account name: %s\n", accountName)
	fmt.Printf("  Collection target: %s\n", target.Hex())
	fmt.Print("\nType 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Transaction cancelled.")
		return
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := contracts.NewEnterpriseWallet(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.CreateCollectionAccount(auth, accountName, target)
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ Collection account created successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteApproveToken(ctx *ExampleContext) {
	fmt.Println("=== Execute Approve Token for Payment ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter token address: ")
	tokenAddrStr := getUserInput()
	if tokenAddrStr == "" {
		fmt.Println("No token address provided. Cancelled.")
		return
	}

	fmt.Print("Enter payment account address: ")
	paymentAddrStr := getUserInput()
	if paymentAddrStr == "" {
		fmt.Println("No payment account address provided. Cancelled.")
		return
	}

	fmt.Print("Enter amount (in smallest unit): ")
	amountStr := getUserInput()
	amount := new(big.Int)
	if amountStr == "" {
		amount = big.NewInt(0)
	} else {
		amount.SetString(amountStr, 10)
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Token: %s\n", tokenAddrStr)
	fmt.Printf("  Payment account: %s\n", paymentAddrStr)
	fmt.Printf("  Amount: %s\n", amount.String())
	fmt.Print("\nType 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Transaction cancelled.")
		return
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := contracts.NewEnterpriseWallet(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.ApproveTokenForPayment(auth, common.HexToAddress(tokenAddrStr), common.HexToAddress(paymentAddrStr), amount)
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ Token approved successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteTransferETH(ctx *ExampleContext) {
	fmt.Println("=== Execute Transfer ETH to Payment ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter payment account address: ")
	paymentAddrStr := getUserInput()
	if paymentAddrStr == "" {
		fmt.Println("No payment account address provided. Cancelled.")
		return
	}

	fmt.Print("Enter amount in wei (e.g., 1000000000000000000 for 1 ETH): ")
	amountStr := getUserInput()
	amount := new(big.Int)
	if amountStr == "" {
		fmt.Println("No amount provided. Cancelled.")
		return
	}
	amount.SetString(amountStr, 10)

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Payment account: %s\n", paymentAddrStr)
	fmt.Printf("  Amount: %s wei\n", amount.String())
	fmt.Print("\nType 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Transaction cancelled.")
		return
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := contracts.NewEnterpriseWallet(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	paymentAddr := common.HexToAddress(paymentAddrStr)
	tx, err := wallet.TransferETHToPayment(auth, common.Address(paymentAddr), amount)
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ ETH transferred successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteSetCollectionTarget(ctx *ExampleContext) {
	fmt.Println("=== Execute Set Collection Target ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter collection account address: ")
	collectionAddrStr := getUserInput()
	if collectionAddrStr == "" {
		fmt.Println("No collection account address provided. Cancelled.")
		return
	}

	fmt.Print("Enter new collection target address: ")
	targetAddrStr := getUserInput()
	if targetAddrStr == "" {
		fmt.Println("No target address provided. Cancelled.")
		return
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Collection account: %s\n", collectionAddrStr)
	fmt.Printf("  New target: %s\n", targetAddrStr)
	fmt.Print("\nType 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Transaction cancelled.")
		return
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := contracts.NewEnterpriseWallet(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.SetCollectionTarget(auth, common.HexToAddress(collectionAddrStr), common.HexToAddress(targetAddrStr))
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ Collection target updated successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteCollectFunds(ctx *ExampleContext) {
	fmt.Println("=== Execute Collect Funds ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter token address (press Enter for ETH): ")
	tokenAddrStr := getUserInput()
	var tokenAddr common.Address
	if tokenAddrStr == "" {
		tokenAddr = common.Address{} // address(0) for ETH
		fmt.Println("Using address(0) for ETH")
	} else {
		tokenAddr = common.HexToAddress(tokenAddrStr)
	}

	fmt.Print("Enter collection account address: ")
	collectionAddrStr := getUserInput()
	if collectionAddrStr == "" {
		fmt.Println("No collection account address provided. Cancelled.")
		return
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Token: %s\n", tokenAddr.Hex())
	fmt.Printf("  Collection account: %s\n", collectionAddrStr)
	fmt.Print("\nType 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Transaction cancelled.")
		return
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := contracts.NewEnterpriseWallet(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.CollectFunds(auth, tokenAddr, common.HexToAddress(collectionAddrStr))
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ Funds collected successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteUpdateMethodController(ctx *ExampleContext) {
	fmt.Println("=== Execute Update Method Controller ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter method selector (e.g., 0x08f25c4a): ")
	methodSigStr := getUserInput()
	if methodSigStr == "" {
		fmt.Println("No method selector provided. Cancelled.")
		return
	}

	// Parse method selector
	var methodSig [4]byte
	methodSigStr = strings.TrimPrefix(methodSigStr, "0x")
	methodBytes := common.FromHex(methodSigStr)
	if len(methodBytes) != 4 {
		fmt.Println("Invalid method selector (must be 4 bytes)")
		return
	}
	copy(methodSig[:], methodBytes)

	fmt.Print("Enter new controller address: ")
	controllerStr := getUserInput()
	if controllerStr == "" {
		fmt.Println("No controller address provided. Cancelled.")
		return
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Method selector: 0x%x\n", methodSig)
	fmt.Printf("  New controller: %s\n", controllerStr)
	fmt.Print("\nType 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Transaction cancelled.")
		return
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := contracts.NewEnterpriseWallet(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.UpdateMethodController(auth, methodSig, common.HexToAddress(controllerStr))
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ Method controller updated successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteUpdatePaymentAccountController(ctx *ExampleContext) {
	fmt.Println("=== Execute Update Payment Account Controller ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter payment account address: ")
	paymentAddrStr := getUserInput()
	if paymentAddrStr == "" {
		fmt.Println("No payment account address provided. Cancelled.")
		return
	}

	fmt.Print("Enter new controller address: ")
	controllerStr := getUserInput()
	if controllerStr == "" {
		fmt.Println("No controller address provided. Cancelled.")
		return
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Payment account: %s\n", paymentAddrStr)
	fmt.Printf("  New controller: %s\n", controllerStr)
	fmt.Print("\nType 'yes' to confirm: ")

	confirmation := getUserInput()
	if confirmation != "yes" {
		fmt.Println("Transaction cancelled.")
		return
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := contracts.NewEnterpriseWallet(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.UpdatePaymentAccountController(auth, common.HexToAddress(paymentAddrStr), common.HexToAddress(controllerStr))
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ Payment account controller updated successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}
