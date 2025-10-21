# go-ethereum

https://goethereumbook.org/zh/smart-contract-compile/

https://docs.soliditylang.org/en/latest/installing-solidity.html

https://geth.ethereum.org/docs/getting-started/installing-geth

```shell
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
# 安装 solidity
sudo apt-get install solc
# 安装 abigen
# These commands install the core Geth software and the following developer tools: clef, devp2p, abigen, bootnode, evm and rlpdump. 
sudo apt-get install ethereum
```

```shell
brew update
brew tap ethereum/ethereum
brew install solidity
brew install ethereum
```

```shell
# 从一个 solidity 文件生成 ABI
solc --abi machine_infos.sol
# 用 abigen 将 ABI 转换为我们可以导入的 Go 文件。
# 这个新文件将包含我们可以用来与 Go 应用程序中的智能合约进行交互的所有可用方法。
abigen --abi=machine_infos.json --pkg=machineinfos --out=machine_infos.go
```

## 合约改动

### 2025.03.27
测试用的 report 合约地址换到: 0xBbD93B4D19D1bA99eae2fB7D14cd6711bb77709B

### 2025.03.07
在主网上部署升级测试网
"MachineInfos": "0xefaF6c5980CCf4CAD5c5Ee0D423BbEA66452be79"
report: 0xa7B9f404653841227AF204a561455113F36d8EC8

### machineinfos 合约更新
0xF9335c71583132d58E5320f73713beEf6da5257D
增加 region 地域字段

### 2025.02.11
report 合约更新了
0x5d72d4f8be9055f519cF49a7B5ED3De07FDDDa39
合约地址用这个吧 合约中其他函数由变动 notify 没变

### 2025.02.10
0xE676096cA8B957e914094c5e044Fcf99f5dbf3C0
更新硬件上报合约，信息里面新加了内存参数

### 2025.01.21
设置注册的机器信息: https://test.dbcscan.io/address/0x73F9467334CD7bD32b3c35f556AeBD7609B77d9B
