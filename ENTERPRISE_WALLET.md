# Enterprise Wallet Integration

本文档说明如何使用Safe Core SDK与企业钱包合约集成。

## 概述

企业钱包是一个基于代理模式的钱包系统，支持：
- 📦 **Payment Accounts** - 用于支出的子账户
- 💰 **Collection Accounts** - 用于收款的子账户
- 🔐 **Method-level Access Control** - 每个方法可由不同的Safe多签钱包控制
- 👑 **Super Admin** - 超级管理员可以更新权限和应急控制

## 已部署合约

- **Factory合约**: `0x19cd09AA77a74f92fC12D4D2f5D63ea61193E157`
- **Implementation合约**: `0x3d6850a4A9790c3aD3924A5d66b4fEEC8cd25bE2`

## 快速开始

### 1. 运行示例

```bash
go run examples/enterprise_wallet_example.go
```

### 2. 主要功能

#### 预测钱包地址

```go
import (
    "github.com/vikkkko/safe-core-sdk-golang/protocol/contracts"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// 连接到工厂合约
factoryContract, err := contracts.NewEnterpriseWalletFactory(
    common.HexToAddress("0x19cd09AA77a74f92fC12D4D2f5D63ea61193E157"),
    client,
)

// 预测地址
var salt [32]byte
copy(salt[:], []byte("my-wallet-v1"))

predictedAddr, err := factoryContract.PredictWalletAddress(
    &bind.CallOpts{},
    common.HexToAddress("0x3d6850a4A9790c3aD3924A5d66b4fEEC8cd25bE2"), // implementation
    salt,
    deployerAddress,
)
```

#### 创建企业钱包

```go
import "github.com/vikkkko/safe-core-sdk-golang/protocol/utils"

// 定义方法权限配置
methodSelectors := [][4]byte{
    utils.CreatePaymentAccountSelector,
    utils.CreateCollectionAccountSelector,
    utils.ApproveTokenForPaymentSelector,
}

// 为每个方法指定控制Safe地址
configs := make([]utils.MethodConfig, len(methodSelectors))
for i := range methodSelectors {
    configs[i] = utils.MethodConfig{
        Controller: safeAddress, // 控制此方法的Safe地址
    }
}

initParams := utils.InitParams{
    Methods:    methodSelectors,
    Configs:    configs,
    SuperAdmin: superAdminSafeAddress, // 超级管理员Safe地址
}

// 生成部署calldata
deployData, err := utils.CreateEnterpriseWalletData(
    implementationAddress,
    salt,
    initParams,
)
```

#### 创建Payment Account

```go
// 生成创建Payment Account的calldata
paymentAccountData, err := utils.CreatePaymentAccountData(
    "Treasury Payment Account",
    controllerSafeAddress,
)

// 通过控制此方法的Safe发送交易到企业钱包
```

#### 创建Collection Account

```go
// 生成创建Collection Account的calldata
collectionAccountData, err := utils.CreateCollectionAccountData(
    "Revenue Collection Account",
    common.Address{}, // address(0) 默认收集到企业钱包
)
```

#### 批准代币用于支付

```go
// 批准USDC代币给Payment Account使用
usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
amount := new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e6)) // 1000 USDC

approveData, err := utils.ApproveTokenForPaymentData(
    usdcAddress,
    paymentAccountAddress,
    amount,
)
```

#### 转账ETH到Payment Account

```go
// 转账1 ETH给Payment Account
ethAmount := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))

ethTransferData, err := utils.TransferETHToPaymentData(
    paymentAccountAddress,
    ethAmount,
)
```

#### 从Collection Account收集资金

```go
// 收集ETH
collectData, err := utils.CollectFundsData(
    common.Address{}, // address(0) 表示收集ETH
    collectionAccountAddress,
)

// 收集ERC20代币
collectTokenData, err := utils.CollectFundsData(
    tokenAddress,
    collectionAccountAddress,
)
```

## 与Safe钱包集成

企业钱包设计用于与Safe多签钱包配合使用：

1. **部署企业钱包**
   - 通过Factory合约部署
   - 配置方法权限给不同的Safe地址

2. **方法级权限控制**
   ```
   - Safe A (2/3) 控制 createPaymentAccount
   - Safe B (3/5) 控制 approveTokenForPayment
   - Safe C (4/7) 作为超级管理员
   ```

3. **交易流程**
   ```
   1. 使用SDK生成交易calldata
   2. 通过Safe Protocol Kit创建Safe交易
   3. 收集Safe owners的签名
   4. 执行交易
   ```

## 方法选择器

SDK提供了预定义的方法选择器：

```go
utils.CreatePaymentAccountSelector    // 0x08f25c4a
utils.CreateCollectionAccountSelector // 0xc8ac06ed
utils.ApproveTokenForPaymentSelector  // 0xa5648c4f
utils.TransferETHToPaymentSelector    // 0x9ce5949e
utils.CollectFundsSelector            // 0xdd6890ef
```

或者自定义计算：

```go
selector := utils.GetMethodSelector("methodName(paramTypes)")
```

## 完整工作流程示例

```go
package main

import (
    "github.com/vikkkko/safe-core-sdk-golang/protocol"
    "github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
    "github.com/vikkkko/safe-core-sdk-golang/protocol/contracts"
)

func main() {
    // 1. 连接到Factory合约
    factoryContract, _ := contracts.NewEnterpriseWalletFactory(factoryAddress, client)

    // 2. 配置企业钱包
    initParams := utils.InitParams{
        Methods: [][4]byte{
            utils.CreatePaymentAccountSelector,
            utils.ApproveTokenForPaymentSelector,
        },
        Configs: []utils.MethodConfig{
            {Controller: safeA},
            {Controller: safeB},
        },
        SuperAdmin: safeC,
    }

    // 3. 通过Safe部署企业钱包
    deployData, _ := utils.CreateEnterpriseWalletData(implAddr, salt, initParams)
    // 将deployData发送到Factory合约...

    // 4. 创建Payment Account (由Safe A控制)
    paymentData, _ := utils.CreatePaymentAccountData("Account 1", safeA)
    // 通过Safe A创建交易发送到企业钱包...

    // 5. 批准代币 (由Safe B控制)
    approveData, _ := utils.ApproveTokenForPaymentData(usdc, paymentAddr, amount)
    // 通过Safe B创建交易发送到企业钱包...
}
```

## 项目结构

```
safe-core-sdk-golang/
├── abi/
│   ├── EnterpriseWallet.json              # 企业钱包ABI
│   ├── EnterpriseWalletFactory.json       # 工厂合约ABI
│   ├── EnterpriseWallet_full.json         # 完整ABI+Bytecode
│   └── EnterpriseWalletFactory_full.json  # 完整ABI+Bytecode
├── protocol/
│   ├── contracts/
│   │   ├── enterprise_wallet.go           # 企业钱包Go绑定
│   │   └── enterprise_wallet_factory.go   # 工厂合约Go绑定
│   └── utils/
│       └── enterprise_wallet.go           # 工具函数
└── examples/
    └── enterprise_wallet_example.go       # 完整示例
```

## API参考

### Factory合约

- `CreateWallet(implementation, salt, initParams)` - 部署新的企业钱包
- `PredictWalletAddress(implementation, salt, deployer)` - 预测钱包地址
- `IsImplementationWhitelisted(implementation)` - 检查实现是否在白名单中

### Enterprise Wallet合约

- `createPaymentAccount(name, controller)` - 创建支付账户
- `createCollectionAccount(name, collectionTarget)` - 创建收款账户
- `approveTokenForPayment(token, paymentAccount, amount)` - 批准代币
- `transferETHToPayment(paymentAccount, amount)` - 转账ETH
- `collectFunds(token, collectionAccount)` - 收集资金
- `getPaymentAccounts()` - 获取所有支付账户
- `getCollectionAccounts()` - 获取所有收款账户

### Utils工具函数

- `CreateEnterpriseWalletData()` - 生成部署calldata
- `CreatePaymentAccountData()` - 生成创建支付账户calldata
- `CreateCollectionAccountData()` - 生成创建收款账户calldata
- `ApproveTokenForPaymentData()` - 生成批准代币calldata
- `TransferETHToPaymentData()` - 生成转账ETH calldata
- `CollectFundsData()` - 生成收集资金calldata
- `GetMethodSelector()` - 计算方法选择器

## 安全考虑

1. **权限隔离** - 不同操作由不同的Safe控制，降低单点风险
2. **应急控制** - Super Admin可以冻结账户或暂停合约
3. **多签验证** - 所有关键操作需要Safe owners多签批准
4. **地址预测** - 使用CREATE2确保地址可预测性

## 相关资源

- [Safe官方文档](https://docs.safe.global/)
- [企业钱包合约代码](https://github.com/user/EnterpriseWalletSuite)
- [Safe Core SDK文档](./README.md)

---

**注意**: 在生产环境使用前，请确保：
1. 合约已经过审计
2. Implementation已添加到Factory白名单
3. 正确配置Safe多签钱包权限
4. 充分测试所有操作流程
