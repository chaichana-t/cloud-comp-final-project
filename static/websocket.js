const loc = window.location
const websocketURI = loc.protocol === "https:" ? "wss://" : "ws://" + loc.host + loc.pathname + "subscribe"
const socket = new WebSocket(websocketURI);

let restaurants = {}

socket.addEventListener("message", (event) => {
    const restaurantID = event.data.slice(0, 36)
    const value = event.data.slice(36)

    if (restaurants[restaurantID] === undefined) {
        createRow(restaurantID, value)
    } else {
        updateRow(restaurantID, value)
    }
})

function createRow(restaurantID, value) {
    const row = document.createElement("tr")
    row.id = restaurantID
    getRestaurantName(restaurantID).then((name) => {
        row.appendChild(createCell(name, "restaurant-id"))
        row.appendChild(createCell(value, "current-customer"))

        const tableBody = document.getElementById("table-body")
        tableBody.appendChild(row)

        restaurants[restaurantID] = true
    })
}

function createCell(value, type) {
    const cell = document.createElement("td")
    cell.className = type
    cell.appendChild(document.createTextNode(value))
    return cell;
}

function updateRow(restaurantID, value) {
    const cell = document.getElementById(restaurantID).getElementsByClassName("current-customer")
    if (cell.length === 0) {
        alert("error")
    }
    cell[0].textContent = value
}

function getRestaurantName(restaurantID) {
    return new Promise((resolve, reject) => {
        let xmlHttp = new XMLHttpRequest();

        xmlHttp.onreadystatechange = () => {
            if (xmlHttp.readyState === 4 && xmlHttp.status === 200) {
                resolve(xmlHttp.responseText);
            }
        }

        xmlHttp.open("GET", "/name?rid=" + restaurantID, true);
        xmlHttp.send(null);
    })
}