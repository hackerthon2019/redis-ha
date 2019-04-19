# 代码结构

1. app 路由及业务逻辑
2. client redis操作
3. cmd main.go
4. static 静态文件

# 如何运行

## 配置redis

此demo没有做配置文件，以代码常量的方式进行配置，定义在在`client/redis.go`中`。

1. `redis-cli` 客户端命令的绝对全路径（仅未进行环境变量配置的windows需要）
2. `netWork` 服务端的网络类型
3. `addr` 服务端IP地址
4. `port` 服务端端口号
5. `expire` 在redis存放数据的有效期（为测试数据加上有效期，可以省去手动删除的工作）

## 启动demo

```
// set up GOPATH
cd hackerthon2019 redis-ha-demo
go run hackerthon2019/redis-ha/redis-ha-demo/cmd/thon/main.go
```
