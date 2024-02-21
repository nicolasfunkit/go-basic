package main

import "fmt"

var (
	Stats = map[string]int{
		"create": 0,
		"update": 0,
	}
)

func CreateUser(user string) {
	fmt.Println("Creating user", user)
	Stats["create"] = Stats["create"] + 1
}

func UpdateUser(user string) {
	fmt.Println("Updating user", user)
	Stats["update"] = Stats["update"] + 1
}

func PurgeStats() {
	Stats = map[string]int{
		"create": 0,
		"update": 0,
	}
}
