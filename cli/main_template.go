/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 09:30
 */
package main

var MainTemplate = `
/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package main

import (
	"log"
	"{{.Package}}/core/router"
	"{{.Package}}/generate"
	middleware2 "{{.Package}}/middleware"
	router2 "{{.Package}}/router"
	
	"github.com/dollarkillerx/vodka"
	"github.com/dollarkillerx/vodka/middleware"
	"github.com/dollarkillerx/vodka/server"
)

func main() {
	v := vodka.New()
	router.ServerAddr = server.Config.Addr
	app := router.New()
	app.Use(middleware2.BasePrometheus)  // 注册全局中间件  基础Prometheus
	router2.Registry(app)
	{{.Pkg}}.Register{{.Server}}Server(v.RegisterServer(), app.RegistryGRPC())

	if server.Config.Prometheus.SwitchOn {
		go middleware.Prometheus.Run(server.Config.Prometheus.Addr)
	}
	log.Println(v.Run(router.ServerAddr))
}
`

var ConfigTemplate = `
ServiceName: "VodkaExample"
Addr: "0.0.0.0:8081"
Log:
  Level: "debug"
  Dir: "log"
  ConsoleLog: true
Prometheus:
  SwitchOn: true
  Addr: "0.0.0.0:8082"
`


//var MainTemplate = `
///**
//*@Program: vodka
//*@MicroServices Framework: https://github.com/dollarkillerx
// */
//package main
//
//import (
//	"google.golang.org/grpc"
//	"{{.GoMod}}/controller"
//	"log"
//	"net"
//)
//
//func main() {
//	server := grpc.NewServer()
//	pb.RegisterServiceServer(server, &controller.{{.ServerName}}{})
//	dial, err := net.Listen("tcp", "0.0.0.0:8080")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	server.Serve(dial)
//}
//`
