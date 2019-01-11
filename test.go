package main

import (
	"encoding/base64"
	"fmt"
	"github.com/twinj/uuid"
	"zskparker.com/foundation/pkg/utils"
)

func main() {

	node := utils.NodeGenerate()
	fmt.Println(node.Generate().Base64())
	fmt.Println(node.Generate().Base64())

	fmt.Println(base64.StdEncoding.EncodeToString([]byte(uuid.NewV4().String())))
}
