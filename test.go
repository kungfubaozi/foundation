package main

import (
	"fmt"
	"zskparker.com/foundation/pkg/match"
)

func main() {

	//h := sha1.New()
	//io.WriteString(h, "aaaaaa")
	//fmt.Printf("%x\n", h.Sum(nil))
	//
	////hmac ,use sha1
	//key := []byte("123456")
	//mac := hmac.New(sha1.New, key)
	//mac.Write([]byte("aaaaaa"))
	//fmt.Printf("%x\n", mac.Sum(nil))
	//
	//fmt.Println(fs_tools_encrypt.Rnd_SHA1_256_512(8))
	//fmt.Println(fs_tools_encrypt.Rnd_SHA1_256_512(8))
	//fmt.Println(fs_tools_encrypt.Rnd_SHA1_256_512(8))
	//fmt.Println(fs_tools_encrypt.Rnd_SHA1_256_512(8))
	//fmt.Println(fs_tools_encrypt.Rnd_SHA1_256_512(8))
	//fmt.Println(fs_tools_encrypt.Rnd_SHA1_256_512(8))

	//s := "1322202107"

	//fmt.Println(s[:1] + s[strings.Index(s, "@")+1:])
	//fmt.Println(s[:3] + s[len(s)-1:])

	fmt.Println(fs_regx_match.Phone("17317278139"))
}
