package main

import (
	"fmt"
	"net"
	"io" // I/O interfaces - used to detect end of stream/file (EOF)
	"log"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // close connection when function exits
	fmt.Printf("New client connected: %v\n", conn.RemoteAddr().String())

	buffer := make([]byte, 1024) // temp bytes to hold incoming data

	for {
		n, err := conn.Read(buffer) // reads bytes and returns number of bytes read (n)
		if err != nil {
			if err == io.EOF { // EOF = client closed connection
				fmt.Printf("Client %s disconnected.\n", conn.RemoteAddr().String())
			} else {
				log.Printf("Read error from %s: %v\n", conn.RemoteAddr().String(), err)
			}
			break // triggers deferred conn.Close()
		}

		message := buffer[:n] // slice buffer to isolate the bytes read (what the client sent)
		fmt.Print(string(message)) // raw bytes -> human-readable string

		_, err = conn.Write([]byte("ACK\n")) // confirmation sent to client
		if err != nil { // if server fails to send to client
			log.Printf("Write error to %s: %v\n", conn.RemoteAddr().String(), err)
			break
		}
	}
}

func main() {
	// 1. Listen (TCP :8080)
	listener, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on :8080")

	// 2. Accept connections
	for {
		conn, err := listener.Accept() // accept connection & initiate stream to read data, returns a net.Conn socket object
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// 3. Handle connection
		go handleConnection(conn) // handle connection in a separate goroutine
	}
}