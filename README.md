# Safe Core SDK for Go

🔐 Go 语言版本的 Safe 多签钱包 SDK，提供完整的钱包创建、交易签名和 API 集成功能。

## ✨ 核心功能

- ✅ **Safe 钱包创建** - CREATE2 地址预测、多签配置、工厂部署
- ✅ **交易管理** - EIP-712 签名、交易哈希计算、多签收集
- ✅ **API 集成** - Safe Transaction Service 完整支持
- ✅ **ERC20 操作** - 基于 ABI 的代币转账、授权、查询
- ✅ **类型安全** - 完整的 Go 类型定义和编译时检查

## 🚀 快速开始

### 1. 安装

```bash
go get github.com/vikkkko/safe-core-sdk-golang
go get github.com/ethereum/go-ethereum
```

### 2. 配置环境

复制 `.env.example` 为 `.env` 并填入配置：

```bash
# Ethereum RPC
RPC_URL=https://sepolia.infura.io/v3/YOUR_INFURA_KEY
CHAIN_ID=11155111

# Safe 配置
SAFE_ADDRESS=0x9aE1311B4c25c9F95b5a5De5AD1b5e6D89dC3e25

# 私钥（仅用于测试网）
DEPLOYER_PRIVATE_KEY=your_private_key_here
OWNER_PRIVATE_KEY=your_private_key_here

# Safe API
SAFE_API_KEY=your_api_key_here
```

### 3. 创建多签钱包

```go
package main

import (
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/vikkkko/safe-core-sdk-golang/protocol/utils"
)

func main() {
    // 配置 2/3 多签钱包
    owners := []common.Address{
        common.HexToAddress("0x9C126aa4Eb6D110D646139969774F2c5b64dD279"),
        common.HexToAddress("0xeB7E951F2D1A38188762dF12E0703aE16F76ab73"),
        common.HexToAddress("0x74f4EFFb0B538BAec703346b03B6d9292f53A4CD"),
    }

    // 准备部署
    callData, _ := utils.PrepareSafeDeployment(utils.DeploySafeConfig{
        Owners:           owners,
        Threshold:        2,
        FactoryAddress:   common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
        SingletonAddress: common.HexToAddress("0x29fcB43b46531BcA003ddC8FCB67FFE91900C762"),
        SaltNonce:        big.NewInt(0),
    })

    // 使用 callData 发送交易到工厂合约即可部署 Safe
}
```

### 4. 管理 Safe 钱包

```go
package main

import (
    "context"
    "github.com/vikkkko/safe-core-sdk-golang/protocol"
    "github.com/vikkkko/safe-core-sdk-golang/api"
)

func main() {
    // 初始化 Safe 客户端
    safeClient, _ := protocol.NewSafe(protocol.SafeConfig{
        SafeAddress: "0x447d4227d88D6A7fB1486879be24Be00418A5fB7",
        RpcURL:      "https://sepolia.infura.io/v3/YOUR_KEY",
        ChainID:     11155111,
    })

    // 初始化 API 客户端
    apiClient, _ := api.NewSafeApiKit(api.SafeApiKitConfig{
        ChainID: 11155111,
        ApiKey:  "YOUR_API_KEY",
    })

    // 获取 Safe 信息
    ctx := context.Background()
    safeInfo, _ := apiClient.GetSafeInfo(ctx, "0x447d4227d88D6A7fB1486879be24Be00418A5fB7")
}
```

## 📁 项目结构

```
safe-core-sdk-golang/
├── api/                    # Safe Transaction Service API 客户端
│   ├── client.go          # HTTP 客户端和 API 调用
│   └── types.go           # API 响应类型定义
├── protocol/              # Safe 协议交互
│   ├── safe.go            # Safe 客户端主要功能
│   ├── contracts/         # 合约绑定（自动生成）
│   ├── managers/          # 交易和签名管理器
│   └── utils/             # 工具函数
│       ├── safe.go            # Safe 初始化工具
│       ├── safe_factory.go    # 工厂部署工具
│       ├── signatures.go      # EIP-712 签名
│       ├── erc20.go           # ERC20 ABI 工具
│       └── address.go         # CREATE2 地址计算
├── types/                 # 核心类型定义
│   └── types.go           # Safe 交易和签名类型
├── examples/              # 示例代码
│   ├── create_multisig_wallet.go  # 创建多签钱包
│   └── transaction_workflow.go    # 交易工作流程
└── tests/                 # 测试
    ├── unit/              # 单元测试
    └── integration/       # 集成测试
```

## 📚 示例代码

### 创建多签钱包

```bash
go run examples/create_multisig_wallet.go
```

演示完整的钱包创建流程：
- 配置多个所有者和签名阈值
- 预测 Safe 合约地址（CREATE2）
- 生成部署交易数据
- 部署并验证

### 交易工作流程

```bash
go run examples/transaction_workflow.go
```

演示完整的 ERC20 转账流程：
- Safe 客户端初始化
- 创建 USDC 转账交易
- EIP-712 交易哈希计算
- Safe owner 签名验证
- 提交到 Safe Transaction Service

## 🛠️ 主要组件

### Protocol Kit (`protocol/`)
- **Safe 客户端** - 连接和管理 Safe 钱包
- **交易创建** - 构建 Safe 兼容的交易
- **EIP-712 签名** - 符合 Safe 标准的交易签名

### API Kit (`api/`)
- **Safe Transaction Service** - 官方 API 集成
- **交易管理** - 提案、确认、查询交易
- **Safe 信息** - 获取配置、所有者、历史记录

### Utils (`protocol/utils/`)
- **Safe 部署** - 钱包创建和初始化工具
- **ERC20 工具** - 标准代币操作 ABI
- **签名工具** - EIP-712 签名和验证
- **地址计算** - CREATE2 地址预测

完整的工具包使用指南请参考 [`protocol/utils/README.md`](./protocol/utils/README.md)

## 🧪 测试

```bash
# 运行所有测试
go test ./...

# 运行单元测试
go test ./tests/unit

# 运行集成测试（需要配置环境变量）
RUN_INTEGRATION_TESTS=true go test ./tests/integration
```

## 📖 文档

- **[多签钱包工作流程](./MULTISIG_WORKFLOW.md)** - 完整的创建和管理指南
- **[Utils 工具包文档](./protocol/utils/README.md)** - Safe 部署工具使用说明
- **[贡献指南](./CONTRIBUTING.md)** - 如何参与项目开发

## 🌐 网络支持

### Sepolia 测试网

```go
const (
    FactoryAddress   = "0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"
    SingletonAddress = "0x29fcB43b46531BcA003ddC8FCB67FFE91900C762"
    FallbackHandler  = "0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99"
)
```

更多网络地址请参考：https://docs.safe.global/safe-smart-account/supported-networks

## ✅ 生产就绪

- ✅ **EIP-712 哈希验证** - 与 TypeScript SDK 计算结果一致
- ✅ **签名验证通过** - Safe 服务成功验证签名
- ✅ **API 提交成功** - 实际交易提案成功提交
- ✅ **多签工作流程** - 完整的 2/3 多签演示
- ✅ **CREATE2 地址预测** - 准确预测部署地址
- ✅ **完整测试覆盖** - 单元测试和集成测试

## 🔗 相关资源

- [Safe 官方文档](https://docs.safe.global/)
- [Safe Transaction Service API](https://safe-transaction-mainnet.safe.global/)
- [TypeScript SDK 源码](https://github.com/safe-global/safe-core-sdk)
- [Go-Ethereum 文档](https://geth.ethereum.org/docs/)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！详见 [CONTRIBUTING.md](./CONTRIBUTING.md)

## 📄 许可证

本项目采用与 Safe 官方 SDK 相同的许可证。

---

🎉 **开始使用 Go 构建安全的多签钱包应用！**
