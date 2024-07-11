package functions

import (
	"fmt"
	"sort"
)

// MoveAnts simulates the movement of ants along the paths
func MoveAnts(farm *Farm, paths []*Path) {
	numAnts := len(farm.Ants)
	antPositions := make([]int, numAnts) // Tracks the position of each ant on its path
	antPaths := make([]int, numAnts)     // Tracks the path assigned to each ant
	roomOccupancy := make(map[string]int) // Tracks the current ant occupying each room

	// Sort paths by their lengths (shortest first) to distribute ants evenly
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i].Rooms) < len(paths[j].Rooms)
	})

	// Calculate how many ants each path should get 
	pathAntsCount := make([]int, len(paths))
	for i := 0; i < numAnts; i++ {
		minIndex := 0
		minSteps := pathAntsCount[minIndex] + len(paths[minIndex].Rooms)
		for j := 1; j < len(paths); j++ {
			steps := pathAntsCount[j] + len(paths[j].Rooms)
			if steps < minSteps {
				minIndex = j
				minSteps = steps
			}
		}
		pathAntsCount[minIndex]++
		antPaths[i] = minIndex
	}
    // Here is a reference snippet of code from extra date/ant2.go:
	for step := 0; ; step++ {
		movements := []string{}
		allFinished := true
		roomOccupancy = make(map[string]int) // Reset room occupancy each turn

		for antIndex := 0; antIndex < numAnts; antIndex++ {
			path := paths[antPaths[antIndex]].Rooms
			currentPos := antPositions[antIndex]
            
			if currentPos < len(path)-1 {
				nextRoom := path[currentPos+1]

				// Check if the next room is free (or is the end room)
				if roomOccupancy[nextRoom.RoomName] == 0 || nextRoom.RoomName == farm.EndRoom.RoomName {
					// Move the ant to the next room
					antPositions[antIndex]++
					movements = append(movements, fmt.Sprintf("L%d-%s", antIndex+1, nextRoom.RoomName))
					roomOccupancy[nextRoom.RoomName] = antIndex + 1
					allFinished = false
				}
			}
		}
        
		if allFinished {
			break
		}

		fmt.Println(movements)
	}
}