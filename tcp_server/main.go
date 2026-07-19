package main

import (
	"fmt"
	"net"
	"log"
	"encoding/gob"

	"system-monitor/models"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // close connection when function exits
	fmt.Printf("New client connected: %v\n", conn.RemoteAddr().String())

//  conn.RemoteAddr()

	decoder := gob.NewDecoder(conn) // reads bytes from connection and decodes into struct

	var systemInfo models.SystemInfo

	err := decoder.Decode(&systemInfo) // & for address to be filled
	if err != nil {
		log.Printf("Failed to decode system information from %s: %v\n", conn.RemoteAddr().String(), err)
		return
	}

	fmt.Println("\nSystem Information:")
	fmt.Println("-------------------")
	fmt.Printf("CPU Usage: %.2f%%\n", systemInfo.CPUUsage)
	fmt.Printf("Memory Usage: %.2f%%\n", systemInfo.MemoryUsage)
	fmt.Printf("Disk Usage: %.2f%%\n", systemInfo.DiskUsage)

	_, err = conn.Write([]byte("ACK\n")) // confirmation sent to client
	if err != nil { // if server fails to send to client
		log.Printf("Write error to %s: %v\n", conn.RemoteAddr().String(), err)
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