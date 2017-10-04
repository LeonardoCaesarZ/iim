package bll

import (
	"encoding/json"
	"iim/auth/dao"
	"iim/crypto"
	"iim/model"
)

// CheckAccontPasswd 检查密码，获取用户信息
func CheckAccontPasswd(email, passwd string) (bool, *model.User, error) {
	user, err := dao.CheckAccontPasswd(email, passwd)
	if err != nil {
		return false, user, err
	}
	if user.ID == 0 { // 登录信息错误
		return false, user, nil
	}
	return true, user, nil
}

// EnAesAfterJSON 结构体转Json后进行AES加密
func EnAesAfterJSON(aesKey string, data interface{}) (string, error) {
	jBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	result, err := crypto.EnAES(jBody, []byte(aesKey))
	if err != nil {
		return "", err
	}

	return string(result), nil
}
