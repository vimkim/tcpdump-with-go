package main

import (
	"bufio"
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
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	// Create a buffered reader for the connection
	reader := bufio.NewReader(conn)

	for {
		// Read data from the connection
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading data or client disconnected:", err)
			break
		}

		// Print the received message
		fmt.Println("Received:", message)

		// Send a response back to the client
		_, err = conn.Write([]byte("Hello from server!\n"))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			break
		}
	}

	fmt.Println("Connection closed by client:", conn.RemoteAddr())
}
