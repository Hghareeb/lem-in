package functions

import (
    "math"
    "sort"
    "fmt"
)

// Edmonds function finds all possible paths from the start room to the end room in the colony.
func Edmonds(colony *Colony) []*Route {
    start := []*Room{colony.StartRoom}
    end := colony.EndRoom
    queue := []*Route{{Rooms: start}} // queue a slice of route pointers with a route containing the start room
    var routes []*Route

    for len(queue) > 0 { // loop till the queue is empty
        route := queue[0] // gets first route
        queue = queue[1:] // updates the queue by removing the first one
        currentRoom := route.Rooms[len(route.Rooms)-1] // gets last room in the route

        if currentRoom == end { // checks if the room is the last room
            newRoute := &Route{Rooms: append([]*Room(nil), route.Rooms...)} // save the route to the struct
            routes = append(routes, newRoute)
            continue // skips to next iteration since a complete route has been found
        }

        for _, link := range currentRoom.Links { // finds all links from current room
            nextRoom := link.Room2 // gets the next room in the link
            if nextRoom == currentRoom { // make sure next room isn't the current room
                nextRoom = link.Room1
            }
            if !containsRoom(route.Rooms, nextRoom) { // ensure that the next room is not already in a route
                newRoute := &Route{Rooms: append(append([]*Room(nil), route.Rooms...), nextRoom)}
                queue = append(queue, newRoute)
            }
        }
    }
    return routes
}

// ChooseOptimalRoutes function selects the optimal routes from the list of all found routes.
func ChooseOptimalRoutes(routes []*Route, numAnts int) []*Route {
    applyMaxFlowAlgorithm(routes)

    var filteredRoutes []*Route
    for _, route := range routes {
        if !route.Skip {
            filteredRoutes = append(filteredRoutes, route)
        }
    }

    sort.Slice(filteredRoutes, func(i, j int) bool {
        return len(filteredRoutes[i].Rooms) < len(filteredRoutes[j].Rooms)
    })

    var optimalRoutes []*Route
    minSteps := math.MaxInt32

    for i := 1; i <= len(filteredRoutes); i++ {
        selectedRoutes := filteredRoutes[:i]
        steps := calculateSteps(selectedRoutes, numAnts)
        if steps < minSteps {
            minSteps = steps
            optimalRoutes = selectedRoutes
        }
    }
    return optimalRoutes
}

// applyMaxFlowAlgorithm applies the Max Flow logic to skip routes with shared rooms.
func applyMaxFlowAlgorithm(routes []*Route) {
    linkedTo := make([][]int, len(routes))
    for i := range linkedTo {
        linkedTo[i] = make([]int, len(routes))
    }

    for i := range routes {
        if routes[i].Skip {
            continue
        }
        for j := i + 1; j < len(routes); j++ {
            if routes[j].Skip {
                continue
            }
            if numOfSameRooms(routes[i].Rooms, routes[j].Rooms) > 2 {
                linkedTo[i][j] = 1
                linkedTo[j][i] = 1
            }
        }
    }

    maxSimilarity, maxRoute := 0, -1
    for i := len(linkedTo) - 1; i >= 0; i-- {
        sumConnections := 0
        for _, conn := range linkedTo[i] {
            sumConnections += conn
        }
        if sumConnections > maxSimilarity {
            maxSimilarity = sumConnections
            maxRoute = i
        }
    }

    if maxSimilarity != 0 {
        routes[maxRoute].Skip = true
        applyMaxFlowAlgorithm(routes)
    }
}

// calculateSteps calculates the total number of steps required to move all ants using the given routes.
func calculateSteps(routes []*Route, numAnts int) int {
    maxRouteLength := 0
    for _, route := range routes {
        if len(route.Rooms) > maxRouteLength {
            maxRouteLength = len(route.Rooms)
        }
    }
    // The total number of steps is the sum of the number of ants and the length of the longest route minus 1.
    totalSteps, remainingAnts := 0, numAnts
    for remainingAnts > 0 {
        totalSteps++
        for range routes {
            if remainingAnts > 0 {
                remainingAnts--
            }
        }
    }
    return totalSteps + maxRouteLength - 1
}

// numOfSameRooms function counts the number of rooms that are the same in two routes.
func numOfSameRooms(route1, route2 []*Room) int {
    count := 0
    for _, room1 := range route1 {
        for _, room2 := range route2 {
            if room1 == room2 {
                count++
            }
        }
    }
    return count
}

// containsRoom function checks if a room is in a slice of rooms.
func containsRoom(rooms []*Room, room *Room) bool {
    for _, r := range rooms {
        if r == room {
            return true
        }
    }
    return false
}

// PrintColonyConfiguration prints the colony configuration.
func PrintColonyConfiguration(colony *Colony) string {
    var result string
    result += fmt.Sprintf("%d\n", colony.TotalAnts)
    for _, room := range colony.Rooms {
        if room.IsStart {
            result += "##start\n"
        }
        if room.IsEnd {
            result += "##end\n"
        }
        result += fmt.Sprintf("%s %s %s\n", room.Name, room.CoordX, room.CoordY)
    }
    for _, link := range colony.Links {
        result += fmt.Sprintf("%s-%s\n", link.Room1.Name, link.Room2.Name)
    }
    return result
}
