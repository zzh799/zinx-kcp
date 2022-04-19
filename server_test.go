package main

import (
	"fmt"
	"github.com/xtaci/kcp-go"
	"google.golang.org/protobuf/proto"
	"liereal.com/zinx-kcp/pb"
	"liereal.com/zinx-kcp/utils"
	"liereal.com/zinx-kcp/ziface"
	"liereal.com/zinx-kcp/znet"
	"net"
	"testing"
	"time"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest, message proto.Message) {
	fmt.Println("Call PingRouter Handle")
	userLoginRequest := message.(*pb.UserLoginRequest)
	fmt.Println("UserLogin-> userName:", userLoginRequest.UserName, "password:", userLoginRequest.Password)

	err := request.GetConnection().SendMsg(
		&pb.Message{
			Response: &pb.Response{
				UserLogin: &pb.UserLoginResponse{
					Result: &pb.Result{
						Error:   "",
						Success: true,
					},
				},
			},
		},
	)

	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func ClientTest() {
	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)
	addr := net.UDPAddr{
		IP:   net.ParseIP(utils.ConfigInstance.Host),
		Port: utils.ConfigInstance.Port,
	}

	conn, err := kcp.Dial(addr.String())
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		//发封包message消息

		dp := &pb.Message{Request: &pb.Request{
			UserLogin: &pb.UserLoginRequest{
				UserName: "aaa",
				Password: "aaa",
			},
		}}
		msg, err := proto.Marshal(dp)
		if err != nil {
			continue
		}
		_, err = conn.Write(msg)
		if err != nil {
			continue
		}
		data := make([]byte, 512)
		n, err := conn.Read(data)
		if n > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := &pb.Message{}

			err := proto.Unmarshal(data[:n], msg)
			if err != nil {
				continue
			}

			fmt.Println("==> Recv UserLogin resut:", msg.Response.UserLogin.Result.Success)
		}

		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	/*
		Test Server
	*/
	utils.ConfigInstance.Reload()
	s := znet.NewServer()
	s.AddRouter("UserLoginRequest", &PingRouter{})
	go ClientTest()
	s.Serve()
	/*
		Test Client
	*/

}
