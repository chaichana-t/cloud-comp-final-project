const loc = window.location
const websocketURI = loc.protocol === "https:" ? "wss://" : "ws://" + loc.host + loc.pathname + "subscribe"
const socket = new WebSocket(websocketURI);

let restaurants = {}

const CLASS_NAME = {
    RESTAURANT_ID: "restaurant-id",
    CUSTOMER_COUNT: "customer-count",
}

socket.addEventListener("message", (event) => {
    const message = parseMessage(event.data)

    if (restaurants[message.restaurantID] === undefined) {
        createNewRestaurant(message.restaurantID, message.value)
        restaurants[message.restaurantID] = true
    } else {
        updateRestaurant(message.restaurantID, message.value)
    }
})

function parseMessage(message) {
    return {
        restaurantID: message.slice(0, 36),
        value: message.slice(36)
    }
}

function createElement(tag, options, children) {
    const e = document.createElement(tag)
    if (options.id !== undefined) e.id = options.id;
    if (options.className !== undefined) e.className = options.className;
    if (options.align !== undefined) e.align = options.align;

    for (const child of children) e.appendChild(child);

    return e;
}

function createNewRestaurant(restaurantID, value) {
    get("info", {rid: restaurantID}).then(info => {
        const restaurantName = createElement("td", {
            className: CLASS_NAME.RESTAURANT_ID
        }, [document.createTextNode(info.name)]);
        const customerCount = createElement("td", {
            className: CLASS_NAME.CUSTOMER_COUNT,
            align: "right",
        }, [document.createTextNode(value + " / " + info.maxCustomer)]);

        const tableBody = document.getElementById("table-body")
        tableBody.appendChild(createElement("tr", {id: restaurantID}, [restaurantName, customerCount]))
    })
}

function updateRestaurant(restaurantID, value) {
    document.getElementById(restaurantID).getElementsByClassName(CLASS_NAME.CUSTOMER_COUNT)[0].textContent = value;
}

function get(uri, query) {
    return new Promise((resolve, reject) => {
        uri += "?"
        for (const key of Object.keys(query)) {
            uri += key + "=" + query[key] + "&";
        }

        let xmlHttp = new XMLHttpRequest();

        xmlHttp.onreadystatechange = () => {
            if (xmlHttp.readyState === 4) {
                if (xmlHttp.status === 200) {
                    resolve(JSON.parse(xmlHttp.responseText))
                } else {
                    reject(xmlHttp.status)
                }
            }
        }

        xmlHttp.open("GET", encodeURI(uri), true);
        xmlHttp.send(null);
    })
}