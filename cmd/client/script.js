const board = document.getElementById('ApiBoard');

function getBoard() {
    fetch('http://localhost:3000/game')
        .then(response => response.json())
        .then(data => {
            console.log(data);
            board.innerHTML = data.message; // Use data.message to access the string
        })
        .catch(error => console.error('Error:', error));
}

document.addEventListener('DOMContentLoaded', (event) => {
    getBoard();
});
