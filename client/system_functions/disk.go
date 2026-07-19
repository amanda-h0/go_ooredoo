package system_functions

import (
	"github.com/shirou/gopsutil/v4/disk"
)

func GetDiskUsage() (float64, error) {
	stats, err := disk.Usage("/") // specify disk path
	if err != nil {
		return 0.0, err
	}

	return stats.UsedPercent, nil
}
