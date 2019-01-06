package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	fmt.Println(bson.NewObjectId().Hex())
}
