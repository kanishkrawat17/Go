package main

import (
	"fmt"
	"GoLearning/utils"
)

const apiName = "messages"
var count int

func main() {
	fmt.Println(apiName, count)
	fmt.Println(GetName())
	fmt.Println(utils.GetAge())
	fmt.Println(x.GetExp())
}