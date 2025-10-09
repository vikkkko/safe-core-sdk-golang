# Safe Protocol Utils

这个包提供了与 Safe 智能合约交互的完整工具集。

## 文件结构

### 核心工具文件
- **`safe.go`** - Safe 初始化和配置工具
- **`safe_factory.go`** - Safe 工厂部署工具
- **`erc20.go`** - ERC20 代币操作工具
- **`signatures.go`** - 签名和验证工具
- **`address.go`** - 地址相关工具
- **`transactions.go`** - 交易处理工具

### 合约绑定文件（自动生成）
- **`safe_contract.go`** (144KB) - 完整的 Safe 合约 Go 绑定
- **`safe_proxy_factory_contract.go`** (40KB) - SafeProxyFactory 合约 Go 绑定

这些绑定是使用 `abigen` 从官方 Safe 合约 ABI 生成的，提供了类型安全的合约调用。

## 快速开始

### 1. 部署 Safe 多签钱包

```go
import "github.com/vikkkko/safe-core-sdk-golang/protocol/utils"

// 准备 owners
owners := []common.Address{
    common.HexToAddress("0x1234..."),
    common.HexToAddress("0x5678..."),
}

// 使用一行代码准备部署数据
factoryCallData, err := utils.PrepareSafeDeployment(utils.DeploySafeConfig{
    Owners:           owners,
    Threshold:        2,
    FactoryAddress:   common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
    SingletonAddress: common.HexToAddress("0x29fcB43b46531BcA003ddC8FCB67FFE91900C762"),
    SaltNonce:        big.NewInt(0),
})

// factoryCallData 可以直接作为交易的 data 发送到工厂合约
```

### 2. 直接使用合约绑定

```go
// 创建 Safe 合约实例
safeContract, err := utils.NewSafeContract(safeAddress, client)

// 调用任何 Safe 合约方法
threshold, err := safeContract.GetThreshold(nil)
owners, err := safeContract.GetOwners(nil)
nonce, err := safeContract.Nonce(nil)

// 计算交易哈希
txHash, err := safeContract.GetTransactionHash(
    nil,
    to,
    value,
    data,
    operation,
    safeTxGas,
    baseGas,
    gasPrice,
    gasToken,
    refundReceiver,
    nonce,
)
```

### 3. 使用工厂合约

```go
// 创建工厂合约实例
factory, err := utils.NewSafeProxyFactoryContract(factoryAddress, client)

// 获取链 ID
chainId, err := factory.GetChainId(nil)

// 获取代理创建代码
proxyCode, err := factory.ProxyCreationCode(nil)

// 部署 Safe（需要发送交易）
tx, err := factory.CreateProxyWithNonce(
    auth,
    singletonAddress,
    initData,
    saltNonce,
)
```

## API 参考

### Safe 初始化

#### `CreateSafeInitData(config SafeSetupConfig) ([]byte, error)`
创建 Safe 的 `setup()` 函数调用数据

#### `CreateSafeInitDataSimple(owners []common.Address, threshold uint) ([]byte, error)`
简化版本，只需要 owners 和 threshold

#### `DefaultSafeSetupConfig(owners []common.Address, threshold uint) SafeSetupConfig`
创建带有合理默认值的配置

### Safe 工厂

#### `CreateSafeFactoryCallData(singleton, initData, saltNonce) ([]byte, error)`
创建工厂合约的 `createProxyWithNonce()` 调用数据

#### `CreateChainSpecificProxyCallData(singleton, initData, saltNonce) ([]byte, error)`
创建链特定的代理部署调用数据（用于 L2 网络）

#### `PrepareSafeDeployment(config DeploySafeConfig) ([]byte, error)`
一站式函数，准备完整的部署数据

### 合约方法

所有 Safe 合约方法都可以通过生成的绑定直接调用：

**Safe 合约方法：**
- `VERSION()` - 获取版本
- `nonce()` - 获取 nonce
- `getThreshold()` - 获取阈值
- `getOwners()` - 获取所有者列表
- `isOwner(address)` - 检查是否是所有者
- `getTransactionHash(...)` - 计算交易哈希
- `domainSeparator()` - 获取域分隔符
- `checkSignatures(...)` - 验证签名
- `execTransaction(...)` - 执行交易
- 等等...完整列表见 `safe_contract.go`

**SafeProxyFactory 合约方法：**
- `createProxyWithNonce(singleton, initializer, saltNonce)` - 创建代理
- `createChainSpecificProxyWithNonce(...)` - 创建链特定代理
- `getChainId()` - 获取链 ID
- `proxyCreationCode()` - 获取代理创建代码
- 等等...完整列表见 `safe_proxy_factory_contract.go`

## 网络地址

### Sepolia 测试网
```go
const (
    SepoliaFactoryAddress   = "0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"
    SepoliaSingletonAddress = "0x29fcB43b46531BcA003ddC8FCB67FFE91900C762"
    SepoliaFallbackHandler  = "0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99"
)
```

### 主网
请参考 Safe 官方文档获取最新的主网合约地址：
https://docs.safe.global/safe-smart-account/supported-networks

## ABI 来源

合约 ABI 从以下官方合约编译：
- **Safe.sol** - `/Users/yinwei/Work/yi/safe-golbal-mock/safe-smart-account/contracts/Safe.sol`
- **SafeProxyFactory.sol** - `/Users/yinwei/Work/yi/safe-golbal-mock/safe-smart-account/contracts/proxies/SafeProxyFactory.sol`

ABI 文件位置：
- `abi/Safe_full.json` (999 行)
- `abi/SafeProxyFactory_full.json` (250 行)

## 生成绑定代码

如果需要重新生成绑定代码：

```bash
# 生成 Safe 合约绑定
abigen --abi abi/Safe_full.json \
       --pkg utils \
       --type SafeContract \
       --out protocol/utils/safe_contract.go

# 生成 SafeProxyFactory 合约绑定
abigen --abi abi/SafeProxyFactory_full.json \
       --pkg utils \
       --type SafeProxyFactoryContract \
       --out protocol/utils/safe_proxy_factory_contract.go
```

## 测试

运行测试：
```bash
go test ./protocol/utils -v
```

所有测试包括：
- 单元测试
- 示例测试（会出现在 godoc 中）
- 集成测试

## 示例程序

完整示例见：
- `examples/create_multisig_wallet.go` - 创建多签钱包
- `examples/transaction_workflow.go` - Safe 交易流程

## 注意事项

1. **地址预测**: CREATE2 地址预测是简化实现，实际地址应从交易收据获取
2. **Gas 估算**: 部署 Safe 通常需要 300,000 - 800,000 gas
3. **签名格式**: 使用 EIP-712 签名，v 值应为 27 或 28
4. **版本兼容**: 绑定基于 Safe 1.4.1 版本

## 高级用法

### 调用任意 Safe 方法

```go
// 使用通用方法调用
result, err := utils.CallSafeMethod(
    client,
    safeAddress,
    "isOwner",
    common.HexToAddress("0x1234..."),
)
```

### 获取代理创建代码

```go
creationCode, err := utils.GetProxyCreationCode(client, factoryAddress)
```

### 链特定部署

```go
// 用于 L2 网络
callData, err := utils.CreateChainSpecificProxyCallData(
    singleton,
    initData,
    saltNonce,
)
```
