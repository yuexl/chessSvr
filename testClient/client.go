package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"

	"chessSvr/proto"
	"chessSvr/utils"

	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

/*
	模拟客户端
*/
func main() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	UserLogin(conn)
	time.Sleep(1 * time.Second)
	GetCard(conn)
}

func SendMsg(conn net.Conn) {
	dp := znet.NewDataPack()
	msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx Client Test Message")))
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}

	//先读出流中的head部分
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	if err != nil {
		fmt.Println("read head error")
	}
	//将headData字节流 拆包到msg中
	msgHead, err := dp.Unpack(headData)
	if err != nil {
		fmt.Println("server unpack err:", err)
		return
	}

	if msgHead.GetDataLen() > 0 {
		//msg 是有data数据的，需要再次读取data数据
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		//根据dataLen从io中读取字节流
		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			fmt.Println("server unpack data err:", err)
			return
		}

		fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
	}
}

func UserLogin(conn net.Conn) {
	chessUser := proto.LoginUser{
		UserId:   1,
		PassWord: utils.MD5("123"),
	}
	fmt.Println(chessUser.PassWord)
	dp := znet.NewDataPack()
	sendlogin, err := json.Marshal(&chessUser)
	if err != nil {
		zlog.Error(err)
		return
	}
	msg, _ := dp.Pack(znet.NewMsgPackage(proto.MsgIdReqLogin, sendlogin))
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}

	//先读出流中的head部分
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	if err != nil {
		fmt.Println("read head error")
	}
	//将headData字节流 拆包到msg中
	msgHead, err := dp.Unpack(headData)
	if err != nil {
		fmt.Println("server unpack err:", err)
		return
	}

	if msgHead.GetDataLen() > 0 {
		//msg 是有data数据的，需要再次读取data数据
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		//根据dataLen从io中读取字节流
		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			fmt.Println("server unpack data err:", err)
			return
		}

		fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
	}
}

func GetCard(conn net.Conn) {
	req := proto.ReqSelCard{
		TechLevel: 8,
	}
	dp := znet.NewDataPack()
	senddata, err := json.Marshal(&req)
	if err != nil {
		zlog.Error(err)
		return
	}

	msg, _ := dp.Pack(znet.NewMsgPackage(proto.MsgIdReqSelectCard, senddata))
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}

	//先读出流中的head部分
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	if err != nil {
		fmt.Println("read head error")
	}
	//将headData字节流 拆包到msg中
	msgHead, err := dp.Unpack(headData)
	if err != nil {
		fmt.Println("server unpack err:", err)
		return
	}

	if msgHead.GetDataLen() > 0 {
		//msg 是有data数据的，需要再次读取data数据
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		//根据dataLen从io中读取字节流
		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			fmt.Println("server unpack data err:", err)
			return
		}

		fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
	}
}
