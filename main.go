package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// room is a linked list of room numbers
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

func main() {

	readFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(readFile), "\n")

	roomsAndCoords := []string{}

	links := []string{}

	for i := 0; i < len(data); i++ {

		if len(strings.Fields(data[i])) == 3 {
			roomsAndCoords = append(roomsAndCoords, data[i])
		}

		if strings.Contains(data[i], "-") {
			links = append(links, data[i])
		}
	}

	var rooms []*Room
	mapRoom := map[string]*Room{}
	var startingRoom *Room
	var endRoom *Room

	for i := 0; i < len(roomsAndCoords); i++ {
		splitRoomsAndCoords := strings.Split(roomsAndCoords[i], " ")
		roomName := splitRoomsAndCoords[0]
		xcord, err := strconv.Atoi(splitRoomsAndCoords[1])
		handleErr(err)
		ycord, err := strconv.Atoi(splitRoomsAndCoords[2])
		handleErr(err)
		room := Room{
			Name:  roomName,
			Xcord: xcord,
			Ycord: ycord,
		}

		mapRoom[roomName] = &room
		rooms = append(rooms, &room)
	}

	// get starting and ending rooms
	for i := 0; i < len(data); i++ {
		if data[i] == "##start" {
			// fmt.Println("you are the start:", data[i+1])
			startingRoom = mapRoom[string(data[i+1][0])]
		}

		if data[i] == "##end" {
			// fmt.Println("you are the start:", data[i+1])
			endRoom = mapRoom[string(data[i+1][0])]
		}
	}

	// check if the room name matches the first link value and add it to the room
	for i := 0; i < len(links); i++ {
		for j := 0; j < len(rooms); j++ {
			splitLinks := strings.Split(links[i], "-")
			if rooms[j].Name == splitLinks[0] {
				room := mapRoom[splitLinks[1]]
				rooms[j].Links = append(rooms[j].Links, room)
			}
		}
	}

	fmt.Println(startingRoom, endRoom)
	// for i := 0; i < len(rooms); i++ {
	// 	fmt.Println(*rooms[i])
	// }
}
