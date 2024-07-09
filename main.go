package main

import (
    "fmt"
    "log"
    "os"
    "Lem-in/functions"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Please provide a filename as an argument")
    }

    filename := os.Args[1]

    if _, err := os.Stat(filename); os.IsNotExist(err) {
        log.Fatalf("File does not exist: %s", filename)
    }

    farm, err := functions.ReadFile(filename)
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    fmt.Printf("Farm loaded with %d ants\n", len(farm.Ants))

    fmt.Println("Rooms:")
    for _, room := range farm.Rooms {
        fmt.Printf("%s %d %d\n", room.RoomName, room.CoordX, room.CoordY)
    }

    fmt.Println("Links:")
    for _, link := range farm.Links {
        fmt.Printf("%s-%s\n", link.Room1.RoomName, link.Room2.RoomName)
    }

    paths, err := functions.FindDisjointPaths(farm)
    if err != nil {
        log.Fatalf("Error finding disjoint paths: %v", err)
    }

    fmt.Println("Found disjoint paths:")
    for i, path := range paths {
        fmt.Printf("Path %d: ", i+1)
        for _, room := range path {
            fmt.Printf("%s ", room.RoomName)
        }
        fmt.Println()
    }

    fmt.Println("Moving ants:")
    functions.MoveAnts(farm, paths)
}