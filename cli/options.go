/**
 * @Author: DollarKillerX
 * @Description: options 命令行参数
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:42 2020/1/6
 */
package main

type Option struct {
	ProtoFileName string // protoFile目录
	Output        string // 输出目录
	GenClientCode bool   // 生成client
	GenServerCode bool   // 生成server
	Prefix        string
}
