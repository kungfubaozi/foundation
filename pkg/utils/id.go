package utils

import (
	"github.com/bwmarrin/snowflake"
	"zskparker.com/foundation/pkg/osenv"
)

func NodeGenerate() *snowflake.Node {
	node, _ := snowflake.NewNode(osenv.GetNodeNumber())
	return node
}
