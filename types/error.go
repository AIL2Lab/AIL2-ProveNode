package types

type ErrorCode uint32

const (
	ErrCodeParam       ErrorCode = iota + 1 // 参数错误
	ErrCodeParse                            // 解析错误
	ErrCodeSign                             // 签名错误
	ErrCodeDatabase                         // 数据库错误
	ErrCodeOnline                           // 上线错误
	ErrCodeMachineInfo                      // 更新机器信息错误
	ErrCodeDbcChain                         // DBC 链错误
	ErrCodeIp2Location                      // ip2location 错误
	ErrCodeUUID
)

var errMsg = map[ErrorCode]string{
	ErrCodeParam:       "parameter error",
	ErrCodeParse:       "parsing error",
	ErrCodeSign:        "sign error",
	ErrCodeDatabase:    "database error",
	ErrCodeOnline:      "online error",
	ErrCodeMachineInfo: "invalid machine info",
	ErrCodeDbcChain:    "dbc chain error",
	ErrCodeIp2Location: "ip2location error",
	ErrCodeUUID:        "UUID error",
}

func (e ErrorCode) String() string {
	return errMsg[e]
}
