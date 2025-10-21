package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
)

const (
	DefaultSafeFactoryAddress   = "0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2" // Sepolia
	DefaultSafeSingletonAddress = "0x29fcB43b46531BcA003ddC8FCB67FFE91900C762" // Sepolia v1.4.1
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	fmt.Println("=== Safe Address Prediction ===\n")
	predictMultiSigSafe()
}

func prompt(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", message)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func predictMultiSigSafe() {
	// Prompt for owner addresses
	fmt.Println("Enter owner addresses (one per line, empty line to finish):")
	var owners []common.Address
	for i := 1; ; i++ {
		ownerAddr := prompt(fmt.Sprintf("Owner %d", i))
		if ownerAddr == "" {
			break
		}

		if !common.IsHexAddress(ownerAddr) {
			log.Printf("Invalid address: %s, skipping", ownerAddr)
			continue
		}

		owners = append(owners, common.HexToAddress(ownerAddr))
	}

	if len(owners) == 0 {
		log.Fatal("Error: At least one owner is required")
	}

	// Prompt for threshold
	thresholdStr := prompt(fmt.Sprintf("Threshold (1-%d)", len(owners)))
	threshold, err := strconv.ParseInt(thresholdStr, 10, 64)
	if err != nil || threshold < 1 || threshold > int64(len(owners)) {
		log.Fatalf("Error: Invalid threshold (must be between 1 and %d)", len(owners))
	}

	// Prompt for salt nonce
	saltNonceStr := prompt("Salt nonce [0]")
	var saltNonce *big.Int
	if saltNonceStr == "" {
		saltNonce = big.NewInt(0)
	} else {
		saltNonce = new(big.Int)
		_, ok := saltNonce.SetString(saltNonceStr, 10)
		if !ok {
			log.Fatal("Error: Invalid salt nonce")
		}
	}

	// Create Safe initialization data
	setupConfig := utils.SafeSetupConfig{
		Owners:          owners,
		Threshold:       big.NewInt(threshold),
		To:              common.Address{},
		Data:            []byte{},
		FallbackHandler: common.Address{},
		PaymentToken:    common.Address{},
		Payment:         big.NewInt(0),
		PaymentReceiver: common.Address{},
	}

	initData, err := utils.CreateSafeInitData(setupConfig)
	if err != nil {
		log.Fatalf("Failed to create init data: %v", err)
	}

	// Predict address
	factory := common.HexToAddress(DefaultSafeFactoryAddress)
	singleton := common.HexToAddress(DefaultSafeSingletonAddress)

	predictedAddress, err := utils.CalculateProxyAddress(factory, singleton, initData, saltNonce)
	if err != nil {
		log.Fatalf("Failed to predict address: %v", err)
	}

	fmt.Printf("\n=== Configuration ===\n")
	fmt.Printf("Factory:   %s\n", factory.Hex())
	fmt.Printf("Singleton: %s\n", singleton.Hex())
	fmt.Printf("Owners (%d):\n", len(owners))
	for i, owner := range owners {
		fmt.Printf("  %d. %s\n", i+1, owner.Hex())
	}
	fmt.Printf("Threshold: %d/%d\n", threshold, len(owners))
	fmt.Printf("Salt:      %s\n", saltNonce.String())

	fmt.Printf("\n=== Result ===\n")
	fmt.Printf("Predicted Safe Address: %s\n", predictedAddress.Hex())
}
