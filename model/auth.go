package model

type AuthRequest struct {
	Account string
	Passwd  string
	AesKey  string
}
