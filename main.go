package main

import (
	"fmt"
	lemin "lemin/packages"
)

func main() {

	numberOfAnts, rooms, start, end := lemin.Rooms()
	// numberOfAnts, rooms, _, _ := lemin.Rooms()

	fmt.Println("number of ants:", numberOfAnts)
	fmt.Println("rooms:", rooms)
	// fmt.Println("start:", start.Name)
	// fmt.Println("end:", end.Name)

	// test all links are connected
	// for i := 0; i < len(rooms); i++ {
	// 	fmt.Println("Room name:", rooms[i].Name)

	// 	for j := 0; j < len(rooms[i].Links); j++ {
	// 		fmt.Println("Link:", rooms[i].Links[j].Name)
	// 	}
	// }

	// find all paths
	// from 1 how do i get to 0 with the current links?
	// starting room has a link to 3,
	// we need to check the rooms linked to 3 to see if it is 0 if not keep checking

	// route := []string{}

	// route = append(route, getRoutes(start, end)[0])

	// fmt.Println(route)
	getRoutes(start, end)
}

func getRoutes(start, end *lemin.Room) {

	routeSoFar := []string{}

	for i := 0; i < len(start.Links); {

		if start.Links[i].Name == end.Name {
			routeSoFar = append(routeSoFar, start.Links[i].Name)
		} else if !start.Links[i].Visited {
			start.Links[i].Visited = true
			routeSoFar = append(routeSoFar, start.Links[i].Name)
			i++
		}
	}

	fmt.Println(routeSoFar)
	// return routeSoFar
}
