package functions

// Room represents a room in the colony.
type Room struct {
    Name      string
    CoordX    string
    CoordY    string
    Visited   bool
    IsEnd     bool
    IsStart   bool
    Links     []*Link
    Neighbors []*Room
    LineNum   int
    Ants      []*Ant
}

// Ant represents an ant in the colony.
type Ant struct {
    ID          int
    Path        *Route
    CurrentRoom *Room
}

// Colony represents the entire ant colony.
type Colony struct {
    Rooms         []*Room
    Ants          []*Ant
    Links         []*Link
    TotalAnts     int
    StartRoom     *Room
    EndRoom       *Room
    AntPositions  map[int]*Room
    Paths         []*Route
    StartRoomLine int
    EndRoomLine   int
}

// Link represents a link between two rooms.
type Link struct {
    Room1 *Room
    Room2 *Room
}

// Route represents a path from the start room to the end room.
type Route struct {
    Rooms    []*Room
    Skip     bool
    Steps    int
    Ants     []*Ant
}
//all good here