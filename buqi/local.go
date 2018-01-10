package buqi

import (
	"log"
	"net"
)

// Local 本地端
type Local struct {
	*Socket
}

// NewLocal 新建一个本地端
func NewLocal(password *Password, listenAddr, remoteAddr *net.TCPAddr) *Local {
	return &Local{
		Socket: &Socket{
			Cipher:     NewCipher(password),
			ListenAddr: listenAddr,
			RemoteAddr: remoteAddr,
		},
	}
}

// Listen 本地端启动监听，接收来自本机浏览器的连接
func (local *Local) Listen(didListen func(listenAddr net.Addr)) {
	listener, err := net.ListenTCP("tcp", local.ListenAddr)
	defer listener.Close()
	if err == nil {
		if didListen != nil {
			didListen(listener.Addr())
		}

		for {
			conn, err := listener.AcceptTCP()
			if err == nil {
				// conn 被断开则直接清空数据
				conn.SetLinger(0)
				go local.handleConn(conn)
			}
		}
	}
}

// handleConn 处理连接
func (local *Local) handleConn(user *net.TCPConn) {
	defer user.Close()
	conn, err := local.DialRemote()
	if err == nil {
		defer conn.Close()
		// conn 被断开则直接清空数据
		conn.SetLinger(0)
		// 读取服务器的数据发送到本地
		go func() {
			local.DecodeCopy(user, conn)
		}()
		// 将客户端的数据发送到服务端
		local.EncodeCopy(conn, user)
	} else {
		log.Fatal(err)
	}
}
