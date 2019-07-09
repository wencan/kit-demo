# kit-demo
go-kit的demo系统

## 已实现特性
* 解偶业务逻辑和协议接口
* GRPC和HTTP并存，并共用业务逻辑
* 比较通用的错误处理
* 结构化日志

## 实现中的特性
* 服务注册/发现、负载均衡、失败重试

## 计划小目标
* 限流、熔断
* 基于mDNS的服务注册/发现
* 分布式跟踪

## 也许大目标
* 根据proto接口定义文件生成协议相关代码