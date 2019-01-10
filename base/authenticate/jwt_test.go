package authenticate

import (
	"fmt"
	"testing"
	"time"
	"zskparker.com/foundation/pkg/utils"
)

func TestDecodeToken(t *testing.T) {

	node := utils.NodeGenerate()

	sa := &SimpleAuthorize{
		UserId:   node.Generate().Base64(),
		ClientId: node.Generate().Base64(),
		UUID:     node.Generate().Base64(),
		Access:   false,
		Relation: node.Generate().Base64(),
	}
	str, err := encodeToken(time.Hour*24*7, sa)
	if err != nil {
		panic(err)
	}

	fmt.Println(str)

}
