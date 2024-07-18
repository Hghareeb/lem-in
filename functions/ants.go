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
	roomOccupancy := make(map[string]bool) // Tracks if a room is occupied

	// Filter out skipped paths
	validPaths := []*Path{}
	for _, path := range paths {
		if !path.Skip {
			validPaths = append(validPaths, path)
		}
	}

	// Sort paths by their lengths (shortest first) to distribute ants evenly
	sort.Slice(validPaths, func(i, j int) bool {
		return len(validPaths[i].Rooms) < len(validPaths[j].Rooms)
	})

	// Calculate the optimal number of ants for each path
	totalTurns := make([]int, len(validPaths))
	for i := 0; i < numAnts; i++ {
		// Assign ants to paths in a way that balances the total turns across paths
		minTurns := totalTurns[0] + len(validPaths[0].Rooms)
		pathIndex := 0
		for j := 1; j < len(validPaths); j++ {
			turns := totalTurns[j] + len(validPaths[j].Rooms)
			if turns < minTurns {
				minTurns = turns
				pathIndex = j
			}
		}
		totalTurns[pathIndex] += len(validPaths[pathIndex].Rooms)
		antPaths[i] = pathIndex
	}

	for step := 0; ; step++ {
		movements := []string{}
		allFinished := true
		roomOccupancy = make(map[string]bool) // Reset room occupancy each turn

		for antIndex := 0; antIndex < numAnts; antIndex++ {
			path := validPaths[antPaths[antIndex]].Rooms
			currentPos := antPositions[antIndex]

			if currentPos < len(path)-1 {
				nextRoom := path[currentPos+1]

				// Check if the next room is free (or is the end room)
				if !roomOccupancy[nextRoom.RoomName] || nextRoom.RoomName == farm.EndRoom.RoomName {
					// Move the ant to the next room
					antPositions[antIndex]++
					movements = append(movements, fmt.Sprintf("L%d-%s", antIndex+1, nextRoom.RoomName))
					roomOccupancy[nextRoom.RoomName] = true
					if currentPos > 0 {
						// Free the previous room
						roomOccupancy[path[currentPos].RoomName] = false
					}
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
