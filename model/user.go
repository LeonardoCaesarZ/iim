package model

// User 包含注册用户的可公开信息
type User struct {
	ID        int64  `db:"id"         json:"id"`
	Name      string `db:"name"       json:"name"`
	Email     string `db:"email"      json:"email"`
	Host      string `db:"host"       json:"host"`
	IsBan     bool   `db:"is_ban"     json:"is_ban"`
	IsVertify bool   `db:"is_vertify" json:"is_vertify"`
	AesKey    string `json:"aes_key"` // 每次登录分配的由认证服务器生成的随机AES密钥
}
