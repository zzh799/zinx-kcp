package ziface

import "google.golang.org/protobuf/reflect/protoreflect"

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgId protoreflect.Name, router IRouter)
	GetConnMgr() IConnManager         //得到链接管理
	SetOnConnStart(func(IConnection)) //设置该Server的连接创建时Hook函数
	SetOnConnStop(func(IConnection))  //设置该Server的连接断开时的Hook函数
	CallOnConnStart(conn IConnection) //调用连接OnConnStart Hook函数
	CallOnConnStop(conn IConnection)  //调用连接OnConnStop Hook函数
}
