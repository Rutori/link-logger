var template = Handlebars.compile(document.getElementById("link-template").innerHTML)

let authPassword
document.addEventListener('DOMContentLoaded', function(){
    authPassword = localStorage.getItem("password")
    if ( authPassword === null) {
        pass = prompt("enter password")
        sha256(pass).then(function (data) {
            authPassword = data
            localStorage.setItem("password", authPassword)
            refresh()
        })
        return
    }

    refresh()
})

document.getElementById("submit-link").onclick = function() {
    submitlinks()
};

document.getElementById("auth-reset").onclick = function() {
    localStorage.removeItem("password")
    window.location.reload()
};

function submitlinks(){
    link = document.getElementById("new-link").value
    fetch('/api/add', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8',
            'Authorization' : authPassword
        },
        body : link
    }).then(function (data) {
        data.text().then(function () {
            refresh()
        })
    })
}

function clearLinks() {
    document.getElementById("links").innerHTML = ""
}

function refresh(){
    clearLinks()
    fetch('/api/list', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json;charset=utf-8',
            'Authorization' : authPassword
        },
    }).then(function (data) {
        data.json().then(function (result) {
            generateLinks(result)
        })
    })
}

function generateLinks(links) {
    for (i = 0; i < links.length; i++){
        links[i].CreatedAt = new Date( links[i].CreatedAt).toDateString()
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
