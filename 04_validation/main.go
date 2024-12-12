package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"username"`
}

func main() {
	http.HandleFunc("/user", UserHandlerReturnJSON)
	http.HandleFunc("/user1", UserHandlerTakeJSON)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func UserHandlerTakeJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}

	var testUser User

	err := json.NewDecoder(r.Body).Decode(&testUser)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("userTest %v", testUser)
	WriteJSON(w, http.StatusOK, map[string]any{
		"ok": true,
	})
}

func UserHandlerReturnJSON(w http.ResponseWriter, r *http.Request) {
	testUser := User{
		Name: "Ivanov Ivan",
		Id:   443490,
	}
	err := WriteJSON(w, http.StatusOK, testUser)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}
}

/*
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
*/
