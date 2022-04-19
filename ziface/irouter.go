package ziface

import "google.golang.org/protobuf/proto"

type IRouter interface {
	PreHandle(request IRequest, message proto.Message)  //在处理conn业务之前的钩子方法
	Handle(request IRequest, message proto.Message)     //处理conn业务的方法
	PostHandle(request IRequest, message proto.Message) //处理conn业务之后的钩子方法
}
