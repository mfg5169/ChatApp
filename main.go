package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

// Define the WebSocket upgrader
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow connections from any origin
    },
}

// This channel will broadcast messages to all connected clients
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// Define a Message struct
type Message struct {
    Username string `json:"username"`
    Content  string `json:"content"`
}

// Handle incoming WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    // Register the new client
    clients[conn] = true
    log.Println("New client connected")

    // Make sure to remove the client when the connection is closed
    defer func() {
        delete(clients, conn)
        conn.Close()
        log.Println("Client disconnected")
    }()

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            log.Println(err)
            return
        }

        // Broadcast the message to all clients
        broadcast <- msg
    }
}

// Broadcast messages to all clients
func handleMessages() {
    for {
        // Get the message from the broadcast channel
        msg := <-broadcast
        // Send the message to all connected clients
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Println(err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main() {
    go handleMessages()

    http.HandleFunc("/ws", handleConnections)

    http.Handle("/", http.FileServer(http.Dir("./static")))

    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
