package abi

import _ "embed"

// Enterprise wallet ABI definitions embedded at compile time.

//go:embed EnterpriseWallet.json
var EnterpriseWallet []byte

//go:embed EnterpriseWalletFactory.json
var EnterpriseWalletFactory []byte
