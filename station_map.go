package main

import "fmt"

type Station struct {
	Id, Name string
	Adj      []Adjacency
}

type Adjacency struct {
	Route string
	Dest  *Station
}

type StationMap struct {
	// A map from the station ID i.e. 940GZZLUSKS to the station record
	Stations map[string]*Station
}

func MakeStationMap() *StationMap {
	m := StationMap{}
	m.Stations = make(map[string]*Station)
	return &m
}

func (m *StationMap) AddStation(s *Tfl20) {
	if DebugLevel() >= 5 {
		fmt.Printf("Adding Station (Tfl20) %s (%s)\n", s.Id, s.Name)
	}
	if current, found := m.Stations[s.Id]; found {
		if DebugLevel() >= 1 {
			fmt.Printf("Duplicate found.  Was: %s/%s, Is: %s/%s\n", current.Id, current.Name, s.Id, s.Name)
		}
		return
	}

	st := Station{Id: s.Id, Name: s.Name, Adj: make([]Adjacency, 0)}
	m.Stations[s.Id] = &st
}

func (m *StationMap) AddStation11(s *Tfl11) {
	if DebugLevel() >= 5 {
		fmt.Printf("Adding Station (Tfl11) %s (%s)\n", s.Id, s.CommonName)
	}
	if current, found := m.Stations[s.Id]; found {
		if DebugLevel() >= 1 {
			fmt.Printf("Duplicate found.  Was: %s/%s, Is: %s/%s\n", current.Id, current.Name, s.Id, s.CommonName)
		}
		return
	}

	st := Station{Id: s.Id, Name: s.CommonName, Adj: make([]Adjacency, 0)}
	m.Stations[s.Id] = &st
}

func (m *StationMap) AddAdjacency(fromId, toId, name string) error {
	if DebugLevel() >= 10 {
		fmt.Printf("Adding Adjacency from %s to %s along %s\n", fromId, toId, name)
	}
	from, foundFrom := m.Stations[fromId]
	to, foundTo := m.Stations[toId]
	if !foundFrom || !foundTo {
		return fmt.Errorf("Missing ID: %s (%t) or %s (%t)", fromId, foundFrom, toId, foundTo)
	}
	from.Adj = append(from.Adj, Adjacency{Route: name, Dest: to})
	return nil
}

func (m *StationMap) AddLine(line *Tfl19) error {
	if DebugLevel() >= 4 {
		fmt.Printf("Adding %s line\n", line.Id)
	}
	route, err := GetLineSequence(line.Id)
	if err != nil {
		return err
	}
	for _, s := range route.Stations {
		m.AddStation(&s)
	}
	// The route's station list seems to use HUB*** IDs rather than the
	// Naplan IDs used in the ordered line route.  Query the full list of
	// Stops and add them to list of known stops first.
	stops, err := GetStopsForLine(line.Id)
	if err != nil {
		return err
	}
	for _, s := range stops {
		m.AddStation11(&s)
	}

	for _, r := range route.OrderedLineRoutes {
		for i := 0; i < (len(r.NaptanIds) - 1); i++ {
			err = m.AddAdjacency(r.NaptanIds[i], r.NaptanIds[i+1], r.Name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
