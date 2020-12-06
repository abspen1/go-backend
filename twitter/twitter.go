package twitter

import (
	"fmt"
	"log"
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
	fmt.Println(data)
	return data
}
