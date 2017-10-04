package http

import (
	"encoding/json"
	"iim/crypto"
	"iim/frame/http/errors"
	"io/ioutil"
	"net/http"
)

// Context 供请求处理函数使用的输入参数
type Context struct {
	R *http.Request
	W http.ResponseWriter
}

// ParseBody 定义了POST Body的解析方式
func (c *Context) ParseBody(dst interface{}) error {
	return c.parseBodyWithChoice(dst, false)
}

// ParseBodyAfterDeRSA 进行RSA解密后才进行json解析
func (c *Context) ParseBodyAfterDeRSA(dst interface{}) error {
	return c.parseBodyWithChoice(dst, true)
}

// 选择是否先进行RSA解密
func (c *Context) parseBodyWithChoice(dst interface{}, deRsa bool) error {
	body, err := ioutil.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	if deRsa {
		body, err = crypto.DeRSA(body)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}
	return nil
}

// RespondErr 响应错误
func (c *Context) RespondErr(err *errors.Error) {
	c.W.WriteHeader(err.Code)

	// 返回包体
	body := struct {
		Msg  string `json:"msg"`
		Text string `json:"text"`
	}{err.Msg, err.Text}

	jbody, _ := json.Marshal(body) // 可认为在errors.go中添加的错误类型不会产生循环引用
	c.W.Write(jbody)
}
