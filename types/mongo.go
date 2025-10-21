package types

import "time"

// machine connection table
type MDBMachineOnline struct {
	MachineKey `bson:",inline"`
	AddTime    time.Time `json:"add_time" bson:"add_time"`
}

// machine info table
type MDBMachineInfo struct {
	MachineKey `bson:",inline"`
	// MachineInfo        `bson:",inline"`
	StakingType        uint8     `json:"staking_type"`
	RegisterTime       time.Time `json:"register_time" bson:"register_time"`
	LastOfflineTime    time.Time `json:"last_offline_time" bson:"last_offline_time"`
	LastDisconnectTime time.Time `json:"last_disconnect_time" bson:"last_disconnect_time"`
	// LastRegisterTime time.Time `json:"last_register_time" bson:"last_register_time"`
	MDBDeepLinkMachineInfoST        `json:"deeplink_st" bson:"deeplink_st,omitempty"`
	MDBDeepLinkMachineInfoBandwidth `json:"deeplink_bw" bson:"deeplink_bw,omitempty"`
}

// time series table
type MDBMachineTM struct {
	Timestamp time.Time  `json:"timestamp" bson:"timestamp"`
	Machine   MachineKey `json:"machine" bson:"machine"`
	// MachineInfo `bson:",inline"`
	DeepLinkMachineInfoST        `json:"deeplink_st" bson:"deeplink_st,omitempty"`
	DeepLinkMachineInfoBandwidth `json:"deeplink_bw" bson:"deeplink_bw,omitempty"`
}

// machine info of deeplink short-term
type MDBDeepLinkMachineInfoST struct {
	DeepLinkMachineInfoST `bson:",inline"`
	CalcPoint             float64 `json:"calc_point" bson:"calc_point,omitempty"`
	Longitude             float32 `json:"longitude" bson:"longitude,omitempty"`
	Latitude              float32 `json:"latitude" bson:"latitude,omitempty"`
	Region                string  `json:"region" bson:"region,omitempty"`
}

// machine info of deeplink bandwidth
type MDBDeepLinkMachineInfoBandwidth struct {
	DeepLinkMachineInfoBandwidth `bson:",inline"`
	Region                       string `json:"region" bson:"region,omitempty"`
}
