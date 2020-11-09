package handler

import (
	"chessSvr/module"
	"chessSvr/proto"

	"encoding/json"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
)

type UseLoginHandler struct {
}

func (u UseLoginHandler) PreHandle(request ziface.IRequest) {
	zlog.Debugf("pre user login, request: %v", request)
	zlog.Info(request.GetConnection().RemoteAddr())
	zlog.Info(request.GetConnection().GetConnID())
	zlog.Info(request.GetMsgID())
	zlog.Info(string(request.GetData()))
}

func (u UseLoginHandler) Handle(request ziface.IRequest) {
	zlog.Debugf("handling user login, request: %v", request)
	zlog.Info(string(request.GetData()))
	var loginUser proto.LoginUser
	err := json.Unmarshal(request.GetData(), &loginUser)
	if err != nil {
		zlog.Error(err)
		request.GetConnection().SendMsg(proto.MsgIdRspLogin, []byte("wrong msg data"))
		return
	}
	if err = module.Login(loginUser.UserId, loginUser.PassWord); err != nil {
		zlog.Error(err)
		request.GetConnection().SendMsg(proto.MsgIdRspLogin, []byte("login failed"))
		return
	}
	request.GetConnection().SendMsg(proto.MsgIdRspLogin, []byte("login success"))
}

func (u UseLoginHandler) PostHandle(request ziface.IRequest) {

}
