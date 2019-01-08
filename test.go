package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().UnixNano()-1546938033630125600 <= 10*60*1e9)
}
