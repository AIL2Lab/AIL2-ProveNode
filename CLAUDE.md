# AIL2-ProveNode — Claude 项目指南

> 本文件是写给 AI 编码助手（如 Claude Code）的项目说明。任何需要读代码或改代码的任务，请先阅读本文件。

## 1. 项目定位

**AIL2-ProveNode** 是 AIL2 网络的 **节点验证与链上报告系统**。它运行在参与计算的物理机上，负责：

- 采集硬件信息（CPU、GPU、内存、磁盘、带宽）
- 根据 GPU 型号计算 **算力积分**（compute score）
- 定期通过 **DBC 智能合约** 向链上汇报节点状态、质押信息与 AI 任务完成证明（"Prove"）
- 通过 WebSocket 实时推送节点状态
- 暴露 Prometheus 指标用于监控

本项目与 DBC（DeepBrainChain）链深度集成，是 AIL2 算力网络质押与任务分配机制的数据来源。

## 2. 仓库结构

```
AIL2-ProveNode/
├── ail2/                          # 核心业务逻辑
│   ├── ai-report/                # AI 任务报告（合约 ABI 自动绑定）
│   ├── calculator/               # GPU 积分计算器
│   │   └── calculator.go        # 加载 gpus.json，按型号评分
│   ├── machine-infos/            # 机器信息合约绑定
│   └── dbc.go                    # DBC 链交互主逻辑（合约读写、私钥签名）
├── app/                           # 可执行程序入口
│   ├── bandwidth/bandwidth.go   # 带宽测速工具 main 包
│   └── ddn/ddn.go                # DDN 节点守护进程 main 包
├── db/
│   ├── mongo.go                  # MongoDB 连接与 CRUD
│   └── ip2location.go            # IP 地理位置查询
├── http/
│   ├── contract.go               # 合约交互 HTTP 端点（Gin）
│   ├── query.go                  # 数据查询端点
│   └── prometheus.go             # /metrics 端点
├── ws/
│   ├── client.go                 # WebSocket 客户端（19KB，主要消息处理）
│   ├── hub.go                    # 消息总线
│   └── echo.go                   # 信息回显
├── types/                         # 共享数据结构
├── log/                           # 日志封装
├── gpus.json                      # GPU 型号 → 算力评分表（数据资产）
├── go.mod / go.sum
├── LICENSE
└── README.md
```

**测试文件**：`ail2/dbc_test.go`、`ail2/calculator_test.go`、`db/mongo_test.go`、`db/ip2location_test.go`、`http/prometheus_test.go`、`ws/*_test.go`。

## 3. 技术栈

| 维度 | 选型 |
|------|------|
| 语言 | **Go 1.22.9** |
| HTTP 框架 | **gin-gonic/gin** |
| WebSocket | **gorilla/websocket** |
| 数据库 | **MongoDB**（`go.mongodb.org/mongo-driver`）|
| 链交互 | **ethereum/go-ethereum**（适配 DBC EVM 兼容链）|
| 系统监控 | **shirou/gopsutil** |
| 带宽测试 | **showwin/speedtest-go** |
| 监控暴露 | **prometheus/client_golang** |

**尚未配置**：Makefile、golangci-lint 配置、`.gitignore`、CI/CD、pre-commit。

## 4. 开发工作流程

### 构建与运行

```bash
go mod download                           # 拉依赖
go build -o bin/ddn ./app/ddn             # 构建 DDN 节点守护进程
go build -o bin/bandwidth ./app/bandwidth # 构建带宽工具
go test ./...                             # 跑全量测试
go test -run TestSpecific ./ail2/...      # 跑单个测试
go vet ./...                              # 静态检查
gofmt -w .                                # 格式化
```

### 配置与密钥

- 节点运行需要 **DBC 链私钥**（签名合约交易），**严禁**硬编码、**严禁**入库。用环境变量或独立的（gitignore 的）配置文件。
- MongoDB 连接串同样通过环境变量注入。

### 合约绑定生成（如需更新）

`ail2/ai-report/`、`ail2/machine-infos/` 目录下的 Go 文件是 `abigen` 从合约 ABI 自动生成的 — **不要手改**。如需更新：

```bash
abigen --abi=<abi.json> --pkg=<pkg> --out=<name>.go
```

## 5. 给 AI 助手的关键约定

1. **中文文字**：文档、commit 正文、日志 user-facing 信息用中文；代码标识符、公开的 JSON 字段、protobuf 类型名用英文。
2. **禁碰链钱包**：任何涉及私钥加载、签名、发送交易的代码改动都要极为谨慎。默认不改，改前先让用户确认。测试用测试网地址，永远不要在测试里用主网私钥。
3. **合约绑定代码免动**：`ail2/ai-report/`、`ail2/machine-infos/` 中明显是 abigen 生成的文件（通常有 "Code generated" 注释头）**不要手动编辑**。改需求要改 ABI 源头重新生成。
4. **`gpus.json` 是数据契约**：这是链上报告的算力评分依据。修改前必须确认是否会影响已上链数据的兼容性。加新 GPU 型号可以，改已有评分需慎重。
5. **Go 风格**：遵循标准 `gofmt`、`go vet` 清洁；错误用 `fmt.Errorf("context: %w", err)` 包装；并发安全优先 `sync` 原语。
6. **测试存在**：改核心模块（`calculator`、`dbc`、`mongo`）时先跑相关 `_test.go`，别让已有测试回归。
7. **无 `.gitignore`**：现状是所有文件都被追踪（包括 `.DS_Store`）。不要新增大型 binary 或本地调试产物；建议在新增文件前考虑是否需要先加一份 `.gitignore`。
8. **WebSocket hub**：`ws/client.go` 体量较大（~19KB），改 hub/client 消息协议前先通读。
9. **Prometheus 指标**：新增指标要遵循 Prometheus 命名规范（`ail2_provenode_<subsystem>_<metric>_<unit>`），且要在 `http/prometheus.go` 注册。
10. **MongoDB schema**：`db/mongo.go` 中的 struct tag 是 schema 的唯一真相，改字段名会破坏历史数据。

## 6. 多仓库协作上下文

Claude 同时管理 `C:\PROJECT\AIL2\` 下 4 个 AIL2Lab 仓库：

| 目录 | 角色 |
|------|------|
| `AIL2DEVOPS/` | DevOps / 部署中枢 |
| `AIL2-Builder/` | 官网前端 + CMS（Next.js 16）|
| `AIL2-ProveNode/` | **本仓库** — 节点验证与链上报告 |
| `AIL2-ComputeNet/` | 分布式 AI 计算网络（Go + libp2p）|

**相关联动**：
- **ProveNode 的节点信息** 常被 **ComputeNet** 调度器用于任务分派 → 两者共享的 schema（机器信息、算力积分）改动要双向同步。
- **Builder 官网**的实时节点数、算力总量通常来自 ProveNode 的 HTTP query 接口 → 改 `http/query.go` 前考虑前端影响。
- 部署脚本应落到 `AIL2DEVOPS/deploy/provenode/`。

## 7. 历史对话总结

**2026-04-19 首次对话**，为 AIL2Lab 建立多仓库工作区：

1. 用户要求创建 SSH 密钥对（邮箱 `official@ail2.org`），生成 `~/.ssh/id_ed25519_ail2`（ed25519，无密码短语）。
2. 公钥被添加到 GitHub（账户 `AIL2-coder`）。
3. 初次通过公开 API 列出了 AIL2-Builder 与 **AIL2-ProveNode** 两个 repo，完成克隆（本仓库 HEAD：`1eaa790 first`）。
4. 随后用户补充了 AIL2-ComputeNet 并完成克隆。
5. `~/.ssh/config` 新增 `Host github-ail2` 别名；所有 4 个 repo 的 origin 切换为 `git@github-ail2:AIL2Lab/<name>.git`，`git pull/push` 自动用新密钥。
6. 额外拉取 `AIL2DEVOPS` 到工作区作为 DevOps 中枢。
7. 用户明确 Claude 今后同时管理 4 个 AIL2Lab 仓库。
8. **本次行动**：为 4 个 repo 各生成一份 CLAUDE.md。

## 8. 下一步建议

- 补 `.gitignore`（至少忽略 `.DS_Store`、`bin/`、`*.log`、本地 `config.yaml`）
- 写 `Makefile`：`make build`、`make test`、`make vet`、`make lint`
- 接入 `golangci-lint` 并加 `.github/workflows/`
- 区分 dev / testnet / mainnet 配置（环境变量或 config 文件）
- 核心合约交互加集成测试（用本地 anvil/ganache 或 DBC 测试网）
- 部署脚本 + systemd unit 文件归档到 `AIL2DEVOPS/deploy/provenode/`
