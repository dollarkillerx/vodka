/**
 * @Author: DollarKiller
 * @Description: test grpc cli
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:05 2019-10-03
 */
package main

import (
	"context"
	"github.com/dollarkillerx/vodka/test/protobuf/demo2/hello"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, e := grpc.Dial("127.0.0.1:50050", grpc.WithInsecure())
	if e != nil {
		panic(e)
	}
	if conn != nil {
		defer conn.Close()
	}

	client := hello.NewHelloServiceClient(conn)

	resp, e := client.SayHello(context.TODO(), &hello.HelloReq{
		Name: "this is 爸爸",
	})

	if e != nil {
		// 如果连接发送错误
		panic(e)
	}

	log.Println(resp)

}
