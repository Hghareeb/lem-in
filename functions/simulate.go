package functions

import (
    "fmt"
    "strings"
)

// SimulateAntsMovement simulates the movement of ants within the colony.
func SimulateAntsMovement(colony *Colony) []string {
    var movements []string
    for !allAntsArrived(colony) { // while there are ants not arrived, if all arrived it will stop
        var currentMove []string
        for _, route := range colony.Paths {
            currentMove = append(currentMove, moveAntsOnRoute(route, colony.EndRoom)...)
        }
        if len(currentMove) > 0 {
            movements = append(movements, strings.Join(currentMove, " "))
        }
    }
    return movements
}

// allAntsArrived checks if all ants have arrived at the end room.
func allAntsArrived(colony *Colony) bool {
    for _, route := range colony.Paths {
        if len(route.Ants) > 0 {
            return false
        }
    }
    return true // if no ants, all of them arrived 
}

// moveAntsOnRoute moves ants on a specific route.
func moveAntsOnRoute(route *Route, endRoom *Room) []string {
    var moves []string
    prefix := "L" 
    for i := len(route.Rooms) - 1; i > 0; i-- {
        currentRoom := route.Rooms[i]
        previousRoom := route.Rooms[i-1]
        for j := 0; j < len(route.Ants); j++ {
            ant := route.Ants[j]
            if ant.CurrentRoom == previousRoom { // if the ant's current room is the previous room
                ant.CurrentRoom = currentRoom // we move it to current room
                moves = append(moves, fmt.Sprintf("%s%d-%s", prefix, ant.ID, currentRoom.Name)) // Use the prefix here
                if currentRoom == endRoom {
                    route.Ants = append(route.Ants[:j], route.Ants[j+1:]...)
                    j--
                }
                break
            }
        }
    }
    return moves
}
