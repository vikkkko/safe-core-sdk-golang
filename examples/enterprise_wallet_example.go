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
	"time"

	"github.com/ethereum/go-ethereum"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/vikkkko/safe-core-sdk-golang/api"
	"github.com/vikkkko/safe-core-sdk-golang/protocol"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
	safetypes "github.com/vikkkko/safe-core-sdk-golang/types"
)

const (
	FactoryAddress       = "0xB67cA0029C0f6DCA816913edBDBdDe8b761C3546"
	ImplementationAddr   = "0x8c09827ebfc844d09145fee5c5214ca6abbd51db"
	SafeFactoryAddress   = "0x72D89c510AFBeC255b81482C8DCC720FC8743175"
	SafeSingletonAddress = "0xeA4156684Bc8AEdfA3062A336cABf8dBe3831ee0"
	GuardFactoryAddress  = "0xYourGuardFactoryAddressHere" // TODO: Replace with actual deployed address
)

// Context holds all necessary data for examples
type ExampleContext struct {
	Client          *ethclient.Client
	RPCURL          string
	PrivateKey      *ecdsa.PrivateKey
	PrivateKeyHex   string
	FromAddress     common.Address
	ChainID         *big.Int
	SafeAPIKey      string
	SafeAPIURL      string
	FactoryContract *EnterpriseWalletFactoryContract
}

// EnterpriseWalletFactoryContract provides typed helpers for the factory contract.
type EnterpriseWalletFactoryContract struct {
	address  common.Address
	abi      ethabi.ABI
	contract *bind.BoundContract
}

// EnterpriseWalletContract provides typed helpers for the wallet contract.
type EnterpriseWalletContract struct {
	address  common.Address
	abi      ethabi.ABI
	contract *bind.BoundContract
}

type WalletAccountInfo struct {
	Account   common.Address
	CreatedAt *big.Int
	IsActive  bool
}

type SuperAdminTransfer struct {
	CurrentSuperAdmin  common.Address
	ProposedSuperAdmin common.Address
	ProposedAt         *big.Int
	Timeout            *big.Int
	IsActive           bool
}

// NewEnterpriseWalletFactoryContract creates a factory helper.
func NewEnterpriseWalletFactoryContract(address common.Address, backend bind.ContractBackend) (*EnterpriseWalletFactoryContract, error) {
	parsed, err := ethabi.JSON(strings.NewReader(utils.EnterpriseWalletFactoryABI))
	if err != nil {
		return nil, fmt.Errorf("parse factory ABI: %w", err)
	}
	bound := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &EnterpriseWalletFactoryContract{
		address:  address,
		abi:      parsed,
		contract: bound,
	}, nil
}

// NewEnterpriseWalletContract creates a wallet helper for a specific address.
func NewEnterpriseWalletContract(address common.Address, backend bind.ContractBackend) (*EnterpriseWalletContract, error) {
	parsed, err := ethabi.JSON(strings.NewReader(utils.EnterpriseWalletABI))
	if err != nil {
		return nil, fmt.Errorf("parse wallet ABI: %w", err)
	}
	bound := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &EnterpriseWalletContract{
		address:  address,
		abi:      parsed,
		contract: bound,
	}, nil
}

type factoryInitParams struct {
	Methods    [][4]byte
	Configs    []utils.MethodConfig
	SuperAdmin common.Address
}

func (f *EnterpriseWalletFactoryContract) IsImplementationWhitelisted(opts *bind.CallOpts, implementation common.Address) (bool, error) {
	var out []interface{}
	if err := f.contract.Call(opts, &out, "isImplementationWhitelisted", implementation); err != nil {
		return false, err
	}
	return *ethabi.ConvertType(out[0], new(bool)).(*bool), nil
}

func (f *EnterpriseWalletFactoryContract) GetWhitelistedImplementations(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	if err := f.contract.Call(opts, &out, "getWhitelistedImplementations"); err != nil {
		return nil, err
	}
	return *ethabi.ConvertType(out[0], new([]common.Address)).(*[]common.Address), nil
}

func (f *EnterpriseWalletFactoryContract) PredictWalletAddress(opts *bind.CallOpts, implementation common.Address, salt [32]byte, deployer common.Address) (common.Address, error) {
	var out []interface{}
	if err := f.contract.Call(opts, &out, "predictWalletAddress", implementation, salt, deployer); err != nil {
		return common.Address{}, err
	}
	return *ethabi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

func (f *EnterpriseWalletFactoryContract) CreateWallet(auth *bind.TransactOpts, implementation common.Address, salt [32]byte, params factoryInitParams) (*gethtypes.Transaction, error) {
	return f.contract.Transact(auth, "createWallet", implementation, salt, params)
}

func (w *EnterpriseWalletContract) GetPaymentAccounts(opts *bind.CallOpts) ([]WalletAccountInfo, error) {
	var out []interface{}
	if err := w.contract.Call(opts, &out, "getPaymentAccounts"); err != nil {
		return nil, err
	}
	type accountInfo struct {
		Account   common.Address
		CreatedAt *big.Int
		IsActive  bool
	}
	outArr := *ethabi.ConvertType(out[0], new([]accountInfo)).(*[]accountInfo)
	accounts := make([]WalletAccountInfo, len(outArr))
	for i, item := range outArr {
		accounts[i] = WalletAccountInfo{Account: item.Account, CreatedAt: item.CreatedAt, IsActive: item.IsActive}
	}
	return accounts, nil
}

func (w *EnterpriseWalletContract) GetCollectionAccounts(opts *bind.CallOpts) ([]WalletAccountInfo, error) {
	var out []interface{}
	if err := w.contract.Call(opts, &out, "getCollectionAccounts"); err != nil {
		return nil, err
	}
	type accountInfo struct {
		Account   common.Address
		CreatedAt *big.Int
		IsActive  bool
	}
	outArr := *ethabi.ConvertType(out[0], new([]accountInfo)).(*[]accountInfo)
	accounts := make([]WalletAccountInfo, len(outArr))
	for i, item := range outArr {
		accounts[i] = WalletAccountInfo{Account: item.Account, CreatedAt: item.CreatedAt, IsActive: item.IsActive}
	}
	return accounts, nil
}

func (w *EnterpriseWalletContract) GetMethodConfig(opts *bind.CallOpts, method [4]byte) (utils.MethodConfig, error) {
	var out []interface{}
	if err := w.contract.Call(opts, &out, "getMethodConfig", method); err != nil {
		return utils.MethodConfig{}, err
	}
	type methodConfig struct {
		Controller common.Address
	}
	result := *ethabi.ConvertType(out[0], new(methodConfig)).(*methodConfig)
	return utils.MethodConfig{Controller: result.Controller}, nil
}

func (w *EnterpriseWalletContract) IsPaymentAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	if err := w.contract.Call(opts, &out, "isPaymentAccount", account); err != nil {
		return false, err
	}
	return *ethabi.ConvertType(out[0], new(bool)).(*bool), nil
}

func (w *EnterpriseWalletContract) IsCollectionAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	if err := w.contract.Call(opts, &out, "isCollectionAccount", account); err != nil {
		return false, err
	}
	return *ethabi.ConvertType(out[0], new(bool)).(*bool), nil
}

func (w *EnterpriseWalletContract) CreatePaymentAccount(auth *bind.TransactOpts, name string, controller common.Address) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "createPaymentAccount", name, controller)
}

func (w *EnterpriseWalletContract) CreateCollectionAccount(auth *bind.TransactOpts, name string, target common.Address) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "createCollectionAccount", name, target)
}

func (w *EnterpriseWalletContract) SetCollectionTarget(auth *bind.TransactOpts, collectionAccount, target common.Address) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "setCollectionTarget", collectionAccount, target)
}

func (w *EnterpriseWalletContract) CollectFunds(auth *bind.TransactOpts, token, collectionAccount common.Address) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "collectFunds", token, collectionAccount)
}

func (w *EnterpriseWalletContract) CreateSafeAndPaymentAccount(
	auth *bind.TransactOpts,
	proxyFactory common.Address,
	safeSingleton common.Address,
	params utils.SafeSetupParams,
	name string,
) (*gethtypes.Transaction, error) {
	setup := convertSafeSetupParams(params)
	return w.contract.Transact(auth, "createSafeAndPaymentAccount", proxyFactory, safeSingleton, setup, name)
}

func (w *EnterpriseWalletContract) CreateSafeAndCollectionAccount(
	auth *bind.TransactOpts,
	proxyFactory common.Address,
	safeSingleton common.Address,
	params utils.SafeSetupParams,
	name string,
	collectionTarget common.Address,
) (*gethtypes.Transaction, error) {
	setup := convertSafeSetupParams(params)
	return w.contract.Transact(auth, "createSafeAndCollectionAccount", proxyFactory, safeSingleton, setup, name, collectionTarget)
}

func (w *EnterpriseWalletContract) UpdateMethodController(auth *bind.TransactOpts, method [4]byte, controller common.Address) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "updateMethodController", method, controller)
}

func (w *EnterpriseWalletContract) UpdateMethodControllers(auth *bind.TransactOpts, methods [][4]byte, controllers []common.Address) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "updateMethodControllers", methods, controllers)
}

func (w *EnterpriseWalletContract) SetMethodController(auth *bind.TransactOpts, methods [][4]byte, controller common.Address) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "setMethodController", methods, controller)
}

func (w *EnterpriseWalletContract) UpdatePaymentAccountController(auth *bind.TransactOpts, paymentAccount, controller common.Address) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "updatePaymentAccountController", paymentAccount, controller)
}

func (w *EnterpriseWalletContract) EmergencyFreeze(auth *bind.TransactOpts, target common.Address, freeze bool) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "emergencyFreeze", target, freeze)
}

func (w *EnterpriseWalletContract) EmergencyPause(auth *bind.TransactOpts, pause bool) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "emergencyPause", pause)
}

func (w *EnterpriseWalletContract) ProposeSuperAdminTransfer(auth *bind.TransactOpts, newSuperAdmin common.Address, timeout *big.Int) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "proposeSuperAdminTransfer", newSuperAdmin, timeout)
}

func (w *EnterpriseWalletContract) ConfirmSuperAdminTransfer(auth *bind.TransactOpts) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "confirmSuperAdminTransfer")
}

func (w *EnterpriseWalletContract) CancelSuperAdminTransfer(auth *bind.TransactOpts) (*gethtypes.Transaction, error) {
	return w.contract.Transact(auth, "cancelSuperAdminTransfer")
}

func (w *EnterpriseWalletContract) GetSuperAdminTransfer(opts *bind.CallOpts) (SuperAdminTransfer, error) {
	var out []interface{}
	if err := w.contract.Call(opts, &out, "getSuperAdminTransfer"); err != nil {
		return SuperAdminTransfer{}, err
	}
	type superAdminTransfer struct {
		CurrentSuperAdmin  common.Address
		ProposedSuperAdmin common.Address
		ProposedAt         *big.Int
		Timeout            *big.Int
		IsActive           bool
	}
	result := *ethabi.ConvertType(out[0], new(superAdminTransfer)).(*superAdminTransfer)
	return SuperAdminTransfer{
		CurrentSuperAdmin:  result.CurrentSuperAdmin,
		ProposedSuperAdmin: result.ProposedSuperAdmin,
		ProposedAt:         result.ProposedAt,
		Timeout:            result.Timeout,
		IsActive:           result.IsActive,
	}, nil
}

func (w *EnterpriseWalletContract) IsValidSuperAdminTransfer(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	if err := w.contract.Call(opts, &out, "isValidSuperAdminTransfer"); err != nil {
		return false, err
	}
	return *ethabi.ConvertType(out[0], new(bool)).(*bool), nil
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

	fmt.Printf("Connected to network (Chain ID: %s)\n", ctx.ChainID.String())
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
	safeAPIKey := os.Getenv("SAFE_API_KEY")
	safeAPIURL := os.Getenv("SAFE_API_BASE_URL")

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

	cleanPrivateKey := strings.TrimPrefix(privateKeyHex, "0x")
	// Parse private key
	privateKey, err := crypto.HexToECDSA(cleanPrivateKey)
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

	// Create factory contract instance
	factoryContract, err := NewEnterpriseWalletFactoryContract(common.HexToAddress(FactoryAddress), client)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to instantiate factory contract: %w", err)
	}

	return &ExampleContext{
		Client:          client,
		RPCURL:          rpcURL,
		PrivateKey:      privateKey,
		PrivateKeyHex:   cleanPrivateKey,
		FromAddress:     fromAddress,
		ChainID:         chainID,
		SafeAPIKey:      safeAPIKey,
		SafeAPIURL:      safeAPIURL,
		FactoryContract: factoryContract,
	}, nil
}

func showMenu() {
	fmt.Println("\n===============================================")
	fmt.Println("   Enterprise Wallet Interactive CLI")
	fmt.Println("===============================================")
	fmt.Println("\nFactory Operations:")
	fmt.Println("  1. Check implementation whitelist")
	fmt.Println("  2. Predict wallet address")
	fmt.Println("  3. Deploy enterprise wallet")
	fmt.Println("\nAccount Management:")
	fmt.Println("  4. Create payment account")
	fmt.Println("  5. Create collection account")
	fmt.Println("  6. Create Safe + Payment account")
	fmt.Println("  7. Create Safe + Collection account")
	fmt.Println("\nAccount Operations:")
	fmt.Println("  8. Set collection target")
	fmt.Println("  9. Collect funds")
	fmt.Println("  10. Update payment account controller")
	fmt.Println("\nPermission Management:")
	fmt.Println("  11. Update method controller (single)")
	fmt.Println("  12. Update method controllers (batch)")
	fmt.Println("  13. Set method controller (batch, same controller)")
	fmt.Println("\nEmergency Controls:")
	fmt.Println("  14. Emergency freeze/unfreeze")
	fmt.Println("  15. Emergency pause/unpause")
	fmt.Println("\nSuperAdmin Transfer:")
	fmt.Println("  16. Propose SuperAdmin transfer")
	fmt.Println("  17. Confirm SuperAdmin transfer")
	fmt.Println("  18. Cancel SuperAdmin transfer")
	fmt.Println("\nQuery Functions:")
	fmt.Println("  19. Query payment accounts")
	fmt.Println("  20. Query collection accounts")
	fmt.Println("  21. Query method config")
	fmt.Println("  22. Check if address is payment/collection account")
	fmt.Println("  23. Query SuperAdmin transfer status")
	fmt.Println("\nUtility:")
	fmt.Println("  24. Show method selectors")
	fmt.Println("\nSafe Transaction Operations:")
	fmt.Println("  25. Payment Account Transfer (Propose transaction)")
	fmt.Println("  26. Payment Account Approve (Propose transaction)")
	fmt.Println("  27. Confirm Safe transaction (using SDK ConfirmTransaction)")
	fmt.Println("\nGuard Factory Operations:")
	fmt.Println("  28. Get Guard Factory Implementation")
	fmt.Println("  29. Predict Guard Address")
	fmt.Println("  30. Get Existing Guard")
	fmt.Println("  31. Create New Guard")
	fmt.Println("  32. Create Safe and Update Payment Account Controller")
	fmt.Println("  0.  Exit")
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

func promptAddress(label, defaultAddr string) common.Address {
	if defaultAddr != "" {
		fmt.Printf("%s [%s]: ", label, defaultAddr)
	} else {
		fmt.Printf("%s: ", label)
	}
	input := getUserInput()
	if input == "" && defaultAddr != "" {
		return common.HexToAddress(defaultAddr)
	}
	return common.HexToAddress(input)
}

func promptBigInt(label string, defaultValue int64) *big.Int {
	fmt.Printf("%s [%d]: ", label, defaultValue)
	input := getUserInput()
	if input == "" {
		return big.NewInt(defaultValue)
	}
	val, ok := new(big.Int).SetString(input, 10)
	if !ok {
		return big.NewInt(defaultValue)
	}
	return val
}

func confirmSend() bool {
	fmt.Print("\nSend transaction? (yes/no) [no]: ")
	return getUserInput() == "yes"
}

func runExample(ctx *ExampleContext, choice string) {
	fmt.Println()
	switch choice {
	case "1":
		checkWhitelist(ctx)
	case "2":
		predictWalletAddress(ctx)
	case "3":
		deployWallet(ctx)
	case "4":
		createPaymentAccount(ctx)
	case "5":
		createCollectionAccount(ctx)
	case "6":
		createSafeAndPaymentAccount(ctx)
	case "7":
		createSafeAndCollectionAccount(ctx)
	case "8":
		setCollectionTarget(ctx)
	case "9":
		collectFunds(ctx)
	case "10":
		updatePaymentAccountController(ctx)
	case "11":
		updateMethodController(ctx)
	case "12":
		updateMethodControllers(ctx)
	case "13":
		setMethodController(ctx)
	case "14":
		emergencyFreeze(ctx)
	case "15":
		emergencyPause(ctx)
	case "16":
		proposeSuperAdminTransfer(ctx)
	case "17":
		confirmSuperAdminTransfer(ctx)
	case "18":
		cancelSuperAdminTransfer(ctx)
	case "19":
		queryPaymentAccounts(ctx)
	case "20":
		queryCollectionAccounts(ctx)
	case "21":
		queryMethodConfig(ctx)
	case "22":
		checkAccountType(ctx)
	case "23":
		querySuperAdminTransfer(ctx)
	case "24":
		showMethodSelectors()
	case "25":
		paymentAccountTransfer(ctx)
	case "26":
		paymentAccountApprove(ctx)
	case "27":
		confirmSafeTransactionSDK(ctx)
	case "28":
		getGuardFactoryImplementation(ctx)
	case "29":
		predictGuardAddress(ctx)
	case "30":
		getExistingGuard(ctx)
	case "31":
		createNewGuard(ctx)
	case "32":
		createSafeAndUpdatePaymentAccountController(ctx)
	default:
		fmt.Println("Invalid choice.")
	}
}

// ============= Operations =============

func checkWhitelist(ctx *ExampleContext) {
	fmt.Println("=== Check Implementation Whitelist ===")

	// Check specific implementation
	isWhitelisted, err := ctx.FactoryContract.IsImplementationWhitelisted(
		&bind.CallOpts{},
		common.HexToAddress(ImplementationAddr),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Implementation %s is whitelisted: %v\n\n", ImplementationAddr, isWhitelisted)

	// Get all whitelisted
	implementations, err := ctx.FactoryContract.GetWhitelistedImplementations(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Total whitelisted implementations: %d\n", len(implementations))
	for i, impl := range implementations {
		fmt.Printf("  %d. %s\n", i+1, impl.Hex())
	}
}

func predictWalletAddress(ctx *ExampleContext) {
	fmt.Println("=== Predict Wallet Address ===")

	saltStr := prompt("Enter salt (or press Enter for default)")
	var salt [32]byte
	if saltStr == "" {
		copy(salt[:], []byte("my-enterprise-wallet-v1"))
	} else {
		copy(salt[:], []byte(saltStr))
	}

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

	fmt.Printf("\nSalt: %s\n", saltStr)
	fmt.Printf("Deployer: %s\n", ctx.FromAddress.Hex())
	fmt.Printf("Predicted address: %s\n", predictedAddr.Hex())
}

func deployWallet(ctx *ExampleContext) {
	fmt.Println("=== Deploy Enterprise Wallet ===")

	var salt [32]byte
	saltStr := prompt("Enter salt")
	copy(salt[:], []byte(saltStr))

	// 手动输入Safe地址作为controller和superAdmin
	safeAddr := promptAddress("Safe address (for controller and superAdmin)", "")
	safeAddress := safeAddr.Hex()

	// Prepare init params
	methodSelectors := [][4]byte{
		utils.CreatePaymentAccountSelector,
		utils.CreateSafeAndPaymentAccountSelector,
		utils.CreateCollectionAccountSelector,
		utils.CollectFundsSelector,
		utils.CreateAndUpdatePaymentAccountControllerSelector,
	}

	configs := make([]utils.MethodConfig, len(methodSelectors))
	for i := range methodSelectors {
		configs[i] = utils.MethodConfig{Controller: common.HexToAddress(safeAddress)}
	}

	contractInitParams := factoryInitParams{
		Methods:    methodSelectors,
		Configs:    configs,
		SuperAdmin: common.HexToAddress(safeAddress),
	}

	// Show calldata
	deployData, err := utils.CreateEnterpriseWalletData(
		common.HexToAddress(ImplementationAddr),
		salt,
		utils.InitParams{
			Methods:    methodSelectors,
			Configs:    configs,
			SuperAdmin: common.HexToAddress(safeAddress),
		},
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", deployData)
	fmt.Printf("Calldata length: %d bytes\n", len(deployData))

	if !confirmSend() {
		fmt.Println("Transaction cancelled.")
		return
	}

	// Send transaction
	factoryAddr := common.HexToAddress(FactoryAddress)
	auth := getAuth(ctx, &factoryAddr, deployData)
	tx, err := ctx.FactoryContract.CreateWallet(auth, common.HexToAddress(ImplementationAddr), salt, contractInitParams)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func createPaymentAccount(ctx *ExampleContext) {
	fmt.Println("=== Create Payment Account (via Multisig) ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	safeAddr := promptAddress("Safe multisig address", "")
	safeAddress := safeAddr.Hex()
	accountName := prompt("Account name")
	controller := promptAddress("Controller address", ctx.FromAddress.Hex())

	// Prepare calldata
	data, err := utils.CreatePaymentAccountData(accountName, controller)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Will be called via Safe multisig: %s\n", safeAddress)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// 创建Safe客户端和API客户端
	fmt.Printf("\n🔧 创建Safe客户端...")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress,
		RpcURL:      os.Getenv("RPC_URL"),
		ChainID:     ctx.ChainID.Int64(),
	})
	if err != nil {
		log.Printf("创建Safe客户端失败: %v", err)
		return
	}

	apiKey := os.Getenv("SAFE_API_KEY")
	apiClient, err := api.NewSafeApiKit(api.SafeApiKitConfig{
		ChainID: ctx.ChainID.Int64(),
		ApiKey:  apiKey,
	})
	if err != nil {
		log.Printf("创建API客户端失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// 选择 channel
	channelStr := prompt("交易 channel [0]")
	channel := uint64(0)
	if channelStr != "" {
		if parsed, err := strconv.ParseUint(channelStr, 10, 64); err == nil {
			channel = parsed
		} else {
			log.Printf("无效的 channel，默认使用 0")
		}
	}

	// 获取Safe信息
	fmt.Printf("📊 获取Safe信息...")
	safeInfo, err := apiClient.GetSafeInfo(context.Background(), safeAddress)
	if err != nil {
		log.Printf("获取Safe信息失败: %v", err)
		return
	}

	// 读取版本，按版本选择 nonce 来源
	version, err := safeClient.GetVersion(context.Background())
	if err != nil {
		log.Printf("读取 Safe 版本失败: %v", err)
		return
	}
	isMulti := strings.Contains(version, "multichannel")

	currentNonce := uint64(0)
	if isMulti {
		onChainNonce, err := safeClient.GetChannelNonce(context.Background(), channel)
		if err != nil {
			log.Printf("读取链上 channel %d nonce 失败: %v", channel, err)
			return
		}
		currentNonce = onChainNonce
	} else {
		onChainNonce, err := safeClient.GetNonce(context.Background())
		if err != nil {
			log.Printf("读取链上 nonce 失败: %v", err)
			return
		}
		currentNonce = onChainNonce
	}

	fmt.Printf(" ✅ (版本: %s, 阈值: %d/%d, channel %d nonce: %d)\n", version, safeInfo.Threshold, len(safeInfo.Owners), channel, currentNonce)

	// 创建Safe交易
	fmt.Printf("📋 创建Safe交易...")
	txData := safetypes.SafeTransactionDataPartial{
		To:      walletAddr.Hex(),
		Value:   "0",
		Data:    "0x" + hex.EncodeToString(data),
		Nonce:   &currentNonce,
		Channel: &channel,
	}

	transaction, err := safeClient.CreateTransaction(context.Background(), txData)
	if err != nil {
		log.Printf("创建交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// 从Safe合约获取交易哈希
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

	// 选择签名私钥
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

	signerPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(selectedPrivateKey, "0x"))
	if err != nil {
		log.Printf("解析 %s 失败: %v", keyLabel, err)
		return
	}
	signerAddress := crypto.PubkeyToAddress(signerPrivateKey.PublicKey)
	fmt.Printf("使用私钥: %s (%s)\n", keyLabel, signerAddress.Hex())

	// 签名交易
	fmt.Printf("\n✍️  签名交易...")
	signature, err := utils.SignMessage(txHash, signerPrivateKey)
	if err != nil {
		log.Printf("签名交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n   签名者: %s\n", signerAddress.Hex())

	// 提交到Safe服务
	fmt.Printf("\n📤 提交交易到Safe服务...")
	proposal := api.ProposeTransactionProps{
		Channel:                 func() *int64 { v := int64(transaction.Data.Channel); return &v }(),
		SafeAddress:             safeAddress,
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
		Sender:                  signerAddress.Hex(),
		Signature:               "0x" + hex.EncodeToString(signature),
		ContractTransactionHash: "0x" + safeTxHash,
	}

	response, err := apiClient.ProposeTransaction(context.Background(), proposal)
	if err != nil {
		log.Printf("提交失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")
	fmt.Printf("   SAFE Transaction Hash: %s\n", response.SafeTxHash)
	fmt.Printf("   当前签名数: %d/%d\n", len(response.Confirmations), safeInfo.Threshold)

	// 查看交易签名状态
	fmt.Printf("\n📋 获取交易签名详情...")
	txDetails, err := apiClient.GetMultisigTransaction(context.Background(), response.SafeTxHash)
	if err != nil {
		log.Printf("获取失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	fmt.Printf("\n=== 交易签名状态 ===\n")
	fmt.Printf("Safe地址: %s\n", txDetails.Safe)
	fmt.Printf("目标合约: %s\n", txDetails.To)
	fmt.Printf("交易哈希: %s\n", txDetails.SafeTxHash)
	fmt.Printf("随机数: %d\n", txDetails.Nonce)
	fmt.Printf("需要签名数: %d\n", txDetails.ConfirmationsRequired)
	fmt.Printf("当前签名数: %d\n", len(txDetails.Confirmations))
	fmt.Printf("已执行: %v\n", txDetails.IsExecuted)

	if len(txDetails.Confirmations) > 0 {
		fmt.Printf("\n已签名地址:\n")
		for i, confirmation := range txDetails.Confirmations {
			fmt.Printf("  %d. %s\n", i+1, confirmation.Owner)
			fmt.Printf("     签名: %s...\n", confirmation.Signature[:20])
			fmt.Printf("     时间: %s\n", confirmation.SubmissionDate.Format("2006-01-02 15:04:05"))
		}
	}

	// 判断是否可以执行
	if len(txDetails.Confirmations) >= txDetails.ConfirmationsRequired {
		fmt.Printf("\n✅ 交易已收集足够签名，可以执行!\n")
		fmt.Printf("\n💡 说明: 这是通过Safe多签调用Enterprise Wallet的createPaymentAccount\n")
		fmt.Printf("   - 当前是1/1多签，已经可以执行\n")
		fmt.Printf("   - 可以在Safe界面或使用SDK执行交易\n")
	} else {
		need := txDetails.ConfirmationsRequired - len(txDetails.Confirmations)
		fmt.Printf("\n⏳ 还需要 %d 个签名才能执行\n", need)

		// 列出待签名地址
		fmt.Printf("\n待签名地址:\n")
		for _, owner := range safeInfo.Owners {
			isSigned := false
			for _, confirmation := range txDetails.Confirmations {
				if strings.EqualFold(confirmation.Owner, owner) {
					isSigned = true
					break
				}
			}
			if !isSigned {
				fmt.Printf("  - %s\n", owner)
			}
		}
	}
}

// createPaymentAccountDirect 是直接调用（EOA方式）的备用方法
func createPaymentAccountDirect(ctx *ExampleContext) {
	fmt.Println("=== Create Payment Account (Direct Call) ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	accountName := prompt("Account name")
	controller := promptAddress("Controller address", ctx.FromAddress.Hex())

	// Prepare calldata
	data, err := utils.CreatePaymentAccountData(accountName, controller)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)
	fmt.Printf("Calldata length: %d bytes\n", len(data))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Send transaction
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.CreatePaymentAccount(auth, accountName, controller)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func createCollectionAccount(ctx *ExampleContext) {
	fmt.Println("=== Create Collection Account ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	accountName := prompt("Account name")
	target := promptAddress("Collection target (or press Enter for wallet itself)", "0x0000000000000000000000000000000000000000")

	// Prepare calldata
	data, err := utils.CreateCollectionAccountData(accountName, target)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)
	fmt.Printf("Calldata length: %d bytes\n", len(data))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Send transaction
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.CreateCollectionAccount(auth, accountName, target)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func createSafeAndPaymentAccount(ctx *ExampleContext) {
	fmt.Println("=== Create Safe + Payment Account (via Multisig) ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	safeAddr := promptAddress("Safe multisig address(Super Admin)", "")
	safeAddress := safeAddr.Hex()
	safeProxyFactory := promptAddress("Safe proxy factory", SafeFactoryAddress)
	safeSingleton := promptAddress("Safe singleton", SafeSingletonAddress)
	accountName := prompt("Payment account name")

	// 从环境变量读取三个 owner 私钥并转换为地址
	var owners []common.Address
	var ownerPrivateKeys []*ecdsa.PrivateKey
	ownerKeyHexes := []string{
		os.Getenv("OWNER_PRIVATE_KEY"),
		os.Getenv("OWNER2_PRIVATE_KEY"),
		os.Getenv("OWNER3_PRIVATE_KEY"),
	}

	for i, keyHex := range ownerKeyHexes {
		if keyHex == "" {
			log.Printf("Warning: OWNER%d_PRIVATE_KEY not set in .env, skipping", i+1)
			continue
		}
		cleanKey := strings.TrimPrefix(keyHex, "0x")
		privateKey, err := crypto.HexToECDSA(cleanKey)
		if err != nil {
			log.Printf("Warning: Failed to parse OWNER%d_PRIVATE_KEY: %v", i+1, err)
			continue
		}
		ownerAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
		owners = append(owners, ownerAddr)
		ownerPrivateKeys = append(ownerPrivateKeys, privateKey)
	}

	if len(owners) == 0 {
		log.Printf("Error: No valid owner private keys found in .env")
		return
	}

	// 询问阈值
	thresholdStr := prompt(fmt.Sprintf("Threshold (%d owners configured) [2]", len(owners)))
	threshold := int64(2)
	if thresholdStr != "" {
		parsed, err := strconv.ParseInt(thresholdStr, 10, 64)
		if err == nil && parsed > 0 && parsed <= int64(len(owners)) {
			threshold = parsed
		}
	}

	fmt.Printf("\n=== Safe 配置 ===\n")
	fmt.Printf("Owners (%d):\n", len(owners))
	for i, owner := range owners {
		fmt.Printf("  %d. %s\n", i+1, owner.Hex())
	}
	fmt.Printf("Threshold: %d/%d\n", threshold, len(owners))

	// 询问是否需要设置 guard
	guardInput := prompt("Guard address (留空则不设置)")
	var guardAddress common.Address
	var setupTo common.Address
	var setupData []byte

	if guardInput != "" {
		if !common.IsHexAddress(guardInput) {
			log.Printf("Error: Invalid guard address")
			return
		}
		guardAddress = common.HexToAddress(guardInput)

		// 询问 SafeGuardSetter 合约地址
		guardSetterInput := prompt("SafeGuardSetter contract address (用于在 setup 时设置 guard)")
		if !common.IsHexAddress(guardSetterInput) {
			log.Printf("Error: Invalid SafeGuardSetter address")
			return
		}
		guardSetterAddress := common.HexToAddress(guardSetterInput)

		// 生成 setGuard calldata
		// 使用 utils.SafeSetGuardData 方法生成标准的 setGuard calldata
		var err error
		setupData, err = utils.SafeSetGuardData(guardAddress)
		if err != nil {
			log.Printf("Error generating setGuard calldata: %v", err)
			return
		}

		setupTo = guardSetterAddress
		fmt.Printf("✅ 将通过 SafeGuardSetter (%s) 设置 Guard: %s\n", guardSetterAddress.Hex(), guardAddress.Hex())
	} else {
		fmt.Println("ℹ️  不设置 Guard")
	}

	params := utils.SafeSetupParams{
		Owners:          owners,
		Threshold:       big.NewInt(threshold),
		To:              setupTo,   // SafeGuardSetter 合约地址（如果需要设置 guard）
		Data:            setupData, // setGuard calldata（如果需要设置 guard）
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: owners[0],                     // 使用第一个 owner 作为 payment receiver
		SaltNonce:       big.NewInt(time.Now().Unix()), // 使用时间戳确保每次部署地址不同
	}

	// 预测 Safe 地址
	fmt.Printf("\n📍 预测 Safe 地址...")
	predictedSafeAddr, err := predictSafeAddress(safeProxyFactory, safeSingleton, params, ctx.Client)
	if err != nil {
		log.Printf("预测 Safe 地址失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")
	fmt.Printf("🎯 将要创建的 Safe 地址: %s\n", predictedSafeAddr.Hex())
	fmt.Printf("📝 SaltNonce: %s\n\n", params.SaltNonce.String())

	// Prepare calldata for enterprise wallet
	data, err := utils.CreateSafeAndPaymentAccountData(safeProxyFactory, safeSingleton, params, accountName)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Will be called via Safe multisig: %s\n", safeAddress)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	log.Printf("ctx.RPCURL: %s", ctx.RPCURL)
	log.Printf("SAFE_API_BASE_URL: %s", ctx.SafeAPIURL)

	// 创建Safe客户端和API客户端
	fmt.Printf("\n🔧 创建Safe客户端...")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress,
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  ctx.PrivateKeyHex,
	})
	if err != nil {
		log.Printf("创建Safe客户端失败: %v", err)
		return
	}

	if ctx.SafeAPIKey == "" {
		fmt.Printf(" ⚠️  SAFE_API_KEY 未配置，将在签名后直接通过 SDK 执行\n")
	}

	apiConfig := api.SafeApiKitConfig{
		ChainID:      ctx.ChainID.Int64(),
		TxServiceURL: ctx.SafeAPIURL,
		// ApiKey:  ctx.SafeAPIKey,
	}
	apiClient, err := api.NewSafeApiKit(apiConfig)
	if err != nil {
		log.Printf("创建API客户端失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// 选择 channel
	channelStr := prompt("交易 channel [0]")
	channel := uint64(0)
	if channelStr != "" {
		if parsed, err := strconv.ParseUint(channelStr, 10, 64); err == nil {
			channel = parsed
		} else {
			log.Printf("无效的 channel，默认使用 0")
		}
	}

	// 获取Safe信息
	fmt.Printf("📊 获取Safe信息...")
	safeInfo, err := apiClient.GetSafeInfo(context.Background(), safeAddress)
	if err != nil {
		log.Printf("获取Safe信息失败: %v", err)
		return
	}

	// 读取版本，按版本选择 nonce 来源
	version, err := safeClient.GetVersion(context.Background())
	if err != nil {
		log.Printf("读取 Safe 版本失败: %v", err)
		return
	}
	isMulti := strings.Contains(version, "multichannel")

	currentNonce := uint64(0)
	if isMulti {
		onChainNonce, err := safeClient.GetChannelNonce(context.Background(), channel)
		if err != nil {
			log.Printf("读取链上 channel %d nonce 失败: %v", channel, err)
			return
		}
		currentNonce = onChainNonce
	} else {
		onChainNonce, err := safeClient.GetNonce(context.Background())
		if err != nil {
			log.Printf("读取链上 nonce 失败: %v", err)
			return
		}
		currentNonce = onChainNonce
	}

	fmt.Printf(" ✅ (版本: %s, 阈值: %d/%d, channel %d nonce: %d)\n", version, safeInfo.Threshold, len(safeInfo.Owners), channel, currentNonce)

	// 创建Safe交易
	fmt.Printf("📋 创建Safe交易...")
	txData := safetypes.SafeTransactionDataPartial{
		To:      walletAddr.Hex(),
		Value:   "0",
		Data:    "0x" + hex.EncodeToString(data),
		Nonce:   &currentNonce,
		Channel: &channel,
	}

	transaction, err := safeClient.CreateTransaction(context.Background(), txData)
	if err != nil {
		log.Printf("创建交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// 从Safe合约获取交易哈希
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

	// 选择签名私钥
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

	signerPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(selectedPrivateKey, "0x"))
	if err != nil {
		log.Printf("解析 %s 失败: %v", keyLabel, err)
		return
	}
	signerAddress := crypto.PubkeyToAddress(signerPrivateKey.PublicKey)
	fmt.Printf("使用私钥: %s (%s)\n", keyLabel, signerAddress.Hex())

	// 签名交易
	fmt.Printf("\n✍️  签名交易...")
	signature, err := utils.SignMessage(txHash, signerPrivateKey)
	if err != nil {
		log.Printf("签名交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n   签名者: %s\n", signerAddress.Hex())

	transaction.AddSignature(safetypes.SafeSignature{
		Signer: signerAddress.Hex(),
		Data:   "0x" + hex.EncodeToString(signature),
	})

	if ctx.SafeAPIKey == "" {
		// 询问用户是否执行交易
		fmt.Println("\n=== 执行交易 ===")
		executeChoice := prompt("签名已添加，是否立即执行交易? (yes/no) [no]")
		if strings.ToLower(executeChoice) == "yes" || strings.ToLower(executeChoice) == "y" {
			result, execErr := safeClient.ExecuteTransaction(context.Background(), transaction)
			if execErr != nil {
				log.Printf("执行Safe交易失败: %v", execErr)
				return
			}
			fmt.Printf("🚀 Safe 交易已执行! Tx hash: %s\n", result.Hash)
			if txObj, ok := result.TransactionResponse.(*gethtypes.Transaction); ok {
				waitForTransaction(ctx, txObj)
			}
		} else {
			fmt.Printf("✅ 签名已添加，但未执行交易\n")
			fmt.Printf("💡 您可以稍后手动执行此交易\n")
		}
		return
	}

	// 提交到Safe服务
	fmt.Printf("\n📤 提交交易到Safe服务...")
	proposal := api.ProposeTransactionProps{
		Channel:                 func() *int64 { v := int64(transaction.Data.Channel); return &v }(),
		SafeAddress:             safeAddress,
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
		Sender:                  signerAddress.Hex(),
		Signature:               "0x" + hex.EncodeToString(signature),
		ContractTransactionHash: "0x" + safeTxHash,
	}

	_, err = apiClient.ProposeTransaction(context.Background(), proposal)
	if err != nil {
		log.Printf("提交失败: %v", err)
		return
	}

	// 等待30s
	fmt.Println("等待30s后获取交易签名详情")
	time.Sleep(30 * time.Second)

	// 查看交易签名状态
	fmt.Printf("\n📋 获取交易签名详情...")
	txDetails, err := apiClient.GetMultisigTransaction(context.Background(), "0x"+safeTxHash)
	if err != nil {
		log.Printf("获取失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	fmt.Printf("\n=== 交易签名状态 ===\n")
	fmt.Printf("Safe地址: %s\n", txDetails.Safe)
	fmt.Printf("目标合约: %s\n", txDetails.To)
	fmt.Printf("交易哈希: %s\n", txDetails.SafeTxHash)
	fmt.Printf("随机数: %d\n", txDetails.Nonce)
	fmt.Printf("需要签名数: %d\n", txDetails.ConfirmationsRequired)
	fmt.Printf("当前签名数: %d\n", len(txDetails.Confirmations))
	fmt.Printf("已执行: %v\n", txDetails.IsExecuted)

	if len(txDetails.Confirmations) > 0 {
		fmt.Printf("\n已签名地址:\n")
		for i, confirmation := range txDetails.Confirmations {
			fmt.Printf("  %d. %s\n", i+1, confirmation.Owner)
			fmt.Printf("     签名: %s...\n", confirmation.Signature[:20])
			fmt.Printf("     时间: %s\n", confirmation.SubmissionDate.Format("2006-01-02 15:04:05"))
		}
	}

	// 判断是否可以执行
	if len(txDetails.Confirmations) >= txDetails.ConfirmationsRequired {
		fmt.Printf("\n✅ 交易已收集足够签名，可以执行!\n")

		// 询问用户是否执行交易
		fmt.Println("\n=== 执行交易 ===")
		executeChoice := prompt("是否立即执行交易? (yes/no) [no]")

		if strings.ToLower(executeChoice) == "yes" || strings.ToLower(executeChoice) == "y" {
			safeTx, buildErr := buildSafeTransactionFromDetails(txDetails)
			if buildErr != nil {
				log.Printf("构建Safe交易失败: %v", buildErr)
				return
			}

			result, execErr := safeClient.ExecuteTransaction(context.Background(), safeTx)
			if execErr != nil {
				log.Printf("执行Safe交易失败: %v", execErr)
				return
			}

			fmt.Printf("🚀 Safe 交易已执行! Tx hash: %s\n", result.Hash)
		} else {
			fmt.Printf("✅ 交易已达到执行阈值，但未执行\n")
			fmt.Printf("💡 您可以稍后手动执行此交易\n")
		}
	} else {
		need := txDetails.ConfirmationsRequired - len(txDetails.Confirmations)
		fmt.Printf("\n⏳ 还需要 %d 个签名才能执行\n", need)

		// 列出待签名地址
		fmt.Printf("\n待签名地址:\n")
		for _, owner := range safeInfo.Owners {
			isSigned := false
			for _, confirmation := range txDetails.Confirmations {
				if strings.EqualFold(confirmation.Owner, owner) {
					isSigned = true
					break
				}
			}
			if !isSigned {
				fmt.Printf("  - %s\n", owner)
			}
		}
	}
}

// createSafeAndPaymentAccountDirect 是直接调用（EOA方式）的备用方法
func createSafeAndPaymentAccountDirect(ctx *ExampleContext) {
	fmt.Println("=== Create Safe + Payment Account (Direct Call) ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	safeProxyFactory := promptAddress("Safe proxy factory", SafeFactoryAddress)
	safeSingleton := promptAddress("Safe singleton", SafeSingletonAddress)
	accountName := prompt("Payment account name")

	// 从环境变量读取三个 owner 私钥并转换为地址
	var owners []common.Address
	ownerKeys := []string{
		os.Getenv("OWNER_PRIVATE_KEY"),
		os.Getenv("OWNER2_PRIVATE_KEY"),
		os.Getenv("OWNER3_PRIVATE_KEY"),
	}

	for i, keyHex := range ownerKeys {
		if keyHex == "" {
			log.Printf("Warning: OWNER%d_PRIVATE_KEY not set in .env, skipping", i+1)
			continue
		}
		cleanKey := strings.TrimPrefix(keyHex, "0x")
		privateKey, err := crypto.HexToECDSA(cleanKey)
		if err != nil {
			log.Printf("Warning: Failed to parse OWNER%d_PRIVATE_KEY: %v", i+1, err)
			continue
		}
		ownerAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
		owners = append(owners, ownerAddr)
	}

	if len(owners) == 0 {
		log.Printf("Error: No valid owner private keys found in .env")
		return
	}

	// 询问阈值
	thresholdStr := prompt(fmt.Sprintf("Threshold (%d owners configured) [2]", len(owners)))
	threshold := int64(2)
	if thresholdStr != "" {
		parsed, err := strconv.ParseInt(thresholdStr, 10, 64)
		if err == nil && parsed > 0 && parsed <= int64(len(owners)) {
			threshold = parsed
		}
	}

	fmt.Printf("\n=== Safe 配置 ===\n")
	fmt.Printf("Owners (%d):\n", len(owners))
	for i, owner := range owners {
		fmt.Printf("  %d. %s\n", i+1, owner.Hex())
	}
	fmt.Printf("Threshold: %d/%d\n", threshold, len(owners))

	params := utils.SafeSetupParams{
		Owners:          owners,
		Threshold:       big.NewInt(threshold),
		To:              common.Address{},
		Data:            []byte{},
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: owners[0],                     // 使用第一个 owner 作为 payment receiver
		SaltNonce:       big.NewInt(time.Now().Unix()), // 使用时间戳确保每次部署地址不同
	}

	// Prepare calldata
	data, err := utils.CreateSafeAndPaymentAccountData(safeProxyFactory, safeSingleton, params, accountName)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)
	fmt.Printf("Calldata length: %d bytes\n", len(data))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Send transaction
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.CreateSafeAndPaymentAccount(auth, safeProxyFactory, safeSingleton, params, accountName)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func createSafeAndCollectionAccount(ctx *ExampleContext) {
	fmt.Println("=== Create Safe + Collection Account ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	safeProxyFactory := promptAddress("Safe proxy factory", SafeFactoryAddress)
	safeSingleton := promptAddress("Safe singleton", SafeSingletonAddress)
	accountName := prompt("Collection account name")
	collectionTarget := promptAddress("Collection target", ctx.FromAddress.Hex())

	// 从环境变量读取三个 owner 私钥并转换为地址
	var owners []common.Address
	ownerKeys := []string{
		os.Getenv("OWNER_PRIVATE_KEY"),
		os.Getenv("OWNER2_PRIVATE_KEY"),
		os.Getenv("OWNER3_PRIVATE_KEY"),
	}

	for i, keyHex := range ownerKeys {
		if keyHex == "" {
			log.Printf("Warning: OWNER%d_PRIVATE_KEY not set in .env, skipping", i+1)
			continue
		}
		cleanKey := strings.TrimPrefix(keyHex, "0x")
		privateKey, err := crypto.HexToECDSA(cleanKey)
		if err != nil {
			log.Printf("Warning: Failed to parse OWNER%d_PRIVATE_KEY: %v", i+1, err)
			continue
		}
		ownerAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
		owners = append(owners, ownerAddr)
	}

	if len(owners) == 0 {
		log.Printf("Error: No valid owner private keys found in .env")
		return
	}

	// 询问阈值
	thresholdStr := prompt(fmt.Sprintf("Threshold (%d owners configured) [2]", len(owners)))
	threshold := int64(2)
	if thresholdStr != "" {
		parsed, err := strconv.ParseInt(thresholdStr, 10, 64)
		if err == nil && parsed > 0 && parsed <= int64(len(owners)) {
			threshold = parsed
		}
	}

	fmt.Printf("\n=== Safe 配置 ===\n")
	fmt.Printf("Owners (%d):\n", len(owners))
	for i, owner := range owners {
		fmt.Printf("  %d. %s\n", i+1, owner.Hex())
	}
	fmt.Printf("Threshold: %d/%d\n", threshold, len(owners))

	params := utils.SafeSetupParams{
		Owners:          owners,
		Threshold:       big.NewInt(threshold),
		To:              common.Address{},
		Data:            []byte{},
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: owners[0],                     // 使用第一个 owner 作为 payment receiver
		SaltNonce:       big.NewInt(time.Now().Unix()), // 使用时间戳确保每次部署地址不同
	}

	// Prepare calldata
	data, err := utils.CreateSafeAndCollectionAccountData(safeProxyFactory, safeSingleton, params, accountName, collectionTarget)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)
	fmt.Printf("Calldata length: %d bytes\n", len(data))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Send transaction
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.CreateSafeAndCollectionAccount(auth, safeProxyFactory, safeSingleton, params, accountName, collectionTarget)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func setCollectionTarget(ctx *ExampleContext) {
	fmt.Println("=== Set Collection Target ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	collectionAccount := promptAddress("Collection account", "")
	newTarget := promptAddress("New target address", "")

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, nil)
	tx, err := wallet.SetCollectionTarget(auth, collectionAccount, newTarget)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func collectFunds(ctx *ExampleContext) {
	fmt.Println("=== Collect Funds ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	tokenAddr := promptAddress("Token address (or 0x0 for ETH)", "0x0000000000000000000000000000000000000000")
	collectionAccount := promptAddress("Collection account", "")

	// Prepare calldata
	data, err := utils.CollectFundsData(tokenAddr, collectionAccount)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.CollectFunds(auth, tokenAddr, collectionAccount)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func updatePaymentAccountController(ctx *ExampleContext) {
	fmt.Println("=== Update Payment Account Controller ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	paymentAccount := promptAddress("Payment account", "")
	newController := promptAddress("New controller", "")

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, nil)
	tx, err := wallet.UpdatePaymentAccountController(auth, paymentAccount, newController)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func updateMethodController(ctx *ExampleContext) {
	fmt.Println("=== Update Method Controller (Single) ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	methodSig := utils.CreatePaymentAccountSelector
	newController := promptAddress("New controller", "")

	fmt.Printf("\nMethod: createPaymentAccount (0x%x)\n", methodSig)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, nil)
	tx, err := wallet.UpdateMethodController(auth, methodSig, newController)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func updateMethodControllers(ctx *ExampleContext) {
	fmt.Println("=== Update Method Controllers (Batch) ===")

	walletAddr := promptAddress("Enterprise wallet address", "")

	methodSigs := [][4]byte{
		utils.CreatePaymentAccountSelector,
		utils.CreateCollectionAccountSelector,
	}
	controllers := []common.Address{
		promptAddress("Controller for createPaymentAccount", ""),
		promptAddress("Controller for createCollectionAccount", ""),
	}

	// Prepare calldata
	data, err := utils.UpdateMethodControllersData(methodSigs, controllers)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.UpdateMethodControllers(auth, methodSigs, controllers)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func setMethodController(ctx *ExampleContext) {
	fmt.Println("=== Set Method Controller (Batch, Same Controller) ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	controller := promptAddress("Controller for all methods", "")

	methodSigs := [][4]byte{
		utils.CreatePaymentAccountSelector,
		utils.CreateCollectionAccountSelector,
		utils.CollectFundsSelector,
	}

	// Prepare calldata
	data, err := utils.SetMethodControllerData(methodSigs, controller)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)
	fmt.Printf("Will update %d methods\n", len(methodSigs))

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.SetMethodController(auth, methodSigs, controller)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func emergencyFreeze(ctx *ExampleContext) {
	fmt.Println("=== Emergency Freeze/Unfreeze ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	target := promptAddress("Target account", "")
	freezeStr := prompt("Freeze? (yes/no)")
	freeze := freezeStr == "yes"

	// Prepare calldata
	data, err := utils.EmergencyFreezeData(target, freeze)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.EmergencyFreeze(auth, target, freeze)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func emergencyPause(ctx *ExampleContext) {
	fmt.Println("=== Emergency Pause/Unpause ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	pauseStr := prompt("Pause? (yes/no)")
	pause := pauseStr == "yes"

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, nil)
	tx, err := wallet.EmergencyPause(auth, pause)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func proposeSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Propose SuperAdmin Transfer ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	newSuperAdmin := promptAddress("New SuperAdmin", "")
	timeout := promptBigInt("Timeout (seconds)", 86400)

	// Prepare calldata
	data, err := utils.ProposeSuperAdminTransferData(newSuperAdmin, timeout)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.ProposeSuperAdminTransfer(auth, newSuperAdmin, timeout)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func confirmSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Confirm SuperAdmin Transfer ===")

	walletAddr := promptAddress("Enterprise wallet address", "")

	// Prepare calldata
	data, err := utils.ConfirmSuperAdminTransferData()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.ConfirmSuperAdminTransfer(auth)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func cancelSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Cancel SuperAdmin Transfer ===")

	walletAddr := promptAddress("Enterprise wallet address", "")

	// Prepare calldata
	data, err := utils.CancelSuperAdminTransferData()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	auth := getAuth(ctx, &walletAddr, data)
	tx, err := wallet.CancelSuperAdminTransfer(auth)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTransaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)
}

func queryPaymentAccounts(ctx *ExampleContext) {
	fmt.Println("=== Query Payment Accounts ===")

	walletAddr := promptAddress("Enterprise wallet address", "")

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	accounts, err := wallet.GetPaymentAccounts(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nFound %d payment account(s):\n", len(accounts))
	for i, acc := range accounts {
		fmt.Printf("  %d. %s (Active: %v, Created: %s)\n", i+1, acc.Account.Hex(), acc.IsActive, acc.CreatedAt.String())
	}
}

func queryCollectionAccounts(ctx *ExampleContext) {
	fmt.Println("=== Query Collection Accounts ===")

	walletAddr := promptAddress("Enterprise wallet address", "")

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	accounts, err := wallet.GetCollectionAccounts(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nFound %d collection account(s):\n", len(accounts))
	for i, acc := range accounts {
		fmt.Printf("  %d. %s (Active: %v, Created: %s)\n", i+1, acc.Account.Hex(), acc.IsActive, acc.CreatedAt.String())
	}
}

func queryMethodConfig(ctx *ExampleContext) {
	fmt.Println("=== Query Method Config ===")

	walletAddr := promptAddress("Enterprise wallet address", "")

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	methodSig := utils.CreatePaymentAccountSelector
	config, err := wallet.GetMethodConfig(&bind.CallOpts{}, methodSig)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nMethod: createPaymentAccount (0x%x)\n", methodSig)
	fmt.Printf("Controller: %s\n", config.Controller.Hex())

	methodSig = utils.CreateSafeAndPaymentAccountSelector
	config, err = wallet.GetMethodConfig(&bind.CallOpts{}, methodSig)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nMethod: createSafeAndPaymentAccount (0x%x)\n", methodSig)
	fmt.Printf("Controller: %s\n", config.Controller.Hex())

}

func checkAccountType(ctx *ExampleContext) {
	fmt.Println("=== Check Account Type ===")

	walletAddr := promptAddress("Enterprise wallet address", "")
	accountAddr := promptAddress("Account to check", "")

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	isPayment, err := wallet.IsPaymentAccount(&bind.CallOpts{}, accountAddr)
	if err != nil {
		log.Printf("Error checking payment account: %v", err)
		return
	}

	isCollection, err := wallet.IsCollectionAccount(&bind.CallOpts{}, accountAddr)
	if err != nil {
		log.Printf("Error checking collection account: %v", err)
		return
	}

	fmt.Printf("\nAccount: %s\n", accountAddr.Hex())
	fmt.Printf("Is Payment Account: %v\n", isPayment)
	fmt.Printf("Is Collection Account: %v\n", isCollection)
}

func querySuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Query SuperAdmin Transfer Status ===")

	walletAddr := promptAddress("Enterprise wallet address", "")

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	transfer, err := wallet.GetSuperAdminTransfer(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	isValid, err := wallet.IsValidSuperAdminTransfer(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nCurrent SuperAdmin: %s\n", transfer.CurrentSuperAdmin.Hex())
	fmt.Printf("Proposed SuperAdmin: %s\n", transfer.ProposedSuperAdmin.Hex())
	fmt.Printf("Proposed At: %s\n", transfer.ProposedAt.String())
	fmt.Printf("Timeout: %s seconds\n", transfer.Timeout.String())
	fmt.Printf("Is Active: %v\n", transfer.IsActive)
	fmt.Printf("Is Valid: %v\n", isValid)
}

func showMethodSelectors() {
	fmt.Println("=== Method Selectors ===")
	fmt.Printf("createPaymentAccount:           0x%x\n", utils.CreatePaymentAccountSelector)
	fmt.Printf("createCollectionAccount:        0x%x\n", utils.CreateCollectionAccountSelector)
	fmt.Printf("collectFunds:                   0x%x\n", utils.CollectFundsSelector)
	fmt.Printf("createSafeAndPaymentAccount:    0x%x\n", utils.CreateSafeAndPaymentAccountSelector)
	fmt.Printf("proposeSuperAdminTransfer:      0x%x\n", utils.ProposeSuperAdminTransferSelector)
	fmt.Printf("confirmSuperAdminTransfer:      0x%x\n", utils.ConfirmSuperAdminTransferSelector)
	fmt.Printf("cancelSuperAdminTransfer:       0x%x\n", utils.CancelSuperAdminTransferSelector)
}

// paymentAccountTransfer creates and proposes a PaymentAccount transfer transaction
// This demonstrates calling PaymentAccount.transfer() through Safe multisig
func paymentAccountTransfer(ctx *ExampleContext) {
	fmt.Println("=== Payment Account Transfer (Propose Transaction) ===")

	// Get PaymentAccount address (controlled by Safe)
	paymentAccountAddr := promptAddress("Payment Account address", "")

	// Get Safe address that controls this PaymentAccount
	safeAddress := promptAddress("Safe multisig address (controller)", "")
	if safeAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		log.Printf("Error: Safe address is required")
		return
	}

	// Get transfer parameters
	fmt.Println("\n=== Transfer Parameters ===")
	tokenAddr := promptAddress("Token address (0x0 for ETH)", "0xAD2B0439ed98F50eDEB0e04F064d492bAFDAd73B")
	fromAddr := promptAddress("From address (PaymentAccount for ETH)", "0x9C126aa4Eb6D110D646139969774F2c5b64dD279")
	toAddr := promptAddress("To address (recipient)", "0xeB7E951F2D1A38188762dF12E0703aE16F76Ab73")
	amountStr := prompt("Amount (in wei)")

	amount := new(big.Int)
	amount, ok := amount.SetString(amountStr, 10)
	if !ok {
		log.Printf("Error: Invalid amount")
		return
	}

	// Generate PaymentAccount.transfer() calldata
	transferData, err := utils.PaymentAccountTransferData(tokenAddr, fromAddr, toAddr, amount)
	if err != nil {
		log.Printf("Error generating transfer calldata: %v", err)
		return
	}

	fmt.Printf("\n=== Transaction Details ===\n")
	fmt.Printf("Target: %s (PaymentAccount)\n", paymentAccountAddr.Hex())
	fmt.Printf("Token: %s\n", tokenAddr.Hex())
	fmt.Printf("From: %s\n", fromAddr.Hex())
	fmt.Printf("To: %s\n", toAddr.Hex())
	fmt.Printf("Amount: %s wei\n", amount.String())
	fmt.Printf("Calldata: 0x%x\n", transferData)
	fmt.Printf("Calldata length: %d bytes\n", len(transferData))
	fmt.Printf("Will be called via Safe: %s\n\n", safeAddress.Hex())

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Create Safe client and API client
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

	// 选择 channel
	channelStr := prompt("交易 channel [0]")
	channel := uint64(0)
	if channelStr != "" {
		if parsed, err := strconv.ParseUint(channelStr, 10, 64); err == nil {
			channel = parsed
		} else {
			log.Printf("无效的 channel，默认使用 0")
		}
	}

	// Get Safe info
	fmt.Printf("📊 获取Safe信息...")
	safeInfo, err := apiClient.GetSafeInfo(context.Background(), safeAddress.Hex())
	if err != nil {
		log.Printf("获取Safe信息失败: %v", err)
		return
	}

	// 读取版本，按版本选择 nonce 来源
	version, err := safeClient.GetVersion(context.Background())
	if err != nil {
		log.Printf("读取 Safe 版本失败: %v", err)
		return
	}
	isMulti := strings.Contains(version, "multichannel")

	currentNonce := uint64(0)
	if isMulti {
		onChainNonce, err := safeClient.GetChannelNonce(context.Background(), channel)
		if err != nil {
			log.Printf("读取链上 channel %d nonce 失败: %v", channel, err)
			return
		}
		currentNonce = onChainNonce
	} else {
		onChainNonce, err := safeClient.GetNonce(context.Background())
		if err != nil {
			log.Printf("读取链上 nonce 失败: %v", err)
			return
		}
		currentNonce = onChainNonce
	}

	fmt.Printf(" ✅ (版本: %s, 阈值: %d/%d, channel %d nonce: %d)\n", version, safeInfo.Threshold, len(safeInfo.Owners), channel, currentNonce)

	// Create Safe transaction
	fmt.Printf("📋 创建Safe交易...")
	txData := safetypes.SafeTransactionDataPartial{
		To:      paymentAccountAddr.Hex(),
		Value:   "0",
		Data:    "0x" + hex.EncodeToString(transferData),
		Nonce:   &currentNonce,
		Channel: &channel,
	}

	transaction, err := safeClient.CreateTransaction(context.Background(), txData)
	if err != nil {
		log.Printf("创建交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// Get transaction hash from Safe contract
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

	// 选择签名私钥
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

	signerPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(selectedPrivateKey, "0x"))
	if err != nil {
		log.Printf("解析 %s 失败: %v", keyLabel, err)
		return
	}
	signerAddress := crypto.PubkeyToAddress(signerPrivateKey.PublicKey)
	fmt.Printf("使用私钥: %s (%s)\n", keyLabel, signerAddress.Hex())

	// Sign transaction
	fmt.Printf("\n✍️  签名交易...")
	signature, err := utils.SignMessage(txHash, signerPrivateKey)
	if err != nil {
		log.Printf("签名交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n   签名者: %s\n", signerAddress.Hex())

	// Submit to Safe service
	fmt.Printf("\n📤 提交交易到Safe服务...")
	proposal := api.ProposeTransactionProps{
		Channel:                 func() *int64 { v := int64(transaction.Data.Channel); return &v }(),
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
		Sender:                  signerAddress.Hex(),
		Signature:               "0x" + hex.EncodeToString(signature),
		ContractTransactionHash: "0x" + safeTxHash,
	}

	_, err = apiClient.ProposeTransaction(context.Background(), proposal)
	if err != nil {
		log.Printf("提交失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")
}

// paymentAccountApprove creates and proposes a PaymentAccount approve transaction
// This demonstrates calling PaymentAccount.approve() through Safe multisig
func paymentAccountApprove(ctx *ExampleContext) {
	fmt.Println("=== Payment Account Approve (Propose Transaction) ===")

	// Get PaymentAccount address (controlled by Safe)
	paymentAccountAddr := promptAddress("Payment Account address", "")

	// Get Safe address that controls this PaymentAccount
	safeAddress := promptAddress("Safe multisig address (controller)", "")
	if safeAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		log.Printf("Error: Safe address is required")
		return
	}

	// Get approve parameters
	fmt.Println("\n=== Approve Parameters ===")
	tokenAddr := promptAddress("Token address (cannot be 0x0)", "0xAD2B0439ed98F50eDEB0e04F064d492bAFDAd73B")
	if tokenAddr.Hex() == "0x0000000000000000000000000000000000000000" {
		log.Printf("Error: Token address cannot be 0x0 for approve")
		return
	}
	spenderAddr := promptAddress("Spender address", "0xeB7E951F2D1A38188762dF12E0703aE16F76Ab73")
	amountStr := prompt("Amount to approve (in wei)")

	amount := new(big.Int)
	amount, ok := amount.SetString(amountStr, 10)
	if !ok {
		log.Printf("Error: Invalid amount")
		return
	}

	// Generate PaymentAccount.approve() calldata
	approveData, err := utils.PaymentAccountApproveData(tokenAddr, spenderAddr, amount)
	if err != nil {
		log.Printf("Error generating approve calldata: %v", err)
		return
	}

	fmt.Printf("\n=== Transaction Details ===\n")
	fmt.Printf("Target: %s (PaymentAccount)\n", paymentAccountAddr.Hex())
	fmt.Printf("Token: %s\n", tokenAddr.Hex())
	fmt.Printf("Spender: %s\n", spenderAddr.Hex())
	fmt.Printf("Amount: %s wei\n", amount.String())
	fmt.Printf("Calldata: 0x%x\n", approveData)
	fmt.Printf("Calldata length: %d bytes\n", len(approveData))
	fmt.Printf("Will be called via Safe: %s\n\n", safeAddress.Hex())

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// Create Safe client and API client
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

	// 选择 channel
	channelStr := prompt("交易 channel [0]")
	channel := uint64(0)
	if channelStr != "" {
		if parsed, err := strconv.ParseUint(channelStr, 10, 64); err == nil {
			channel = parsed
		} else {
			log.Printf("无效的 channel，默认使用 0")
		}
	}

	// Get Safe info
	fmt.Printf("📊 获取Safe信息...")
	safeInfo, err := apiClient.GetSafeInfo(context.Background(), safeAddress.Hex())
	if err != nil {
		log.Printf("获取Safe信息失败: %v", err)
		return
	}

	// 读取版本，按版本选择 nonce 来源
	version, err := safeClient.GetVersion(context.Background())
	if err != nil {
		log.Printf("读取 Safe 版本失败: %v", err)
		return
	}
	isMulti := strings.Contains(version, "multichannel")

	currentNonce := uint64(0)
	if isMulti {
		onChainNonce, err := safeClient.GetChannelNonce(context.Background(), channel)
		if err != nil {
			log.Printf("读取链上 channel %d nonce 失败: %v", channel, err)
			return
		}
		currentNonce = onChainNonce
	} else {
		onChainNonce, err := safeClient.GetNonce(context.Background())
		if err != nil {
			log.Printf("读取链上 nonce 失败: %v", err)
			return
		}
		currentNonce = onChainNonce
	}

	fmt.Printf(" ✅ (版本: %s, 阈值: %d/%d, channel %d nonce: %d)\n", version, safeInfo.Threshold, len(safeInfo.Owners), channel, currentNonce)

	// Create Safe transaction
	fmt.Printf("📋 创建Safe交易...")
	txData := safetypes.SafeTransactionDataPartial{
		To:      paymentAccountAddr.Hex(),
		Value:   "0",
		Data:    "0x" + hex.EncodeToString(approveData),
		Nonce:   &currentNonce,
		Channel: &channel,
	}

	transaction, err := safeClient.CreateTransaction(context.Background(), txData)
	if err != nil {
		log.Printf("创建交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// Get transaction hash from Safe contract
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

	// 选择签名私钥
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

	signerPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(selectedPrivateKey, "0x"))
	if err != nil {
		log.Printf("解析 %s 失败: %v", keyLabel, err)
		return
	}
	signerAddress := crypto.PubkeyToAddress(signerPrivateKey.PublicKey)
	fmt.Printf("使用私钥: %s (%s)\n", keyLabel, signerAddress.Hex())

	// Sign transaction
	fmt.Printf("\n✍️  签名交易...")
	signature, err := utils.SignMessage(txHash, signerPrivateKey)
	if err != nil {
		log.Printf("签名交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n   签名者: %s\n", signerAddress.Hex())

	// Submit to Safe service
	fmt.Printf("\n📤 提交交易到Safe服务...")
	proposal := api.ProposeTransactionProps{
		Channel:                 func() *int64 { v := int64(transaction.Data.Channel); return &v }(),
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
		Sender:                  signerAddress.Hex(),
		Signature:               "0x" + hex.EncodeToString(signature),
		ContractTransactionHash: "0x" + safeTxHash,
	}

	_, err = apiClient.ProposeTransaction(context.Background(), proposal)
	if err != nil {
		log.Printf("提交失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")
}

// confirmSafeTransactionSDK 使用SDK的高级ConfirmTransaction方法确认Safe交易
// 这是一个简化的方法，自动处理签名、提交和执行
func confirmSafeTransactionSDK(ctx *ExampleContext) {
	fmt.Println("=== Confirm Safe Transaction (SDK Method) ===")

	safeTxHash := prompt("Safe transaction hash (0x...)")
	if safeTxHash == "" {
		log.Printf("Error: Safe transaction hash is required")
		return
	}

	// 确保有0x前缀
	if !strings.HasPrefix(safeTxHash, "0x") && !strings.HasPrefix(safeTxHash, "0X") {
		safeTxHash = "0x" + safeTxHash
	}

	// 选择使用哪个私钥进行签名
	fmt.Println("\n=== 选择签名私钥 ===")
	fmt.Println("1. DEPLOYER_PRIVATE_KEY (当前默认)")
	fmt.Println("2. OWNER_PRIVATE_KEY")
	fmt.Println("3. OWNER2_PRIVATE_KEY")
	fmt.Println("4. OWNER3_PRIVATE_KEY")
	keyChoice := prompt("选择私钥 [1]")
	if keyChoice == "" {
		keyChoice = "1"
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
		selectedPrivateKey = os.Getenv("DEPLOYER_PRIVATE_KEY")
		keyLabel = "DEPLOYER_PRIVATE_KEY"
	}

	if selectedPrivateKey == "" {
		log.Printf("Error: %s not set in .env", keyLabel)
		return
	}

	// 解析私钥获取签名者地址
	cleanPrivateKey := strings.TrimPrefix(selectedPrivateKey, "0x")
	privateKey, err := crypto.HexToECDSA(cleanPrivateKey)
	if err != nil {
		log.Printf("Error parsing private key: %v", err)
		return
	}
	signerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Printf("使用私钥: %s\n", keyLabel)
	fmt.Printf("签名者地址: %s\n", signerAddress.Hex())

	// 创建 API 客户端
	fmt.Printf("\n🔧 创建API客户端...")
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

	// 获取交易详情以获取Safe地址和展示信息
	// 注意: SDK的ConfirmTransaction内部也会调用GetMultisigTransaction
	// 这里提前调用是为了：
	// 1. 获取Safe地址（创建Safe客户端必需）
	// 2. 向用户展示交易状态（用户体验）
	fmt.Printf("📋 获取交易详情...")
	txDetails, err := apiClient.GetMultisigTransaction(context.Background(), safeTxHash)
	if err != nil {
		log.Printf("获取交易详情失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	safeAddress := txDetails.Safe
	fmt.Printf("\n=== 交易信息 ===\n")
	fmt.Printf("Safe 地址: %s\n", safeAddress)
	fmt.Printf("目标地址: %s\n", txDetails.To)
	fmt.Printf("当前签名数: %d/%d\n", len(txDetails.Confirmations), txDetails.ConfirmationsRequired)
	fmt.Printf("已执行: %v\n", txDetails.IsExecuted)

	// true代表owner签名完就发送上链了。如果使用了代付，不要用true。
	autoExecuteStr := prompt("\nAuto-execute if threshold met? (yes/no) [no]")
	autoExecute := true
	if autoExecuteStr == "no" || autoExecuteStr == "n" {
		autoExecute = false
	}

	fmt.Printf("autoExecute: %v\n", autoExecute)

	fmt.Printf("\n🔧 创建Safe客户端...")
	// 创建Safe客户端
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress,
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  selectedPrivateKey,//ctx.PrivateKeyHex,
	})
	if err != nil {
		log.Printf("创建Safe客户端失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// 使用SDK的高级ConfirmTransaction方法
	fmt.Printf("\n🚀 调用SDK的ConfirmTransaction方法...\n")
	result, err := safeClient.ConfirmTransaction(context.Background(), protocol.ConfirmTransactionConfig{
		SafeTxHash:  safeTxHash,
		APIClient:   apiClient,
		AutoExecute: autoExecute,
	})
	if err != nil {
		log.Printf("❌ 确认交易失败: %v", err)
		return
	}

	// 显示结果
	fmt.Printf("\n=== 确认结果 ===\n")
	fmt.Printf("已经签名: %v\n", result.AlreadySigned)
	fmt.Printf("提交了新签名: %v\n", result.SignatureSubmitted)
	fmt.Printf("当前签名数: %d/%d\n", result.CurrentSignatures, result.RequiredSignatures)
	fmt.Printf("达到阈值: %v\n", result.ThresholdMet)
	fmt.Printf("已执行: %v\n", result.TransactionExecuted)

	if result.AlreadySigned {
		fmt.Printf("\n💡 您的地址已经签名过此交易\n")
	}

	if result.SignatureSubmitted {
		fmt.Printf("\n✅ 成功提交签名到Safe Transaction Service\n")
	}

	if result.ThresholdMet && !result.TransactionExecuted && !autoExecute {
		fmt.Printf("\n💡 交易已达到执行阈值，但未自动执行（AutoExecute=false）\n")
		fmt.Printf("   您可以手动执行或重新运行此功能并启用自动执行\n")
	}

	if result.TransactionExecuted && result.ExecutionResult != nil {
		fmt.Printf("\n🎉 交易已成功执行!\n")
		fmt.Printf("   执行交易哈希: %s\n", result.ExecutionResult.Hash)

		if txObj, ok := result.ExecutionResult.TransactionResponse.(*gethtypes.Transaction); ok {
			waitForTransaction(ctx, txObj)
		}
	}

	if !result.ThresholdMet {
		need := result.RequiredSignatures - result.CurrentSignatures
		fmt.Printf("\n⏳ 还需要 %d 个签名才能执行\n", need)
	}
}

// ============= Helpers =============

// getAuth creates a transaction auth with automatic gas estimation
// If calldata is provided, it will estimate gas; otherwise uses default 800000
func getAuth(ctx *ExampleContext, to *common.Address, calldata []byte) *bind.TransactOpts {
	nonce, err := ctx.Client.PendingNonceAt(context.Background(), ctx.FromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := ctx.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ctx.PrivateKey, ctx.ChainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice

	// Estimate gas limit
	var gasLimit uint64
	if to != nil && calldata != nil {
		estimated, err := ctx.Client.EstimateGas(context.Background(), ethereum.CallMsg{
			From: ctx.FromAddress,
			To:   to,
			Data: calldata,
		})
		if err != nil {
			// If estimation fails, use a safe default
			fmt.Printf("⚠️  Warning: gas estimation failed (%v), using default 800000\n", err)
			gasLimit = 800000
		} else {
			// Add 20% buffer to estimated gas
			gasLimit = estimated * 120 / 100
			fmt.Printf("Estimated gas: %d (with 20%% buffer: %d)\n", estimated, gasLimit)
		}
	} else {
		// No calldata provided, use default
		gasLimit = 800000
	}

	auth.GasLimit = gasLimit

	return auth
}

func waitForTransaction(ctx *ExampleContext, tx *gethtypes.Transaction) {
	fmt.Println("Waiting for confirmation...")

	receipt, err := bind.WaitMined(context.Background(), ctx.Client, tx)
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Printf("✓ Transaction successful!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func convertSafeSetupParams(params utils.SafeSetupParams) struct {
	Owners          []common.Address
	Threshold       *big.Int
	To              common.Address
	Data            []byte
	FallbackHandler common.Address
	PaymentToken    common.Address
	Payment         *big.Int
	PaymentReceiver common.Address
	SaltNonce       *big.Int
} {
	return struct {
		Owners          []common.Address
		Threshold       *big.Int
		To              common.Address
		Data            []byte
		FallbackHandler common.Address
		PaymentToken    common.Address
		Payment         *big.Int
		PaymentReceiver common.Address
		SaltNonce       *big.Int
	}{
		Owners:          params.Owners,
		Threshold:       safeBigInt(params.Threshold),
		To:              params.To,
		Data:            params.Data,
		FallbackHandler: params.FallbackHandler,
		PaymentToken:    params.PaymentToken,
		Payment:         safeBigInt(params.Payment),
		PaymentReceiver: params.PaymentReceiver,
		SaltNonce:       safeBigInt(params.SaltNonce),
	}
}

func safeBigInt(value *big.Int) *big.Int {
	if value == nil {
		return big.NewInt(0)
	}
	return value
}

func normalizeDataHex(value string) string {
	if value == "" {
		return "0x"
	}
	if strings.HasPrefix(value, "0x") || strings.HasPrefix(value, "0X") {
		return value
	}
	return "0x" + value
}

func normalizeAddressHex(value string) string {
	if value == "" {
		return common.Address{}.Hex()
	}
	if !strings.HasPrefix(value, "0x") && !strings.HasPrefix(value, "0X") {
		value = "0x" + value
	}
	return common.HexToAddress(value).Hex()
}

func buildSafeTransactionFromDetails(details *api.SafeMultisigTransactionResponse) (*safetypes.SafeTransaction, error) {
	if details == nil {
		return nil, fmt.Errorf("transaction details missing")
	}

	tx := &safetypes.SafeTransaction{
		Data: safetypes.SafeTransactionData{
			To:             normalizeAddressHex(details.To),
			Value:          details.Value,
			Data:           normalizeDataHex(details.Data),
			Operation:      safetypes.OperationType(details.Operation),
			SafeTxGas:      strconv.FormatInt(details.SafeTxGas, 10),
			BaseGas:        strconv.FormatInt(details.BaseGas, 10),
			GasPrice:       details.GasPrice,
			GasToken:       normalizeAddressHex(details.GasToken),
			RefundReceiver: normalizeAddressHex(details.RefundReceiver),
			Nonce:          uint64(details.Nonce),
		},
		Signatures: make(map[string]safetypes.SafeSignature),
	}

	for _, confirmation := range details.Confirmations {
		sig := confirmation.Signature
		if !strings.HasPrefix(sig, "0x") && !strings.HasPrefix(sig, "0X") {
			sig = "0x" + sig
		}
		lowerOwner := strings.ToLower(confirmation.Owner)
		tx.Signatures[lowerOwner] = safetypes.SafeSignature{
			Signer: confirmation.Owner,
			Data:   sig,
		}
	}

	return tx, nil
}

// ============= Guard Factory Operations =============

func getGuardFactoryImplementation(ctx *ExampleContext) {
	fmt.Println("=== Get Guard Factory Implementation ===")

	// Prompt for factory address
	factoryAddr := promptAddress("Guard Factory address", GuardFactoryAddress)
	fmt.Printf("\nUsing Guard Factory: %s\n", factoryAddr.Hex())

	// Create factory contract instance
	guardFactory, err := utils.NewSimpleUniversalGuardFactoryContract(factoryAddr, ctx.Client)
	if err != nil {
		log.Printf("Failed to create Guard Factory instance: %v", err)
		return
	}

	// Get implementation address
	fmt.Printf("\nQuerying implementation address...")
	implAddr, err := guardFactory.GetImplementation(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get implementation: %v", err)
		return
	}

	fmt.Printf(" ✅\n")
	fmt.Printf("\nImplementation Address: %s\n", implAddr.Hex())
}

func predictGuardAddress(ctx *ExampleContext) {
	fmt.Println("=== Predict Guard Address ===")

	// Prompt for factory address
	factoryAddr := promptAddress("Guard Factory address", GuardFactoryAddress)
	fmt.Printf("\nUsing Guard Factory: %s\n", factoryAddr.Hex())

	// Prompt for required signer
	requiredSigner := promptAddress("Required signer address", "")
	if requiredSigner == (common.Address{}) {
		log.Printf("Error: Required signer address is required")
		return
	}

	// Create factory contract instance
	guardFactory, err := utils.NewSimpleUniversalGuardFactoryContract(factoryAddr, ctx.Client)
	if err != nil {
		log.Printf("Failed to create Guard Factory instance: %v", err)
		return
	}

	// Predict guard address
	fmt.Printf("\nPredicting guard address for signer: %s...", requiredSigner.Hex())
	predictedAddr, err := guardFactory.PredictGuardAddress(&bind.CallOpts{}, requiredSigner)
	if err != nil {
		log.Printf("Failed to predict guard address: %v", err)
		return
	}

	fmt.Printf(" ✅\n")
	fmt.Printf("\nPredicted Guard Address: %s\n", predictedAddr.Hex())
	fmt.Printf("\nNote: This is the deterministic address where the guard will be deployed.\n")
	fmt.Printf("The guard contract is not yet deployed at this address.\n")
}

func getExistingGuard(ctx *ExampleContext) {
	fmt.Println("=== Get Existing Guard ===")

	// Prompt for factory address
	factoryAddr := promptAddress("Guard Factory address", GuardFactoryAddress)
	fmt.Printf("\nUsing Guard Factory: %s\n", factoryAddr.Hex())

	// Prompt for required signer
	requiredSigner := promptAddress("Required signer address", "")
	if requiredSigner == (common.Address{}) {
		log.Printf("Error: Required signer address is required")
		return
	}

	// Create factory contract instance
	guardFactory, err := utils.NewSimpleUniversalGuardFactoryContract(factoryAddr, ctx.Client)
	if err != nil {
		log.Printf("Failed to create Guard Factory instance: %v", err)
		return
	}

	// Get guard address
	fmt.Printf("\nQuerying guard for signer: %s...", requiredSigner.Hex())
	guardAddr, err := guardFactory.GetGuard(&bind.CallOpts{}, requiredSigner)
	if err != nil {
		log.Printf("Failed to get guard: %v", err)
		return
	}

	fmt.Printf(" ✅\n")
	if guardAddr == (common.Address{}) {
		fmt.Printf("\n⚠️  No guard exists for this signer\n")
		fmt.Printf("You can create one using option 31\n")
	} else {
		fmt.Printf("\nGuard Address: %s\n", guardAddr.Hex())
		fmt.Printf("Required Signer: %s\n", requiredSigner.Hex())
	}
}

func createNewGuard(ctx *ExampleContext) {
	fmt.Println("=== Create New Guard ===")

	// Prompt for factory address
	factoryAddr := promptAddress("Guard Factory address", GuardFactoryAddress)
	fmt.Printf("\nUsing Guard Factory: %s\n", factoryAddr.Hex())

	// Prompt for required signer
	requiredSigner := promptAddress("Required signer address", "")
	if requiredSigner == (common.Address{}) {
		log.Printf("Error: Required signer address is required")
		return
	}

	// Create factory contract instance
	guardFactory, err := utils.NewSimpleUniversalGuardFactoryContract(factoryAddr, ctx.Client)
	if err != nil {
		log.Printf("Failed to create Guard Factory instance: %v", err)
		return
	}

	// Check if guard already exists
	fmt.Printf("\nChecking if guard already exists...")
	existingGuard, err := guardFactory.GetGuard(&bind.CallOpts{}, requiredSigner)
	if err != nil {
		log.Printf("Failed to check existing guard: %v", err)
		return
	}

	if existingGuard != (common.Address{}) {
		fmt.Printf("\n\n⚠️  Guard already exists!\n")
		fmt.Printf("Guard Address: %s\n", existingGuard.Hex())
		fmt.Printf("Required Signer: %s\n", requiredSigner.Hex())
		fmt.Printf("\nThe factory will return the existing guard address if you proceed.\n")
	} else {
		fmt.Printf(" ✅ No existing guard found\n")
	}

	// Predict the guard address
	predictedAddr, err := guardFactory.PredictGuardAddress(&bind.CallOpts{}, requiredSigner)
	if err != nil {
		log.Printf("Failed to predict guard address: %v", err)
		return
	}
	fmt.Printf("Guard will be deployed at: %s\n", predictedAddr.Hex())

	// Confirm transaction
	if !confirmSend() {
		fmt.Println("Transaction cancelled")
		return
	}

	// Create transaction auth
	auth := getAuth(ctx, &factoryAddr, nil)

	// Create guard
	fmt.Printf("\nCreating guard...\n")
	tx, err := guardFactory.CreateGuard(auth, requiredSigner)
	if err != nil {
		log.Printf("Failed to create guard: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	waitForTransaction(ctx, tx)

	// Verify guard was created
	fmt.Printf("\nVerifying guard creation...")
	createdGuard, err := guardFactory.GetGuard(&bind.CallOpts{}, requiredSigner)
	if err != nil {
		log.Printf("Failed to verify guard: %v", err)
		return
	}

	if createdGuard != (common.Address{}) {
		fmt.Printf(" ✅\n")
		fmt.Printf("\nGuard successfully created!\n")
		fmt.Printf("Guard Address: %s\n", createdGuard.Hex())
		fmt.Printf("Required Signer: %s\n", requiredSigner.Hex())
		fmt.Printf("\nThis guard can now be set on any Safe wallet.\n")
		fmt.Printf("All transactions from Safes using this guard will require a signature from: %s\n", requiredSigner.Hex())
	} else {
		fmt.Printf("\n⚠️  Guard creation verification failed\n")
	}
}

// predictSafeAddress 使用CREATE2预测Safe地址
func predictSafeAddress(
	factoryAddress common.Address,
	singletonAddress common.Address,
	params utils.SafeSetupParams,
	client *ethclient.Client,
) (common.Address, error) {
	// 创建 Safe 初始化数据
	setupConfig := utils.SafeSetupConfig{
		Owners:          params.Owners,
		Threshold:       params.Threshold,
		To:              params.To,
		Data:            params.Data,
		FallbackHandler: params.FallbackHandler,
		PaymentToken:    params.PaymentToken,
		Payment:         params.Payment,
		PaymentReceiver: params.PaymentReceiver,
	}

	initData, err := utils.CreateSafeInitData(setupConfig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create init data: %w", err)
	}

	// 优先从工厂合约读取 proxy creation code hash，避免本地硬编码与链上不一致
	// if client != nil {
	// 	factory, err := utils.NewSafeProxyFactoryContract(factoryAddress, client)
	// 	if err == nil {
	// 		if codeHash, err := factory.ProxyCreationCodehash(&bind.CallOpts{}, singletonAddress); err == nil {
	// 			fmt.Printf("codeHash: %#x\n", codeHash)
	// 			return utils.CalculateProxyAddressWithCodeHash(
	// 				factoryAddress,
	// 				initData,
	// 				params.SaltNonce,
	// 				codeHash[:],
	// 			)
	// 		}
	// 	}
	// }

	return utils.CalculateProxyAddress(
		factoryAddress,
		singletonAddress,
		initData,
		params.SaltNonce,
	)
}

func createSafeAndUpdatePaymentAccountController(ctx *ExampleContext) {
	fmt.Println("=== Create Safe and Update Payment Account Controller ===")
	fmt.Println("This operation creates a new Safe and sets it as the controller for an existing payment account")
	fmt.Println()

	walletAddr := promptAddress("Enterprise wallet address", "")
	paymentAccount := promptAddress("Payment account address", "")
	safeAddr := promptAddress("Safe multisig address (for proposing transaction)", "")
	safeAddress := safeAddr.Hex()
	safeProxyFactory := promptAddress("Safe proxy factory", SafeFactoryAddress)
	safeSingleton := promptAddress("Safe singleton", SafeSingletonAddress)

	// 从环境变量读取三个 owner 私钥并转换为地址
	var owners []common.Address
	var ownerPrivateKeys []*ecdsa.PrivateKey
	ownerKeyHexes := []string{
		os.Getenv("OWNER_PRIVATE_KEY"),
		os.Getenv("OWNER2_PRIVATE_KEY"),
		os.Getenv("OWNER3_PRIVATE_KEY"),
	}

	for i, keyHex := range ownerKeyHexes {
		if keyHex == "" {
			log.Printf("Warning: OWNER%d_PRIVATE_KEY not set in .env, skipping", i+1)
			continue
		}
		cleanKey := strings.TrimPrefix(keyHex, "0x")
		privateKey, err := crypto.HexToECDSA(cleanKey)
		if err != nil {
			log.Printf("Warning: Failed to parse OWNER%d_PRIVATE_KEY: %v", i+1, err)
			continue
		}
		ownerAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
		owners = append(owners, ownerAddr)
		ownerPrivateKeys = append(ownerPrivateKeys, privateKey)
	}

	if len(owners) == 0 {
		log.Printf("Error: No valid owner private keys found in .env")
		return
	}

	// 询问阈值
	thresholdStr := prompt(fmt.Sprintf("Threshold (%d owners configured) [2]", len(owners)))
	threshold := int64(2)
	if thresholdStr != "" {
		parsed, err := strconv.ParseInt(thresholdStr, 10, 64)
		if err == nil && parsed > 0 && parsed <= int64(len(owners)) {
			threshold = parsed
		}
	}

	fmt.Printf("\n=== New Safe Configuration ===\n")
	fmt.Printf("Owners (%d):\n", len(owners))
	for i, owner := range owners {
		fmt.Printf("  %d. %s\n", i+1, owner.Hex())
	}
	fmt.Printf("Threshold: %d/%d\n", threshold, len(owners))

	// 询问是否需要设置 guard
	guardInput := prompt("Guard address (leave empty to skip)")
	var guardAddress common.Address
	var setupTo common.Address
	var setupData []byte

	if guardInput != "" {
		if !common.IsHexAddress(guardInput) {
			log.Printf("Error: Invalid guard address")
			return
		}
		guardAddress = common.HexToAddress(guardInput)

		// 询问 SafeGuardSetter 合约地址
		guardSetterInput := prompt("SafeGuardSetter contract address")
		if !common.IsHexAddress(guardSetterInput) {
			log.Printf("Error: Invalid SafeGuardSetter address")
			return
		}
		guardSetterAddress := common.HexToAddress(guardSetterInput)

		// 生成 setGuard calldata
		var err error
		setupData, err = utils.SafeSetGuardData(guardAddress)
		if err != nil {
			log.Printf("Error generating setGuard calldata: %v", err)
			return
		}

		setupTo = guardSetterAddress
		fmt.Printf("✅ Will set Guard via SafeGuardSetter (%s): %s\n", guardSetterAddress.Hex(), guardAddress.Hex())
	} else {
		fmt.Println("ℹ️  No Guard will be set")
	}

	params := utils.SafeSetupParams{
		Owners:          owners,
		Threshold:       big.NewInt(threshold),
		To:              setupTo,
		Data:            setupData,
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: owners[0],
		SaltNonce:       big.NewInt(time.Now().Unix()),
	}

	// 预测新 Safe 地址
	fmt.Printf("\n📍 Predicting new Safe address...")
	predictedSafeAddr, err := predictSafeAddress(safeProxyFactory, safeSingleton, params, ctx.Client)
	if err != nil {
		log.Printf("Failed to predict Safe address: %v", err)
		return
	}
	fmt.Printf(" ✅\n")
	fmt.Printf("🎯 New Safe address will be: %s\n", predictedSafeAddr.Hex())
	fmt.Printf("📝 SaltNonce: %s\n\n", params.SaltNonce.String())

	// Prepare calldata
	data, err := utils.CreateAndUpdatePaymentAccountControllerData(
		paymentAccount,
		safeProxyFactory,
		safeSingleton,
		params,
	)
	if err != nil {
		log.Printf("Error generating calldata: %v", err)
		return
	}

	fmt.Printf("\nCalldata: 0x%x\n", data)
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Will be called via Safe multisig: %s\n", safeAddress)
	fmt.Printf("Target: EnterpriseWallet at %s\n", walletAddr.Hex())
	fmt.Printf("Payment Account: %s\n", paymentAccount.Hex())
	fmt.Printf("New Controller: %s\n\n", predictedSafeAddr.Hex())

	if !confirmSend() {
		fmt.Println("Cancelled.")
		return
	}

	// 创建Safe客户端和API客户端
	fmt.Printf("\n🔧 Creating Safe client...")
	safeClient, err := protocol.NewSafe(protocol.SafeConfig{
		SafeAddress: safeAddress,
		RpcURL:      ctx.RPCURL,
		ChainID:     ctx.ChainID.Int64(),
		PrivateKey:  ctx.PrivateKeyHex,
	})
	if err != nil {
		log.Printf("Failed to create Safe client: %v", err)
		return
	}

	if ctx.SafeAPIKey == "" {
		fmt.Printf(" ⚠️  SAFE_API_KEY not configured\n")
	}

	apiConfig := api.SafeApiKitConfig{
		ChainID:      ctx.ChainID.Int64(),
		TxServiceURL: ctx.SafeAPIURL,
	}
	apiClient, err := api.NewSafeApiKit(apiConfig)
	if err != nil {
		log.Printf("Failed to create API client: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// 选择 channel
	channelStr := prompt("Transaction channel [0]")
	channel := uint64(0)
	if channelStr != "" {
		if parsed, err := strconv.ParseUint(channelStr, 10, 64); err == nil {
			channel = parsed
		} else {
			log.Printf("Invalid channel, using 0")
		}
	}

	// 获取Safe信息
	fmt.Printf("📊 Getting Safe info...")
	safeInfo, err := apiClient.GetSafeInfo(context.Background(), safeAddress)
	if err != nil {
		log.Printf("Failed to get Safe info: %v", err)
		return
	}

	// 读取版本
	version, err := safeClient.GetVersion(context.Background())
	if err != nil {
		log.Printf("Failed to get Safe version: %v", err)
		return
	}
	isMulti := strings.Contains(version, "multichannel")

	currentNonce := uint64(0)
	if isMulti {
		onChainNonce, err := safeClient.GetChannelNonce(context.Background(), channel)
		if err != nil {
			log.Printf("Failed to get channel %d nonce: %v", channel, err)
			return
		}
		currentNonce = onChainNonce
	} else {
		onChainNonce, err := safeClient.GetNonce(context.Background())
		if err != nil {
			log.Printf("Failed to get nonce: %v", err)
			return
		}
		currentNonce = onChainNonce
	}

	fmt.Printf(" ✅ (版本: %s, 阈值: %d/%d, channel %d nonce: %d)\n", version, safeInfo.Threshold, len(safeInfo.Owners), channel, currentNonce)

	// 创建Safe交易
	fmt.Printf("📋 创建Safe交易...")
	txData := safetypes.SafeTransactionDataPartial{
		To:      walletAddr.Hex(),
		Value:   "0",
		Data:    "0x" + hex.EncodeToString(data),
		Nonce:   &currentNonce,
		Channel: &channel,
	}

	transaction, err := safeClient.CreateTransaction(context.Background(), txData)
	if err != nil {
		log.Printf("创建交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// 获取Safe交易哈希
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

	// 选择签名私钥
	fmt.Println("\n=== Select Signing Private Key ===")
	fmt.Println("1. DEPLOYER_PRIVATE_KEY")
	fmt.Println("2. OWNER_PRIVATE_KEY")
	fmt.Println("3. OWNER2_PRIVATE_KEY")
	fmt.Println("4. OWNER3_PRIVATE_KEY")
	keyChoice := prompt("Select private key [2]")
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

	signerPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(selectedPrivateKey, "0x"))
	if err != nil {
		log.Printf("Error parsing %s: %v", keyLabel, err)
		return
	}
	signerAddress := crypto.PubkeyToAddress(signerPrivateKey.PublicKey)
	fmt.Printf("使用私钥: %s (%s)\n", keyLabel, signerAddress.Hex())

	// 签名交易
	fmt.Printf("\n✍️  签名交易...")
	signature, err := utils.SignMessage(txHash, signerPrivateKey)
	if err != nil {
		log.Printf("签名交易失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n   签名者: %s\n", signerAddress.Hex())

	// 提交到Safe服务
	fmt.Printf("\n📤 提交交易到Safe服务...")
	proposal := api.ProposeTransactionProps{
		Channel:                 func() *int64 { v := int64(transaction.Data.Channel); return &v }(),
		SafeAddress:             safeAddress,
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
		Sender:                  signerAddress.Hex(),
		Signature:               "0x" + hex.EncodeToString(signature),
		ContractTransactionHash: "0x" + safeTxHash,
	}

	_, err = apiClient.ProposeTransaction(context.Background(), proposal)
	if err != nil {
		log.Printf("提交失败: %v", err)
		return
	}
	fmt.Printf(" ✅\n")

	// 检查签名数是否足够，如果足够则自动执行
	fmt.Printf("🔍 检查签名数...")
	txDetail, err := apiClient.GetMultisigTransaction(context.Background(), "0x"+safeTxHash)
	if err != nil {
		log.Printf("获取交易详情失败: %v", err)
		return
	}

	confirmationsCount := len(txDetail.Confirmations)
	fmt.Printf(" (%d/%d)\n", confirmationsCount, txDetail.ConfirmationsRequired)

	if confirmationsCount >= txDetail.ConfirmationsRequired {
		// 签名数足够，执行交易
		fmt.Printf("✅ 签名数已满足，执行交易...\n")

		// 处理 channel 字段
		txChannel := uint64(0)
		if txDetail.Channel != nil {
			txChannel = uint64(*txDetail.Channel)
		}

		// 构建SafeTransaction对象
		safeTransaction := &safetypes.SafeTransaction{
			Data: safetypes.SafeTransactionData{
				Channel:        txChannel,
				To:             txDetail.To,
				Value:          txDetail.Value,
				Data:           txDetail.Data,
				Operation:      safetypes.OperationType(txDetail.Operation),
				SafeTxGas:      fmt.Sprintf("%d", txDetail.SafeTxGas),
				BaseGas:        fmt.Sprintf("%d", txDetail.BaseGas),
				GasPrice:       txDetail.GasPrice,
				GasToken:       txDetail.GasToken,
				RefundReceiver: txDetail.RefundReceiver,
				Nonce:          uint64(txDetail.Nonce),
			},
			Signatures: make(map[string]safetypes.SafeSignature),
		}

		// 添加所有签名
		for _, conf := range txDetail.Confirmations {
			sigBytes := common.FromHex(conf.Signature)
			if len(sigBytes) >= 65 {
				safeTransaction.AddSignature(safetypes.SafeSignature{
					Signer: conf.Owner,
					Data:   conf.Signature,
				})
			}
		}

		// 执行交易
		result, err := safeClient.ExecuteTransaction(context.Background(), safeTransaction)
		if err != nil {
			log.Printf("执行交易失败: %v", err)
			return
		}

		fmt.Printf("✅ 交易已执行!\n")
		fmt.Printf("Transaction Hash: %s\n", result.Hash)
	} else {
		fmt.Printf("📋 交易已提交，等待更多签名 (需要 %d 个签名)\n", txDetail.ConfirmationsRequired)
		fmt.Printf("Safe Transaction Hash: 0x%s\n", safeTxHash)
	}
}
