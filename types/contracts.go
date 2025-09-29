package types

// ContractInfo represents information about a deployed contract
type ContractInfo struct {
	Version         SafeVersion `json:"version"`         // Contract version
	Address         string      `json:"address"`         // Contract address
	Deployed        bool        `json:"deployed"`        // Whether the contract is deployed
	NetworkId       string      `json:"networkId"`       // Network ID where the contract is deployed
	TransactionHash string      `json:"transactionHash"` // Deployment transaction hash
}

// SafeDeploymentConfig represents configuration for deploying a new Safe
type SafeDeploymentConfig struct {
	SafeVersion     SafeVersion     `json:"safeVersion"`              // Safe contract version to deploy
	SafeSetupConfig SafeSetupConfig `json:"safeSetupConfig"`         // Safe setup configuration
	SaltNonce       *string         `json:"saltNonce,omitempty"`     // Optional salt nonce for deterministic deployment
}

// PredictedSafeProps represents properties of a predicted Safe address
type PredictedSafeProps struct {
	SafeAddress         string              `json:"safeAddress"`         // Predicted Safe address
	SafeDeploymentConfig SafeDeploymentConfig `json:"safeDeploymentConfig"` // Deployment configuration used for prediction
}

// SafeConfig represents configuration for connecting to an existing Safe
type SafeConfig struct {
	SafeAddress *string               `json:"safeAddress,omitempty"` // Existing Safe address
	Predicted   *PredictedSafeProps   `json:"predicted,omitempty"`   // Predicted Safe properties (for undeployed Safes)
}

// ConnectSafeConfig represents configuration for connecting to a Safe
type ConnectSafeConfig struct {
	SafeAddress string `json:"safeAddress"` // Safe address to connect to
	RpcURL      string `json:"rpcUrl"`      // RPC URL for blockchain connection
	ChainID     int64  `json:"chainId"`     // Chain ID
}

// SafeAccountConfig represents the account configuration for a Safe
type SafeAccountConfig struct {
	Owners           []string `json:"owners"`           // List of owner addresses
	Threshold        uint     `json:"threshold"`        // Required number of confirmations
	To               string   `json:"to,omitempty"`     // Optional setup call target
	Data             string   `json:"data,omitempty"`   // Optional setup call data
	FallbackHandler  string   `json:"fallbackHandler,omitempty"`  // Fallback handler address
	PaymentToken     string   `json:"paymentToken,omitempty"`     // Payment token address
	Payment          string   `json:"payment,omitempty"`          // Payment amount
	PaymentReceiver  string   `json:"paymentReceiver,omitempty"`  // Payment receiver address
}

// SafeInfo represents information about a Safe account
type SafeInfo struct {
	Address         string   `json:"address"`         // Safe address
	Nonce           uint64   `json:"nonce"`           // Current nonce
	Threshold       uint     `json:"threshold"`       // Required confirmations
	Owners          []string `json:"owners"`          // List of owners
	MasterCopy      string   `json:"masterCopy"`      // Master copy address
	Modules         []string `json:"modules"`         // Enabled modules
	FallbackHandler string   `json:"fallbackHandler"` // Fallback handler address
	Guard           string   `json:"guard"`           // Transaction guard address
	Version         string   `json:"version"`         // Safe version
}

// CompatibilityFallbackHandlerContractType represents the fallback handler contract type
type CompatibilityFallbackHandlerContractType struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// CreateCallBaseContract represents the CreateCall contract
type CreateCallBaseContract struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// MultiSendBaseContract represents the MultiSend contract
type MultiSendBaseContract struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// MultiSendCallOnlyBaseContract represents the MultiSendCallOnly contract
type MultiSendCallOnlyBaseContract struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// SafeBaseContract represents the Safe master copy contract
type SafeBaseContract struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// SafeProxyFactoryBaseContract represents the Safe proxy factory contract
type SafeProxyFactoryBaseContract struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// SignMessageLibBaseContract represents the SignMessageLib contract
type SignMessageLibBaseContract struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// SafeWebAuthnSignerFactoryContract represents the WebAuthn signer factory contract
type SafeWebAuthnSignerFactoryContract struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// SafeWebAuthnSharedSignerContract represents the WebAuthn shared signer contract
type SafeWebAuthnSharedSignerContract struct {
	Address string      `json:"address"` // Contract address
	Version SafeVersion `json:"version"` // Contract version
}

// PasskeyCredential represents a passkey credential
type PasskeyCredential struct {
	ID              string `json:"id"`              // Credential ID
	PublicKey       string `json:"publicKey"`       // Public key
	RawID           []byte `json:"rawId"`           // Raw credential ID
	Response        string `json:"response"`        // Credential response
	AuthenticatorData []byte `json:"authenticatorData"` // Authenticator data
	ClientDataJSON  []byte `json:"clientDataJSON"`  // Client data JSON
	Signature       []byte `json:"signature"`       // Signature
}

// PasskeyArgType represents arguments for passkey operations
type PasskeyArgType struct {
	Credential PasskeyCredential `json:"credential"` // Passkey credential
}

// DefaultSafeVersion represents the default Safe version to use
const DefaultSafeVersion = SafeVersion141

// PredeterminedSaltNonce represents the predetermined salt nonce for consistent deployment addresses
const PredeterminedSaltNonce = "0xcfe33a586323e7325be6aa6ecd8b4600d232a9037e83c8ece69413b777dabe65"