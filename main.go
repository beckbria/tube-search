package main

import (
	"fmt"
	"log"
)

func DebugLevel() uint {
	// 0: Disable debugging
	// 1: Logging
	return 10
}

func main() {
	m := MakeStationMap()
	lines, err := GetAllLines()
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		if DebugLevel() >= 1 {
			fmt.Println(l.Id)
		}
		err = m.AddLine(&l)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, st := range m.Stations {
		fmt.Printf("%s (%s) -> [", st.Id, st.Name)
		for _, adj := range st.Adj {
			fmt.Printf("%s (%s),", adj.Dest.Id, adj.Dest.Name)
		}
		fmt.Print("]\n")
	}
}
