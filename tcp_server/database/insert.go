package database

import (
	"fmt"
	"time"
	"system-monitor/models"
)


func InsertSystemInfo(info models.SystemInfo) error {

	fmt.Println("Inserting into database...")

	readableTime := time.Unix(info.Timestamp, 0).Format("2006-01-02 15:04:05")

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
		readableTime,
	)

	return err
}