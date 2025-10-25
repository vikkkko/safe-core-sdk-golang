# GraphQL 使用指南

## 快速开始

### 方式 1: 交互式 CLI（推荐）

```bash
go build -o graphql_example examples/graphql_example.go
./graphql_example
```

详见：[GraphQL Example README](../examples/README_GRAPHQL.md)

### 方式 2: 代码集成

```go
import "github.com/vikkkko/safe-core-sdk-golang/graphql"

client := graphql.NewDefaultClient()
defer client.Close()

ctx := context.Background()
allowances, err := client.GetPaymentAllowances(ctx, "0x...")
```

详见：[GraphQL Module README](../graphql/README.md)

## 环境变量

```bash
# .env 文件
GRAPH_ENDPOINT=https://api.studio.thegraph.com/query/103887/mvp/version/latest
RPC_URL=https://your-rpc-endpoint  # 可选，用于丰富交易信息
```

## 可用查询

| 查询 | 方法 | 说明 |
|------|------|------|
| 授权记录 | `GetPaymentAllowances` | 查询哪些地址授权了付款账户 |
| 交易信息 | `GetTransactionInfo` | 通过交易哈希查询详情 |
| 付款账户余额 | `GetPaymentAccounts` | 查询付款账户代币余额 |
| 收款账户余额 | `GetCollectionAccounts` | 查询收款账户代币余额 |

## 注意事项

1. **地址格式**: GraphQL 查询要求小写地址
2. **大数处理**: Amount 字段使用自定义 `BigInt` 类型
3. **网络连接**: 需要网络访问 The Graph API
4. **RPC 配置**: 配置 `RPC_URL` 可以丰富交易信息（Gas Used, Status 等）

## 扩展

添加新查询的步骤见 [GraphQL Example README](../examples/README_GRAPHQL.md)
