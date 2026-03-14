package main

import (
	"encoding/json"
	"fmt"
	"GoLearning/utils"
	"GoLearning/closures"
)

const apiName = "messages"
var count int

// Struct for {"name": "Kanishk", "age": 25, "skills": ["React", "Go"]}

type User struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Skills []string `json:"skills"`
}

func main() {
	fmt.Println(apiName, count)
	fmt.Println(GetName())
	fmt.Println(utils.GetAge())
	// fmt.Println(x.GetExp())

	var user User

	user = User{
		Name: "Kanishk",
		Age: "25",
		Skills: []string{"React", "Go"},
	}
	// user := User{
	// 	Name: "Kanishk",
	// 	Age: "25",
	// 	Skills: []string{"React", "Go"},
	// }

	fmt.Println(user)

	jsonData, err := json.Marshal(user)

	fmt.Println(string(jsonData), "##", err);


	// Pointers
	var x int = 2;
	var y int = 2;
	Double(2)
	DoubleUsingPointer(&y)
	fmt.Println("Value of X: ",  x)

	fmt.Println(&y,"Val of y: ", y)

	// var z int = 10;

	c := &Calculator{ result: 100 }
	c.Add(10);

	fmt.Println(c, )

	counter := closures.MakeCounter();
	counter()
	counter()
	fmt.Println(counter())

}