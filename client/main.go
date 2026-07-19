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
	// Connect to TCP server
	conn, err := net.Dial("tcp", "192.168.10.111:8080")
	if err != nil {
		log.Fatalf("Failed to connect to TCP server: %v", err)
	}
	defer conn.Close()

	// fmt.Println("\nSystem Information:")
	// fmt.Println("-------------------")

	// CPU usage	
	CPUUsage, err := system_functions.GetCPUInfo()
	if err != nil {
		log.Fatalf("Failed to get CPU usage: %v", err)
	}
	// fmt.Printf("CPU Usage: %.2f%%\n", CPUUsage)

	// memory usage
	MemoryUsage, err := system_functions.GetMemoryUsage()
	if err != nil {
		log.Fatalf("Failed to get memory usage: %v", err)
	}
	// fmt.Printf("Memory Usage: %.2f%%\n", MemoryUsage)

	// disk usage
	DiskUsage, err := system_functions.GetDiskUsage()
	if err != nil {
		log.Fatalf("Failed to get disk usage: %v", err)
	}
	// fmt.Printf("Disk Usage: %.2f%%\n\n", DiskUsage)

	systemInfo := models.SystemInfo{
		CPU:    CPUUsage,
		Mem: MemoryUsage,
		Disk:   DiskUsage,
		Time: time.Now().Unix(),
	}

	// payload
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer) // gob encoder that writes into buffer
	err = encoder.Encode(systemInfo) // turns to binary
	if err != nil {
		log.Fatalf("Failed to encode system information: %v", err)
	}
	
	// send data to server
	_, err = conn.Write(buffer.Bytes()) // sends bytes
	if err != nil {
		log.Fatalf("Failed to send data to server: %v", err)
	}

	fmt.Println("System information sent to server.")
}