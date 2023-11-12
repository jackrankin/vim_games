const arr = new Array();
const x=0,y=0;
const typing=0;

function getLetters(){

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

var map = {}; // You could also use an array
onkeydown = onkeyup = function(e){
    e = e || event; // to deal with IE
    map[e.keyCode] = e.type == 'keydown';
    console.log(map);
    /* insert conditional here */
}

//document.addEventListener("keydown", function (event) {
  //  if (!typing){
    //    return;
    //}


    //const key = event.key;
    //if (key)
//});



