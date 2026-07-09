package system_functions

import (
	"github.com/shirou/gopsutil/v4/mem"
)

func GetMemoryUsage() (float64, error) {
	// fetch virtual memory statistics
	stats, err := mem.VirtualMemory()

	if err != nil {
		return 0.0, err
	}

	//  used percent = (used memory / total memory) * 100
	return stats.UsedPercent, nil
}
