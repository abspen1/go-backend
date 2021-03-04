package tracker

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

// CoveredCall struct
type CoveredCall struct {
	Stock   string
	Open    string
	IV      float64
	STD     float64
	DTE     float64
	ROI     float64
	Returns float64
	Close   string
}

// Post func to add todos to the database
func Post(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Todo not added due to an error")
}

func AddCoveredCall(cc CoveredCall) {
	var stats Stats
	var currStats Stats
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

	wins, _ := redis.Float64(client.Do("GET", "CC-wins"))
	if wins > 0 {
		wins++
		client.Do("SET", "CC-wins", wins)
	}

	currStats.getCurrentCC(client)
	stats.adjustStats(currStats, wins)

}

func (stats *Stats) adjustStats(currStats Stats, wins float64) {
	stats.Strategy = "Covered Call"
	stats.Count = currStats.Count + 1
	stats.Returns = currStats.Returns + stats.Returns
	stats.ROI = (currStats.ROI + stats.ROI) / stats.Count
	stats.WinLoss = stats.Count / wins
}

func (stats *Stats) getCurrentCC(client redis.Conn) {
	stats.Count, _ = redis.Float64(client.Do("GET", "CC-count"))
	stats.ROI, _ = redis.Float64(client.Do("GET", "CC-roi"))

	stats.Returns, _ = redis.Float64(client.Do("GET", "CC-returns"))
}
