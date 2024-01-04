const arr = new Array();
const keys = new Set();
const visited = new Set(); 
const words = new Array();
const wordSet = new Set();
const keyMap = {72 : 'h', 74 : 'j', 75 : 'k', 76 : 'l'}

let x=0, y=0;
let X=0, Y=0; 
let typing=0;
let motion=0;
let word = "";
let down=0;
let right=0;
let left=0;
let up=0;
let keyCounter=0;
let score=0;
let gameOver=0;


var urlParams = new URLSearchParams(window.location.search);
var name = urlParams.get('name');
var gameString = urlParams.get('gameString');
var gameId = urlParams.get('gameId');

async function checkFinish() {
    const result = await fetch("http://localhost:8000/checkFinish/" + name + '/' + gameId)
        .then(response => {
            return response.json();
        })
        .then(data => {
            return data
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
    
    if (result) {
        gameOver = 1;
        updateScoreBoard(result)
    }

    console.log(result)
}

checkFinish();
setLetters()

async function setLetters(){
    document.getElementById("cell-" + ((4 * y) + x).toString()).style.border = "2px solid blue";

    for (let i = 0; i < 16; i++) {
        let g = gameString[i];
        if (i%4 == 0) arr.push(new Array());
        arr.slice(-1)[0].push(g);
        document.getElementById("hitbox-"+i.toString()).innerText = g;
    }
}

let endTime = new Date();
endTime.setSeconds(endTime.getSeconds() + 20);

function checkTime() {
    let currentTime = new Date();

    let remainingTime = Math.floor((endTime - currentTime) / 1000);
    document.getElementById("your_time").innerText = "TIME: " + remainingTime;

    if (remainingTime <= 0) {
        endGame();
    } else {
        requestAnimationFrame(checkTime);
    }
}

async function endGame() {
    const result = await fetch("http://localhost:8000/finish/" + name + '/' + gameId + '/' + score)
        .then(response => {
            return response.json();
        })
        .then(data => {
            return data
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });

    gameOver = 1;
    updateScoreBoard(result)
}

function updateScoreBoard(result){
    let res = "--LEADERBOARD--\n"
    
    for (let i = 0; i < result.length; i++) {
        res += result[i]['Username'] + " - " + result[i]['Score'] + "\n"
    }

    document.getElementById("wList").innerText = res

    for (let i = 0; i < 16; i++)
        document.getElementById("cell-" + (i).toString()).style.backgroundColor = "white";

}

checkTime();

async function checkWord(){
    if (word.length == 0) {
        return;
    }

    for (let i = 0; i < 16; i++)
        document.getElementById("cell-" + (i).toString()).style.backgroundColor = "white";
     
    const valid = await fetch("http://localhost:8000/validateWord/" + word)
        .then(response => {
            return response.json(); 
        })
        .then(data => {
            return data
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });

    if (valid && !wordSet.has(word) && word.length > 2) {
        words.push(word + " >> " + (200*word.length).toString());
        wordSet.add(word);
        score += 200 * word.length;
        console.log(score);
        document.getElementById("your_score").innerText = "SCORE: " + score.toString()
    }

    word="";
    visited.clear();
    document.getElementById("wList").innerText = "--------WORDS--------\n" + words.join("\n");
}

function makeHighlightMove(l, r, d, u){
    let dy = r - l;
    let dx = d - u;

    document.getElementById("cell-" + (4*x + y).toString()).style.border = "1px solid black"
    document.getElementById("cell-" + (4*X + Y).toString()).style.border = "2px solid blue"

    if (x + dx != -1 && x + dx != 4 && y + dy != -1 && y + dy != 4) {
        x = x + dx;
        y = y + dy;
    }

    console.log("HOME SQUARE", X, Y)
    console.log("DEST SQUARE", x, y)

    document.getElementById("cell-" + (4*x + y).toString()).style.border = "2px solid red"
}

function makeNullMove(l, r, d, u){
    let dy = r - l;
    let dx = d - u;

    document.getElementById("cell-" + (4*x + y).toString()).style.border = "1px solid black"
    
    if (x + dx != -1 && x + dx != 4 && y + dy != -1 && y + dy != 4) {
        x += dx;
        y += dy;
        X=x,Y=y;
    }

    document.getElementById("cell-" + (4*x + y).toString()).style.border = "2px solid blue"
}

function addLetter() {
    if (visited.has(4*X + Y)) return;
    visited.add(4*X + Y)
    word += document.getElementById("cell-" + (4*X + Y).toString()).innerText;
}

addEventListener("keydown", function(e) {
    if (gameOver == 1) {
        return 
    }
    console.log(keys)
    if (keys.has(e.keyCode)) return;
    if (typing) {keys.add(e.keyCode)}

    let l=0,r=0,d=0,u=0

    if (e.keyCode == 72) {
        if (motion) keys.add(e.keyCode);
        left=l=1;
    } else if (e.keyCode == 75){
        if (motion) keys.add(e.keyCode);
        up=u=1;
    } else if (e.keyCode == 76){
        if (motion) keys.add(e.keyCode);
        right=r=1;
    } else if(e.keyCode == 74){
        if (motion) keys.add(e.keyCode);
        down=d=1;
    } else if (e.keyCode == 86) {
        addLetter();
        typing = 1;
    } else if (e.keyCode == 68) {
        typing = 0;
        checkWord();
        keys.clear();
        visited.clear();
    }

    if (typing) {
        makeHighlightMove(l,r,d,u);
    } else if (!typing) {
        makeNullMove(l,r,d,u);
    }

    console.log(up, down, left, right, X, Y, x, y)
});


addEventListener("keyup", function(e) {
    if (gameOver)
        return;

    if (e.keyCode == 72) {
        left=0;
    } else if (e.keyCode == 75){
        up=0;
    } else if (e.keyCode == 76){
        right=0;
    } else if(e.keyCode == 74){
        down=0;
    }

    if (!left && !right && !up && !down && keys.has(e.keyCode)) {
        if (!visited.has(4*x + y)){
            document.getElementById("cell-" + (4*X + Y).toString()).style.border = "1px solid black"
            X=x,Y=y;
            document.getElementById("cell-" + (4*X + Y).toString()).style.border = "2px solid blue"
            addLetter();
        } else {
            document.getElementById("cell-" + (4*x + y).toString()).style.border = "1px solid black"
            x=X,y=Y;
            document.getElementById("cell-" + (4*X + Y).toString()).style.border = "2px solid blue"
        }
        keys.clear();
        console.log(word);
    }
});

