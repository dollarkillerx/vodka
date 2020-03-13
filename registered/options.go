/**
 * @Author: DollarKillerX
 * @Description: options 注册中心选项
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:26 2020/1/3
 */
package registered

type Options struct {
	Addr         []string `json:"addr"`          // 注册中心 持久化节点的地址
	RegistryPath string   `json:"registry_path"` // 注册根路径

	Timeout   int         `json:"timeout"`    // 超时时间
	HeartBeat int         `json:"heart_beat"` // 服务心跳
	Debug     bool        `json:"debug"`      // 是否是debug
	Config    interface{} `json:"config"`     // 一些自定义配置
}

type SetOption func(options *Options)

func WithAddr(addr []string) SetOption {
	return func(options *Options) {
		options.Addr = addr
	}
}

func WithRegistryPath(path string) SetOption {
	return func(options *Options) {
		options.RegistryPath = path
	}
}

func WithTimeout(timeout int) SetOption {
	return func(options *Options) {
		options.Timeout = timeout
	}
}

func WithHeartBeat(timeout int) SetOption {
	return func(options *Options) {
		options.HeartBeat = timeout
	}
}

func WithDebug(debug bool) SetOption {
	return func(options *Options) {
		options.Debug = debug
	}
}

func WithConfig(config interface{}) SetOption {
	return func(options *Options) {
		options.Config = config
	}
}
