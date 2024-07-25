document.addEventListener("DOMContentLoaded", function() {
    const visualizationContainer = document.getElementById("visualization-container");
    const pathsContainer = document.getElementById("paths-container");
    const fileInput = document.getElementById("file-input");
    const startButton = document.getElementById("start-button");

    let ants = {};
    let antMovements = [];
    let rooms = [];
    let links = [];
    let numAnts = 0;

    fileInput.addEventListener("change", handleFileSelect);

    function handleFileSelect(event) {
        const file = event.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = function(e) {
                const content = e.target.result;
                parseFileContent(content);
                createVisualization();
                startButton.disabled = false;
            };
            reader.readAsText(file);
        }
    }

    function parseFileContent(content) {
        const lines = content.split('\n').map(line => line.trim()).filter(line => line !== '');
        let currentSection = 'ants';

        rooms = [];
        links = [];
        antMovements = [];

        lines.forEach(line => {
            if (line.startsWith('##')) {
                if (line === '##start' || line === '##end') {
                    currentSection = line.substring(2);
                } else {
                    currentSection = 'rooms';
                }
            } else if (line.includes(' ')) {
                const [name, col, row] = line.split(' ');
                rooms.push({
                    name,
                    col: parseInt(col),
                    row: parseInt(row),
                    type: currentSection === 'start' ? 'start' : currentSection === 'end' ? 'end' : ''
                });
            } else if (line.includes('-')) {
                const [from, to] = line.split('-');
                links.push({ from, to });
            } else if (!isNaN(parseInt(line))) {
                numAnts = parseInt(line);
                antMovements = Array.from({ length: numAnts }, (_, i) => [`L${i + 1}-start`]);
            }
        });

        // Assign dummy movements for testing
        for (let i = 0; i < numAnts; i++) {
            antMovements[i].push(`L${i + 1}-end`);
        }
    }

    function createVisualization() {
        const container = document.getElementById('visualization-container');
        container.innerHTML = ''; // Clear any existing visualization
    
        // Create a grid of cells for the rooms
        rooms.forEach(room => {
            const cell = document.createElement('div');
            cell.style.gridRowStart = room.row;
            cell.style.gridColumnStart = room.col;
            cell.textContent = room.name;
            cell.className = 'room ' + room.type; // Use the room type as a CSS class
            container.appendChild(cell);
        });
    
        // Create lines for the links
        links.forEach(link => {
            const fromRoom = rooms.find(room => room.name === link.from);
            const toRoom = rooms.find(room => room.name === link.to);
            const line = document.createElement('div');
            line.style.gridRowStart = fromRoom.row;
            line.style.gridColumnStart = fromRoom.col;
            line.style.gridRowEnd = toRoom.row;
            line.style.gridColumnEnd = toRoom.col;
            line.className = 'link';
            container.appendChild(line);
        });
    
        // Create elements for the ants
        antMovements.forEach((antMovement, i) => {
            const ant = document.createElement('div');
            const startRoom = rooms.find(room => room.name === antMovement[0].split('-')[1]);
            ant.style.gridRowStart = startRoom.row;
            ant.style.gridColumnStart = startRoom.col;
            ant.textContent = 'L' + (i + 1);
            ant.className = 'ant';
            container.appendChild(ant);
        });
    }

    startButton.addEventListener("click", () => {
        let interval = setInterval(() => {
            let allFinished = true;
            Object.keys(ants).forEach(antId => {
                const ant = ants[antId];
                if (ant.position < ant.movements.length) {
                    allFinished = false;
                    const currentRoomId = ant.movements[ant.position].split('-')[1];
                    const nextRoomId = ant.movements[ant.position + 1]?.split('-')[1];

                    if (nextRoomId) {
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
                }
            });
            if (allFinished) {
                clearInterval(interval);
            }
        }, 1000);
    });
});
function parseFileContent(content) {
    const lines = content.split('\n').map(line => line.trim()).filter(line => line !== '');
    let currentSection = 'ants';

    rooms = [];
    links = [];
    antMovements = [];

    lines.forEach(line => {
        if (line.startsWith('##')) {
            if (line === '##start' || line === '##end') {
                currentSection = line.slice(2);
            }
        } else {
            switch (currentSection) {
                case 'ants':
                    numAnts = parseInt(line);
                    currentSection = '';
                    break;
                case 'start':
                case 'end':
                case 'rooms':
                    rooms.push(line);
                    break;
                case 'links':
                    links.push(line);
                    break;
                default:
                    antMovements.push(line);
                    break;
            }
        }
    });
}
// doesnt work well