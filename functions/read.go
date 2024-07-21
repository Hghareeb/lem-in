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

	// Create a map to check for duplicate coordinates
	coordsMap := make(map[string]bool)

	// Read rooms and links from the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line == "##start" {
			if !scanner.Scan() {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			startRoom, err := parseRoom(scanner.Text())
			if err != nil {
				return nil, err
			}
			startRoom.Start = true
			if coordsMap[fmt.Sprintf("%d,%d", startRoom.CoordX, startRoom.CoordY)] {
				return nil, fmt.Errorf("ERROR: duplicate coordinates for rooms")
			}
			coordsMap[fmt.Sprintf("%d,%d", startRoom.CoordX, startRoom.CoordY)] = true
			farm.StartRoom = startRoom
			farm.Rooms = append(farm.Rooms, startRoom)
		} else if line == "##end" {
			if !scanner.Scan() {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			endRoom, err := parseRoom(scanner.Text())
			if err != nil {
				return nil, err
			}
			endRoom.End = true
			if coordsMap[fmt.Sprintf("%d,%d", endRoom.CoordX, endRoom.CoordY)] {
				return nil, fmt.Errorf("ERROR: duplicate coordinates for rooms")
			}
			coordsMap[fmt.Sprintf("%d,%d", endRoom.CoordX, endRoom.CoordY)] = true
			farm.EndRoom = endRoom
			farm.Rooms = append(farm.Rooms, endRoom)
		} else if strings.HasPrefix(line, "##") {
			// Error on unrecognized sections starting with ##
			return nil, fmt.Errorf("ERROR: invalid section %s", line)
		} else if strings.HasPrefix(line, "#") {
			// Skip lines starting with a single #
			continue
		} else if strings.Contains(line, "-") {
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
		} else {
			// Room definition
			room, err := parseRoom(line)
			if err != nil {
				return nil, err
			}
			if coordsMap[fmt.Sprintf("%d,%d", room.CoordX, room.CoordY)] {
				return nil, fmt.Errorf("ERROR: duplicate coordinates for rooms")
			}
			coordsMap[fmt.Sprintf("%d,%d", room.CoordX, room.CoordY)] = true
			farm.Rooms = append(farm.Rooms, room)
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

// mustAtoi converts a string to an integer and returns an error if it fails
func mustAtoi(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("ERROR: invalid data format")
	}
	return i, nil
}

// parseRoom parses a room definition from a string
func parseRoom(line string) (*Room, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return nil, fmt.Errorf("ERROR: invalid data format")
	}
	roomName := fields[0]
	if strings.HasPrefix(roomName, "L") || strings.HasPrefix(roomName, "#") || strings.Contains(roomName, " ") {
		return nil, fmt.Errorf("ERROR: invalid data format")
	}
	coordX, err := mustAtoi(fields[1])
	if err != nil {
		return nil, err
	}
	coordY, err := mustAtoi(fields[2])
	if err != nil {
		return nil, err
	}
	return &Room{
		RoomName: roomName,
		CoordX:   coordX,
		CoordY:   coordY,
	}, nil
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
