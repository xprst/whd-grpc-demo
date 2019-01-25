package main

import (
	"flag"
	"fmt"
	"github.com/xprst/whd-grpc-base/app"
	"github.com/xprst/whd-grpc-base/server"
	pb "whd-grpc-demo/proto"
	"whd-grpc-demo/service"
)

func main() {
	// 从参数中读取配置文件地址
	configPath := flag.String("conf","./config/app.json", "config file path")
	flag.Parse()

	// 使用配置文件初始化a程序
	s := server.NewServer(server.WithConfig(*configPath))

	// 从配置文件获取配置，使用app提供的全局对象
	hello := app.Config.GetString("userDefineConfig")
	fmt.Println("hello " + hello )

	// 注册自己实现的grpc服务
	pb.RegisterUserServiceServer(s.GrpcServer(), service.NewUserService())
	// 启动grpc服务
	err := s.StartServer()
	if err != nil {
		panic(err)
	}
}
