package znet

import (
	"google.golang.org/protobuf/proto"
	"liereal.com/zinx-kcp/ziface"
)

type Request struct {
	conn ziface.IConnection //已经和客户端建立好的 链接
	msg  proto.Message      //客户端请求的数据
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() proto.Message {
	return r.msg
}
