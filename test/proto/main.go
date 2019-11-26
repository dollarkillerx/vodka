/**
 * @Author: DollarKiller
 * @Description: proto read file test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:39 2019-10-04
 */
package main

import (
	"log"
	"os"

	"github.com/emicklei/proto"
)

func main() {
	file, e := os.Open("tools/vodka/h1.proto")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	parser := proto.NewParser(file)
	parse, e := parser.Parse()
	if e != nil {
		log.Fatal(e)
	}
	proto.Walk(parse,
		proto.WithOption(func(option *proto.Option) {
			log.Println(option)
		}),
		proto.WithEnum(func(enum *proto.Enum) {
			log.Println(enum)
		}),
		proto.WithOneof(func(oneof *proto.Oneof) {
			log.Println(oneof)
		}),
		proto.WithMessage(func(message *proto.Message) {
			//log.Println(message.Name)
			//log.Println(message.Comment)
			//log.Println(message.Elements)
			//log.Println(message.IsExtend)
			//log.Println(message.Parent)
			//log.Println(message.Position)
			//log.Println("===============")
		}),
		proto.WithService(func(service *proto.Service) {
			//log.Println(service)
		}),
		proto.WithRPC(func(rpc *proto.RPC) {
			log.Println(rpc.Name)
			log.Println(rpc.RequestType)
			log.Println(rpc.ReturnsType)

			log.Println("=================")
		}),
	)
}
