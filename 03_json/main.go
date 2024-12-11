package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"username"`
}

var JSONData = `
{
	"id": 555,
	"price": 4000,
	"items": [
	{
		"name": "shoes",
		"number": 1
	},
	{
		"name": "pants",
		"number": 2
	},
	{
		"name": "t-short",
		"number": 3
	}]
}
`

type Order struct {
	Id    int    `json:"id"`
	Price int    `json:"price"`
	Items []Item `json: "items"`
}

type Item struct {
	Name   string `json: "name"`
	Number int    `json: "number"`
}

func main() {
	//Data >>> JSON
	user1 := User{
		555,
		"Ivaov Ivan",
	}
	jsonMarshal, _ := json.Marshal(user1)
	jsonMarshalIndent, _ := json.MarshalIndent(user1, "", "----")
	fmt.Println(string(jsonMarshal))
	fmt.Println(string(jsonMarshalIndent))

	/// JSON >>> DATA
	var order1 Order
	err := json.Unmarshal([]byte(JSONData), &order1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", order1)
}
