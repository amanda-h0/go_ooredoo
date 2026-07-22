package main

import (
	"fmt"
	"log"
	"net" // for TCP connection
	"encoding/gob" // converts structs - bytes
	"bytes"
	"time"

	"system-monitor/client/system_functions"
	"system-monitor/models"
)

func main() {
	for {
		// Connect to TCP server
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("Failed to connect to TCP server: %v", err)

			// try again in 30 secs
			time.Sleep(30 * time.Second)
			continue
		}

		// CPU usage	
		CPUUsage, err := system_functions.GetCPUInfo()
		if err != nil {
			log.Fatalf("Failed to get CPU usage: %v", err)
			conn.Close()
			time.Sleep(30 * time.Second)
			continue
		}

		// memory usage
		MemoryUsage, err := system_functions.GetMemoryUsage()
		if err != nil {
			log.Fatalf("Failed to get memory usage: %v", err)
			conn.Close()
			time.Sleep(30 * time.Second)
			continue
		}

		// disk usage
		DiskUsage, err := system_functions.GetDiskUsage()
		if err != nil {
			log.Fatalf("Failed to get disk usage: %v", err)
			conn.Close()
			time.Sleep(30 * time.Second)
			continue
		}

		systemInfo := models.SystemInfo{
			CPUUsage:    CPUUsage,
			MemoryUsage: MemoryUsage,
			DiskUsage:   DiskUsage,
			Timestamp:   time.Now().Unix(),
		}

		// payload
		var buffer bytes.Buffer
		encoder := gob.NewEncoder(&buffer) // gob encoder that writes into buffer

		err = encoder.Encode(systemInfo) // turns to binary
		if err != nil {
			log.Fatalf("Failed to encode system information: %v", err)
			conn.Close()
			time.Sleep(30 * time.Second)
			continue
		}
		
		// send data to server
		_, err = conn.Write(buffer.Bytes()) // sends bytes
		if err != nil {
			log.Fatalf("Failed to send data to server: %v", err)
			conn.Close()
			time.Sleep(30 * time.Second)
			continue
		}

		fmt.Println("System information sent to server.")

		conn.Close()

		time.Sleep(30 * time.Second)
	}
}