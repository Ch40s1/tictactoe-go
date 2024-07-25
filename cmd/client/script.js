const cells = document.getElementsByClassName("cell");
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
      name: player.name,
      symbol: player.symbol,
      cord: cellId,
    };

    // Remove event listeners after the move
    Array.from(cells).forEach((cell) => {
      cell.removeEventListener("click", handleClick);
    });

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

        if (data.winner === "winner") {
          alert(`Game over! ${player.name} wins!`);
          endGame();
          return;
        }

        // alert(`${player.name}'s turn ends`);
        nextTurn();
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  Array.from(cells).forEach((cell) => {
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
    Array.from(cells).forEach((cell) => {
      cell.removeEventListener("click", handleClick);
    });
  }

  nextTurn();
}

document.addEventListener("DOMContentLoaded", (event) => {
  getBoard();
  game();
});
