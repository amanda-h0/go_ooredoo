package database

type Metric struct {
	ID          int     `json:"id"`
	IPAddress   string  `json:"ip_address"`
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
	Timestamp   string  `json:"timestamp"`
}

func GetUniqueIPs() ([]string, error) {
	// unique IPs
	rows, err := DB.Query("SELECT DISTINCT ip_address FROM system_metrics")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ips []string
	for rows.Next() {
		var ip string
		err := rows.Scan(&ip) // reads and writes into the var
		if err != nil {
			return nil, err
		}
		ips = append(ips, ip)
	}
	return ips, nil
}

func GetMetricsByIP(ip string) ([]Metric, error) {
	rows, err := DB.Query(`
		SELECT id, ip_address, cpu_usage, memory_usage, disk_usage, timestamp
		FROM system_metrics
		WHERE ip_address = ?
		ORDER BY id DESC
	`, ip)

	if err != nil {
			return nil, err
		}

	var metrics []Metric
	for rows.Next() {
		var m Metric

		err := rows.Scan(&m.ID, &m.IPAddress, &m.CPUUsage, &m.MemoryUsage, &m.DiskUsage, &m.Timestamp)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, m)
	}
	return metrics, nil
}