package functions

// DistributeAnts distributes ants to the paths with the least cost.
func DistributeAnts(colony *Colony) {
    for i := 1; i <= colony.TotalAnts; i++ {
        ant := &Ant{ID: i, CurrentRoom: colony.StartRoom}
        colony.Ants = append(colony.Ants, ant)

        // Find the path with the least total "cost"
        minCost := len(colony.Paths[0].Rooms) + len(colony.Paths[0].Ants)
        minPath := colony.Paths[0]

        for _, path := range colony.Paths[1:] {
            cost := len(path.Rooms) + len(path.Ants)
            if cost < minCost {
                minCost = cost
                minPath = path
            }
        }

        // Assign the ant to the path with the least cost
        minPath.Ants = append(minPath.Ants, ant)
    }
}
