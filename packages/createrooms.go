package lemin

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name  string
	Xcord int
	Ycord int
	Links []*Room
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetRoomsAndLinks(data []string) (roomsAndCoords, links []string) {
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

func CreateRooms(roomsAndCoords, links []string) (rooms []*Room, mapRoom map[string]*Room) {
	mapRoom = map[string]*Room{}

	for i := 0; i < len(roomsAndCoords); i++ {
		splitRoomsAndCoords := strings.Split(roomsAndCoords[i], " ")
		roomName := splitRoomsAndCoords[0]
		xcord, err := strconv.Atoi(string(splitRoomsAndCoords[1]))
		handleErr(err)
		ycord, err := strconv.Atoi(string(splitRoomsAndCoords[2]))
		handleErr(err)
		room := Room{
			Name:  roomName,
			Xcord: xcord,
			Ycord: ycord,
		}

		mapRoom[roomName] = &room
		rooms = append(rooms, &room)
	}
	return rooms, mapRoom
}

func AddLinkToRooms(rooms []*Room, links []string, mapRoom map[string]*Room) []*Room {

	for i := 0; i < len(links); i++ {
		for j := 0; j < len(rooms); j++ {
			splitLinks := strings.Split(links[i], "-")
			// check if the room name matches the first link value
			if rooms[j].Name == splitLinks[0] {
				room := mapRoom[splitLinks[1]]
				rooms[j].Links = append(rooms[j].Links, room)
			}
		}
	}
	return rooms
}

func StartAndEndRooms(data []string, mapRoom map[string]*Room) (startingRoom, endRoom *Room) {

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

func Rooms() (linkedRooms []*Room, startRoom, endRoom *Room) {

	data := ReadFile(os.Args[1])

	roomsAndCoords, links := GetRoomsAndLinks(data)
	rooms, mapRoom := CreateRooms(roomsAndCoords, links)
	linkedRooms = AddLinkToRooms(rooms, links, mapRoom)

	startRoom, endRoom = StartAndEndRooms(data, mapRoom)

	return linkedRooms, startRoom, endRoom
}
