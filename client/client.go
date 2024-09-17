package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to the server. Type your message and press Enter.")
	fmt.Println("Type 'close' to close the connection.")

	reader := bufio.NewReader(os.Stdin)

	for {
		// Read user input
		fmt.Print("Enter message: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Trim newline characters
		input = strings.TrimSpace(input)

		// If the user types "close", break the loop to close the connection
		if input == "close" {
			fmt.Println("Closing the connection.")
			break
		}

		// Send the message to the server
		_, err = conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// Read the server's response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Println("Server response:", string(buffer[:n]))
	}
}
