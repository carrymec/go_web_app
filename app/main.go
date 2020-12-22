package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {

	id := primitive.NewObjectID()
	fmt.Println(id.Hex())
	fmt.Println("v1.0")
	fmt.Println("v2")
	fmt.Println("v3")
}
