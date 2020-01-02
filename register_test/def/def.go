/**
 * @Author: DollarKillerX
 * @Description: def.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:33 2019/12/31
 */
package def

import "time"

type Option struct {
	Hosts   []string
	Timeout time.Duration
}

type setOption func(option *Option)

func SetHosts(host ...string) setOption {
	return func(option *Option) {
		option.Hosts = host
	}
}

func SetTimeout(time time.Duration) setOption {
	return func(option *Option) {
		option.Timeout = time
	}
}

func (o *Option) Init(seti ...setOption) {
	for _, v := range seti {
		v(o)
	}
}
