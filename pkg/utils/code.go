package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

func GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+{}|:?><~"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func RandomMD5(key string) string {
	return Md5(GetRandomString() + string(time.Now().UnixNano()) + key)
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//日期格式：模仿java中的结构体
type DateStyle string

const (
	MM_DD                   = "MM-dd"
	YYYYMM                  = "yyyyMM"
	YYYY_MM                 = "yyyy-MM"
	YYYY_MM_DD              = "yyyy-MM-dd"
	YYYYMMDD                = "yyyyMMdd"
	YYYYMMDDHHMMSS          = "yyyyMMddHHmmss"
	YYYYMMDDHHMM            = "yyyyMMddHHmm"
	YYYYMMDDHH              = "yyyyMMddHH"
	YYMMDDHHMM              = "yyMMddHHmm"
	MM_DD_HH_MM             = "MM-dd HH:mm"
	MM_DD_HH_MM_SS          = "MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM        = "yyyy-MM-dd HH:mm"
	YYYY_MM_DD_HH_MM_SS     = "yyyy-MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS = "yyyy-MM-dd HH:mm:ss.SSS"

	MM_DD_EN                   = "MM/dd"
	YYYY_MM_EN                 = "yyyy/MM"
	YYYY_MM_DD_EN              = "yyyy/MM/dd"
	MM_DD_HH_MM_EN             = "MM/dd HH:mm"
	MM_DD_HH_MM_SS_EN          = "MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_EN        = "yyyy/MM/dd HH:mm"
	YYYY_MM_DD_HH_MM_SS_EN     = "yyyy/MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS_EN = "yyyy/MM/dd HH:mm:ss.SSS"

	MM_DD_CN               = "MM月dd日"
	YYYY_MM_CN             = "yyyy年MM月"
	YYYY_MM_DD_CN          = "yyyy年MM月dd日"
	MM_DD_HH_MM_CN         = "MM月dd日 HH:mm"
	MM_DD_HH_MM_SS_CN      = "MM月dd日 HH:mm:ss"
	YYYY_MM_DD_HH_MM_CN    = "yyyy年MM月dd日 HH:mm"
	YYYY_MM_DD_HH_MM_SS_CN = "yyyy年MM月dd日 HH:mm:ss"

	HH_MM       = "HH:mm"
	HH_MM_SS    = "HH:mm:ss"
	HH_MM_SS_MS = "HH:mm:ss.SSS"
)

//日期转字符串
func FormatDate(date time.Time, dateStyle DateStyle) string {
	layout := string(dateStyle)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSS", "000", -1)

	return date.Format(layout)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(data string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	var origData []byte
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding([]byte(data), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(data string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var crypted []byte
	crypted, e := base64.StdEncoding.DecodeString(data)
	if e != nil {
		return nil, e
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)

	origData = PKCS5UnPadding(origData)

	return origData, nil
}
