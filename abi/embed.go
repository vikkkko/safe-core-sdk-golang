package abi

import _ "embed"

// Enterprise wallet ABI definitions embedded at compile time.

//go:embed EnterpriseWallet.json
var EnterpriseWallet []byte

//go:embed EnterpriseWalletFactory.json
var EnterpriseWalletFactory []byte

//go:embed PaymentAccount.json
var PaymentAccount []byte

// Safe ABI definitions embedded at compile time.

//go:embed Safe.json
var Safe []byte

//go:embed SafeProxyFactory_full.json
var SafeProxyFactory []byte
