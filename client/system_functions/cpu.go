package system_functions

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUInfo() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	// check for errors
	if err != nil {
		fmt.Println("An error occured.")
		return 0.0, err
	}

	// ensure percentages slice is not empty
	if len(percentages) == 0 {
		return 0.0, nil
	}

	// return the first value
	return percentages[0], nil
}
