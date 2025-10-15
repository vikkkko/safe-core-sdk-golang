package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
)

const (
	// Enterprise Wallet deployed contracts on your network
	FactoryAddress     = "0xC5473e192d07420B09b684086d3631830b268bE7"
	ImplementationAddr = "0x5D92e1c1B4F8fB2a291B9A44451dBE4eAAe2b286"
)

// Context holds all necessary data for examples
type ExampleContext struct {
	Client          *ethclient.Client
	PrivateKey      *ecdsa.PrivateKey
	FromAddress     common.Address
	ChainID         *big.Int
	Auth            *bind.TransactOpts
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

func (f *EnterpriseWalletFactoryContract) CreateWallet(auth *bind.TransactOpts, implementation common.Address, salt [32]byte, params factoryInitParams) (*types.Transaction, error) {
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

func (w *EnterpriseWalletContract) CreatePaymentAccount(auth *bind.TransactOpts, name string, controller common.Address) (*types.Transaction, error) {
	return w.contract.Transact(auth, "createPaymentAccount", name, controller)
}

func (w *EnterpriseWalletContract) CreateCollectionAccount(auth *bind.TransactOpts, name string, target common.Address) (*types.Transaction, error) {
	return w.contract.Transact(auth, "createCollectionAccount", name, target)
}

func (w *EnterpriseWalletContract) ApproveTokenForPayment(auth *bind.TransactOpts, token, paymentAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.contract.Transact(auth, "approveTokenForPayment", token, paymentAccount, amount)
}

func (w *EnterpriseWalletContract) TransferETHToPayment(auth *bind.TransactOpts, paymentAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.contract.Transact(auth, "transferETHToPayment", paymentAccount, amount)
}

func (w *EnterpriseWalletContract) SetCollectionTarget(auth *bind.TransactOpts, collectionAccount, target common.Address) (*types.Transaction, error) {
	return w.contract.Transact(auth, "setCollectionTarget", collectionAccount, target)
}

func (w *EnterpriseWalletContract) CollectFunds(auth *bind.TransactOpts, token, collectionAccount common.Address) (*types.Transaction, error) {
	return w.contract.Transact(auth, "collectFunds", token, collectionAccount)
}

func (w *EnterpriseWalletContract) CreateSafeAndPaymentAccount(
	auth *bind.TransactOpts,
	proxyFactory common.Address,
	safeSingleton common.Address,
	params utils.SafeSetupParams,
	name string,
) (*types.Transaction, error) {
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
) (*types.Transaction, error) {
	setup := convertSafeSetupParams(params)
	return w.contract.Transact(auth, "createSafeAndCollectionAccount", proxyFactory, safeSingleton, setup, name, collectionTarget)
}

func (w *EnterpriseWalletContract) UpdateMethodController(auth *bind.TransactOpts, method [4]byte, controller common.Address) (*types.Transaction, error) {
	return w.contract.Transact(auth, "updateMethodController", method, controller)
}

func (w *EnterpriseWalletContract) UpdateMethodControllers(auth *bind.TransactOpts, methods [][4]byte, controllers []common.Address) (*types.Transaction, error) {
	return w.contract.Transact(auth, "updateMethodControllers", methods, controllers)
}

func (w *EnterpriseWalletContract) SetMethodController(auth *bind.TransactOpts, methods [][4]byte, controller common.Address) (*types.Transaction, error) {
	return w.contract.Transact(auth, "setMethodController", methods, controller)
}

func (w *EnterpriseWalletContract) UpdatePaymentAccountController(auth *bind.TransactOpts, paymentAccount, controller common.Address) (*types.Transaction, error) {
	return w.contract.Transact(auth, "updatePaymentAccountController", paymentAccount, controller)
}

func (w *EnterpriseWalletContract) EmergencyFreeze(auth *bind.TransactOpts, target common.Address, freeze bool) (*types.Transaction, error) {
	return w.contract.Transact(auth, "emergencyFreeze", target, freeze)
}

func (w *EnterpriseWalletContract) EmergencyPause(auth *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return w.contract.Transact(auth, "emergencyPause", pause)
}

func (w *EnterpriseWalletContract) ProposeSuperAdminTransfer(auth *bind.TransactOpts, newSuperAdmin common.Address, timeout *big.Int) (*types.Transaction, error) {
	return w.contract.Transact(auth, "proposeSuperAdminTransfer", newSuperAdmin, timeout)
}

func (w *EnterpriseWalletContract) ConfirmSuperAdminTransfer(auth *bind.TransactOpts) (*types.Transaction, error) {
	return w.contract.Transact(auth, "confirmSuperAdminTransfer")
}

func (w *EnterpriseWalletContract) CancelSuperAdminTransfer(auth *bind.TransactOpts) (*types.Transaction, error) {
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
	factoryContract, err := NewEnterpriseWalletFactoryContract(common.HexToAddress(FactoryAddress), client)
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
	fmt.Println("  15. Prepare createSafeAndPaymentAccount data")
	fmt.Println("  16. Prepare createSafeAndCollectionAccount data")
	fmt.Println("\nSuperAdmin Transfer - Prepare Data:")
	fmt.Println("  31. Prepare proposeSuperAdminTransfer data")
	fmt.Println("  32. Prepare confirmSuperAdminTransfer data")
	fmt.Println("  33. Prepare cancelSuperAdminTransfer data")
	fmt.Println("\nEnterprise Wallet - Execute on Chain (TX):")
	fmt.Println("  23. Execute createPaymentAccount")
	fmt.Println("  24. Execute createCollectionAccount")
	fmt.Println("  25. Execute createSafeAndPaymentAccount")
	fmt.Println("  26. Execute createSafeAndCollectionAccount")
	fmt.Println("  27. Execute approveTokenForPayment")
	fmt.Println("  28. Execute transferETHToPayment")
	fmt.Println("  29. Execute setCollectionTarget")
	fmt.Println("  30. Execute collectFunds")
	fmt.Println("  31. Execute updateMethodController")
	fmt.Println("  32. Execute updatePaymentAccountController")
	fmt.Println("\nSuperAdmin Transfer - Execute on Chain (TX):")
	fmt.Println("  34. Execute proposeSuperAdminTransfer")
	fmt.Println("  35. Execute confirmSuperAdminTransfer")
	fmt.Println("  36. Execute cancelSuperAdminTransfer")
	fmt.Println("\nBatch Operations - Prepare Data:")
	fmt.Println("  39. Prepare updateMethodControllers (batch with different controllers)")
	fmt.Println("  40. Prepare setMethodController (batch with same controller)")
	fmt.Println("  41. Prepare emergencyFreeze")
	fmt.Println("\nBatch Operations - Execute on Chain (TX):")
	fmt.Println("  42. Execute updateMethodControllers")
	fmt.Println("  43. Execute setMethodController")
	fmt.Println("  44. Execute emergencyFreeze")
	fmt.Println("\nQuery Functions (View):")
	fmt.Println("  17. Query payment accounts")
	fmt.Println("  18. Query collection accounts")
	fmt.Println("  19. Query method config")
	fmt.Println("  20. Check if address is payment account")
	fmt.Println("  21. Check if address is collection account")
	fmt.Println("\nSuperAdmin Transfer - Query:")
	fmt.Println("  37. Query SuperAdmin transfer proposal")
	fmt.Println("  38. Check if transfer is valid")
	fmt.Println("\nUtility:")
	fmt.Println("  22. Run all read-only examples")
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
		examplePrepareCreateSafeAndPaymentAccount(ctx)
	case "16":
		examplePrepareCreateSafeAndCollectionAccount(ctx)
	case "17":
		exampleQueryPaymentAccounts(ctx)
	case "18":
		exampleQueryCollectionAccounts(ctx)
	case "19":
		exampleQueryMethodConfig(ctx)
	case "20":
		exampleCheckIsPaymentAccount(ctx)
	case "21":
		exampleCheckIsCollectionAccount(ctx)
	case "22":
		runAllReadOnlyExamples(ctx)
	case "23":
		exampleExecuteCreatePaymentAccount(ctx)
	case "24":
		exampleExecuteCreateCollectionAccount(ctx)
	case "25":
		exampleExecuteCreateSafeAndPaymentAccount(ctx)
	case "26":
		exampleExecuteCreateSafeAndCollectionAccount(ctx)
	case "27":
		exampleExecuteApproveToken(ctx)
	case "28":
		exampleExecuteTransferETH(ctx)
	case "29":
		exampleExecuteSetCollectionTarget(ctx)
	case "30":
		exampleExecuteCollectFunds(ctx)
	case "31":
		exampleExecuteUpdateMethodController(ctx)
	case "32":
		exampleExecuteUpdatePaymentAccountController(ctx)
	case "33":
		examplePrepareProposeSuperAdminTransfer(ctx)
	case "34":
		examplePrepareConfirmSuperAdminTransfer(ctx)
	case "35":
		examplePrepareCancelSuperAdminTransfer(ctx)
	case "36":
		exampleExecuteProposeSuperAdminTransfer(ctx)
	case "37":
		exampleExecuteConfirmSuperAdminTransfer(ctx)
	case "38":
		exampleExecuteCancelSuperAdminTransfer(ctx)
	case "39":
		examplePrepareUpdateMethodControllers(ctx)
	case "40":
		examplePrepareSetMethodController(ctx)
	case "41":
		examplePrepareEmergencyFreeze(ctx)
	case "42":
		exampleExecuteUpdateMethodControllers(ctx)
	case "43":
		exampleExecuteSetMethodController(ctx)
	case "44":
		exampleExecuteEmergencyFreeze(ctx)
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
	configs := make([]utils.MethodConfig, len(methodSelectors))
	for i := range methodSelectors {
		configs[i] = utils.MethodConfig{Controller: ctx.FromAddress}
	}

	contractInitParams := factoryInitParams{
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

func examplePrepareCreateSafeAndPaymentAccount(ctx *ExampleContext) {
	fmt.Println("=== Prepare createSafeAndPaymentAccount Data ===")

	proxyFactory := common.HexToAddress(FactoryAddress)
	safeSingleton := common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552") // Safe v1.4.1 mainnet copy (replace per network)

	params := utils.SafeSetupParams{
		Owners:          []common.Address{ctx.FromAddress},
		Threshold:       big.NewInt(1),
		To:              common.Address{},
		Data:            []byte{},
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: ctx.FromAddress,
		SaltNonce:       big.NewInt(0),
	}

	data, err := utils.CreateSafeAndPaymentAccountData(proxyFactory, safeSingleton, params, "Treasury Payment Account")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Proxy factory: %s\n", proxyFactory.Hex())
	fmt.Printf("Safe singleton: %s\n", safeSingleton.Hex())
	fmt.Printf("Owners: %s\n", ctx.FromAddress.Hex())
	fmt.Printf("Threshold: %s\n", params.Threshold.String())
	fmt.Printf("Payment receiver: %s\n", params.PaymentReceiver.Hex())
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareCreateSafeAndCollectionAccount(ctx *ExampleContext) {
	fmt.Println("=== Prepare createSafeAndCollectionAccount Data ===")

	proxyFactory := common.HexToAddress(FactoryAddress)
	safeSingleton := common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552")
	collectionTarget := ctx.FromAddress

	params := utils.SafeSetupParams{
		Owners:          []common.Address{ctx.FromAddress},
		Threshold:       big.NewInt(1),
		To:              common.Address{},
		Data:            []byte{},
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: ctx.FromAddress,
		SaltNonce:       big.NewInt(0),
	}

	data, err := utils.CreateSafeAndCollectionAccountData(proxyFactory, safeSingleton, params, "Revenue Collection Account", collectionTarget)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Proxy factory: %s\n", proxyFactory.Hex())
	fmt.Printf("Safe singleton: %s\n", safeSingleton.Hex())
	fmt.Printf("Owners: %s\n", ctx.FromAddress.Hex())
	fmt.Printf("Threshold: %s\n", params.Threshold.String())
	fmt.Printf("Collection target: %s\n", collectionTarget.Hex())
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
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
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
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
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
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
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

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
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

	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
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
		{"Prepare createSafeAndPaymentAccount", examplePrepareCreateSafeAndPaymentAccount},
		{"Prepare createSafeAndCollectionAccount", examplePrepareCreateSafeAndCollectionAccount},
	}

	for i, ex := range examples {
		fmt.Printf("\n[%d/%d] %s\n", i+1, len(examples), ex.name)
		fmt.Println(strings.Repeat("-", 50))
		ex.fn(ctx)
	}

	fmt.Println("\n=== All Read-Only Examples Completed ===")
}

func promptWithDefault(label, def string) string {
	if def == "" {
		fmt.Printf("%s: ", label)
	} else {
		fmt.Printf("%s [%s]: ", label, def)
	}
	input := getUserInput()
	if input == "" {
		return def
	}
	return input
}

func parseAddressList(value string) ([]common.Address, error) {
	parts := strings.Split(value, ",")
	addresses := make([]common.Address, 0, len(parts))
	for _, raw := range parts {
		trimmed := strings.TrimSpace(raw)
		if trimmed == "" {
			continue
		}
		if !common.IsHexAddress(trimmed) {
			return nil, fmt.Errorf("invalid address: %s", trimmed)
		}
		addresses = append(addresses, common.HexToAddress(trimmed))
	}
	return addresses, nil
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
	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
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

func exampleExecuteCreateSafeAndPaymentAccount(ctx *ExampleContext) {
	fmt.Println("=== Execute createSafeAndPaymentAccount ===")
	fmt.Println("WARNING: This will send an actual transaction!\n")

	walletAddrStr := promptWithDefault("Enterprise wallet address", "")
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	proxyFactoryStr := promptWithDefault("Proxy factory address", FactoryAddress)
	safeSingletonStr := promptWithDefault("Safe singleton address", "0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552")
	ownersInput := promptWithDefault("Safe owners (comma separated)", ctx.FromAddress.Hex())
	owners, err := parseAddressList(ownersInput)
	if err != nil || len(owners) == 0 {
		fmt.Println("Invalid owners list. Cancelled.")
		return
	}

	thresholdInput := promptWithDefault("Safe threshold", "1")
	threshold, err := strconv.Atoi(thresholdInput)
	if err != nil || threshold <= 0 || threshold > len(owners) {
		fmt.Println("Invalid threshold. Cancelled.")
		return
	}

	name := promptWithDefault("Payment account name", "Treasury Payment Account")
	saltInput := promptWithDefault("Salt nonce (uint, default 0)", "0")
	saltNonce := new(big.Int)
	saltNonce, ok := saltNonce.SetString(saltInput, 10)
	if !ok {
		fmt.Println("Invalid salt nonce. Cancelled.")
		return
	}

	params := utils.SafeSetupParams{
		Owners:          owners,
		Threshold:       big.NewInt(int64(threshold)),
		To:              common.Address{},
		Data:            []byte{},
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: ctx.FromAddress,
		SaltNonce:       saltNonce,
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.CreateSafeAndPaymentAccount(
		auth,
		common.HexToAddress(proxyFactoryStr),
		common.HexToAddress(safeSingletonStr),
		params,
		name,
	)
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
		fmt.Printf("✓ createSafeAndPaymentAccount executed. Block %d Gas %d\n", receipt.BlockNumber.Uint64(), receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteCreateSafeAndCollectionAccount(ctx *ExampleContext) {
	fmt.Println("=== Execute createSafeAndCollectionAccount ===")
	fmt.Println("WARNING: This will send an actual transaction!\n")

	walletAddrStr := promptWithDefault("Enterprise wallet address", "")
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	proxyFactoryStr := promptWithDefault("Proxy factory address", FactoryAddress)
	safeSingletonStr := promptWithDefault("Safe singleton address", "0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552")
	collectionTargetStr := promptWithDefault("Collection target (address)", ctx.FromAddress.Hex())
	ownersInput := promptWithDefault("Safe owners (comma separated)", ctx.FromAddress.Hex())
	owners, err := parseAddressList(ownersInput)
	if err != nil || len(owners) == 0 {
		fmt.Println("Invalid owners list. Cancelled.")
		return
	}

	thresholdInput := promptWithDefault("Safe threshold", "1")
	threshold, err := strconv.Atoi(thresholdInput)
	if err != nil || threshold <= 0 || threshold > len(owners) {
		fmt.Println("Invalid threshold. Cancelled.")
		return
	}

	name := promptWithDefault("Collection account name", "Revenue Collection Account")
	saltInput := promptWithDefault("Salt nonce (uint, default 0)", "0")
	saltNonce := new(big.Int)
	saltNonce, ok := saltNonce.SetString(saltInput, 10)
	if !ok {
		fmt.Println("Invalid salt nonce. Cancelled.")
		return
	}

	params := utils.SafeSetupParams{
		Owners:          owners,
		Threshold:       big.NewInt(int64(threshold)),
		To:              common.Address{},
		Data:            []byte{},
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: ctx.FromAddress,
		SaltNonce:       saltNonce,
	}

	auth, err := getFreshAuth(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.CreateSafeAndCollectionAccount(
		auth,
		common.HexToAddress(proxyFactoryStr),
		common.HexToAddress(safeSingletonStr),
		params,
		name,
		common.HexToAddress(collectionTargetStr),
	)
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
		fmt.Printf("✓ createSafeAndCollectionAccount executed. Block %d Gas %d\n", receipt.BlockNumber.Uint64(), receipt.GasUsed)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
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

// ============= SuperAdmin Transfer Examples - Prepare Data =============

func examplePrepareProposeSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Prepare Propose SuperAdmin Transfer ===")
	newSuperAdmin := common.HexToAddress("0x4444444444444444444444444444444444444444")
	timeout := big.NewInt(86400) // 1 day

	data, err := utils.ProposeSuperAdminTransferData(newSuperAdmin, timeout)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("New SuperAdmin: %s\n", newSuperAdmin.Hex())
	fmt.Printf("Timeout: %s seconds (1 day)\n", timeout.String())
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareConfirmSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Prepare Confirm SuperAdmin Transfer ===")
	data, err := utils.ConfirmSuperAdminTransferData()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareCancelSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Prepare Cancel SuperAdmin Transfer ===")
	data, err := utils.CancelSuperAdminTransferData()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

// ============= SuperAdmin Transfer Examples - Execute on Chain =============

func exampleExecuteProposeSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Execute Propose SuperAdmin Transfer ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter new SuperAdmin address: ")
	newSuperAdminStr := getUserInput()
	if newSuperAdminStr == "" {
		fmt.Println("No new SuperAdmin address provided. Cancelled.")
		return
	}

	fmt.Print("Enter timeout in seconds (press Enter for default 1 day): ")
	timeoutStr := getUserInput()
	timeout := big.NewInt(0) // 0 means use default
	if timeoutStr != "" {
		timeout.SetString(timeoutStr, 10)
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  New SuperAdmin: %s\n", newSuperAdminStr)
	if timeout.Cmp(big.NewInt(0)) == 0 {
		fmt.Printf("  Timeout: default (1 day)\n")
	} else {
		fmt.Printf("  Timeout: %s seconds\n", timeout.String())
	}
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.ProposeSuperAdminTransfer(auth, common.HexToAddress(newSuperAdminStr), timeout)
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
		fmt.Printf("✓ SuperAdmin transfer proposed successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
		fmt.Println("\nNote: Check transaction logs for the proposal ID")
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteConfirmSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Execute Confirm SuperAdmin Transfer ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.ConfirmSuperAdminTransfer(auth)
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
		fmt.Printf("✓ SuperAdmin transfer confirmed successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteCancelSuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Execute Cancel SuperAdmin Transfer ===")
	fmt.Println("WARNING: This will send an actual transaction!")

	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.CancelSuperAdminTransfer(auth)
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
		fmt.Printf("✓ SuperAdmin transfer cancelled successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

// ============= SuperAdmin Transfer Examples - Query =============

func exampleQuerySuperAdminTransfer(ctx *ExampleContext) {
	fmt.Println("=== Query SuperAdmin Transfer Proposal ===")
	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()

	if walletAddrStr == "" {
		fmt.Println("Using example address...")
		walletAddrStr = "0xCD6c4962346F5680C765127ED29A8F5cc53a6B66"
	}

	walletAddr := common.HexToAddress(walletAddrStr)
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	transfer, err := wallet.GetSuperAdminTransfer(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error querying transfer: %v", err)
		return
	}

	fmt.Printf("Current SuperAdmin: %s\n", transfer.CurrentSuperAdmin.Hex())
	fmt.Printf("Proposed SuperAdmin: %s\n", transfer.ProposedSuperAdmin.Hex())
	fmt.Printf("Proposed At: %s\n", transfer.ProposedAt.String())
	fmt.Printf("Timeout: %s seconds\n", transfer.Timeout.String())
	fmt.Printf("Is Active: %v\n", transfer.IsActive)
}

func exampleCheckSuperAdminTransferValid(ctx *ExampleContext) {
	fmt.Println("=== Check if SuperAdmin Transfer is Valid ===")
	fmt.Print("Enter enterprise wallet address: ")
	walletAddrStr := getUserInput()

	if walletAddrStr == "" {
		fmt.Println("Using example address...")
		walletAddrStr = "0xCD6c4962346F5680C765127ED29A8F5cc53a6B66"
	}

	walletAddr := common.HexToAddress(walletAddrStr)
	wallet, err := NewEnterpriseWalletContract(walletAddr, ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	isValid, err := wallet.IsValidSuperAdminTransfer(&bind.CallOpts{})
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Current proposal valid: %v\n", isValid)
}

// ============= Batch Operations Examples - Prepare Data =============

func examplePrepareUpdateMethodControllers(ctx *ExampleContext) {
	fmt.Println("=== Prepare Update Method Controllers (Batch with Different Controllers) ===")

	// Example: Update 3 methods with different controllers
	methodSigs := [][4]byte{
		utils.CreatePaymentAccountSelector,
		utils.CreateCollectionAccountSelector,
		utils.ApproveTokenForPaymentSelector,
	}

	controllers := []common.Address{
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
	}

	data, err := utils.UpdateMethodControllersData(methodSigs, controllers)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("Updating 3 methods with different controllers:")
	for i, sig := range methodSigs {
		fmt.Printf("  Method %d: 0x%x -> Controller: %s\n", i+1, sig, controllers[i].Hex())
	}
	fmt.Printf("\nCalldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareSetMethodController(ctx *ExampleContext) {
	fmt.Println("=== Prepare Set Method Controller (Batch with Same Controller) ===")

	// Example: Set the same controller for multiple methods
	methodSigs := [][4]byte{
		utils.TransferETHToPaymentSelector,
		utils.CollectFundsSelector,
	}

	controller := common.HexToAddress("0x4444444444444444444444444444444444444444")

	data, err := utils.SetMethodControllerData(methodSigs, controller)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("Setting same controller for multiple methods:")
	fmt.Printf("  Controller: %s\n", controller.Hex())
	fmt.Println("  Methods:")
	for i, sig := range methodSigs {
		fmt.Printf("    %d. 0x%x\n", i+1, sig)
	}
	fmt.Printf("\nCalldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

func examplePrepareEmergencyFreeze(ctx *ExampleContext) {
	fmt.Println("=== Prepare Emergency Freeze ===")
	fmt.Println("Note: This now requires METHOD CONTROLLER permission (changed from superAdmin)")

	target := common.HexToAddress("0x1111111111111111111111111111111111111111")
	freeze := true

	data, err := utils.EmergencyFreezeData(target, freeze)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nTarget account: %s\n", target.Hex())
	fmt.Printf("Freeze: %v\n", freeze)
	fmt.Printf("Calldata length: %d bytes\n", len(data))
	fmt.Printf("Calldata: 0x%x\n", data)
}

// ============= Batch Operations Examples - Execute on Chain =============

func exampleExecuteUpdateMethodControllers(ctx *ExampleContext) {
	fmt.Println("=== Execute Update Method Controllers (Batch) ===")
	fmt.Println("WARNING: This will send an actual transaction!")
	fmt.Println("This allows updating multiple methods with different controllers in one transaction.")

	fmt.Print("\nEnter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("How many methods do you want to update? ")
	countStr := getUserInput()
	if countStr == "" {
		fmt.Println("No count provided. Cancelled.")
		return
	}

	var count int
	fmt.Sscanf(countStr, "%d", &count)
	if count <= 0 {
		fmt.Println("Invalid count. Cancelled.")
		return
	}

	methodSigs := make([][4]byte, count)
	controllers := make([]common.Address, count)

	for i := 0; i < count; i++ {
		fmt.Printf("\nMethod %d:\n", i+1)
		fmt.Print("  Enter method selector (e.g., 0x08f25c4a): ")
		methodSigStr := getUserInput()
		if methodSigStr == "" {
			fmt.Println("  No method selector provided. Cancelled.")
			return
		}

		methodSigStr = strings.TrimPrefix(methodSigStr, "0x")
		methodBytes := common.FromHex(methodSigStr)
		if len(methodBytes) != 4 {
			fmt.Println("  Invalid method selector (must be 4 bytes). Cancelled.")
			return
		}
		copy(methodSigs[i][:], methodBytes)

		fmt.Print("  Enter controller address: ")
		controllerStr := getUserInput()
		if controllerStr == "" {
			fmt.Println("  No controller address provided. Cancelled.")
			return
		}
		controllers[i] = common.HexToAddress(controllerStr)
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Updating %d methods:\n", count)
	for i := 0; i < count; i++ {
		fmt.Printf("    %d. Method 0x%x -> Controller %s\n", i+1, methodSigs[i], controllers[i].Hex())
	}
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.UpdateMethodControllers(auth, methodSigs, controllers)
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
		fmt.Printf("✓ Method controllers updated successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteSetMethodController(ctx *ExampleContext) {
	fmt.Println("=== Execute Set Method Controller (Batch with Same Controller) ===")
	fmt.Println("WARNING: This will send an actual transaction!")
	fmt.Println("This allows setting the same controller for multiple methods in one transaction.")

	fmt.Print("\nEnter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter controller address (will be set for all methods): ")
	controllerStr := getUserInput()
	if controllerStr == "" {
		fmt.Println("No controller address provided. Cancelled.")
		return
	}
	controller := common.HexToAddress(controllerStr)

	fmt.Print("How many methods do you want to update? ")
	countStr := getUserInput()
	if countStr == "" {
		fmt.Println("No count provided. Cancelled.")
		return
	}

	var count int
	fmt.Sscanf(countStr, "%d", &count)
	if count <= 0 {
		fmt.Println("Invalid count. Cancelled.")
		return
	}

	methodSigs := make([][4]byte, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter method selector %d (e.g., 0x08f25c4a): ", i+1)
		methodSigStr := getUserInput()
		if methodSigStr == "" {
			fmt.Println("No method selector provided. Cancelled.")
			return
		}

		methodSigStr = strings.TrimPrefix(methodSigStr, "0x")
		methodBytes := common.FromHex(methodSigStr)
		if len(methodBytes) != 4 {
			fmt.Println("Invalid method selector (must be 4 bytes). Cancelled.")
			return
		}
		copy(methodSigs[i][:], methodBytes)
	}

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Controller: %s\n", controller.Hex())
	fmt.Printf("  Updating %d methods:\n", count)
	for i, sig := range methodSigs {
		fmt.Printf("    %d. 0x%x\n", i+1, sig)
	}
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.SetMethodController(auth, methodSigs, controller)
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
		fmt.Printf("✓ Method controller set successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}

func exampleExecuteEmergencyFreeze(ctx *ExampleContext) {
	fmt.Println("=== Execute Emergency Freeze ===")
	fmt.Println("WARNING: This will send an actual transaction!")
	fmt.Println("Note: This now requires METHOD CONTROLLER permission (changed from superAdmin)")

	fmt.Print("\nEnter enterprise wallet address: ")
	walletAddrStr := getUserInput()
	if walletAddrStr == "" {
		fmt.Println("No wallet address provided. Cancelled.")
		return
	}

	fmt.Print("Enter target account address to freeze/unfreeze: ")
	targetStr := getUserInput()
	if targetStr == "" {
		fmt.Println("No target address provided. Cancelled.")
		return
	}
	target := common.HexToAddress(targetStr)

	fmt.Print("Freeze (true) or Unfreeze (false)? Enter 'true' or 'false': ")
	freezeStr := getUserInput()
	freeze := freezeStr == "true"

	fmt.Printf("\nReview transaction:\n")
	fmt.Printf("  Wallet: %s\n", walletAddrStr)
	fmt.Printf("  Target: %s\n", target.Hex())
	fmt.Printf("  Action: %v\n", map[bool]string{true: "FREEZE", false: "UNFREEZE"}[freeze])
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

	wallet, err := NewEnterpriseWalletContract(common.HexToAddress(walletAddrStr), ctx.Client)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	tx, err := wallet.EmergencyFreeze(auth, target, freeze)
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
		fmt.Printf("✓ Emergency freeze executed successfully!\n")
		fmt.Printf("  Block: %d\n", receipt.BlockNumber.Uint64())
		fmt.Printf("  Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("✗ Transaction failed")
	}
}
