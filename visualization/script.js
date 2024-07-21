document.addEventListener("DOMContentLoaded", function() {
    const visualizationContainer = document.getElementById("visualization-container");
    const pathsContainer = document.getElementById("paths-container");

    const rooms = [
        { name: "start", col: 1, row: 6, type: "start" },
        { name: "0", col: 4, row: 8 },
        { name: "o", col: 6, row: 8 },
        { name: "n", col: 6, row: 6 },
        { name: "e", col: 8, row: 4 },
        { name: "t", col: 1, row: 9 },
        { name: "E", col: 5, row: 9 },
        { name: "a", col: 8, row: 9 },
        { name: "m", col: 8, row: 6 },
        { name: "h", col: 4, row: 6 },
        { name: "A", col: 5, row: 2 },
        { name: "c", col: 8, row: 1 },
        { name: "k", col: 11, row: 2 },
        { name: "end", col: 11, row: 6, type: "end" }
    ];

    const links = [
        { from: "start", to: "t" },
        { from: "n", to: "e" },
        { from: "a", to: "m" },
        { from: "A", to: "c" },
        { from: "0", to: "o" },
        { from: "E", to: "a" },
        { from: "k", to: "end" },
        { from: "start", to: "h" },
        { from: "o", to: "n" },
        { from: "m", to: "end" },
        { from: "t", to: "E" },
        { from: "start", to: "0" },
        { from: "h", to: "A" },
        { from: "e", to: "end" },
        { from: "c", to: "k" },
        { from: "n", to: "m" },
        { from: "h", to: "n" }
    ];

    // Draw rooms
    rooms.forEach(room => {
        const roomDiv = document.createElement("div");
        roomDiv.className = "room";
        if (room.type) {
            roomDiv.classList.add(room.type);
        }
        roomDiv.style.gridColumn = room.col;
        roomDiv.style.gridRow = room.row;
        roomDiv.innerText = room.name;
        roomDiv.id = room.name;
        visualizationContainer.appendChild(roomDiv);
    });

    // Draw links using SVG
    links.forEach(link => {
        const fromRoom = document.getElementById(link.from);
        const toRoom = document.getElementById(link.to);

        const fromX = fromRoom.offsetLeft + fromRoom.offsetWidth / 2;
        const fromY = fromRoom.offsetTop + fromRoom.offsetHeight / 2;
        const toX = toRoom.offsetLeft + toRoom.offsetWidth / 2;
        const toY = toRoom.offsetTop + toRoom.offsetHeight / 2;

        const path = document.createElementNS("http://www.w3.org/2000/svg", "line");
        path.setAttribute("x1", fromX);
        path.setAttribute("y1", fromY);
        path.setAttribute("x2", toX);
        path.setAttribute("y2", toY);
        path.classList.add("path");

        pathsContainer.appendChild(path);
    });

    // Dynamic ant movements
    const antMovements = [
        ["L1-start", "L1-t", "L1-E", "L1-a", "L1-m", "L1-end"],
        ["L2-start", "L2-h", "L2-A", "L2-c", "L2-k", "L2-end"],
        ["L3-start", "L3-0", "L3-o", "L3-n", "L3-e", "L3-end"]
    ];

    const ants = {};

    antMovements.forEach((movements, index) => {
        const antId = movements[0].split('-')[0];
        const initialRoomId = movements[0].split('-')[1];
        const ant = document.createElement("div");
        ant.className = "ant";
        ant.id = antId;
        document.getElementById(initialRoomId).appendChild(ant);
        ants[antId] = { element: ant, movements: movements.slice(1), position: 0 };
    });

    document.getElementById("start-button").addEventListener("click", () => {
        let interval = setInterval(() => {
            let allFinished = true;
            Object.keys(ants).forEach(antId => {
                const ant = ants[antId];
                if (ant.position < ant.movements.length) {
                    allFinished = false;
                    const currentRoomId = ant.movements[ant.position - 1]?.split('-')[1];
                    const nextRoomId = ant.movements[ant.position].split('-')[1];
                    const currentRoom = document.getElementById(currentRoomId);
                    const nextRoom = document.getElementById(nextRoomId);

                    if (currentRoom) {
                        currentRoom.classList.remove("occupied");
                    }

                    nextRoom.appendChild(ant.element);
                    nextRoom.classList.add("occupied");

                    // Animate the path (line) between the rooms
                    const path = pathsContainer.querySelector(`line[x1="${currentRoom.offsetLeft + currentRoom.offsetWidth / 2}"][y1="${currentRoom.offsetTop + currentRoom.offsetHeight / 2}"][x2="${nextRoom.offsetLeft + nextRoom.offsetWidth / 2}"][y2="${nextRoom.offsetTop + nextRoom.offsetHeight / 2}"]`);
                    if (path) path.classList.add("active");

                    ant.position++;
                }
            });
            if (allFinished) {
                clearInterval(interval);
            }
        }, 1000);
    });
});
