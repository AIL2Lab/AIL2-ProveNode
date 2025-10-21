package types

type WsHeader struct {
	Version   uint32 `json:"version"`   // 协议版本，暂时用 0
	Timestamp int64  `json:"timestamp"` // 时间戳
	Id        uint64 `json:"id"`        // 消息 ID
	Type      uint32 `json:"type"`      // 消息类型 WsMessageType
	PubKey    []byte `json:"pub_key"`   // 公钥，验证消息安全完整，暂时不需要
	Sign      []byte `json:"sign"`      // 签名，验证消息安全完整，暂时不需要
}

type WsRequest struct {
	WsHeader
	Body []byte `json:"body"` // 消息体，由 WsOnlineRequest 或者 WsMachineInfoRequest 编码后的字节流
}

type WsResponse struct {
	WsHeader
	Code    uint32 `json:"code"`    // 返回的错误码
	Message string `json:"message"` // 返回的错误描述
	Body    []byte `json:"body"`    // 消息体，暂时用不到
}

type WsMessageType uint32

const (
	WsMtOnline WsMessageType = iota + 1
	WsMtDeepLinkMachineInfoST
	WsMtNotify
	WsMtDeepLinkMachineInfoBW
)

type WsConnInfo struct {
	MachineKey
	StakingType StakingType
	ClientIP    string
}

type WsOnlineRequest struct {
	MachineKey
	StakingType StakingType `json:"staking_type"`
}

// type ModelInfo struct {
// 	Model string `json:"model" bson:"model"`
// }

// type WsMachineInfoRequest MachineInfo

type WsNotifyMessage struct {
	Unregister WsUnregisterNotify `json:"unregister,omitempty"`
}

type WsUnregisterNotify struct {
	Message string `json:"message"`
}
