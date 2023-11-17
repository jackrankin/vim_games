
const arr = new Array();
const a = new Set();
var x=0,y=0;
var typing=0;

var word = "";
var down=0;
var right=0;
var left=0;
var up=0;
var keyCounter=0;

function getLetters(){ // in reality this should read from the DB

    for (var i = 0; i < 16; i++) {
        var letter = String.fromCharCode('A'.charCodeAt() + Math.floor(Math.random() * 26));

        if (i%4 == 0){
            arr.push(new Array());
        }

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

function checkWord(){
    console.log(word);
    word="";
    document.getElementById("currentWord").innerText = word;
    for (let i = 0; i < 16; i++) {

        document.getElementById("cell-" + (i).toString()).style.backgroundColor = "white";
    }
}


function makeMove(){

    let ch=0;
    if (up && y > 0) {
        ch=1;
        y--;
    }

    if (down && y < 3){
        ch=1;
        y++;
    }

    if (right && x < 3){
        ch=1;
        x++;
    }

    if (left && x > 0) {
        ch=1;
        x--;
    }


    word += document.getElementById("hitbox-" + ((4 * y) + x).toString()).innerText;
    
    document.getElementById("currentWord").innerText = word;
    document.getElementById("cell-" + ((4 * y) + x).toString()).style.backgroundColor = "red";

}

addEventListener("keydown", function(e) {
    if (e.keyCode == 37) {
        left=1;
    } else if(e.keyCode == 38){
        up=1;
    } else if(e.keyCode == 39){
        right=1;
    } else if(e.keyCode == 40){
        down=1;
    } 
    a.add(e.keyCode); 
});


addEventListener("keyup", function(e) {
    if (e.keyCode == 32 && !typing) {
        typing = 1;
    } else if (e.keyCode == 32 && typing) {
        typing = 0;
        resetKeys();
        checkWord();
    }

    if(a.size == 1 && typing && (e.keyCode == 32 || e.keyCode == 37 || e.keyCode == 38 || e.keyCode == 39 || e.keyCode == 40)) {
        makeMove();
        resetKeys();
    }

    a.delete(e.keyCode);
});

