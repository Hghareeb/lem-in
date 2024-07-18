package functions

/*
import (
	"errors"
	"fmt"
)

type Path struct {
	Rooms []*Room
}

// FindAllPaths finds all possible paths from start to end.
func FindAllPaths(farm *Farm) ([]*Path, error) {
	if farm.StartRoom == nil || farm.EndRoom == nil {
		return nil, errors.New("start or end room not defined")
	}

	var allPaths []*Path
	var currentPath []*Room
	visited := make(map[*Room]bool)

	// Depth-first search function to find all possible paths
	var dfs func(*Room)
	dfs = func(room *Room) {
		fmt.Printf("Visiting room: %s\n", room.RoomName)
		if room == farm.EndRoom {
			fmt.Println("Found a path to the end room!")
			path := make([]*Room, len(currentPath))
			copy(path, currentPath)
			allPaths = append(allPaths, &Path{Rooms: path})
			fmt.Printf("Added path: %v\n", path)
			return
		}
		// Iterate through the links of the current room and explore unvisited rooms
		for _, link := range room.Links {
			nextRoom := link.Room2
			if nextRoom == room {
				nextRoom = link.Room1
			}
			// If the next room has not been visited, explore it
			if !visited[nextRoom] {
				fmt.Printf("Exploring unvisited room: %s\n", nextRoom.RoomName)
				visited[nextRoom] = true
				currentPath = append(currentPath, nextRoom)
				dfs(nextRoom)
				currentPath = currentPath[:len(currentPath)-1]
				visited[nextRoom] = false
			}
		}
	}
	// Start the search from the start room
	fmt.Printf("Starting search from room: %s\n", farm.StartRoom.RoomName)
	visited[farm.StartRoom] = true
	currentPath = append(currentPath, farm.StartRoom)
	dfs(farm.StartRoom)

	if len(allPaths) == 0 {
		return nil, errors.New("no paths found between start and end rooms")
	}

	fmt.Println("All paths found:")
	for _, path := range allPaths {
		fmt.Printf("%v\n", path.Rooms)
	}

	return allPaths, nil
}
/*