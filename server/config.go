/**
*@program: vodka
*@description: 配置文件解析
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-13 20:31
 */
package server

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	log2 "log"
)

type config struct {
	ServiceName string `yaml:"ServiceName"`
	Addr string `yaml:"Addr"`
	Log  log `yaml:"Log"`
	Prometheus prometheus `yaml:"Prometheus"`
}
type log struct {
	Level string `yaml:"Level"`
	Dir string `yaml:"Dir"`
	ConsoleLog bool `yaml:"ConsoleLog"`
}
type prometheus struct {
	SwitchOn bool `yaml:"SwitchOn"`
	Addr string `yaml:"Addr"`
}

var Config *config

func init() {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		createConfig()
		log2.Fatalln("config.yml err: ",err)
	}
	Config = &config{}
	err = yaml.Unmarshal(file, Config)
	if err != nil {
		log2.Fatalln("Config Unmarshal Error: ",err)
	}
}

func createConfig() {
	err := ioutil.WriteFile("config.yaml", []byte(configTemplate), 00755)
	if err != nil {
		log2.Fatalln(err)
	}
}


var configTemplate = `
ServiceName: "VodkaExample"
Addr: "0.0.0.0:8081"
Log:
  Level: "debug"
  Dir: "log"
  ConsoleLog: true
Prometheus:
  SwitchOn: true
  Addr: "0.0.0.0:8082"
`


