// +build OMIT
package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	id int
	name string
}

func main() {
	data := []byte(`{"id": 123, "name": "jill"}`)
	user := User{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s", user.name)
}