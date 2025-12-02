**Safe 多通道升级指南（Go SDK）**

本文帮助已有基于 Safe 1.4/1.5 的开发者快速迁移到多通道版本（例如 `1.5.0+multichannel.1`），涵盖核心改动点、示例代码与注意事项。

---

### 1. 版本识别（必改）
- 判断逻辑：`strings.Contains(version, "multichannel")` 即视为多通道。
- 读取链上版本：`Safe.VERSION()`，不要依赖本地默认值。
- 示例：
```go
version, _ := safeClient.GetVersion(ctx)
isMulti := strings.Contains(version, "multichannel")
```

### 2. 交易数据结构（必改）
- 新增字段 `Channel`：
  - `types.SafeTransactionData{Channel uint64}`
  - `types.SafeTransactionDataPartial{Channel *uint64}`
- 多通道 Safe **必须显式传入 channel**；旧版可默认 0。
- 示例：
```go
channel := uint64(1) // 指定通道
tx := types.SafeTransactionDataPartial{
    To: "0x...", Value: "0", Data: "0x...",
    Channel: &channel,
    // Nonce: 可留空，SDK 自动按 channel 读取
}
```

### 3. Nonce 获取（必改）
- 旧版：`nonce()`。
- 多通道：`channelNonces(channel)`。
- SDK：
  - 多通道：`safeClient.GetChannelNonce(ctx, channel)`
  - 旧版：`safeClient.GetNonce(ctx)`（等价 channel 0）
- 示例：
```go
if isMulti {
    n, _ := safeClient.GetChannelNonce(ctx, channel)
    fmt.Println("channel nonce:", n)
} else {
    n, _ := safeClient.GetNonce(ctx)
    fmt.Println("nonce:", n)
}
```

### 4. 哈希计算 / ABI（必改）
- 多通道 SafeTx typehash：
  `keccak256("SafeTx(uint256 channel,address to,uint256 value,bytes data,uint8 operation,uint256 safeTxGas,uint256 baseGas,uint256 gasPrice,address gasToken,address refundReceiver,uint256 nonce)")`
  常量：`0x71678f49e7d9069b5052963a4bd4fe1effcf63d36fb57041d4499a96c1785d84`
- 合约函数签名增加 `channel`：
  - `execTransaction(uint256 channel, ...)`
  - `getTransactionHash(uint256 channel, ...)`
  - 新增只读 `channelNonces(uint256)`
- 示例（计算链上哈希）：
```go
txHash, _ := safeClient.GetTransactionHash(
    ctx,
    channel, // 新增
    to, value, data, operation,
    safeTxGas, baseGas, gasPrice, gasToken, refundReceiver,
    big.NewInt(int64(nonce)),
)
```
- 确保 ABI 来自多通道部署的 Safe.json/Safe_full.json，并重新生成绑定。

### 5. API 交互（必改）
- 请求/响应新增 `channel`：
  - `ProposeTransactionProps.Channel *int64`
  - `SafeMultisigTransactionResponse.Channel *int64`
  - `SafeMultisigTransactionEstimate.Channel *int64`
- Nonce API：使用 `/api/v1/safes/{addr}/nonces/`，返回 `{ "nonces": { "<channel>": "<nonce>" } }`。
- 示例（提交提案）：
```go
chPtr := func(c uint64)*int64{ v:=int64(c); return &v }
_, err := apiClient.ProposeTransaction(ctx, api.ProposeTransactionProps{
    Channel: chPtr(channel),
    SafeAddress: safeAddr,
    SafeTxHash: "0x"+safeTxHash,
    To: tx.Data.To,
    Value: tx.Data.Value,
    Data: tx.Data.Data,
    Operation: int(tx.Data.Operation),
    GasToken: tx.Data.GasToken,
    SafeTxGas: 0, BaseGas: 0,
    GasPrice: tx.Data.GasPrice,
    RefundReceiver: tx.Data.RefundReceiver,
    Nonce: int64(tx.Data.Nonce),
    Sender: signer.Hex(),
    Signature: "0x"+hex.EncodeToString(sig),
    ContractTransactionHash: "0x"+safeTxHash,
})
```
- 调试：设置 `SAFE_API_DEBUG=1` 查看完整 URL/status。

### 6. 签名与执行（必改）
- 签名对象需包含 channel 参与的哈希（见第 4 点）。
- `ExecuteTransaction` 调用必须传入 channel，与 `GetTransactionHash` 保持一致。
- 确保 `sender` 是 Safe owner，否则 tx-service 返回 422。

### 7. 示例变更速览
- `examples/create_multisig_wallet.go`：使用多通道 Safe/Factory/Handler 地址，预测地址时读取工厂 `proxyCreationCodehash`，部署/验证走本地 tx-service。
- `examples/safe_management_example.go`：
  - 支持选择 channel；按版本从链上读取对应 nonce。
  - 提交 tx-service 时携带 channel。
  - 支持选择签名私钥（避免 sender 非 owner）。
- `protocol/confirm.go`：从 API 响应重建交易时写入 channel，确保执行参数正确。

### 8. 常见错误与排查
- **422 contract-tx-hash mismatch**：未传 channel 或 nonce 来源错误（用错 channel）。按第 3/5/6 点检查。
- **revert on getTransactionHash/execTransaction**：仍在用旧 ABI/旧 typehash，或版本判断失败。确认 VERSION 包含 “multichannel” 并使用多通道 ABI。
- **sender 不是 owner**：提交提案时 Sender 应为 Safe owner；在示例中可选择正确私钥。

### 9. 环境变量建议
```
RPC_URL=...
CHAIN_ID=...
SAFE_API_URL=http://localhost:8000   # 本地 tx-service
SAFE_API_KEY=                        # 如无可留空
DEPLOYER_PRIVATE_KEY=...
OWNER_PRIVATE_KEY=...
OWNER2_PRIVATE_KEY=...
OWNER3_PRIVATE_KEY=...
```

### 10. 升级步骤清单
1) 更新 ABI/绑定至多通道 Safe。  
2) 数据结构加 `Channel`，调用处显式传入。  
3) 版本检测基于 `VERSION()`，`strings.Contains(version, "multichannel")`。  
4) Nonce 获取用 `channelNonces`（多通道）；构造哈希/执行传 channel。  
5) API 请求/响应带 `channel`，使用 `/nonces/`。  
6) 示例/业务代码按上述示例修改，测试提交与执行。  
7) 如有自建 tx-service，设置 `SAFE_API_URL` 并可用 `SAFE_API_DEBUG=1` 调试。
