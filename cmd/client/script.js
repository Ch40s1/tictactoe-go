const board = document.getElementById("ApiBoard");
let firstCell = document.getElementById("0,0");
let secondCell = document.getElementById("0,1");
let thirdCell = document.getElementById("0,2");
let fourthCell = document.getElementById("1,0");
let fifthCell = document.getElementById("1,1");
let sixthCell = document.getElementById("1,2");
let seventhCell = document.getElementById("2,0");
let eighthCell = document.getElementById("2,1");
let ninthCell = document.getElementById("2,2");

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

document.addEventListener("DOMContentLoaded", (event) => {
  getBoard();
});
