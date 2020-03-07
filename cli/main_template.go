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
	"net"
	"%s/core/router"
	"%s/generate"
	router2 "%s/router"
	
	"google.golang.org/grpc"
)

func init() {
	log.Println("Vodka is initialized")
}

func main() {
	server := grpc.NewServer()
	router := router.New()
	router2.Registry(router)  
	pb.RegisterServiceServer(server, router.RegistryGRPC())
	dial, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve(dial)
}
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
