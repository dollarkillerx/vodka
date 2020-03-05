/**
 * @Author: DollarKillerX
 * @Description: main
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:26 2020/1/6
 */
package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	opt := Option{}

	app := cli.NewApp()

	app.Name = "Vodka MicroServices Framework Cli"
	app.Author = "DollarKiller"
	app.Email = "dollarkiller@dollarkiller.com"
	app.Copyright = "https://github.com/dollarkillerx"
	app.Version = "v0.0.1"

	app.Flags = []cli.Flag{
		// 输入idl 路径
		cli.StringFlag{
			Name:        "f",
			Usage:       "idl file path",
			Required:    true,
			Value:       "vodka.proto",
			Destination: &opt.ProtoFileName,
		},
		// 输出代码路径
		cli.StringFlag{
			Name:        "o",
			Usage:       "output directory",
			Required:    false,
			Value:       "./output/",
			Destination: &opt.Output,
		},
		// 生成客户端代码
		cli.BoolFlag{
			Name:        "c",
			Usage:       "generate grpc client code",
			Required:    false,
			Destination: &opt.GenClientCode,
		},
		// 生成服务端代码
		cli.BoolFlag{
			Name:        "s",
			Usage:       "generate grpc server code",
			Required:    false,
			Destination: &opt.GenServerCode,
		},
		// go mod
		cli.StringFlag{
			Name:        "m",
			Usage:       "go mod name",
			Required:    true,
			Value:       "vodka",
			Destination: &opt.GoMod,
		},
	}

	// 执行逻辑入口
	app.Action = func(ctx *cli.Context) error {
		err := genMgr.Run(&opt)
		if err != nil {
			return err
		}
		return nil
	}

	// 容错
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

//  controller: 存在服务的方法实现
//  idl: 存放服务的idl定义
//  main: 存放服务的入口代码
//  scripts: 存放服务的脚本
//  conf: 存放服务的配置文件
//  app/router: 存放服务的路由
//  app/config: 存放服务的一些配置
//  datamodels: 存放服务的实体代码
//  generate: grpc生成的代码
