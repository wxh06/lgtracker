package main

import (
	"fmt"

	"github.com/wxh06/lgtracker/internal/datafile"
)

func main() {
	var users = map[int]map[int]map[string]int{}

	if err := datafile.Read("public/users.json", &users); err != nil {
		panic(err)
	}

	// 存储每个用户的通过情况
	for user, passed := range users {
		if err := datafile.Write(fmt.Sprintf("public/users/%d.json", user), passed); err != nil {
			panic(err)
		}
	}
}
