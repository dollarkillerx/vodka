# Vodka MicroServices Framework 
重构vodka

### 依赖
- Sonyflake  (分布式ID)
- ETCD (持久化)
- Redis (持久化)
- json-iterator (json序列化)
- grpc (通讯)
- easylog (日志库)


### 注册中心
采用插件式开发
内部实现了2种插件  etcd,redis

### 负载均衡
- 随机
- 轮询
- 加权
    - 加权随机
    - 加权轮询
- 一致性hash

### 脚手架
- 目录规范  (包管理 按照官方GO MOD)
```
controller: 存在服务的方法实现
idl: 存放服务的idl定义
main: 存放服务的入口代码
scripts: 存放服务的脚本
conf: 存放服务的配置文件
app/router: 存放服务的路由
app/config: 存放服务的一些配置
datamodels: 存放服务的实体代码
generate: grpc生成的代码
```
- 命令行参数设计
    - -f 指定idl文件
    - -o 指定代码生成路径
    - -c 指定客户端调用代码
    - -s 指定服务端调用代码
