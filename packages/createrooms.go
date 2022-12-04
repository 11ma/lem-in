package lemin

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name    string
	Xcord   int
	Ycord   int
	Visited bool
	Links   []*Room
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getRoomsAndLinks(data []string) (roomsAndCoords, links []string) {
	for i := 0; i < len(data); i++ {

		if len(strings.Fields(data[i])) == 3 {
			roomsAndCoords = append(roomsAndCoords, data[i])
		}

		if strings.Contains(data[i], "-") {
			links = append(links, data[i])
		}
	}
	return roomsAndCoords, links
}

func createRooms(roomsAndCoords, links []string) (rooms []*Room, mapRoom map[string]*Room) {
	mapRoom = map[string]*Room{}

	for i := 0; i < len(roomsAndCoords); i++ {
		splitRoomsAndCoords := strings.Split(roomsAndCoords[i], " ")
		roomName := splitRoomsAndCoords[0]
		xcord, err := strconv.Atoi(string(splitRoomsAndCoords[1]))
		handleErr(err)
		ycord, err := strconv.Atoi(string(splitRoomsAndCoords[2]))
		handleErr(err)
		room := Room{
			Name:    roomName,
			Xcord:   xcord,
			Ycord:   ycord,
			Visited: false,
		}

		mapRoom[roomName] = &room
		rooms = append(rooms, &room)
	}
	return rooms, mapRoom
}

func addLinkToRooms(rooms []*Room, links []string, mapRoom map[string]*Room) []*Room {
	// check if the room name matches the link and add it to the corresponding room
	for i := 0; i < len(links); i++ {
		for j := 0; j < len(rooms); j++ {
			splitLinks := strings.Split(links[i], "-")
			if rooms[j].Name == splitLinks[0] {
				room := mapRoom[splitLinks[1]]

				// append the second room the first room's links of rooms
				rooms[j].Links = append(rooms[j].Links, room)

				// also append the first the room to the second room's links of rooms
				room.Links = append(room.Links, rooms[j])
			}
		}
	}
	return rooms
}

func startAndEndRooms(data []string, mapRoom map[string]*Room) (startingRoom, endRoom *Room) {

	for i := 0; i < len(data); i++ {
		if data[i] == "##start" {
			startingRoom = mapRoom[string(data[i+1][0])]
		}

		if data[i] == "##end" {
			endRoom = mapRoom[string(data[i+1][0])]
		}
	}

	return startingRoom, endRoom
}

func Rooms() (numberOfAnts int, linkedRooms []*Room, startRoom, endRoom *Room) {

	data := ReadFile(os.Args[1])

	numberOfAnts, err := strconv.Atoi(data[0])
	handleErr(err)

	roomsAndCoords, links := getRoomsAndLinks(data)
	rooms, mapRoom := createRooms(roomsAndCoords, links)
	linkedRooms = addLinkToRooms(rooms, links, mapRoom)

	startRoom, endRoom = startAndEndRooms(data, mapRoom)

	return numberOfAnts, linkedRooms, startRoom, endRoom
}
