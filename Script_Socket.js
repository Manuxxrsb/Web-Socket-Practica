const socket = new WebSocket("http://localhost:8080/Web_Socket");

socket.onopen = (event) => { // Mientras la conexion este abierta
    console.log("WebSocket connection opened");
};

socket.onmessage = (event) => { // Cuando se recibe un mensaje
    const output = document.getElementById("output");
    output.innerHTML += `<p>Received: ${event.data}</p>`;
};

function sendMessage() { // Enviar mensaje
    const messageInput = document.getElementById("messageInput");
    const message = messageInput.value;
    socket.send(message);
    messageInput.value = "";
}