const arr = new Array();
const a = new Set();
const visited = new Set(); 
const words = new Array();
const wordSet = new Set();

var x=0,y=0;
var typing=0;
var word = "";
var down=0;
var right=0;
var left=0;
var up=0;
var keyCounter=0;
var score=0;

async function getLetters(){
    document.getElementById("cell-" + ((4 * y) + x).toString()).style.border = "2px solid blue";
    
    const result = await fetch('http://localhost:8000/generateRandom')
        .then(response => {
            return response.text(); // Returns a promise
        })
        .then(data => {
            return data
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });

    console.log(result)

    for (var i = 0; i < 16; i++) {
        var letter = result[i];
        if (i%4 == 0) arr.push(new Array());
        arr.slice(-1)[0].push(letter);
        document.getElementById("hitbox-"+i.toString()).innerText = letter;
    }
}

getLetters();

function resetKeys(){
    up=0;
    down=0;
    left=0;
    right=0;
}

async function checkWord(){
    console.log(word);

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

    console.log("is", word, "?", valid)

    if (valid && !wordSet.has(word) && word.length > 2) {
        words.push(word);
        wordSet.add(word);
        score += 200 * word.length;
        console.log(score);
        document.getElementById("your_score").innerText = score.toString();
    }
    
    word="";
    visited.clear();
    document.getElementById("wList").innerText = words.join("\n");
}

function makeMove(){
    let prev_x=x, prev_y=y;

    if (up && y > 0) y--;
    if (down && y < 3) y++;
    if (right && x < 3) x++;
    if (left && x > 0) x--;
    
    if (visited.has((4*y)+x)){
        x=prev_x;
        y=prev_y;
        return; 
    } else if (typing) {
        visited.add((4*y)+x);
        word += document.getElementById("hitbox-" + ((4 * y) + x).toString()).innerText;
        document.getElementById("cell-" + ((4 * y) + x).toString()).style.backgroundColor = "red";
    }

    document.getElementById("cell-" + ((4 * prev_y) + prev_x).toString()).style.border = "1px solid black";
    document.getElementById("cell-" + ((4 * y) + x).toString()).style.border = "2px solid blue";
}

addEventListener("keydown", function(e) {

    if (e.keyCode == 72) {
        left=1;
    } else if(e.keyCode == 75){
        up=1;
    } else if(e.keyCode == 76){
        right=1;
    } else if(e.keyCode == 74){
        down=1;
    } 

    a.add(e.keyCode);
});

addEventListener("keyup", function(e) {
    if (e.keyCode == 86 && !typing) {
        typing = 1;
        makeMove();
    } else if (e.keyCode == 68 && typing) {
        typing = 0;
        resetKeys();
        checkWord();
    } else if(a.size == 1) {
        makeMove();
        resetKeys();
    } 

    a.delete(e.keyCode);
});

