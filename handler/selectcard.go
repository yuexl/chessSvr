package handler

import (
	"encoding/json"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"

	"chessSvr/logic"
	"chessSvr/proto"
)

type SelectCardHandler struct {
}

func (s SelectCardHandler) PreHandle(request ziface.IRequest) {
	zlog.Debug(request)
	zlog.Debug(request.GetMsgID())
	zlog.Debug(string(request.GetData()))
}

func (s SelectCardHandler) Handle(request ziface.IRequest) {
	reqSelCard := proto.ReqSelCard{}
	err := json.Unmarshal(request.GetData(), &reqSelCard)
	if err != nil {
		zlog.Error(err)
		request.GetConnection().SendMsg(proto.MsgIdRspSelectCard, []byte("error req data"))
		return
	}
	zlog.Debug(reqSelCard.TechLevel)

	selectCards := logic.GetSelectCards(reqSelCard.TechLevel)
	zlog.Debug(selectCards)
	request.GetConnection().SendMsg(proto.MsgIdRspSelectCard, []byte(selectCards))
}

func (s SelectCardHandler) PostHandle(request ziface.IRequest) {
	zlog.Debug("SelectCardHandler post")
}
