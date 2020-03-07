/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-07 13:03
 */
package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU. ",
	})

	// rate（hd_errors_total{service="hh"}[10m]) 统计10m增加数量
	hdFailures = prometheus.NewCounterVec( // counter 计数器
		prometheus.CounterOpts{
			Name: "hd_errors_total", // 硬盘采样
			Help: "Number of hard-disk errors.",
		},
		[]string{"device", "service"}, // 自定义tag
	)

	rpcDurations = prometheus.NewSummaryVec( // 统计分布
		prometheus.SummaryOpts{
			Name:       "rpc_durations_seconds",
			Help:       "RPC latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}, // 50%请求耗时0.05  90%请求0.01 99%
		},
		[]string{"service"},
	)
)

func init() {
	// 注册
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
	prometheus.MustRegister(rpcDurations)
}

func main() {
	// 这个协议程序 作定时上报
	go func() {
		for {
			cpuTemp.Set(165.3)
			hdFailures.With(prometheus.Labels{
				"device":  "/dev/sda",    // 磁盘
				"service": "hello.world", // 服务
			}).Inc() // 递增
			rpcDurations.WithLabelValues("uniform").Observe(0.05) // 计算出请求耗时
			time.Sleep(time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Println(http.ListenAndServe(":8081", nil))
	//time.Sleep(time.Second * 2)
}
