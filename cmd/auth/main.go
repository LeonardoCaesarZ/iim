package main

import (
	"iim/auth/api"
	"iim/crypto"
	"iim/db"
	"iim/frame/http"
	"iim/session"
)

func main() {
	crypto.Init()  // 加解密模块初始化，加载私钥
	session.Init() // 会话模块初始化，伪随机数种子初始化
	db.Init()      // 数据库模块初始化，导入参数，测试连接

	server := http.NewServer()
	server.Base("/auth/")              // 设定Get、Post的基础路径，Down不生效
	server.Post("login", api.Login)    // 登录接口
	server.Down("/", "./data/public/") // 开放静态资源下载
	server.Serve(9999)                 // 绑定端口并开始服务
}
