package types

type OfflineNotifyRequest struct {
	MachineId string `json:"machine_id"`
	IsOnline  bool   `json:"is_online"`
}

type OfflineNotifyResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Success  bool   `json:"success"`
	IsRented bool   `json:"is_rented"`
}
