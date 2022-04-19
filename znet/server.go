package znet

import (
	"errors"
	"fmt"
	"github.com/xtaci/kcp-go"
	"liereal.com/zinx-kcp/utils"
	"liereal.com/zinx-kcp/ziface"
	"net"
	"time"
)

type Server struct {
	Name       string
	Host       net.Addr
	msgHandler ziface.IMsgHandler  //当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	ConnMgr    ziface.IConnManager //当前Server的链接管理器
	//该Server的连接创建时Hook函数
	OnConnStart func(conn ziface.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn ziface.IConnection)
}

// SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
}

func CallBackToClient(conn net.Conn, data []byte, cnt int) error {
	//Send back
	cnt, err := conn.Write(data[:cnt])
	if err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at %s, is starting\n", s.Host.String())
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPacketSize: %d\n",
		utils.ConfigInstance.Version,
		utils.ConfigInstance.MaxConn,
		utils.ConfigInstance.MaxPacketSize)

	go func() {
		//监听服务器地址
		listen, err := kcp.Listen(s.Host.String())
		if err != nil {
			fmt.Println("listen:", s.Host.String(), "err", err)
			return
		}

		//已经监听成功
		fmt.Println("start Zinx-kcp server  ", s.Name, " succ, now listenning...")

		//TODO server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid = 0

		//启动server网络连接业务
		for {
			s.msgHandler.StartWorkerPool()

			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= utils.ConfigInstance.MaxConn {
				conn.Close()
				continue
			}

			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnection(s, conn, cid, s.msgHandler)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx-kcp server , name ", s.Name)
	// Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()

}

func (s Server) Serve() {
	s.Start()
	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加
	//阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer() ziface.IServer {
	s := &Server{
		Name: utils.ConfigInstance.Name,
		Host: &net.UDPAddr{
			IP:   net.ParseIP(utils.ConfigInstance.Host),
			Port: utils.ConfigInstance.Port,
		},
		msgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(), //创建ConnManager
	}
	return s
}
