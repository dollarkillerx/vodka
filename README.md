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
