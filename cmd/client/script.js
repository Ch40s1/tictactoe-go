const board = document.getElementById("ApiBoard");
const cells = document.getElementsByClassName("cell");
let playerName = document.getElementById("PlayerName");
let playerSymbol = document.getElementById("PlayerSymbol");

// create json player with their values
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

// sends get request to /board
function getBoard() {
  fetch("http://localhost:3000/board")
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      // uses the data and gets the key and loop over the cells
      Object.keys(data).forEach((key) => {
        const cell = document.getElementById(key);
        if (cell) {
          cell.innerHTML = data[key];
        }
      });
    })
    .catch((error) => console.error("Error:", error));
}

async function playerTurn(player, nextTurn, endGame) {
  playerName.innerHTML = player.name;
  playerSymbol.innerHTML = player.symbol;

  function handleClick(event) {
    let cellId = event.target.id;
    console.log(cellId);
    const data = {
      symbol: player.symbol,
      cord: cellId,
    };

    // Remove event listeners after the move
    Array.from(cells).forEach((cell) => {
      cell.removeEventListener("click", handleClick);
    });

    // sends post to /turn
    fetch("http://localhost:3000/turn", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Success:", data);
        getBoard();
        if (cellId === "2,2") {
          alert("Game over!");
          endGame();
          return;
        }
        alert(`${player.name}'s turn ends`);
        nextTurn();
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  // Add click event listeners to all cells
  Array.from(cells).forEach((cell) => {
    cell.addEventListener("click", handleClick);
  });
}

function game() {
  // creates a tuen based game
  let turn = 0;
  // array of players
  const players = [player1, player2];
  // game state
  let gameStatus = "ongoing";

  function nextTurn() {
    // moves on only when the game is ongoing
    if (gameStatus === "ongoing") {
      playerTurn(players[turn], nextTurn, endGame);
      turn = 1 - turn; // Switch turn
    }
  }

  function endGame() {
    gameStatus = "over";
    // remove the event listeners for each cell
    Array.from(cells).forEach((cell) => {
      cell.removeEventListener("click", handleClick);
    });
  }

  nextTurn(); // Start the game with player1's turn
}

// ensures the code runs when the html is fully loaded
document.addEventListener("DOMContentLoaded", (event) => {
  getBoard();
  game();
});
