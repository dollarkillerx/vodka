/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 16:41
 */
package main

import (
	"context"
	"fmt"
	"github.com/dollarkillerx/vodka/cli/test/grpc_test/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	dial, err := grpc.Dial("0.0.0.0:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer dial.Close()
	client := pb.NewServiceClient(dial)
	go Run1Client(client)
	go Run2Client(client)
	Run3Client(client)
}

func Run1Client(ser pb.ServiceClient) {
	resp, err := ser.Run1(context.TODO(), &pb.Req{Msg: "Hello"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}

func Run2Client(ser pb.ServiceClient) {

}

func Run3Client(ser pb.ServiceClient) {
	//ser.Run3(context.TODO(),)
}
