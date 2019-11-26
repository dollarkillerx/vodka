/**
 * @Author: DollarKiller
 * @Description: vodka cli tools
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:06 2019-10-04
 */
package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

/**
命令行参数
-f 指定idl文件
-o 指定代码生成的路径
-c 指定生成客户端调用代码
-s 指定生成服务端框架代码
*/

func main() {
	var opt Option

	app := cli.NewApp()

	app.Name = "Lightweight golang microservices framework Cli"
	app.Author = "dollarkiller"
	app.Email = "dollarkiller@dollarkiller.com"
	app.Copyright = "https://github.com/dollarkillerx"
	app.Version = "v0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{ // 输出idl
			Name:        "f",
			Value:       "./vodka.proto", // 默认值
			Usage:       "idl filename",  // 注释
			Destination: &opt.Proto3Filename,
		},
		cli.StringFlag{ // 输出代码路径
			Name:        "o",
			Value:       "./output/",
			Usage:       "output directory",
			Destination: &opt.Output,
		},
		cli.BoolFlag{
			Name:        "c",
			Usage:       "generate grpc cli code",
			Destination: &opt.GenClientCode,
		},
		cli.BoolFlag{
			Name:        "s",
			Usage:       "generate grpc server code",
			Destination: &opt.GenServerCode,
		},
	}

	app.Action = func(c *cli.Context) error {
		err := genMgr.Run(&opt)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
