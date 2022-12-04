package main

import (
	"fmt"
	lemin "lemin/packages"
)

func main() {

	// numberOfAnts, rooms, start, end := lemin.Rooms()
	numberOfAnts, rooms, _, _ := lemin.Rooms()

	fmt.Println("number of ants:", numberOfAnts)
	fmt.Println("rooms:", rooms)
	// fmt.Println("start:", start.Name)
	// fmt.Println("end:", end.Name)

	for i := 0; i < len(rooms); i++ {
		fmt.Println("Room name:", rooms[i].Name, "Links:", rooms[i].Links)
	}

	// find all valid paths
	// from 1 how do i get to 0 with the current links?
}

func getRoute(start, end *lemin.Room) []string {

	routeSoFar := []string{}
	// starting room has a link to 3,
	// we need to check the rooms linked to 3 to see if it is 0 if not keep checking
	for i := 0; i < len(start.Links); i++ {
		if start.Links[i].Name != end.Name {
			routeSoFar = append(routeSoFar, start.Links[i].Name)
		} else if i == len(start.Links)-1 {
			routeSoFar = append(routeSoFar, start.Links[i].Name)
			getRoute(start.Links[i], end)
		}
	}
	return routeSoFar
}
