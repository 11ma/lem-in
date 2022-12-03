package main

import (
	"fmt"
	lemin "lemin/packages"
)

func main() {

	rooms, start, end := lemin.Rooms()

	// for i := 0; i < len(rooms); i++ {
	// 	fmt.Println(*rooms[i])
	// }

	fmt.Println("rooms:", rooms)
	fmt.Println("start:", start)
	fmt.Println("end:", end)
}
