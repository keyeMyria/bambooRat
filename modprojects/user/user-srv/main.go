package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/xuyiwenak/bambooRat/modprojects/user/user-srv/base"
	"github.com/xuyiwenak/bambooRat/modprojects/user/user-srv/base/config"
	"github.com/xuyiwenak/bambooRat/modprojects/user/user-srv/handler"
	"github.com/xuyiwenak/bambooRat/modprojects/user/user-srv/model"
	proto "github.com/xuyiwenak/bambooRat/modprojects/user/user-srv/proto/user"
	"time"
)

func main() {

	// 初始化配置、数据库等信息
	base.Init()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("bambooRat.micro.sdk.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	proto.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
