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

	// Read number of ants
	if scanner.Scan() {
		numAnts, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("invalid number of ants: %v", err)
		}
		farm.Ants = make([]*Ant, numAnts)
		for i := 0; i < numAnts; i++ {
			farm.Ants[i] = &Ant{AntID: i + 1}
		}
	}

	// Read rooms and links
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line == "##start" {
			if scanner.Scan() {
				farm.StartRoom = parseRoom(scanner.Text())
				farm.StartRoom.Start = true
				farm.Rooms = append(farm.Rooms, farm.StartRoom)
			}
		} else if line == "##end" {
			if scanner.Scan() {
				farm.EndRoom = parseRoom(scanner.Text())
				farm.EndRoom.End = true
				farm.Rooms = append(farm.Rooms, farm.EndRoom)
			}
		} else if strings.Contains(line, "-") {
			// Link definition
			fields := strings.Split(line, "-")
			if len(fields) != 2 {
				return nil, fmt.Errorf("invalid link definition: %s", line)
			}
			room1 := findRoom(farm, fields[0])
			room2 := findRoom(farm, fields[1])
			if room1 != nil && room2 != nil {
				if !areRoomsLinked(room1, room2) {
					link := &Link{Room1: room1, Room2: room2}
					farm.Links = append(farm.Links, link)
					room1.Links = append(room1.Links, link)
					room2.Links = append(room2.Links, link)
				} else {
					return nil, fmt.Errorf("duplicate link between %s and %s", room1.RoomName, room2.RoomName)
				}
			} else {
				return nil, fmt.Errorf("one or both rooms in link not found: %s - %s", fields[0], fields[1])
			}
		} else {
			// Room definition
			room := parseRoom(line)
			if room != nil {
				farm.Rooms = append(farm.Rooms, room)
			} else {
				return nil, fmt.Errorf("invalid room definition: %s", line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return farm, nil
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parseRoom(line string) *Room {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		panic(fmt.Sprintf("invalid room definition: %s", line))
	}
	roomName := fields[0]
	if strings.HasPrefix(roomName, "L") || strings.HasPrefix(roomName, "#") || strings.Contains(roomName, " ") {
		panic(fmt.Sprintf("invalid room name: %s", roomName))
	}
	return &Room{
		RoomName: roomName,
		CoordX:   mustAtoi(fields[1]),
		CoordY:   mustAtoi(fields[2]),
	}
}

func findRoom(farm *Farm, name string) *Room {
	for _, room := range farm.Rooms {
		if room.RoomName == name {
			return room
		}
	}
	return nil
}

func areRoomsLinked(room1, room2 *Room) bool {
	for _, link := range room1.Links {
		if (link.Room1 == room1 && link.Room2 == room2) || (link.Room1 == room2 && link.Room2 == room1) {
			return true
		}
	}
	return false
}