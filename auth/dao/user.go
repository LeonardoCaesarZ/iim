package dao

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"iim/db"
	"iim/model"
)

// CheckAccontPasswd 依据用户密码从Mysql获取用户信息
func CheckAccontPasswd(email, passwd string) (*model.User, error) {
	user := &model.User{}
	m := md5.New()
	m.Write([]byte(passwd))
	query := fmt.Sprintf("select id, name, email, host, is_ban, is_vertify from user where email = '%s' and passwd = '%s'", email, hex.EncodeToString(m.Sum(nil)))
	err := db.Master.Get(user, query)

	if err != nil {
		return user, err
	}

	return user, nil
}
