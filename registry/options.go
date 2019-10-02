/**
 * @Author: DollarKiller
 * @Description: 注册中心的设置
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:28 2019-09-23
 */
package registry

import "time"

/**
注册多注册中心
支持可扩展
基于接口 && 插件
*/

type Options struct {
	Address      []string      // 注册地址
	Timeout      time.Duration // 超时设置
	RegistryPath string        // 注册路径  用于层级遍历
	HeartBeat    int64         // 心跳时间
}

type Option func(*Options)

func WithTimeout(timeout time.Duration) Option {
	return func(options *Options) {
		options.Timeout = timeout
	}
}

func WithAddrs(addrs []string) Option {
	return func(options *Options) {
		options.Address = addrs
	}
}
