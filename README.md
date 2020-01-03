# Vodka MicroServices Framework 
技术验证

经历过一些分布方式 项目 新的思路  重构vodka

### 注册中心
采用插件式开发
内部实现了三种插件  etcd,redis,mem


### 依赖
- Sonyflake  (分布式ID)
- ETCD (持久化)
- Redis (持久化)
- json-iterator (json序列化)
- grpc (通讯)
- easylog (日志库)
