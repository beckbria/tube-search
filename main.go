package tube_search

import (
	"fmt"
	"log"
)

func main() {
	lines, err := GetAllLines()
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		fmt.Println(l.Id)
	}
}
