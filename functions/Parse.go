package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// File parses the input file to create a colony.
func File(filename string) *Colony {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format, unable to open file")
		return nil
	}
	defer f.Close()

	var colony Colony
	colony.AntPositions = make(map[int]*Room)
	lineNum := 1

	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		colony.TotalAnts, err = strconv.Atoi(scanner.Text())
		if err != nil || colony.TotalAnts <= 0 {
			fmt.Println("ERROR: invalid data format, invalid number of ants")
			return nil
		}
		lineNum++
	}
	// Read the rest of the file
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		switch {
		case len(fields) == 0:
			fmt.Printf("Warning: Empty line on line %d: %s\n", lineNum, line)
		case len(fields) == 1:
			handleSingleField(&colony, line, lineNum)
		case len(fields) == 3:
			room := NewRoom(line, lineNum)
			if room != nil {
				colony.Rooms = append(colony.Rooms, room)
			} else {
				return nil // Terminate if there's an error creating a room
			}
		default:
			fmt.Printf("ERROR: invalid data format on line %d: %s\n", lineNum, line)
			return nil
		}
		lineNum++
	}
	// Check for errors reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR: invalid data format, error reading file")
		return nil
	}

	// Check for duplicate coordinates
	if !CheckDuplicateCoordinates(colony.Rooms) {
		return nil
	}

	for _, link := range colony.Links {
		link.Room1.Links = append(link.Room1.Links, link)
		link.Room2.Links = append(link.Room2.Links, link)
	}

	// Set the StartRoom and EndRoom after reading all rooms
	for _, room := range colony.Rooms {
		if colony.StartRoomLine > 0 && colony.StartRoomLine == room.LineNum {
			colony.StartRoom = room
			room.IsStart = true
		}
		if colony.EndRoomLine > 0 && colony.EndRoomLine == room.LineNum {
			colony.EndRoom = room
			room.IsEnd = true
		}
	}

	if colony.StartRoom == nil || colony.EndRoom == nil {
		fmt.Println("ERROR: invalid data format, no start or end room found")
		return nil
	}

	return &colony
}

// handleSingleField processes lines with a single field.
func handleSingleField(colony *Colony, line string, lineNum int) {
	switch line {
	case "##start":
		colony.StartRoomLine = lineNum + 1 // Set StartRoomLine to the next line
	case "##end":
		colony.EndRoomLine = lineNum + 1 // Set EndRoomLine to the next line
	default:
		link := NewLink(line, lineNum, colony)
		if link != nil {
			colony.Links = append(colony.Links, link)
		}
	}
}

// NewRoom creates a new room from a line in the file.
func NewRoom(line string, lineNum int) *Room {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		fmt.Printf("ERROR: invalid data format on line %d: %s\n", lineNum, line)
		return nil
	}

	name := fields[0]
	coordX := fields[1]
	coordY := fields[2]

	// Validate coordinates
	if !isNumber(coordX) || !isNumber(coordY) {
		fmt.Printf("ERROR: invalid data format on line %d: coordinates must be numbers: %s\n", lineNum, line)
		return nil
	}

	newRoom := &Room{
		Name:    name,
		CoordX:  coordX,
		CoordY:  coordY,
		Visited: false,
		IsEnd:   false,
		IsStart: false,
		Links:   make([]*Link, 0),
		LineNum: lineNum, // Set the LineNum field
	}

	return newRoom
}

// isNumber checks if a string is a valid number.
func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

// NewLink creates a new link between two rooms from a line in the file.
func NewLink(line string, lineNum int, colony *Colony) *Link {
	linkSplit := strings.Split(line, "-")
	if len(linkSplit) != 2 {
		fmt.Printf("ERROR: invalid data format on line %d: %s\n", lineNum, line)
		return nil
	}

	roomName1 := linkSplit[0]
	roomName2 := linkSplit[1]

	room1 := findRoomByName(roomName1, colony.Rooms)
	room2 := findRoomByName(roomName2, colony.Rooms)

	if room1 == nil || room2 == nil {
		fmt.Printf("ERROR: invalid data format, invalid room name(s) on line %d: %s\n", lineNum, line)
		return nil
	} else if room1 == room2 {
		fmt.Printf("ERROR: invalid data format, link to itself on line %d: %s\n", lineNum, line)
		return nil
	}

	newLink := &Link{
		Room1: room1,
		Room2: room2,
	}

	return newLink
}

// findRoomByName finds a room by its name.
func findRoomByName(name string, rooms []*Room) *Room {
	for _, room := range rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

// CheckDuplicateCoordinates checks for duplicate coordinates among the rooms.
func CheckDuplicateCoordinates(rooms []*Room) bool {
	coordinates := make(map[string]int)
	for _, room := range rooms {
		coord := room.CoordX + "," + room.CoordY
		if lineNum, exists := coordinates[coord]; exists {
			if lineNum == room.LineNum {
				fmt.Printf("ERROR: invalid data format, duplicate coordinates on line %d\n", room.LineNum)
				return false
			}
			fmt.Println("ERROR: invalid data format, duplicate coordinates")
			return false
		}
		coordinates[coord] = room.LineNum
	}
	return true
}
