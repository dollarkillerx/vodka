/**
 * @Author: DollarKillerX
 * @Description: def.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:23 2019/12/31
 */
package main

type Options struct {
	host  string
	ip    string
	debug bool
}

type Option func(opts *Options)

func (o *Options) InitOptions(ops ...Option) {
	for _, v := range ops {
		v(o)
	}
}

func SetIp(ip string) Option {
	return func(opts *Options) {
		opts.ip = ip
	}
}

func SetHost(host string) Option {
	return func(opts *Options) {
		opts.host = host
	}
}

func SetDebug(debug bool) Option {
	return func(opts *Options) {
		opts.debug = debug
	}
}

func main() {
	options := Options{}
	options.InitOptions(SetIp("0.0.0.0"), SetHost("dollarkiller.com"), SetDebug(false))
}

