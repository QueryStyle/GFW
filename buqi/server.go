package buqi

import (
	"encoding/binary"
	"net"
)

// Server 服务端
type Server struct {
	*Socket
}

// NewServer 新建一个服务端
func NewServer(password *Password, listenAddr *net.TCPAddr) *Server {
	return &Server{
		Socket: &Socket{
			Cipher:     NewCipher(password),
			ListenAddr: listenAddr,
		},
	}
}

// Listen 服务端监听
func (server *Server) Listen(didListen func(listenAddr net.Addr)) {
	listener, err := net.ListenTCP("tcp", server.ListenAddr)
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
				go server.handleConn(conn)
			}
		}
	}
}

// handleConn 处理 SOCKS5 连接
func (server *Server) handleConn(localConn *net.TCPConn) {
	defer localConn.Close()
	buf := make([]byte, 256)

	_, err := server.DecodeRead(localConn, buf)
	// VER 只支持SOCKS5
	if err != nil || buf[0] != 0x05 {
		return
	}

	// 不验证，直接通过
	server.EncodeWrite(localConn, []byte{0x05, 0x00})

	// 获取远程服务地址
	n, err := server.DecodeRead(localConn, buf)
	// n 最短的长度为7 情况为 ATYP=3 DST.ADDR占用1字节 值为0x0
	if err != nil || n < 7 {
		return
	}

	// CMD 客户端请求类型 CONNECT： 0x01；BIND： 0x02；UDP： ASSOCIATE 0x03
	if buf[1] != 0x01 {
		return
	}

	var dIP []byte
	// aType 服务器地址类型
	switch buf[3] {
	case 0x01:
		//	IPV4类型
		dIP = buf[4 : 4+net.IPv4len]
	case 0x03:
		//	域名类型
		ipAddr, err := net.ResolveIPAddr("ip", string(buf[5:n-2]))
		if err != nil {
			return
		}
		dIP = ipAddr.IP
	case 0x04:
		//	IPV6类型
		dIP = buf[4 : 4+net.IPv6len]
	default:
		return
	}

	dPort := buf[n-2:]
	dstAddr := &net.TCPAddr{
		IP:   dIP,
		Port: int(binary.BigEndian.Uint16(dPort)),
	}

	// 连接远程服务
	dstServer, err := net.DialTCP("tcp", nil, dstAddr)
	if err != nil {
		return
	}
	defer dstServer.Close()
	// conn 被断开则直接清空数据
	dstServer.SetLinger(0)

	// 响应客户端连接成功
	server.EncodeWrite(localConn, []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	// 将客户端的数据发送到服务端
	go func() {
		err := server.DecodeCopy(dstServer, localConn)
		if err != nil {
			// 一旦抛出错误就直接退出
			localConn.Close()
			dstServer.Close()
		}
	}()
	// 将服务端的数据发送到客户端
	server.EncodeCopy(localConn, dstServer)
}
