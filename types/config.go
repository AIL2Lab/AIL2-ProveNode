package types

import (
	"encoding/json"
	"os"
)

type MongoDB struct {
	URI        string `json:"URI"`
	Database   string `json:"Database"`
	ExpireTime int64  `json:"ExpireTime"`
}

type IP2LocationDB struct {
	DatabasePath string `json:"DatabasePath"`
}

type Prometheus struct {
	JobName        string `json:"JobName"`
	RemoteWriteURL string `json:"RemoteWriteURL"`
}

type ContractConfig struct {
	AbiFile         string `json:"AbiFile"`
	ContractAddress string `json:"ContractAddress"`
}

type ChainConfig struct {
	Rpc                 string         `json:"Rpc"`
	ChainId             int64          `json:"ChainId"`
	PrivateKey          string         `json:"PrivateKey"`
	ReportContract      ContractConfig `json:"ReportContract"`
	MachineInfoContract ContractConfig `json:"MachineInfoContract"`
}

type Certificate struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type NotifyThirdParty struct {
	OfflineNotify string `json:"OfflineNotify"`
}

type Config struct {
	Addr             string           `json:"Addr"`
	LogLevel         string           `json:"LogLevel"`
	LogFile          string           `json:"LogFile"`
	MongoDB          MongoDB          `json:"MongoDB"`
	IP2LDB           IP2LocationDB    `json:"IP2LDB"`
	Prometheus       Prometheus       `json:"Prometheus"`
	Chain            ChainConfig      `json:"Chain"`
	Certificate      Certificate      `json:"Certificate"`
	NotifyThirdParty NotifyThirdParty `json:"NotifyThirdParty"`
}

func LoadConfig(configPath string) (*Config, error) {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
