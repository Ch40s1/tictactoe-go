const board = document.getElementById("ApiBoard");
const cells = document.getElementsByClassName("cell");
let firstCell = document.getElementById("0,0");
let secondCell = document.getElementById("0,1");
let thirdCell = document.getElementById("0,2");
let fourthCell = document.getElementById("1,0");
let fifthCell = document.getElementById("1,1");
let sixthCell = document.getElementById("1,2");
let seventhCell = document.getElementById("2,0");
let eighthCell = document.getElementById("2,1");
let ninthCell = document.getElementById("2,2");
let playerName = document.getElementById("PlayerName");
let playerSymbol = document.getElementById("PlayerSymbol");
let player1 = {
  id: 0,
  name: "Daniel",
  symbol: "X",
};
let player2 = {
  id: 1,
  name: "John",
  symbol: "O",
};

function getBoard() {
  fetch("http://localhost:3000/board")
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      Object.keys(data).forEach(key => {
        const cell = document.getElementById(key);
        if (cell) {
          cell.innerHTML = data[key];
        }
      });
    })
    .catch((error) => console.error("Error:", error));
}

function playerTurn(player, nextTurn, endGame) {
  playerName.innerHTML = player.name;
  playerSymbol.innerHTML = player.symbol;

  function handleClick(event) {
    let cellId = event.target.id;
    console.log(cellId);
    // Remove event listeners after the move
    Array.from(cells).forEach(cell => {
      cell.removeEventListener("click", handleClick);
    });
    if (cellId === "2,2") {
      alert("Game over!");
      endGame();
      return;
    }
    alert(`${player.name}'s turn ends`);
    nextTurn();
  }

  // Add click event listeners to all cells
  Array.from(cells).forEach(cell => {
    cell.addEventListener("click", handleClick);
  });
}

function game() {
  let turn = 0;
  const players = [player1, player2];
  let gameStatus = "ongoing";

  function nextTurn() {
    if (gameStatus === "ongoing") {
      playerTurn(players[turn], nextTurn, endGame);
      turn = 1 - turn; // Switch turn
    }
  }

  function endGame() {
    gameStatus = "over";
    Array.from(cells).forEach(cell => {
      cell.removeEventListener("click", handleClick);
    });
  }

  nextTurn(); // Start the game with player1's turn
}


document.addEventListener("DOMContentLoaded", (event) => {
  game();
});
