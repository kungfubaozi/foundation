package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().UnixNano()-1546997543346679100 <= 10*60*1e9)
}
