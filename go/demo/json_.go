// +build OMIT

package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func main() {
	data := []byte(`{"id": 123, "name": "jill"}`)
	user := User{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s", user.Name)
}