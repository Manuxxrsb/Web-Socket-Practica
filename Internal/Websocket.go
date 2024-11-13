package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Actualizar la conexión HTTP a una conexión WebSocket.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Manejar mensajes de WebSocket aquí
	for {
		messageType, data, err := conn.ReadMessage() // Leer el mensaje del cliente
		if err != nil {                              // Manejar errores
			fmt.Println(err)
			return
		}
		fmt.Printf("Mensaje recibido: %s\n", data)

		// Devolver el mensaje al cliente
		if err := conn.WriteMessage(messageType, data); err != nil { // Enviar el mensaje al cliente
			fmt.Println(err)
			return
		}
	}
}

func RunSocket() { //Exposicion del servidor WebSocket en el puerto 8080
	http.HandleFunc("/Web_Socket", handleWebSocket)
	fmt.Println("El servidor WebSocket está ejecutándose en :8080/ws")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
