/**
 * @Author: DollarKiller
 * @Description: 服务存储源信息
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:39 2019-09-23
 */
package registry

// 服务抽象
type Service struct {
	Name  string  `json:"name"`
	Nodes []*Node `json:"nodes"`
}

// 服务节点抽象
type Node struct {
	Id     string `json:"id"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"` // 权重
}
