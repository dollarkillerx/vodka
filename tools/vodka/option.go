/**
 * @Author: DollarKiller
 * @Description: Option命令行参数
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:11 2019-10-04
 */
package main

type Option struct {
	Proto3Filename string
	Output         string
	GenClientCode  bool
	GenServerCode  bool
	Prefix         string
}

