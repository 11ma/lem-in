package main

import (
	"fmt"
	lemin "lemin/packages"
)

func main() {

	numberOfAnts, rooms, start, end := lemin.Rooms()

	fmt.Println("numberOfAnts: ", numberOfAnts)
	fmt.Println("rooms: ", rooms)
	fmt.Println("start: ", start)
	fmt.Println("end: ", end)
}
