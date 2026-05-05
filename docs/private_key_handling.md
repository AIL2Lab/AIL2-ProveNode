# Private key handling

ProveNode 里的链上交易私钥（`Chain.PrivateKey` 字段，hex-encoded secp256k1）现在支持三种来源，按推荐顺序：

## 推荐：环境变量（PrivateKeyEnv）

适合 systemd / Docker / Kubernetes Secret 等编排器注入凭据的场景。私钥不落盘也不进配置文件。

```json
{
  "Chain": {
    "Rpc": "...",
    "ChainId": ...,
    "PrivateKeyEnv": "AIL2_DBC_PRIVATE_KEY",
    "ReportContract": { ... },
    "MachineInfoContract": { ... }
  }
}
```

启动时设置：

```bash
# systemd
LoadCredential=dbc-key:/run/secrets/dbc-key
Environment=AIL2_DBC_PRIVATE_KEY=...

# Docker
docker run -e AIL2_DBC_PRIVATE_KEY="$(cat /secrets/dbc-key)" ...

# Kubernetes
env:
  - name: AIL2_DBC_PRIVATE_KEY
    valueFrom:
      secretKeyRef:
        name: dbc-secrets
        key: private-key
```

## 也可以：磁盘文件（PrivateKeyFile）

适合 systemd `LoadCredential=` 或 secret volume 挂载到固定路径。文件应当 `chmod 0400`，与 ProveNode 二进制 / 配置分目录存放，并加进 `.gitignore`。

```json
{
  "Chain": {
    "PrivateKeyFile": "/etc/ail2/dbc.key",
    ...
  }
}
```

文件内容：单行 hex 字符串（前后空白会被 trim）。

```text
abcdef0123456789... (64 hex chars)
```

## 不推荐：内联（PrivateKey，已 deprecated）

旧配置格式仍然能跑，但每次启动会打 warn 日志：

```
Chain.PrivateKey is set inline in the config — migrate to PrivateKeyEnv or PrivateKeyFile ...
```

```json
{
  "Chain": {
    "PrivateKey": "abcdef...",
    ...
  }
}
```

风险：

- 备份系统会把整个 config 一起拉走，私钥跟着泄漏
- 误提交到版本控制
- 远程查看 / 日志聚合时可能被截图记录

## 解析顺序

`InitDbcChain` 按下面顺序解析，谁先非空用谁：

1. `PrivateKeyEnv` → `os.Getenv(name)`
2. `PrivateKeyFile` → 读文件，trim 空白
3. `PrivateKey` → 直接用，记 warn

任何一步取到的值都会过 `crypto.HexToECDSA` 校验格式，非法 hex 直接 fail-fast。

## 密钥轮换流程

1. 在受控环境用 `geth account new` / `cast wallet new` 生成新密钥
2. 把新地址加入合约 ACL（如有）；旧地址保留一段过渡期
3. 更新对应的 env 变量或文件，重启 ProveNode
4. 监控合约事件验证新密钥能正常签 tx
5. 从合约 ACL 移除旧地址

ProveNode 进程自身不持久化任何与私钥相关的状态（除了内存中），所以"重启即生效"。
