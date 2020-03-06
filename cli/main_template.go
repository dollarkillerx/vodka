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
	"%s/generate"
	"%s/router"
	"log"
	"net"
	
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	pb.RegisterServiceServer(server, &router.%sRouter{})
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
