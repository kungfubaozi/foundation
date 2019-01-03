package main

import (
	"fmt"
	"github.com/pborman/uuid"
)

func main() {
	//a := "9s8df098kllfd"
	str := fmt.Sprintf("%s.%s", "e6875cefb1a057adcb37adae3e53359a", uuid.New())
	fmt.Println(str[0:32])
}
