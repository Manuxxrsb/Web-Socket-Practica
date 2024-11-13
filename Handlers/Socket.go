package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var actualizador = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ejecuta_Socket() gin.HandlerFunc {
	return func(Request *gin.Context) {
		conn, err := actualizador.Upgrade(Request.Writer, Request.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		for {
			mt, mensaje, err := conn.ReadMessage() // Leer el mensaje del cliente
			if err != nil {                        // Manejar errores
				fmt.Println(err)
				return
			}
			fmt.Printf("Recibido: %s\n", mensaje)
			err = conn.WriteMessage(mt, mensaje)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	}
}
