var template = Handlebars.compile(document.getElementById("link-template").innerHTML)

var authPassword = "not_a_real_password_duh"


document.addEventListener('DOMContentLoaded', function(){
    refresh()
})

document.getElementById("submit-link").onclick = function() {
    submitlinks()
};

function submitlinks(){
    link = document.getElementById("new-link").value
    sha256(authPassword).then(function (auth) {
        fetch('/api/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8',
                'Authorization' : auth
            },
            body : link
        }).then(function (data) {
            data.text().then(function () {
                refresh()
            })
        })
    })
}

function clearLinks() {
    document.getElementById("links").innerHTML = ""
}

function refresh(){
    clearLinks()
    sha256(authPassword).then(function (auth) {
        fetch('/api/list', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json;charset=utf-8',
                'Authorization' : auth
            },
        }).then(function (data) {
            data.json().then(function (result) {
                generateLinks(result)
            })
        })
    })
}

function generateLinks(links) {
    for (i = 0; i < links.length; i++){
        document.getElementById("links").innerHTML += template(links[i])
    }
}

async function sha256(message) {
    // encode as UTF-8
    const msgBuffer = new TextEncoder().encode(message);

    // hash the message
    const hashBuffer = await crypto.subtle.digest('SHA-256', msgBuffer);

    // convert ArrayBuffer to Array
    const hashArray = Array.from(new Uint8Array(hashBuffer));

    // convert bytes to hex string
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
    return hashHex;
}
