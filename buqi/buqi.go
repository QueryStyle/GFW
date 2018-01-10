package buqi

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

// BufSize 缓冲区大小
const BufSize = 1024

// Socket 用于传输的 TCP Socket
type Socket struct {
	Cipher     *Cipher
	ListenAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

func init() {
	// 更新随机种子
	rand.Seed(time.Now().Unix())
}

// Start 启动 
func Start() (config *Config) {
	config = &Config{}
	if !config.ReadConfig() {
		pwd := RandPassword().String()
		config = &Config{
			Local:    "127.0.0.1:12345",
			Server:   ":59386",
			Current:  "Server",
			Password: pwd,
		}
		config.SaveConfig()
		log.Fatal("已为您创建配置文件模板，请检查！")
	}
	return
}

// DecodeRead 读取并解码
func (s *Socket) DecodeRead(conn *net.TCPConn, bs []byte) (n int, err error) {
	n, err = conn.Read(bs)
	if err != nil {
		return
	}
	s.Cipher.decrypt(bs[:n])
	return
}

// EncodeWrite 编码并输出
func (s *Socket) EncodeWrite(conn *net.TCPConn, bs []byte) (int, error) {
	s.Cipher.encrypt(bs)
	return conn.Write(bs)
}

// EncodeCopy 读取并加密src的数据再写入到dst
func (s *Socket) EncodeCopy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BufSize)
	for {
		readCount, errRead := src.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
			return nil
		}
		if readCount > 0 {
			writeCount, errWrite := s.EncodeWrite(dst, buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// DecodeCopy 读取并解码src的数据再写入到dst
func (s *Socket) DecodeCopy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BufSize)
	for {
		readCount, errRead := s.DecodeRead(src, buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
			return nil
		}
		if readCount > 0 {
			writeCount, errWrite := dst.Write(buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// DialRemote 连接远程服务器
func (s *Socket) DialRemote() (*net.TCPConn, error) {
	conn, err := net.DialTCP("tcp", nil, s.RemoteAddr)
	if err != nil {
		return nil, errors.New("警告：连接远程服务器失败")
	}
	return conn, nil
}
