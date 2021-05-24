package hgenid

import (
	"github.com/hlib-go/snowflake"
	"math/rand"
	"strconv"
)

var g, _ = snowflake.NewNode(rand.Int63n(31))

func GenId() string {
	return g.Generate().String() + strconv.Itoa(rand.Intn(10))
}
