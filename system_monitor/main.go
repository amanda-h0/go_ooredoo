package main

import (
	"fmt"
	"log"

	"system-monitor/system_functions"
)

func main() {
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
}