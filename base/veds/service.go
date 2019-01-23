package veds

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/veds/pb"
	"zskparker.com/foundation/pkg/errno"
)

type Service interface {
	Encrypt(ctx context.Context, in *fs_base_veds.CryptRequest) (*fs_base_veds.CryptResponse, error)

	Decrypt(ctx context.Context, in *fs_base_veds.CryptRequest) (*fs_base_veds.CryptResponse, error)
}

type vedsService struct {
	key string
}

func (svc *vedsService) Decrypt(ctx context.Context, in *fs_base_veds.CryptRequest) (*fs_base_veds.CryptResponse, error) {
	resp := func(s *fs_base.State) (*fs_base_veds.CryptResponse, error) {
		return &fs_base_veds.CryptResponse{State: s}, nil
	}

	if len(in.Value) == 0 && len(svc.key) != 32 {
		return resp(errno.ErrRequest)
	}

	s, e := decrypt(in.Value, []byte(svc.key))

	if e != nil {
		return resp(errno.ErrDecrypt)
	}

	return &fs_base_veds.CryptResponse{
		State: errno.Ok,
		Value: s,
	}, nil

}

func (svc *vedsService) Encrypt(ctx context.Context, in *fs_base_veds.CryptRequest) (*fs_base_veds.CryptResponse, error) {
	resp := func(s *fs_base.State) (*fs_base_veds.CryptResponse, error) {
		return &fs_base_veds.CryptResponse{State: s}, nil
	}

	if len(in.Value) == 0 && len(svc.key) != 32 {
		return resp(errno.ErrRequest)
	}

	s, e := encrypt([]byte(in.Value), []byte(svc.key))

	if e != nil {
		return resp(errno.ErrEncrypt)
	}

	return &fs_base_veds.CryptResponse{
		State: errno.Ok,
		Value: s,
	}, nil

}

func NewService(key string) Service {
	var svc Service
	{
		svc = &vedsService{key: key}
	}
	return svc
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func cbcEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//填充原文
	blockSize := block.BlockSize()
	rawData = pkcs7Padding(rawData, blockSize)
	//初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	//block大小 16
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	//block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func cbcDncrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	//解填充
	encryptData = pkcs7UnPadding(encryptData)
	return encryptData, nil
}

func encrypt(rawData, key []byte) (string, error) {
	data, err := cbcEncrypt(rawData, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func decrypt(rawData string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dnData, err := cbcDncrypt(data, key)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}
