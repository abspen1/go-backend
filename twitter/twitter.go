package twitter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

// Data struct
type Data struct {
	Followers   int
	TweetsLiked int
	Tweets      int
	TweetsRead  int
	LatestTweet string
	Accuracy    string
}

// Get func to get the twitter data for Bottimus
func Get(w http.ResponseWriter, r *http.Request) {
	data := GetTwitterData()

	json.NewEncoder(w).Encode(data)
}

// GetTwitterData function
func GetTwitterData() Data {
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
	var data Data
	data.Followers, _ = redis.Int(client.Do("GET", "followers"))
	data.TweetsLiked, _ = redis.Int(client.Do("GET", "favorites"))
	data.Tweets, _ = redis.Int(client.Do("GET", "statuses"))
	data.TweetsRead, _ = redis.Int(client.Do("GET", "read"))
	data.LatestTweet, _ = redis.String(client.Do("GET", "recent"))
	// data.Accuracy, _ = redis.String(client.Do("GET", "pct"))
	// April 30th tweets read = 5458196
	fmt.Println("Just pulled Twitter data for CloudBot")
	return data
}
