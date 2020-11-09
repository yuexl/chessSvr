package proto

const (
	MsgIdReqLogin      = 0
	MsgIdRspLogin      = 1
	MsgIdReqSelectCard = 2
	MsgIdRspSelectCard = 3
)

type LoginUser struct {
	UserId   int    `json:"user_id"`
	PassWord string `json:"pass_word"`
}

type ReqSelCard struct {
	TechLevel int8 `json:"tech_level"`
}

type ResSelCard struct {
	TechLevel int8 `json:"tech_level"`
}
