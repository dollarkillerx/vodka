/**
 * @Author: DollarKillerX
 * @Description: service 服务定义
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:26 2020/1/3
 */
package registered

// 服务
type Service struct {
	Name string  `json:"name"` // 服务名称
	Node []*Node `json:"node"` // 服务名下的节点
}

// 节点
type Node struct {
	Name string `json:"name"` // 服务的名称  这里有点冗余
	ID   string `json:"id"`   // 服务的IP
	Addr string `json:"addr"` // 服务的地址

	Weight int         `json:"weight"` // 服务的权重
	Load   int         `json:"load"`   // 服务的负载
	Data   interface{} `json:"data"`   // 服务的一些自定义数据
}
