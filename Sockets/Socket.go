package sockets

import (
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
)

type MessageLocal struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Is_group string `json:"is_group"`
}

type Client struct {
	Conn     *websocket.Conn
	Username string
}

var clients = make(map[string]*Client)

var mu sync.Mutex

func HandleWebSocket(c *websocket.Conn) {
	// Leer el nombre de usuario
	var data struct {
		Username string `json:"username"`
	}

	if err := c.ReadJSON(&data); err != nil {
		// fmt.Println("Error al leer el nombre de usuario:", err)
		return
	}
	username := data.Username // Obtener el nombre de usuario

	mu.Lock()
	clients[username] = &Client{Conn: c, Username: username}
	mu.Unlock()

	fmt.Println("Usuario conectado:", username)

	defer func() {
		mu.Lock()
		delete(clients, username)
		mu.Unlock()
		c.Close()
		fmt.Println("Usuario desconectado:", username)
	}()

	for {
		var msg struct {
			Recipient string `json:"recipient"`
			Message   string `json:"message"`
		}
		err := c.ReadJSON(&msg)
		if err != nil {
			// fmt.Println("Error al leer el mensaje:", err)
			break
		}

		// fmt.Printf("Mensaje recibido de %s a %s: %s\n", username, msg.Recipient, msg.Message)
		SendPrivateMessage(msg.Recipient, msg.Message, username)
	}
}

func SendPrivateMessage(recipient, message, sender string) {
	mu.Lock()
	defer mu.Unlock()

	if client, ok := clients[recipient]; ok {
		msg := struct {
			Sender  string `json:"sender"`
			Message string `json:"message"`
		}{
			Sender:  sender,
			Message: message,
		}
		err := client.Conn.WriteJSON(msg)
		if err != nil {
			// fmt.Println("Error al enviar el mensaje:", err)
		} else {
			// fmt.Printf("Mensaje enviado de %s a %s: %s\n", sender, recipient, message)
		}
	}
}
