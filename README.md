# kit-demo
go-kit的demo系统

目demo的开发目的：
* 学习和分享交流
* 为微服务开发，尤其是基于kit的微服务开发，提供参考

没有单元测试。正式项目开发应至少为基础逻辑提供单元测试

几个组件是因为找不到合适的开源项目，自己实现的，但也是自己思考已久的结果。因为是新实现的，功能简单，质量强差人意，待以后慢慢优化完善。

## 特性
### 已实现特性
* 解偶业务逻辑和接口逻辑
* GRPC和HTTP并存，并共用业务逻辑
* 比较通用的错误处理
* 结构化日志

### 实现中的特性
* 服务注册/发现、负载均衡、失败重试

### 计划小目标
* 限流、熔断
* 基于mDNS的服务注册/发现
* 分布式跟踪

### 也许大目标
* 根据proto接口定义文件生成接口逻辑（endpoint和transport）代码

## 目录结构
```
|-- go-cli                      // cli客户端
|   |-- endpoint                // endpoint
|   `-- transport               // 传输层
|       |-- grpc                // grpc传输层
|       `-- http                // http传输层
|-- go-service                  // 服务
|   |-- cmd                     // 接口逻辑
|   |   |-- endpoint            // endpoint
|   |   `-- transport           // 传输层
|   |       |-- grpc            // grpc传输层
|   |       `-- http            // http传输层
|   `-- service                 // 业务逻辑实现
`-- protocol                    // 协议
    |-- *.proto                 // 接口定义
    |-- github.com/...          // 生成的grpc接口
    |-- google.golang.org/...   // 生成的grpc接口
    `-- model                   // 协议请求/响应数据模型
```