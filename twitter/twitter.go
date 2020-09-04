package twitter

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
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

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// GetTwitterData function
func GetTwitterData() Data {
	secret := goDotEnvVariable("REDIS")

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
