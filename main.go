package main

import (
	"fmt"
	"log"
	"server/internal/register/domain/user"
)

func main() {
	user, err := user.CreateUser("1", "Tiago", "tiago@gmail.com", "bolaverde")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
