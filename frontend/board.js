
const arr = new Array();

function getLetters(){

    for (var i = 0; i < 16; i++) {
        var letter = String.fromCharCode('A'.charCodeAt() + Math.floor(Math.random() * 26));

        if (i%4 == 0){
            arr.push(new Array());
        }

        arr.slice(-1)[0].push(letter);
        document.getElementById("cell-"+i.toString()).innerHTML = letter;
    }
}

function handleDrag(){
    
}

getLetters();

console.log(arr);
