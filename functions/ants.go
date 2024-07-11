package functions

import (
    "fmt"
)

// MoveAnts simulates the movement of ants along the paths
func MoveAnts(farm *Farm, paths [][]*Room) {
    numAnts := len(farm.Ants)
    antPositions := make([]int, numAnts)  // Tracks the position of each ant on its path
    antPaths := make([]int, numAnts)      // Tracks the path assigned to each ant
    antInRoom := make(map[*Room]int)      // Tracks which ant is in which room

    // Assign paths to ants in a round-robin fashion and track their positions on each path and their rooms 
    pathIndex := 0
    for i := 0; i < numAnts; i++ {
        antPaths[i] = pathIndex
        pathIndex = (pathIndex + 1) % len(paths)
    }

    step := 0
    for {
        movements := []string{}
        occupiedRooms := make(map[*Room]bool)
        allFinished := true

        for antIndex := 0; antIndex < numAnts; antIndex++ {
            path := paths[antPaths[antIndex]]
            currentPos := antPositions[antIndex]

            if currentPos < len(path)-1 {
                nextRoom := path[currentPos+1]

                if !occupiedRooms[nextRoom] && (currentPos == 0 || antInRoom[path[currentPos]] == antIndex+1) {
                    if currentPos > 0 {
                        delete(antInRoom, path[currentPos])
                    }
                    occupiedRooms[nextRoom] = true
                    antInRoom[nextRoom] = antIndex + 1
                    antPositions[antIndex]++
                    movements = append(movements, fmt.Sprintf("L%d-%s", antIndex+1, nextRoom.RoomName))
                    allFinished = false
                } else {
                    occupiedRooms[path[currentPos]] = true
                }
            }
        }

        if allFinished {
            break
        }

        fmt.Println(movements)
        step++
    }
}