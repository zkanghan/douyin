package main

import (
	"douyin/cmd/user/infra/db"
	"douyin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/spf13/viper"
	"log"
	"net"
)

func Init() {
	// init config file
	viper.SetConfigFile("cmd/user/conf.yaml")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic("Read ConfigFile failed: " + err.Error())
	}
	// init db
	db.Init()
}

func main() {
	Init()

	addr, err := net.ResolveTCPAddr("tcp", "localhost:8889")
	if err != nil {
		panic(err)
	}
	srv := userservice.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
	)

	err = srv.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
