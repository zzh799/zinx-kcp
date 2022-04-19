package znet

import (
	"google.golang.org/protobuf/proto"
	"liereal.com/zinx-kcp/ziface"
)

type BaseRouter struct {
}

func (r *BaseRouter) PreHandle(req ziface.IRequest, message proto.Message) {

}

func (r *BaseRouter) Handle(req ziface.IRequest, message proto.Message) {

}

func (r *BaseRouter) PostHandle(req ziface.IRequest, message proto.Message) {

}
