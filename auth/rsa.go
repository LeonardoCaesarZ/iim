package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

var (
	privatePath = "pem/private.pem"
	publicPath  = "pem/public.pem"
	privateKey  *rsa.PrivateKey
)

func readPriKeyIntoMemroy() error {
	privatePem, err := ioutil.ReadFile(privatePath)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(privatePem)
	if block == nil {
		return errors.New("private key error")
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	return nil
}

func RSAEncrypt(data []byte) ([]byte, error) { // just for test, will not be used in im
	publicKey, err := ioutil.ReadFile(publicPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pubInterface.(*rsa.PublicKey), data)
}

func RSADecrypt(data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
}

func GenerateRSAKeyPair(bits int) error {
	// 生成rsa密钥对
	keyPair, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	// 导出私钥证书
	bs := x509.MarshalPKCS1PrivateKey(keyPair)
	block := &pem.Block{
		Type:  "im private",
		Bytes: bs,
	}
	file, err := os.Create(privatePath)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	// 导出公钥证书
	bs, err = x509.MarshalPKIXPublicKey(&keyPair.PublicKey)
	block = &pem.Block{
		Type:  "im public",
		Bytes: bs,
	}
	if err != nil {
		return err
	}
	file, err = os.Create(publicPath)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	return nil
}
