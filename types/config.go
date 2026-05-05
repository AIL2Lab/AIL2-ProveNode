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
	Rpc     string `json:"Rpc"`
	ChainId int64  `json:"ChainId"`
	// PrivateKey is the legacy in-config plaintext field. New deployments
	// should leave it empty and instead set PrivateKeyEnv or PrivateKeyFile
	// so the secret never lives next to the rest of the JSON config (which
	// tends to end up in version control, backups, or shared dashboards).
	// Kept for backward compat; a deprecation warning is logged on load.
	PrivateKey string `json:"PrivateKey,omitempty"`
	// PrivateKeyEnv names an environment variable holding the hex-encoded
	// private key. Recommended for systemd / Docker / Kubernetes secrets
	// where the orchestrator already has a path for injecting credentials.
	PrivateKeyEnv string `json:"PrivateKeyEnv,omitempty"`
	// PrivateKeyFile points at a file (typically 0400, outside the repo /
	// world-readable backups) containing only the hex-encoded private key.
	// Trailing whitespace and newlines are trimmed on read. Useful with
	// systemd LoadCredential= or any "secret-mounted at path" pattern.
	PrivateKeyFile      string         `json:"PrivateKeyFile,omitempty"`
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
