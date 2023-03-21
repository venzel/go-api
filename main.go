package main

import (
	"fmt"
	"log"
	user "server/internal/user/domain"
)

func main() {
	user, err := user.CreateUser("1", "Tiago", "tiago@gmail.com", "bolaverde")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
