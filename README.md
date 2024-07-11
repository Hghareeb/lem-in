Lem-in
Lem-in is a Go implementation of a program that finds the shortest path for ants to travel from a start room to an end room in a farm. The program reads a file containing the farm configuration, including the number of ants, room definitions, and links between rooms. It then finds the shortest path(s) and simulates the movement of ants along those paths.

Usage
To run the program, provide the filename containing the farm configuration as a command-line argument:

go run main.go example1.txt



The program will output the following:

The number of ants loaded from the file
A list of rooms with their names and coordinates
A list of links between rooms
The disjoint paths found from the start room to the end room
A simulation of the ants moving along the paths, with each step showing the movements made
File Format
The farm configuration file should follow this format:

<number of ants>
<room definition>
...
<room definition>
<link definition>
...
<link definition>



<number of ants> is an integer representing the number of ants in the farm.
<room definition> is a line with three fields: <room name> <x coordinate> <y coordinate>.
<link definition> is a line with two room names separated by a hyphen (-), representing a link between those rooms.
Additionally, the file should contain the following special lines:

##start: Followed by a room definition, this line indicates the start room.
##end: Followed by a room definition, this line indicates the end room.
Example file:

10
##start
start 1 6
0 4 8
o 6 8
n 6 6
e 8 4
t 1 9
E 5 9
a 8 9
m 8 6
h 4 6
A 5 2
c 8 1
k 11 2
##end
end 11 6
start-t
n-e
a-m
A-c
0-o
E-a
k-end
start-h
o-n
m-end
t-E
start-0
h-A
e-end
c-k
n-m
h-n



Project Structure
The project is organized into the following packages and files:

main.go: The entry point of the program.
functions/read.go: Contains functions for reading the farm configuration from a file.
functions/path.go: Contains functions for finding paths between the start and end rooms.
functions/ants.go: Contains functions for simulating the movement of ants along the paths.
functions/struct.go: Defines the structs used in the program.