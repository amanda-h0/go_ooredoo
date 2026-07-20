package models

type SystemInfo struct {
	IPAddress string
	CPUUsage float64
	MemoryUsage float64
	DiskUsage float64
	Timestamp int64
}