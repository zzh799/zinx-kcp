package znet

import (
	"liereal.com/zinx-kcp/pb"
	"liereal.com/zinx-kcp/ziface"
)

func (m *MsgHandle) DispatchRequset(sender ziface.IRequest, request *pb.Request) {
	if request.UserRegister != nil {
		m.DoMsgHandler(sender, request.UserRegister)
	}

	if request.UserLogin != nil {
		m.DoMsgHandler(sender, request.UserLogin)
	}
}

func (m *MsgHandle) DispatchResponse(sender ziface.IRequest, request *pb.Response) {
	if request.UserRegister != nil {
		m.DoMsgHandler(sender, request.UserRegister)
	}

	if request.UserLogin != nil {
		m.DoMsgHandler(sender, request.UserLogin)
	}
}
