var host = "http://localhost:32001"

function setHost(newHost) {
    host = newHost
}

function setStatus(status) {
    document.getElementById('status').innerText = status;
}

function setPanic(panic) {
    document.getElementById('panic').innerText = panic;
    showPage('panicPage');
}

function startMining() {
    var address = document.getElementById('walletAddress').value;
    fetch(host + "/startmining?address=" + address);
}

function stopMining() {
    fetch(host + "/stopmining");
}

function showPage(id) {
    var containers = document.getElementsByClassName("container")
    for(var i = 0; i < containers.length; i++) {
        containers[i].style.display = (containers[i].id === id ? '' : 'none')
    }
}