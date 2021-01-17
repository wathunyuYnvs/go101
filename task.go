package task

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type Users []User

var myStore = Store{"test.json"}

func Add(args1 string, args2 string) {
	i, _ := strconv.ParseInt(strings.Split(args2, "=")[1], 10, 64)
	AddToUsers(strings.Split(args1, "=")[1], i)
}
func List() {
	users := GetUsers()
	b, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println(string(b))
	}
}
func AddToUsers(name string, age int64) {
	Users := GetUsers()
	a := len(Users)
	u := User{
		ID:   int64(a + 1),
		Name: name,
		Age:  age,
	}
	Users = append(Users, u)
	SaveToJson(Users)
}

func SaveToJson(users Users) {
	b, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		myStore.write(b)
	}
}

func GetUsers() Users {
	var out Users
	if jsonFile := myStore.read(); jsonFile != nil {
		err := json.Unmarshal(jsonFile, &out)
		if err != nil {
			fmt.Println("Error ", err)
		}
	} else {
		out = Users{}
	}
	return out
}
