package api

import (
	"encoding/json"
	"iim/auth/bll"
	"iim/crypto"
	"iim/frame/http"
	"iim/frame/http/errors"
	"iim/session"
)

// Login POST /auth/login
func Login(c *http.Context) interface{} {
	// 节省空间，RSA 1024加解密长度限制为117
	params := struct {
		Account string `json:"a"` // 6-40
		Passwd  string `json:"p"` // 10-20
		AesKey  string `json:"k"` // 16, 24, 32
	}{}

	// 先RSA解密再JSON解析
	err := c.ParseBodyAfterDeRSA(&params)
	if err != nil {
		return err
	}

	// 检查参数
	accLen := len(params.Account)
	passLen := len(params.Passwd)
	aesLen := len(params.AesKey)
	if accLen < 6 || accLen > 40 ||
		passLen < 6 || passLen > 20 ||
		(aesLen != 16 && aesLen != 24 && aesLen != 32) {
		return errors.ErrParamIsWrong
	}

	// 通过数据库检查用户登录信息
	isLogin, user, err := bll.CheckAccontPasswd(params.Account, params.Passwd)
	if err != nil {
		return err
	}

	// 登录信息错误
	if !isLogin {
		return errors.ErrAuthNotPass
	}

	// 生成随机AES密钥
	// aesKey := base64.StdEncoding.EncodeToString(crypto.GenerateAESKey())
	aesKey := string(crypto.GenerateAESKey())
	user.AesKey = aesKey

	// 新建session
	sessionID, err := session.CreateSession(c.R, &user)
	if err != nil {
		return err
	}

	result := struct {
		SessionID string `json:"session_id"`
		AesKey    string `json:"aes_key"`
		Host      string `json:"host"`
	}{sessionID, aesKey, user.Host}

	jBody, err := json.Marshal(&result)
	if err != nil {
		return err
	}

	enAesBody, err := crypto.EnAES([]byte(params.AesKey), jBody)
	if err != nil {
		return err
	}

	return enAesBody
}
