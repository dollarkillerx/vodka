/**
 * @Author: DollarKillerX
 * @Description: test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:54 2020/1/7
 */
package main

import (
	test "github.com/dollarkillerx/vodka/cli/output/generate"
	"google.golang.org/grpc"
)

func getClient() (test.ServiceClient, error) {
	conn, e := grpc.Dial("0.0.0.0:8083", grpc.WithInsecure())
	if e != nil {
		return nil, e
	}
	client := test.NewServiceClient(conn)
	return client, nil
}
