package main

import (
	"encoding/json"
	"log"
	"net/http"

	"system-monitor/tcp_server/database"
)

func serversHandler(w http.ResponseWriter, r *http.Request) {
	ips, err := database.GetUniqueIPs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ips) // returns json for /api/servers
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	server := r.URL.Query().Get("server")
	if server == "" {
		http.Error(w, "missing 'server' query parameter", http.StatusBadRequest)
		return
	}

	metrics, err := database.GetMetricsByIP(server)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func withCORS(next http.HandlerFunc) http.HandlerFunc { // to accept frontend from diff port
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func main() {
	database.InitDB()

	http.HandleFunc("/api/servers", withCORS(serversHandler)) // cross origin resource sharing
	http.HandleFunc("/api/metrics", withCORS(metricsHandler))

	log.Println("API server listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil)) // starts the HTTP server and blocks if the server crashes
}