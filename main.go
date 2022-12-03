package main

import (
	"fmt"
	lemin "lemin/packages"
)

func main() {

	rooms := lemin.Rooms()

	for i := 0; i < len(rooms); i++ {
		fmt.Println(*rooms[i])
	}
}
