Prometheus
===
- 数据类型
```
cpu_usage 14.04
key     val
cpu_usage{core="1",ip="128.0.0.1"}  14.04
key        lobels                val
DataModelFiltering
```
- 度量类型
- 计数器Counter采样 处理一个请求计数器加1
- 比如 http错误次数 or pv
- Gauges采样
- 用于处理可能随时间减少的值。 可以上升下降 正 or 负
- 记录瞬间的值，比如内存变化 温度变化 连接池连接数
- Histogram采样
- 对每个采样殿进行统计(并不是一段时间的统计) 打到各个bucket中去
- 对每个采样点值累计和sum
- 对采样点计次累计和count
- Summary采样
- 对客户端在一段时间内(默认10min) 的每个采样点进行统计，并形成分位图

### 中间件开发
- 当前服务元信息封装
    - 当前服务名
    - 当前方法名
    - 当前环境 (生产环境or测试环境)
    - 当前服务集群
    - 当前服务的机房
    - 当前请求的trace_id
    - 当前服务器ip
    - 客户端请求ip

// rate（hd_errors_total{service="hh"}[10m]) 统计10m增加数量
