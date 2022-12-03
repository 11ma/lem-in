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
	Links []Room
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

	var rooms []Room

	for i := 0; i < len(roomsAndCoords); i++ {
		splitRoomsAndCoords := strings.Split(roomsAndCoords[i], " ")
		roomName := splitRoomsAndCoords[0]
		xcord, err := strconv.Atoi(string(splitRoomsAndCoords[1]))
		ycord, err := strconv.Atoi(string(splitRoomsAndCoords[2]))
		handleErr(err)
		room := Room{
			Name:  roomName,
			Xcord: xcord,
			Ycord: ycord,
		}

		rooms = append(rooms, room)
	}

	// var roomLinks [][]string

	for i := 0; i < len(links); i++ {
		// fmt.Println(links[i])
		for j := 0; j < len(rooms); j++ {
			splitLinks := strings.Split(links[i], "-")
			// check if the room name matches the first link value
			if rooms[j].Name == splitLinks[0] {
				// fmt.Println(rooms[j].Name, splitLinks[0])
				connectedRoom := splitLinks[1]
				if rooms[j].Name == connectedRoom {
					rooms[j].Links = append(rooms[j].Links, rooms[j])
				}
			}
		}
	}

	// fmt.Println(roomLinks)
	fmt.Println(rooms)
}
