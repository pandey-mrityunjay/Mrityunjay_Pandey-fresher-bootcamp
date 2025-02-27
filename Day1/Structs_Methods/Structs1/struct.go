package main

import (
	"fmt"
	"time"
)

//Value type that groups together the related data

// defining a struct
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	var appUser user
	//instance of a struct
	appUser = user{
		firstName: "Mrityunjay",
		lastName:  "Pandey",
		birthdate: "28th Nov 2003",
		createdAt: time.Now(),
	}
   fmt.Println(appUser)
}
