package database

import (
	"fmt"
	"system-monitor/models"
)


func InsertSystemInfo(info models.SystemInfo) error {

	fmt.Println("Inserting into database...")

	query := `
	INSERT INTO system_metrics
	(
	ip_address,
	cpu_usage,
	memory_usage,
	disk_usage,
	timestamp
	)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		info.IPAddress,
		info.CPUUsage,
		info.MemoryUsage,
		info.DiskUsage,
		info.Timestamp,
	)

	return err
}