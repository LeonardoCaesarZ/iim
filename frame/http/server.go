package http

import (
	"encoding/json"
	"iim/frame/http/errors"
	"net/http"
	"strconv"
)

// Server 框架核心
type Server struct {
	base       string              // 所有注册URL的前部增加部分，减少重复填写
	router     map[string]HttpFunc // string: URL+method, HttpFunc: 处理函数
	registered []string            // 已注册的URL
}

// HttpFunc 使用框架的所有请求处理函数必须按照此格式
type HttpFunc func(context *Context) interface{}

// NewServer 返回新的框架核心
func NewServer() *Server {
	h := &Server{}
	h.base = ""
	h.router = make(map[string]HttpFunc)
	h.registered = []string{}
	return h
}

// Base 为所有注册URL的增加前部，减少重复填写
func (h *Server) Base(base string) {
	h.base = base
}

// Get 注册支持GET模式的URL
func (h *Server) Get(pattern string, handler HttpFunc) {
	h.handleFunc("GET", pattern, handler)
}

// Post 注册支持POST模式的URL
func (h *Server) Post(pattern string, handler HttpFunc) {
	h.handleFunc("POST", pattern, handler)
}

// Down 为pattern路径提供HTTP协议的文件下载功能
func (h *Server) Down(pattern, res string) {
	http.Handle(pattern, http.FileServer(http.Dir(res)))
}

// Serve 开始接受服务，代码阻塞
func (h *Server) Serve(port int) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}
}

func (h *Server) handleFunc(method string, pattern string, handler HttpFunc) {
	pattern = h.base + pattern
	// 相同URL相同method会覆盖原先注册的处理函数
	h.router[pattern+method] = handler

	// 防止二次HandleFunc造成的panic，同一注册URL可能同时接受GET、POST模式
	isok := true
	for _, url := range h.registered {
		if url == pattern {
			isok = false
		}
	}
	if isok {
		http.HandleFunc(pattern, h.workFunc)
		h.registered = append(h.registered, pattern)
	}
}

func (h *Server) workFunc(w http.ResponseWriter, r *http.Request) {
	context := &Context{}
	context.R = r
	context.W = w

	handler, ok := h.router[r.URL.Path+r.Method]
	if !ok { // 相同URL不同method时会触发，即该URL上的该方法未注册
		context.RespondErr(errors.ErrMethodNotSupport)
		return
	}

	result := handler(context)
	if result != nil {
		err, ok := result.(error)
		if ok {
			// 未注册的错误
			context.RespondErr(errors.Err(err.Error()))
			return
		}

		Err, ok := result.(*errors.Error)
		if ok {
			// 已注册错误类型
			context.RespondErr(Err)
			return
		}

		bs, ok := result.([]byte)
		if ok {
			w.Write(bs)
			return
		}

		// 处理函数成功响应
		jbody, err := json.Marshal(result)
		if err != nil {
			context.RespondErr(errors.Err(err.Error()))
		}
		w.Write(jbody)
	}
	// result 为nil时默认响应码200，包体为空，处理函数成功响应
}
