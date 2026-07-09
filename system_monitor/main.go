package main

import (
	"fmt"
	"log"
	"net" // for TCP connection

	"system-monitor/system_functions"
)

func main() {
	// Connect to TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to connect to TCP server: %v", err)
	}
	defer conn.Close()

	fmt.Println("\nSystem Information:")
	fmt.Println("-------------------")

	// CPU usage	
	CPUUsage, err := system_functions.GetCPUInfo()
	if err != nil {
		log.Fatalf("Failed to get CPU usage: %v", err)
	}
	fmt.Printf("CPU Usage: %.2f%%\n", CPUUsage)

	// memory usage
	MemoryUsage, err := system_functions.GetMemoryUsage()
	if err != nil {
		log.Fatalf("Failed to get memory usage: %v", err)
	}
	fmt.Printf("Memory Usage: %.2f%%\n", MemoryUsage)

	// disk usage
	DiskUsage, err := system_functions.GetDiskUsage()
	if err != nil {
		log.Fatalf("Failed to get disk usage: %v", err)
	}
	fmt.Printf("Disk Usage: %.2f%%\n\n", DiskUsage)

	// format payload
	payload := fmt.Sprintf("CPU Usage: %.2f%%\nMemory Usage: %.2f%%\nDisk Usage: %.2f%%\n", CPUUsage, MemoryUsage, DiskUsage)
	
	// send data to server
	_, err = conn.Write([]byte(payload))
	if err != nil {
		log.Fatalf("Failed to send data to server: %v", err)
	}
}