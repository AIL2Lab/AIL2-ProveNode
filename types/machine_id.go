package types

import "errors"

type MachineKey struct {
	MachineId   string `json:"machine_id" bson:"machine_id"`
	Project     string `json:"project" bson:"project"`
	ContainerId string `json:"container_id" bson:"container_id"`
}

// type MachineInfo struct {
// 	// Project string `json:"project" bson:"project"`
// 	// Models         []ModelInfo `json:"models" bson:"models"`
// 	GPUNames       []string `json:"gpu_names" bson:"gpu_names"`
// 	GPUMemoryTotal []int32  `json:"gpu_memory_total" bson:"gpu_memory_total"`
// 	// UtilizationGPU int      `json:"utilization_gpu" bson:"utilization_gpu"`
// 	MemoryTotal int64 `json:"memory_total" bson:"memory_total"`
// 	// MemoryUsed     int64    `json:"memory_used" bson:"memory_used"`
// 	CpuType  string `json:"cpu_type" bson:"cpu_type"`
// 	CpuRate  int32  `json:"cpu_rate" bson:"cpu_rate"`
// 	Wallet   string `json:"wallet" bson:"wallet"`
// 	ClientIP string `json:"client_ip" bson:"client_ip"`
// }

// machine info of deeplink short-term
type DeepLinkMachineInfoST struct {
	GPUNames       []string `json:"gpu_names" bson:"gpu_names,omitempty"`
	GPUMemoryTotal []int32  `json:"gpu_memory_total" bson:"gpu_memory_total,omitempty"` // GB
	MemoryTotal    int64    `json:"memory_total" bson:"memory_total,omitempty"`         // GB
	CpuType        string   `json:"cpu_type" bson:"cpu_type,omitempty"`
	CpuRate        int32    `json:"cpu_rate" bson:"cpu_rate,omitempty"`
	Wallet         string   `json:"wallet" bson:"wallet,omitempty"`
	ClientIP       string   `json:"client_ip" bson:"client_ip,omitempty"`
}

// machine info of deeplink bandwidth
type DeepLinkMachineInfoBandwidth struct {
	CpuCores    int32  `json:"cpu_cores" bson:"cpu_cores,omitempty"`
	MemoryTotal int64  `json:"memory_total" bson:"memory_total,omitempty"` // GB
	Hdd         int64  `json:"hdd" bson:"hdd,omitempty"`
	Bandwidth   int32  `json:"bandwidth" bson:"bandwidth,omitempty"`
	Wallet      string `json:"wallet" bson:"wallet,omitempty"`
	ClientIP    string `json:"client_ip" bson:"client_ip,omitempty"`
}

func (info *DeepLinkMachineInfoST) Validate() error {
	if len(info.GPUNames) == 0 {
		return errors.New("empty gpu names")
	}
	if len(info.GPUMemoryTotal) == 0 {
		return errors.New("empty gpu memory total")
	}
	if info.MemoryTotal == 0 {
		return errors.New("invalid memory total")
	}
	if info.CpuType == "" {
		return errors.New("empty cpu type")
	}
	if info.CpuRate == 0 {
		return errors.New("invalid cpu rate")
	}
	if info.Wallet == "" {
		return errors.New("empty wallet")
	}
	return nil
}

func (info *DeepLinkMachineInfoBandwidth) Validate() error {
	if info.CpuCores == 0 {
		return errors.New("invalid cpu cores")
	}
	if info.MemoryTotal == 0 {
		return errors.New("invalid memory total")
	}
	if info.Hdd == 0 {
		return errors.New("invalid hdd")
	}
	if info.Bandwidth == 0 {
		return errors.New("invalid bandwidth")
	}
	if info.Wallet == "" {
		return errors.New("empty wallet")
	}
	return nil
}
