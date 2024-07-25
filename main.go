package main

import (
    "Lem-in/functions"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <filename>")
        return
    }

    // Parse the colony from the input file
    colony := functions.File(os.Args[1])
    if colony == nil {
        fmt.Println("Failed to parse colony from input file.")
        return
    }

    // Print the colony configuration (this echoes the input)
    inputText := functions.PrintColonyConfiguration(colony)
    fmt.Print(inputText)
    fmt.Println()

    // Find all possible routes from start to end
    routes := functions.Edmonds(colony)
    if len(routes) == 0 {
        fmt.Printf("ERROR: invalid data format\nNo path from start room to end room\n")
        return
    }
    // Print all possible paths
    //just to check
    // for _, route := range routes {
    // 	fmt.Println(route)

    // Choose optimal routes
    colony.Paths = functions.ChooseOptimalRoutes(routes, colony.TotalAnts)

    // Distribute ants to routes
    functions.DistributeAnts(colony)

    // Simulate ant movements
    movements := functions.SimulateAntsMovement(colony)

    // Print the movements
    for _, move := range movements {
        if move != "" {
            fmt.Println(move)
        }
    }
    fmt.Println("$")
}
