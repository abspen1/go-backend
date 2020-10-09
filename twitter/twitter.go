package twitter

import (
	"fmt"
	"log"

	"github.com/abspen1/restful-go/auth"
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
	secret := auth.GoDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var data Data
	data.Followers, _ = redis.Int(client.Do("GET", "tendie_followers"))
	data.TweetsLiked, _ = redis.Int(client.Do("GET", "tendie_favorites"))
	data.Tweets, _ = redis.Int(client.Do("GET", "tendie_statuses"))
	data.TweetsRead, _ = redis.Int(client.Do("GET", "tendie_read"))
	data.LatestTweet, _ = redis.String(client.Do("GET", "tendie_recent"))
	data.Accuracy, _ = redis.String(client.Do("GET", "tendie_pct"))
	fmt.Println(data)
	return data
}
