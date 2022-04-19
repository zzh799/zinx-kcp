package ziface

import (
	"google.golang.org/protobuf/proto"
	"net"
)

type IConnection interface {
	Start()            //启动连接，让当前连接开始工作
	Stop()             //停止连接，结束当前连接状态M
	GetConnID() uint32 //获取当前连接ID
	RemoteAddr() net.Addr
	GetConnection() net.Conn                     //从当前连接获取原始的socket conn
	SendMsg(data proto.Message) error            //将Message数据发送数据给远程的客户端
	SendBuffMsg(data proto.Message) error        //将Message数据发送给远程的TCP客户端(有缓冲)
	SetProperty(key string, value interface{})   //设置链接属性
	GetProperty(key string) (interface{}, error) //获取链接属性
	RemoveProperty(key string)                   //移除链接属性

}

type HandleFunc func(net.Conn, []byte, int) error
