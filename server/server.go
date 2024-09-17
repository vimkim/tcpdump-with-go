package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Start listening on a specific port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		// Accept new connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))

	// Send a response back to the client
	conn.Write([]byte("Hello from server!"))
}
