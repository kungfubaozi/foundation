package main

import (
	"fmt"
	"zskparker.com/foundation/pkg/tool/encrypt"
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

	fmt.Println(fs_tools_encrypt.SHA256_1_512("123442342356"))
}
