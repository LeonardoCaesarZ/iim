package tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

// todo:
// 1. 半包问题

// Context 封装了TCP连接的输入输出，包头包体模式
type Context struct {
	Conn   *net.TCPConn
	Addr   string
	Header *Header
	Body   []byte
}

// Header 包头
type Header struct {
	Code    int // 包的种类
	BodyLen int // 包体长度
	Time    int // 时间戳
	Vert    int // 校验位
}

// 起始符 终止符
var (
	START = []byte{'a', 'a', 'a', 'a'}
	END   = []byte{'b', 'b', 'b', 'b'}
)

// NewContext 返回按包头包体处理好的上下文
func NewContext(conn *net.TCPConn) *Context {
	context := &Context{conn, conn.RemoteAddr().String(), nil, nil}
	err := context.parse()
	if err != nil { // 包解析错误，则结束连接
		conn.Close()
		return nil
	}
	return context
}

func (c *Context) parse() error {
	err := c.readHeader()
	if err != nil {
		return err
	}

	err = c.readBody()
	if err != nil {
		return err
	}

	return nil
}

func (c *Context) readHeader() error {
	// 检查起始符
	start, err := c.read(4)
	if err != nil {
		return err
	}

	if c.bytesToInt(start) != c.bytesToInt(START) {
		return errors.New("起始符错误")
	}

	// 读取包种类码
	code, err := c.read(4)
	if err != nil {
		return err
	}

	// 读取包体长度
	len, err := c.read(4)
	if err != nil {
		return err
	}

	// 读取时间戳
	ts, err := c.read(4)
	if err != nil {
		return err
	}

	// 读取校验码
	check, err := c.read(4)
	if err != nil {
		return err
	}

	header := &Header{c.bytesToInt(code), c.bytesToInt(len), c.bytesToInt(ts), c.bytesToInt(check)}
	c.Header = header
	return nil
}

func (c *Context) readBody() error {
	body, err := c.read(c.Header.BodyLen)
	if err != nil {
		return err
	}

	// 检查结束符
	end, err := c.read(4)
	if err != nil {
		return err
	}

	if c.bytesToInt(end) != c.bytesToInt(END) {
		return errors.New("起始符错误")
	}

	c.Body = body
	return nil
}

func (c *Context) read(len int) ([]byte, error) {
	tmp := make([]byte, len)
	length, err := c.Conn.Read(tmp)
	if err != nil || length < len {
		return tmp, errors.New("包头或包体解析错误")
	}
	return tmp, nil
}

func (c *Context) bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}
