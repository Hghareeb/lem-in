package functions

import (
    "container/heap"
    "errors"
    "fmt"
    "math"
)

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// Item is something we manage in a priority queue.
type Item struct {
    room     *Room
    priority int
    index    int
}

// Len is the number of elements in the collection.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less reports whether the element with index i should sort before the element with index j.
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}

// Swap swaps the elements with indexes i and j.
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

// Push adds an element to the heap.
func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

// Pop removes and returns the minimum element (according to Less).
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil  // avoid memory leak
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, priority int) {
    item.priority = priority
    heap.Fix(pq, item.index)
}

// ShortestPath finds the shortest path from start to end using Dijkstra's algorithm.
func ShortestPath(farm *Farm) ([]*Room, error) {
    distances := make(map[*Room]int)
    previous := make(map[*Room]*Room)
    pq := make(PriorityQueue, 0)

    for _, room := range farm.Rooms {
        distances[room] = math.MaxInt32
        previous[room] = nil
    }
    distances[farm.StartRoom] = 0

    heap.Push(&pq, &Item{room: farm.StartRoom, priority: 0})

    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        currentRoom := item.room

        if currentRoom == farm.EndRoom {
            break
        }

        for _, link := range currentRoom.Links {
            neighbor := link.Room2
            if neighbor == currentRoom {
                neighbor = link.Room1
            }
            alt := distances[currentRoom] + 1 // Assuming each link has equal weight
            if alt < distances[neighbor] {
                distances[neighbor] = alt
                previous[neighbor] = currentRoom
                heap.Push(&pq, &Item{room: neighbor, priority: alt})
            }
        }
    }

    path := []*Room{}
    for u := farm.EndRoom; u != nil; u = previous[u] {
        path = append([]*Room{u}, path...)
    }

    if len(path) == 0 || path[0] != farm.StartRoom {
        return nil, fmt.Errorf("no path found from %s to %s", farm.StartRoom.RoomName, farm.EndRoom.RoomName)
    }

    return path, nil
}

// BFS to find disjoint paths
func FindDisjointPaths(farm *Farm) ([][]*Room, error) {
    paths := [][]*Room{}
    visited := make(map[*Room]bool)
    
    // Helper function to perform BFS
    bfs := func(start, end *Room) ([]*Room, error) {
        queue := [][]*Room{{start}}
        visited[start] = true

        for len(queue) > 0 {
            path := queue[0]
            queue = queue[1:]
            node := path[len(path)-1]

            if node == end {
                return path, nil
            }

            for _, link := range node.Links {
                next := link.Room2
                if next == node {
                    next = link.Room1
                }

                if !visited[next] {
                    newPath := append([]*Room{}, path...)
                    newPath = append(newPath, next)
                    queue = append(queue, newPath)
                    visited[next] = true
                }
            }
        }

        return nil, errors.New("no path found")
    }

    // Find multiple disjoint paths
    for {
        path, err := bfs(farm.StartRoom, farm.EndRoom)
        if err != nil {
            break
        }
        paths = append(paths, path)
        // Mark all rooms in this path as visited
        for _, room := range path {
            visited[room] = true
        }
    }

    if len(paths) == 0 {
        return nil, errors.New("no disjoint paths found from start to end")
    }

    return paths, nil
}