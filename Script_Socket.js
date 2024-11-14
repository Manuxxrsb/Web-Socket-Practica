const socket = new WebSocket("http://localhost:8080/Web_Socket");

socket.onopen = (event) => { // Mientras la conexion este abierta
    console.log("WebSocket connection opened");
};

socket.onmessage = (event) => { // Cuando se recibe un mensaje
    console.log("Message received: ", event.data);
    const Respuesta = JSON.parse(event.data);
    const output = document.getElementById("output");
    output.innerHTML += `<p>${Respuesta.username}: ${Respuesta.contenido}</p>`;
};

function sendMessage() { // Enviar mensaje
    const messageInput = document.getElementById("messageInput");
    const userInput = document.getElementById("userInput");

    const message = messageInput.value;
    const user = userInput.value;

    const Packet = {
        username: user,
        contenido: message
    }

    let paquete = JSON.stringify(Packet);
    console.log(paquete);

    socket.send(paquete);
    messageInput.value = "";
}