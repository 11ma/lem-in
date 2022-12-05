package lemin

import (
	"reflect"
)

func IsRoomContainedInRoute(element Room, slice []Room) bool {
	for _, potentialelement := range slice {
		if element.Name == potentialelement.Name {
			return true
		}
	}
	return false
}

func IsRouteConatinedInRoutes(route []Room, routes [][]Room) bool {
	for _, potentialroute := range routes {
		if reflect.DeepEqual(route, potentialroute) {
			return true
		}
	}
	return false
}

// takes in a list of routes and returns the same list sorted by length from shortest to longest
func RouteSorter(routes [][]Room) (routesSortedByLength [][]Room) {
	noOfRoutes := len(routes)
	for length := 0; len(routesSortedByLength) < noOfRoutes; length++ {
		for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
			currentRoute := routes[routeIndex]
			lengthOfRoute := len(currentRoute)
			if lengthOfRoute == length {
				routesSortedByLength = append(routesSortedByLength, currentRoute)
			}
		}
	}
	return routesSortedByLength
}

// takes a list of routes and returns a list of routes with routes that go through the same middle rooms removed, preferring to keep the shortest route
func RemoveDuplicates(existingRoutes [][]Room) [][]Room {
	for route1Index := 0; route1Index < len(existingRoutes); route1Index++ {
		route1 := existingRoutes[route1Index]
		for route2Index := 0; route2Index < len(existingRoutes); route2Index++ {
			route2 := existingRoutes[route2Index]
			for roomIndex := 1; roomIndex < len(route1)-1; roomIndex++ {
				room := route1[roomIndex]
				if IsRoomContainedInRoute(room, route2) && route2Index != route1Index && len(route1) <= len(route2) {
					existingRoutes = append(existingRoutes[:route2Index], existingRoutes[route2Index+1:]...)
					route2Index--
					if route2Index < route1Index {
						route1Index--
					}
					break
				}
			}
		}
	}
	return existingRoutes
}

// finds a route that hasn't already been found and put in existingRoutes if possible, or returns an empty route
func FindRoute(startingRoom Room, endingRoom Room, allRooms []Room, existingRoute []Room, existingRoutes *[][]Room) (routeNames []Room) {
	existingRoute = append(existingRoute, startingRoom)
	if startingRoom.Name == endingRoom.Name {
		return existingRoute
	}

	for i := 0; i < len(startingRoom.Links); i++ {
		currentLinkedRoom := startingRoom.Links[i]

		if !IsRoomContainedInRoute((*currentLinkedRoom), existingRoute) && len(FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute, existingRoutes)) != 0 {
			potentialRoute := FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute, existingRoutes)
			if !IsRouteConatinedInRoutes(potentialRoute, *existingRoutes) && len(potentialRoute) != 0 {

				*existingRoutes = append(*existingRoutes, potentialRoute)
				existingRoute = potentialRoute

				return existingRoute
			}
		}
	}

	if existingRoute[len(existingRoute)-1].Name == endingRoom.Name {
		return existingRoute
	}

	return []Room{}
}

// finds all possible valid routes, and sorts them from shortest to longest
func FindAllRoutes(startingRoom Room, endingRoom Room, allRooms []Room, existingRoutes [][]Room) (allRoutesNames [][]Room) {
	potentialRoute := FindRoute(startingRoom, endingRoom, allRooms, []Room{}, &existingRoutes)
	for len(potentialRoute) != 0 {
		if !IsRouteConatinedInRoutes(potentialRoute, existingRoutes) {
			existingRoutes = append(existingRoutes, potentialRoute)
		}
		potentialRoute = FindRoute(startingRoom, endingRoom, allRooms, []Room{}, &existingRoutes)
	}
	return RouteSorter(RemoveDuplicates(existingRoutes))
}
