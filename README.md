# iim
an experimental instant messaging system

## 设计

iim系统由认证、通讯、资源三个子系统组成。

客户端A、B与服务器C之间的大致运作流程为：

1. A、B分别使用账号密码向C的**认证系统**进行登录操作

   该请求附带随机AES(1)，使用RSA公钥加密；响应为Cookie与随机AES(2)，使用AES(1)解密

2. A、B分别使用Cookie向C的**通讯系统**进行连接，随后使用该连接进行通讯

   A、B通过已建立连接向C相互发送多种信息包体，全程使用AES(2)进行信息体的加解密

3. A、B利用**资源系统**进行资源文件的传输

   A、B在通讯过程中收发资源文件包时，会进行资源文件的下载和上传

### 认证系统

> 基于HTTP协议

1. 访问MySQL核对登录信息
2. 生成session文件
3. 返回sessionID

### 通讯系统

> 基于TCP协议

### 资源系统

> 基于HTTP协议

## 目录结构

```
iim
├── auth	// 认证系统
├── cmd		// 各子系统的main.go
├── db		// 数据库工具
├── model	// 结构体定义
├── pem		// RSA证书
├── session	// session模块
├── sql		// 数据库初始化脚本
└── vendor	// 依赖
```

