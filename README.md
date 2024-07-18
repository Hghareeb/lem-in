# Lem-in Project

## Overview

The Lem-in project simulates the movement of ants from a start room to an end room through a series of interconnected rooms. The goal is to optimize the movement so that the ants reach the end room with the least number of turns, ensuring no two ants occupy the same room simultaneously.

## Features

- **Pathfinding**: Finds all possible paths from the start room to the end room.
- **Ant Movement Optimization**: Distributes ants across paths to minimize the number of turns.
- **Flow Optimization**: Applies a max-flow algorithm to optimize the path selection.

## Project Structure

- **main.go**: The entry point of the application.
- **functions/structs.go**: Contains the definitions of the main data structures used in the project.
- **functions/ants.go**: Contains the logic for moving the ants along the paths.
- **functions/path.go**: Contains the pathfinding logic and the flow optimization algorithm.
- **functions/read.go**: Contains the logic for reading the input file and parsing the farm configuration.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Running the Project

1. Clone the repository:
   ```sh
   git clone <https://github.com/Hghareeb/lem-in.git>
   cd lem-in
