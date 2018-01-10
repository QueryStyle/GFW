package main

import (
	"https://github.com/QueryStyle/GFW/buqi"
	"log"
	"net"
)

func main() {
	// 生成密码
	config := buqi.Start()
	// 解析密码
	password, err := buqi.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	switch config.Current {
	case "Local":
		// 本地端
		local, err := net.ResolveTCPAddr("tcp", config.Local)
		if err != nil {
			log.Fatalln(err)
		}
		// 服务端
		server, err := net.ResolveTCPAddr("tcp", config.Server)
		if err != nil {
			log.Fatalln(err)
		}

		// 监听
		listen := buqi.NewLocal(password, local, server)
		listen.Listen(func(local net.Addr) {
			log.Println(config.Password)
		})
	case "Server":
		server, err := net.ResolveTCPAddr("tcp", config.Server)
		if err != nil {
			log.Fatalln(err)
		}
		// 监听
		listen := buqi.NewServer(password, server)
		listen.Listen(func(server net.Addr) {
			log.Println("端口：" + config.Server)
			log.Println("密码：" + config.Password)
		})
	}
}
