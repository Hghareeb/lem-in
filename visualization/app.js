// Function to animate the ants
function animateAnts(farmData, paths) {
    const antElements = [];

    // Create ant elements
    farmData.ants.forEach((ant, index) => {
        const antElement = document.createElement('div');
        antElement.className = 'ant';
        antElement.setAttribute('data-id', ant.antID);
        antElements.push(antElement);
        farm.appendChild(antElement);
    });

    // Move ants along the paths
    let step = 0;
    const moveAnts = () => {
        // Clear previous positions
        antElements.forEach(antElement => {
            antElement.style.left = '';
            antElement.style.top = '';
        });

        // Get movements for the current step
        const movements = farmData.moveAnts(paths);

        // Update ant positions
        movements.forEach(movement => {
            const [antID, roomName] = movement.split('-');
            const ant = antElements.find(antElement => antElement.getAttribute('data-id') === antID);
            const room = farmData.rooms.find(room => room.roomName === roomName);
            ant.style.left = `${room.coordX}px`;
            ant.style.top = `${room.coordY}px`;
        });

        step++;
        requestAnimationFrame(moveAnts);
    };

    requestAnimationFrame(moveAnts);
}
