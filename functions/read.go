package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads the farm configuration from a file and returns a Farm struct.
func ReadFile(filename string) (*Farm, error) {
	// Open the file for reading
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	farm := &Farm{}
	scanner := bufio.NewScanner(f)

	// Read the number of ants
	if !scanner.Scan() {
		return nil, fmt.Errorf("ERROR: invalid data format")
	}
	numAnts, err := strconv.Atoi(scanner.Text())
	if err != nil || numAnts <= 0 {
		return nil, fmt.Errorf("ERROR: invalid data format")
	}
	farm.Ants = make([]*Ant, numAnts)
	for i := 0; i < numAnts; i++ {
		farm.Ants[i] = &Ant{AntID: i + 1}
	}

	// Read rooms and links from the file line by line 
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Check for start room 
		if line == "##start" {
			if !scanner.Scan() {
				return nil, fmt.Errorf("ERROR: invalid data format")
			} 
			farm.StartRoom = parseRoom(scanner.Text())
			if farm.StartRoom == nil {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			farm.StartRoom.Start = true
			farm.Rooms = append(farm.Rooms, farm.StartRoom)
		} else if line == "##end" { // Check for end room
			if !scanner.Scan() {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			farm.EndRoom = parseRoom(scanner.Text())
			if farm.EndRoom == nil {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			farm.EndRoom.End = true
			farm.Rooms = append(farm.Rooms, farm.EndRoom)
		} else if strings.Contains(line, "-") { // Check for links
			// Link definition
			fields := strings.Split(line, "-")
			if len(fields) != 2 {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			room1 := findRoom(farm, fields[0])
			room2 := findRoom(farm, fields[1])
			if room1 == nil || room2 == nil {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			if !areRoomsLinked(room1, room2) {
				link := &Link{Room1: room1, Room2: room2}
				farm.Links = append(farm.Links, link)
				room1.Links = append(room1.Links, link)
				room2.Links = append(room2.Links, link)
			} else {
				return nil, fmt.Errorf("ERROR: duplicate link between %s and %s", room1.RoomName, room2.RoomName)
			}
		} else { // Room definition 
			room := parseRoom(line)
			if room != nil {
				farm.Rooms = append(farm.Rooms, room)
			} else { 
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
		}
	}
	// Check for errors during file reading 
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Ensure start and end rooms are defined 
	if farm.StartRoom == nil || farm.EndRoom == nil {
		return nil, fmt.Errorf("ERROR: start or end room not defined")
	}

	return farm, nil
}

// mustAtoi converts a string to an integer and panics if it fails
func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// parseRoom parses a room definition from a string
func parseRoom(line string) *Room {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return nil
	}
	roomName := fields[0]
	if strings.HasPrefix(roomName, "L") || strings.HasPrefix(roomName, "#") || strings.Contains(roomName, " ") {
		return nil
	}
	return &Room{
		RoomName: roomName,
		CoordX:   mustAtoi(fields[1]),
		CoordY:   mustAtoi(fields[2]),
	}
}

// findRoom finds a room by name in the farm
func findRoom(farm *Farm, name string) *Room {
	for _, room := range farm.Rooms {
		if room.RoomName == name {
			return room
		}
	}
	return nil
}

// areRoomsLinked checks if two rooms are already linked
func areRoomsLinked(room1, room2 *Room) bool {
	for _, link := range room1.Links {
		if (link.Room1 == room1 && link.Room2 == room2) || (link.Room1 == room2 && link.Room2 == room1) {
			return true
		}
	}
	return false
}
