package botsffl

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

// Leaders struct
type Leaders struct {
	Standings string
	Points    string
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// SetLeaders function
func SetLeaders() Leaders {
	var status Leaders
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

	status.Standings, _ = redis.String(client.Do("GET", "mw_standings"))
	status.Points, _ = redis.String(client.Do("GET", "mw_points"))

	return status
}
