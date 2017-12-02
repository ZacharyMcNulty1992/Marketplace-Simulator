package main

import (
	"actor"
	"fmt"
)

func main() {
	fmt.Printf("Initializing World\n")
	//TODO: create a world map

	//TODO: change this so that it is creating a lot of new people.
	person := actor.NewPerson("merchant")
	//TODO: have the main thread be the thread that controls world events.
	person.Init()
}
