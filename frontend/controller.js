var gameId = "";
var name = "";
var gameString = "";

async function enterGame() {
    gameId = document.getElementById("gameField").value
    name = document.getElementById("nameField").value
    const result = await fetch("http://localhost:8000/joinRoom/" + name + "/" + gameId)
        .then(response => {
            return response.text();
        })
        .then(data => {
            return data
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
    gameString = result
    updatePage();
}

async function makeGame() {
    name = document.getElementById("_nameField").value
    console.log(name)

    const result = await fetch("http://localhost:8000/makeRoom")
        .then(response => {
            return response.json();
        })
        .then(data => {
            return data
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });

    gameId = result['GameId']
    gameString = result['GameString']

    const r2 = await fetch("http://localhost:8000/joinRoom/" + name + "/" + gameId)
        .then(response => {
            return response.text();
        })
        .then(data => {
            return data
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
    //updatePage();
}

function updatePage() {
    var url = "./boggle.html?name=" + encodeURIComponent(name) +
    "&gameString=" + encodeURIComponent(gameString) +
    "&gameId=" + encodeURIComponent(gameId);
    location.assign(url);
}





