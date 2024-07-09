package functions

type Room struct {
    RoomName  string
    CoordX    int
    CoordY    int
    Start     bool
    End       bool
    IsVisited bool
    Links     []*Link
}

type Link struct {
    Room1 *Room
    Room2 *Room
}

type Farm struct {
    Rooms     []*Room
    Ants      []*Ant
    Links     []*Link
    StartRoom *Room
    EndRoom   *Room
}

type Ant struct {
    AntID int
}

type Path struct {
    Rooms []*Room
}
