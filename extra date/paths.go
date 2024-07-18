package functions
/*

import (
	"errors"
)

type Path struct {
	Rooms []*Room
}

// FindAllPaths finds all possible paths from start to end.
func FindAllPaths(farm *Farm) ([]*Path, error) {
	// Check if start or end room is not defined
	if farm.StartRoom == nil || farm.EndRoom == nil {
		return nil, errors.New("start or end room not defined")
	}

	var allPaths []*Path
	var currentPath []*Room
	visited := make(map[*Room]bool)

	// Depth-first search function
	var dfs func(*Room)
	dfs = func(room *Room) {
		// If we reach the end room, add the current path to allPaths
		if room == farm.EndRoom {
			path := make([]*Room, len(currentPath))
			copy(path, currentPath)
			allPaths = append(allPaths, &Path{Rooms: path})
			return
		}

		// Iterate through the links of the current room
		for _, link := range room.Links {
			nextRoom := link.Room2
			if nextRoom == room {
				nextRoom = link.Room1
			}

			// If the next room has not been visited, explore it
			if !visited[nextRoom] {
				visited[nextRoom] = true
				currentPath = append(currentPath, nextRoom)
				dfs(nextRoom)
				currentPath = currentPath[:len(currentPath)-1]
				visited[nextRoom] = false
			}
		}
	}

	visited[farm.StartRoom] = true
	currentPath = append(currentPath, farm.StartRoom)
	dfs(farm.StartRoom)

	// If no paths are found between start and end rooms, return an error
	if len(allPaths) == 0 {
		return nil, errors.New("no paths found between start and end rooms")
	}

	return allPaths, nil
}
*/
// Here is a reference snippet of code from functions/read.go:
