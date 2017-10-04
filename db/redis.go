package db

import (
	"errors"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// Redis 包含连接Redis必备的参数，后续开发连接池
type Redis struct {
	Host string
	Port string
}

// NewRedis 获得结构体
func NewRedis(host string, port string) *Redis {
	return &Redis{host, port}
}

func (r *Redis) ConnShort(command string, args ...interface{}) (interface{}, error) {
	conn, err := redis.Dial("tcp", r.Host+":"+r.Port)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err := conn.Do(command, args...)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *Redis) Set(key string, value interface{}) error {
	_, err := r.ConnShort("set", key, value)
	return err
}

func (r *Redis) Get(key string) (interface{}, error) {
	return r.ConnShort("get", key)
}

func (r *Redis) GetStr(key string) (string, error) {
	reply, err := r.Get(key)
	if err != nil {
		return "", err
	}

	value, ok := reply.([]byte)
	if !ok {
		return "", errors.New("wrong type")
	}

	return string(value), nil
}

func (r *Redis) GetInt(key string) (int, error) {
	str, err := r.GetStr(key)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return value, nil
}
