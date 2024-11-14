package handlers

import (
	models "backend/Models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)


var actualizador = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ejecuta_Socket(db *gorm.DB) gin.HandlerFunc {
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

			// Deserializar el mensaje del cliente
            var wsMessage models.Mensaje
            if err := json.Unmarshal(mensaje, &wsMessage); err != nil {
                fmt.Println("Error al deserializar mensaje:", err)
                continue
            }

            fmt.Printf("Usuario: %s, Mensaje: %s\n", wsMessage.Username, wsMessage.Contenido)

			if err := db.Create(&wsMessage).Error; err != nil {
				Request.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}


			err = conn.WriteMessage(mt, mensaje)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	}
}
