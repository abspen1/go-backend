package botsffl

import (
	"log"
	"os"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

// Leaders struct
type Leaders struct {
	Standings1  string
	Standings2  string
	Standings3  string
	Standings4  string
	Standings5  string
	Standings6  string
	Standings7  string
	Standings8  string
	Standings9  string
	Standings10 string
	Standings11 string
	Standings12 string
	Points1     string
	Points2     string
	Points3     string
	Points4     string
	Points5     string
	Points6     string
	Points7     string
	Points8     string
	Points9     string
	Points10    string
	Points11    string
	Points12    string
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func (status *Leaders) setStandings(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.Standings1, _ = redis.String(client.Do("GET", "mw_standings_1"))
		case 2:
			status.Standings2, _ = redis.String(client.Do("GET", "mw_standings_2"))
		case 3:
			status.Standings3, _ = redis.String(client.Do("GET", "mw_standings_3"))
		case 4:
			status.Standings4, _ = redis.String(client.Do("GET", "mw_standings_4"))
		case 5:
			status.Standings5, _ = redis.String(client.Do("GET", "mw_standings_5"))
		case 6:
			status.Standings6, _ = redis.String(client.Do("GET", "mw_standings_6"))
		case 7:
			status.Standings7, _ = redis.String(client.Do("GET", "mw_standings_7"))
		case 8:
			status.Standings8, _ = redis.String(client.Do("GET", "mw_standings_8"))
		case 9:
			status.Standings9, _ = redis.String(client.Do("GET", "mw_standings_9"))
		case 10:
			status.Standings10, _ = redis.String(client.Do("GET", "mw_standings_10"))
		case 11:
			status.Standings11, _ = redis.String(client.Do("GET", "mw_standings_11"))
		case 12:
			status.Standings12, _ = redis.String(client.Do("GET", "mw_standings_12"))
		}
		i++
	}
}

func (status *Leaders) setPoints(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.Points1, _ = redis.String(client.Do("GET", "mw_points_1"))
		case 2:
			status.Points2, _ = redis.String(client.Do("GET", "mw_points_2"))
		case 3:
			status.Points3, _ = redis.String(client.Do("GET", "mw_points_3"))
		case 4:
			status.Points4, _ = redis.String(client.Do("GET", "mw_points_4"))
		case 5:
			status.Points5, _ = redis.String(client.Do("GET", "mw_points_5"))
		case 6:
			status.Points6, _ = redis.String(client.Do("GET", "mw_points_6"))
		case 7:
			status.Points7, _ = redis.String(client.Do("GET", "mw_points_7"))
		case 8:
			status.Points8, _ = redis.String(client.Do("GET", "mw_points_8"))
		case 9:
			status.Points9, _ = redis.String(client.Do("GET", "mw_points_9"))
		case 10:
			status.Points10, _ = redis.String(client.Do("GET", "mw_points_10"))
		case 11:
			status.Points11, _ = redis.String(client.Do("GET", "mw_points_11"))
		case 12:
			status.Points12, _ = redis.String(client.Do("GET", "mw_points_12"))
		}
		i++
	}
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
	status.setStandings(client)
	status.setPoints(client)

	return status
}
