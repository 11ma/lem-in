package main

import (
	"fmt"
	lemin "lemin/packages"
)

func createGraph() (graph map[string][]string) {
	graph = map[string][]string{}
	_, rooms, _, _ := lemin.Rooms()
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i].Links); j++ {
			graph[rooms[i].Name] = append(graph[rooms[i].Name], rooms[i].Links[j].Name)
		}

	}
	return graph
}

func main() {

	// numberOfAnts, rooms, mapRoom, start, end := lemin.Rooms()

	// print al rooms are connected
	// fmt.Println("number of ants:", numberOfAnts)
	// fmt.Println("rooms:", rooms)
	// fmt.Println("start:", start.Name)
	// fmt.Println("end:", end.Name)

	graph := createGraph()
	fmt.Println(graph)

	// find all paths
	// from 1 how do i get to 0 with the current links?
	// starting room has a link to 3,
	// we need to check the rooms linked to 3 to see if it is 0 if not keep checking

	// route := []string{}

	// route = append(route, getRoutes(start, end)[0])

	// fmt.Println(route)
	// getRoutes(start, end)
}

func getRoutes(graph map[string][]string, start, end *lemin.Room) {

	routeSoFar := []string{start.Name}
	startIndex := 0
	endIndex := len(start.Links)

	for startIndex < endIndex {
		if start.Links[startIndex].Name != end.Name && !start.Links[startIndex].Visited {
			start.Links[startIndex].Visited = true
			routeSoFar = append(routeSoFar, start.Links[startIndex].Name)
		}
		startIndex++
	}
	fmt.Println(routeSoFar)
	// return routeSoFar
}
