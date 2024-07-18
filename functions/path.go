package functions

import (
	"errors"
)

type Path struct {
	Rooms []*Room
	Skip  bool
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
		if room == farm.EndRoom {
			path := make([]*Room, len(currentPath))
			copy(path, currentPath)
			allPaths = append(allPaths, &Path{Rooms: path})
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
				visited[nextRoom] = true
				currentPath = append(currentPath, nextRoom)
				dfs(nextRoom)
				currentPath = currentPath[:len(currentPath)-1]
				visited[nextRoom] = false
			}
		}
	}
	// Start the search from the start room
	visited[farm.StartRoom] = true
	currentPath = append(currentPath, farm.StartRoom)
	dfs(farm.StartRoom)

	if len(allPaths) == 0 {
		return nil, errors.New("no paths found between start and end rooms")
	}

	applyFindMaxFlow(allPaths)

	return allPaths, nil
}

// applyFindMaxFlow skips paths with maximum similarity to optimize flow
func applyFindMaxFlow(paths []*Path) {
	linkedTo := make([][]int, len(paths))
	for i := range linkedTo {
		linkedTo[i] = make([]int, len(paths))
	}
	for i := range paths {
		if paths[i].Skip {
			continue
		}
		for j := i + 1; j < len(paths); j++ {
			if paths[j].Skip {
				continue
			}
			if numOfSameRooms(paths[i].Rooms, paths[j].Rooms) > 2 {
				linkedTo[i][j] = 1
				linkedTo[j][i] = 1
			}
		}
	}
	maxSimilarity, maxPath := 0, -1
	for i := len(linkedTo) - 1; i >= 0; i-- {
		sumConnections := 0
		for _, conn := range linkedTo[i] {
			sumConnections += conn
		}
		if sumConnections > maxSimilarity {
			maxSimilarity = sumConnections
			maxPath = i
		}
	}
	if maxSimilarity != 0 {
		paths[maxPath].Skip = true
		applyFindMaxFlow(paths)
	}
}

// numOfSameRooms calculates the number of same rooms between two paths
func numOfSameRooms(rooms1, rooms2 []*Room) int {
	count := 0
	roomSet := make(map[string]bool)
	for _, room := range rooms1 {
		roomSet[room.RoomName] = true
	}
	for _, room := range rooms2 {
		if roomSet[room.RoomName] {
			count++
		}
	}
	return count
}
