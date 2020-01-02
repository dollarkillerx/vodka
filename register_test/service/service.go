/**
 * @Author: DollarKillerX
 * @Description: service.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:34 2019/12/31
 */
package service

type Service struct {
	Name  string
	Nodes []*Node
}

type Node struct {
	Id   string
	Host string

	Load int // 负载
	Data interface{}
}
