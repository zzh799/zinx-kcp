package ziface

import (
	"liereal.com/zinx-kcp/pb"
)

type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetData() pb.Message
}
