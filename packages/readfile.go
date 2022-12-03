package lemin

import (
	"log"
	"os"
	"strings"
)

func ReadFile(file string) []string {
	readFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(readFile), "\n")

	return data
}
