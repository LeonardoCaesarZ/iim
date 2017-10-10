package tcp

import (
	"net"
)

// todo
// 1. 优雅退出

// Server TCP服务器结构
type Server struct {
	listen  *net.TCPListener
	handler TcpFunc
	// alive   bool
	// smooth  bool
}

type TcpFunc func(context *Context) interface{} // TcpFunc 连接逻辑处理函数

// NewServer 新建Tcp服务器类
func NewServer(port int, handler TcpFunc) *Server {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(""), port, ""}) // 创建、绑定、监听
	if err != nil {
		panic("fail to listen")
	}

	// server := &Server{listen, handler, logger, true, false}
	server := &Server{listen, handler}
	// go server.sigHandler()

	return server
}

// func (s *Server) sigHandler() {
// 	c := make(chan os.Signal)
// 	signal.Notify(c)
// 	sig := <-c
// 	switch sig {
// 	case syscall.SIGTERM: // 平缓终止
// 		s.alive = false
// 		s.smooth = true
// 	}
// }

// Loop Tcp服务器工作主循环
func (s *Server) Loop() {
	for {
		conn, err := s.listen.AcceptTCP()
		if err != nil { // 发生错误不跳出循环
			// s.logger(conn.RemoteAddr().String(), err) // 记录日志
			continue
		}
		go s.work(conn) // 处理函数
	}
}

// 封装handler与logger于同一协程
func (s *Server) work(conn *net.TCPConn) {
	context := NewContext(conn)
	if context == nil {
		return
	}
	result := s.handler(context)
	msg, ok := result.(string)
	if ok {
		conn.Write([]byte(msg))
	}
	conn.Close()
	// go s.logger(context.Addr, result)
}

// func (s *Server) logger(addr string, info interface{}) {
// 	return
// }
