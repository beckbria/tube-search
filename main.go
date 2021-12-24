package main

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
		stops, err := GetStopsForLine(l.Id)
		if err != nil {
			log.Fatal(err)
		}
		for _, s := range stops {
			fmt.Printf("\t%s -- %s\n", s.Id, s.CommonName)
		}
	}
}
