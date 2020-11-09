package router

import (
	"github.com/aceld/zinx/ziface"

	"chessSvr/handler"
	"chessSvr/proto"
)

func SetupRouter(svr ziface.IServer) {
	svr.AddRouter(proto.MsgIdReqLogin, &handler.UseLoginHandler{})
	svr.AddRouter(proto.MsgIdReqSelectCard, &handler.SelectCardHandler{})
}
