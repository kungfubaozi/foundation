package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"zskparker.com/foundation/pkg/utils"
)

func main() {

	node := utils.NodeGenerate()

	uid, _ := uuid.NewV1()
	fmt.Println(uuid.NewV5(uid, utils.NodeGenerate().Generate().Base64()).String())
	fmt.Println(uuid.NewV5(uid, node.Generate().Base64()).String())
	fmt.Println(uuid.NewV5(uid, node.Generate().Base64()).String())
	fmt.Println(uuid.NewV5(uid, node.Generate().Base64()).String())

	fmt.Println(node.Generate().Base64())
	fmt.Println(node.Generate().Base64())
	fmt.Println(node.Generate().Base64())
	fmt.Println(node.Generate().Base64())

	fmt.Println(uuid.NewV5(uid, "3242").String()[24:])
}
