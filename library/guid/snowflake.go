package guid

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/frame/g"
)

var (
	workerId int64 = 1
	node *snowflake.Node = nil
)

func init() {
	var err error = nil
	node, err = snowflake.NewNode(workerId)
	if err != nil {
		g.Log().Error("init snowflake error", err.Error())
	}
	fmt.Println("init snowflake")
}

func Next() int64 {
	return node.Generate().Int64()
}
