package main

import (
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"

	"chessSvr/router"
)

func main() {
	zlog.Info("svr start ...")

	iServer := znet.NewServer()

	router.SetupRouter(iServer)

	iServer.Serve()

}

//tcpAddr, err := net.ResolveTCPAddr("tcp", conf.GConfig.Server.Host+":"+conf.GConfig.Server.Port)
//if err != nil {
//	logrus.Errorln(err)
//}
//listener, err := net.ListenTCP("tcp", tcpAddr)
//defer listener.Close()
//if err != nil {
//	logrus.Errorln(err)
//}
//
//for {
//	conn, err := listener.Accept()
//	if err != nil {
//		logrus.Errorln(err)
//		continue
//	}
//
//	go func(conn net.Conn) {
//		//dosomething
//
//		defer func() {
//			if err := recover(); err != nil {
//				logrus.Errorln(err)
//			}
//		}()
//	}(conn)
//
//	var reqBytes []byte
//	_, err = conn.Read(reqBytes)
//	if err != nil {
//		logrus.Errorln(err)
//	}
//	logrus.Infoln(string(reqBytes))
//}
