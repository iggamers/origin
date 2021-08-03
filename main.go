package main

import (
	"log"

	"github.com/asim/go-micro/v3/web"
	"github.com/i9code/xcho"

	"origin/api"
)

func main() {

	x, _ := xcho.Start(xcho.Routes(func(group *xcho.Group) {
		api.Mount(group.Group("/api"))
	}))
	// 创建 service
	service := web.NewService(
		web.Name("abc"),      //服务名称
		web.Address(":8071"), //端口
		web.Version("0.0.1"), //版本
		//web.RegisterTTL(time.Second*30),
		//web.RegisterInterval(time.Second*15),
		web.Handler(x),
	)
	//log.Infof("ops=%v", service.Options())
	// 初始化 service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	//
	//运行
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

//
//func main() {
//
//	srv := httpServer.NewServer(
//		server.Name(SERVER_NAME),
//		server.Address(":8080"),
//	)
//
//	 engine,_ := xcho.Start(xcho.Routes(func(group *xcho.Group) {
//		 api.Mount(group.Group("/api"))
//	 }))
//
//
//	hd := srv.NewHandler(engine)
//	if err := srv.Handle(hd); err != nil {
//		log.Fatalln(err)
//	}
//
//	service := micro.NewService(
//		micro.Server(srv),
//		micro.Registry(registry.NewRegistry()),
//	)
//	service.Init()
//	service.Run()
//}
