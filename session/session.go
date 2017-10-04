package session

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// todo
// 1. session文件夹下的多级目录

const (
	sessionDir = "./data/private/session/"
)

// Init session module
// 每次启动都会初始化伪随机数种子，使伪随机序列不同于前一次生成的
func Init() {
	fmt.Println("[modeul session]")
	rand.Seed(time.Now().Unix())
	fmt.Println("initialize random seed... [OK]")
}

// CreateSession 在sessionDir目录下创建session文件
// 根据输入参数决定文件名，即sessionID
func CreateSession(r *http.Request, data interface{}) (string, error) {
	sessionID := generateSessionID(r.RemoteAddr)
	return sessionID, UpdateSession(sessionID, data)
}

// UpdateSession 更新session
func UpdateSession(sessionID string, data interface{}) error {
	jBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(sessionDir+sessionID, []byte(jBody), 777)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSession 删除session
func DeleteSession(sessionID string) error {
	err := os.Remove(sessionDir + sessionID)
	if err != nil {
		return err
	}
	return nil
}

// 截取或用'0'补全字符串至指定长度
func formatString(x string, returnLen int) string {
	length := len(x)
	if length > returnLen {
		return x[length-returnLen:]
	}
	for i := 0; i < returnLen-length; i++ {
		x = "0" + x
	}
	return x
}

// generateSessionID generate a sessionID by ip, timestamp, random number,
func generateSessionID(ip string) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10) // 时间戳
	random := strconv.Itoa(rand.Intn(999999))             // 随机数6位以内

	m := md5.New()
	m.Write([]byte(ip + timestamp))
	return hex.EncodeToString(m.Sum(nil)) + formatString(timestamp, 4) + formatString(random, 4)
}
