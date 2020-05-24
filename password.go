package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type user struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//get the user account
func checkUser() user {
	jsonFile, err := os.Open("account.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users user

	json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func changePassword(myuser user) {
	file, _ := json.Marshal(myuser)
	ioutil.WriteFile("account.json", file, 0644)
}
