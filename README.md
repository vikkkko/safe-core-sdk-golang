# Safe Core SDK for Go

🔐 一个完整的Go语言版本的Safe多签钱包SDK，移植自官方TypeScript版本。

## ✨ 功能特性

### 🏗️ 核心功能
- ✅ **Safe钱包创建** - 使用CREATE2预测地址，支持多签配置
- ✅ **交易签名** - EIP-712标准签名，与Safe服务完全兼容
- ✅ **API集成** - Safe Transaction Service完整API支持
- ✅ **ERC20支持** - 基于ABI的代币操作，类型安全

### 🔧 技术特性
- 🎯 **EIP-712兼容** - 正确的类型化数据签名
- 🛡️ **类型安全** - 完整的Go类型定义和验证
- 📦 **ABI编码** - 使用go-ethereum标准ABI包
- 🌐 **多网络** - 支持主网、测试网等多个网络

## 🚀 快速开始

### 安装依赖

```bash
go mod init your-project
go get github.com/ethereum/go-ethereum
# 注意：实际使用时需要将import路径替换为实际的包路径
```

### 基本使用

```go
package main

import (
    "context"
    "github.com/vikkkko/safe-core-sdk-golang/protocol"
    "github.com/vikkkko/safe-core-sdk-golang/api"
)

func main() {
    // 创建Safe客户端
    safeClient, _ := protocol.NewSafe(protocol.SafeConfig{
        SafeAddress: "0x447d4227d88D6A7fB1486879be24Be00418A5fB7",
        RpcURL:      "https://sepolia.infura.io/v3/YOUR_INFURA_KEY",
        ChainID:     11155111,
    })

    // 创建API客户端
    apiClient, _ := api.NewSafeApiKit(api.SafeApiKitConfig{
        ChainID: 11155111,
        ApiKey:  "YOUR_SAFE_API_KEY",
    })

    // 获取Safe信息
    safeInfo, _ := apiClient.GetSafeInfo(context.Background(), "0x447d4227d88D6A7fB1486879be24Be00418A5fB7")
    fmt.Printf("Owners: %v\n", safeInfo.Owners)
    fmt.Printf("Threshold: %d\n", safeInfo.Threshold)
}
```

## 📁 项目结构

```
safe-core-sdk-golang/
├── api/                    # Safe Transaction Service API客户端
│   ├── client.go          # HTTP客户端和API调用
│   └── types.go           # API响应类型定义
├── protocol/              # Safe协议交互
│   ├── safe.go            # Safe客户端主要功能
│   └── utils/             # 工具函数
│       ├── signatures.go  # EIP-712签名和验证
│       ├── erc20.go       # ERC20 ABI工具
│       └── transactions.go # 交易处理工具
├── types/                 # 核心类型定义
│   └── types.go           # Safe交易和签名类型
├── examples/              # 示例代码
│   ├── create_multisig_wallet.go      # 多签钱包创建
│   └── transaction_workflow.go       # 交易工作流程
├── MULTISIG_WORKFLOW.md   # 多签钱包完整工作流程
└── README.md              # 项目说明
```

## 🧪 示例代码

### 1. 创建多签钱包

```bash
go run ./examples/create_multisig_wallet.go
```

演示如何配置和创建新的Safe多签钱包：
- 📋 配置多个所有者和签名阈值
- 🔮 预测Safe合约地址 (CREATE2)
- 📝 生成部署交易数据
- ⛽ Gas费用估算

### 2. 完整交易工作流程

```bash
go run ./examples/transaction_workflow.go
```

演示完整的ERC20转账流程：
- 🔧 Safe客户端初始化
- 💰 创建USDC转账交易
- 🔐 EIP-712交易哈希计算
- ✍️  Safe owner签名验证
- 📤 提交到Safe Transaction Service


## 🔧 核心组件

### Protocol Kit
- **Safe客户端管理** - 连接和配置Safe钱包
- **交易创建** - 构建Safe兼容的交易数据
- **EIP-712签名** - 符合Safe标准的交易哈希和签名

### API Kit
- **Safe Transaction Service** - 与官方API服务集成
- **交易提交** - 提案、确认、查询交易状态
- **Safe信息** - 获取钱包配置、所有者、历史记录

### Types Kit
- **类型定义** - Safe交易、签名、配置的Go类型
- **数据验证** - 确保类型安全和数据完整性

### Utils
- **ERC20工具** - 基于ABI的标准代币操作
- **签名工具** - EIP-712签名验证和恢复
- **交易工具** - 数据编码、哈希计算等

## 🎯 实际应用

### ✅ 已验证功能
- 🔐 **EIP-712哈希兼容** - 与TypeScript版本计算结果一致
- ✍️  **签名验证通过** - Safe服务成功验证签名
- 📤 **API提交成功** - 实际交易提案成功提交
- 🔄 **多签工作流程** - 完整的2/3多签演示

### 🚀 生产就绪特性
- **网络兼容** - 支持主网、Sepolia等多个网络
- **错误处理** - 完整的错误处理和调试信息
- **类型安全** - 编译时类型检查，避免运行时错误
- **标准兼容** - 使用以太坊生态标准工具

## 📚 文档

- **[多签钱包工作流程](./MULTISIG_WORKFLOW.md)** - 完整的创建和管理指南
- **示例代码** - 详细的使用示例和最佳实践
- **API文档** - 函数签名和参数说明

## 🔗 相关资源

- [Safe官方文档](https://docs.safe.global/)
- [Safe Transaction Service API](https://safe-transaction-mainnet.safe.global/)
- [TypeScript SDK源码](https://github.com/safe-global/safe-core-sdk)
- [Go-Ethereum文档](https://geth.ethereum.org/docs/)

## 🤝 贡献

欢迎提交Issue和Pull Request来改进这个SDK！

## 📄 许可证

本项目采用与Safe官方SDK相同的许可证。

---

🎉 **Ready to Build!** 开始使用Go构建安全的多签钱包应用吧！