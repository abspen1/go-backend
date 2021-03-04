package tracker

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

// Stats struct
type Stats struct {
	Strategy string
	Count    float64
	ROI      float64
	WinLoss  float64
	Returns  float64
}

// Get func to get the current todos
func Get(w http.ResponseWriter, r *http.Request) {
	stats := getStats()

	json.NewEncoder(w).Encode(stats)
}

func getStats() Stats {
	var stats Stats
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	stats.setStatsStruct(client)

	return stats
}

// Set Stats
func (stats *Stats) setStatsStruct(client redis.Conn) {
	stats.Count, _ = redis.Float64(client.Do("GET", "CC-count"))
	stats.ROI, _ = redis.Float64(client.Do("GET", "CC-roi"))
	stats.Returns, _ = redis.Float64(client.Do("GET", "CC-returns"))
	stats.Strategy, _ = redis.String(client.Do("GET", "CC"))
}
