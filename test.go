package main

import (
	"fmt"
	"time"
)

func main() {
	//a := "9s8df098kllfd"
	fmt.Println(100100 << 1)
	fmt.Println(100100 << 2)

	fmt.Println(time.Now().UnixNano()-1542125745114871000 <= 60*1000*1e6)
}
