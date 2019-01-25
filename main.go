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
	// 从参数中读取配置文件
	configPath := flag.String("conf","./config/app.json", "config file path")
	flag.Parse()

	// 使用配置文件初始化app
	err := app.InitWithConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("InitWithConfig failed ： %v",err))
	}

	// 从配置文件获取配置
	hello := app.Config.GetString("userDefineConfig")
	fmt.Println("hello " + hello )
	// 启动grpc服务
	s := server.NewServer(server.WithPort(app.GrpcPort))
	pb.RegisterUserServiceServer(s.GrpcServer(), service.NewUserService())
	_ = s.StartServer()
}
