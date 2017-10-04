package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"strconv"
)

// aes 加解密参数
const (
	AesBits = 16 // 16x8bit=128bit, can choose: 16 or 24 or 32, AesBits对应的base64长度，16->24, 24->32, 32->48
)

// GenerateAESKey 生成随机AES密钥
func GenerateAESKey() []byte {
	// 方案1
	// 生成的随机byte数组，每byte取值0-255
	// 转码前必须先base64，base64会造成长度增加
	// bs := make([]byte, AesBits)
	// for i := 0; i < AesBits; i += 8 {
	// 	r := rand.New(rand.NewSource(time.Now().UnixNano())) // rand seed
	// 	bits := math.Float64bits(r.Float64())
	// 	binary.LittleEndian.PutUint64(bs[i:i+8], bits)
	// }
	// return bs

	// 方案2
	// 每byte取值48-57，纯数字
	// 不用考虑转码问题
	key := ""
	for i := 0; i < AesBits; i += 8 {
		key += strconv.Itoa(10000000 + rand.Intn(89999999))
	}
	return []byte(key)
}

// EnAES 加密
func EnAES(key []byte, data []byte) ([]byte, error) {
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

// DeAES 解密
func DeAES(key []byte, data []byte) ([]byte, error) {
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
