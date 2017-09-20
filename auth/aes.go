package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"math"
	"math/rand"
	"time"
)

var (
	aesBits = 16 // 16x8bit=128bit, can choose: 16 or 24 or 32
)

func GenerateAESKey() []byte {
	bs := make([]byte, aesBits)
	for i := 0; i < aesBits; i += 8 {
		r := rand.New(rand.NewSource(time.Now().UnixNano())) // rand seed
		bits := math.Float64bits(r.Float64())
		binary.LittleEndian.PutUint64(bs[i:i+8], bits)
	}
	return bs
}

func AESEncrypt(data []byte, key []byte) ([]byte, error) {
	// 生成加密用的block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 填充，保证位长度为aesBits的倍数
	blockSize := block.BlockSize()
	npadding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{0}, npadding)
	data = append(data, padtext...)

	// CBC模式加密
	blockMode := cipher.NewCBCEncrypter(block, key)
	encrypted := make([]byte, len(data))
	blockMode.CryptBlocks(encrypted, data)

	return encrypted, nil
}

func AESDecrypt(data []byte, key []byte) ([]byte, error) {
	// 生成加密用的block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// CBC模式解密
	blockMode := cipher.NewCBCDecrypter(block, key)
	decrypted := make([]byte, len(data))
	blockMode.CryptBlocks(decrypted, data)

	return decrypted, nil
}
