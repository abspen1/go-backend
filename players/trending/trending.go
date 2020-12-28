package trending

import (
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// DayAdd struct
type DayAdd struct {
	DayAdd1  string
	DayAdd2  string
	DayAdd3  string
	DayAdd4  string
	DayAdd5  string
	DayAdd6  string
	DayAdd7  string
	DayAdd8  string
	DayAdd9  string
	DayAdd10 string
}

// DayDrop struct
type DayDrop struct {
	DayDrop1  string
	DayDrop2  string
	DayDrop3  string
	DayDrop4  string
	DayDrop5  string
	DayDrop6  string
	DayDrop7  string
	DayDrop8  string
	DayDrop9  string
	DayDrop10 string
}

// WeekAdd struct
type WeekAdd struct {
	WeekAdd1  string
	WeekAdd2  string
	WeekAdd3  string
	WeekAdd4  string
	WeekAdd5  string
	WeekAdd6  string
	WeekAdd7  string
	WeekAdd8  string
	WeekAdd9  string
	WeekAdd10 string
}

// WeekDrop struct
type WeekDrop struct {
	WeekDrop1  string
	WeekDrop2  string
	WeekDrop3  string
	WeekDrop4  string
	WeekDrop5  string
	WeekDrop6  string
	WeekDrop7  string
	WeekDrop8  string
	WeekDrop9  string
	WeekDrop10 string
}

//DailyAdd function will return the DayAdd trending players from sleeper API
func DailyAdd() DayAdd {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")
	client, err := redis.Dial("tcp", host, redis.DialDatabase(12), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var daily DayAdd

	i := 0
	for i < 11 {
		i++
		key := "daily_trendingA_" + strconv.Itoa(i)
		player, _ := redis.String(client.Do("GET", key))
		switch i {
		case 1:
			daily.DayAdd1 = player
			break
		case 2:
			daily.DayAdd2 = player
			break
		case 3:
			daily.DayAdd3 = player
			break
		case 4:
			daily.DayAdd4 = player
			break
		case 5:
			daily.DayAdd5 = player
			break
		case 6:
			daily.DayAdd6 = player
			break
		case 7:
			daily.DayAdd7 = player
			break
		case 8:
			daily.DayAdd8 = player
			break
		case 9:
			daily.DayAdd9 = player
			break
		case 10:
			daily.DayAdd10 = player
			break
		}
	}
	return daily
}

//DailyDrop function will return the DayDrop trending players from sleeper API
func DailyDrop() DayDrop {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")
	client, err := redis.Dial("tcp", host, redis.DialDatabase(12), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var daily DayDrop

	i := 0
	for i < 11 {
		i++
		key := "daily_trendingD_" + strconv.Itoa(i)
		player, _ := redis.String(client.Do("GET", key))
		switch i {
		case 1:
			daily.DayDrop1 = player
			break
		case 2:
			daily.DayDrop2 = player
			break
		case 3:
			daily.DayDrop3 = player
			break
		case 4:
			daily.DayDrop4 = player
			break
		case 5:
			daily.DayDrop5 = player
			break
		case 6:
			daily.DayDrop6 = player
			break
		case 7:
			daily.DayDrop7 = player
			break
		case 8:
			daily.DayDrop8 = player
			break
		case 9:
			daily.DayDrop9 = player
			break
		case 10:
			daily.DayDrop10 = player
			break
		}
	}
	return daily
}

// WeeklyAdd function that returns the top trending players past 5 days on sleeper API
func WeeklyAdd() WeekAdd {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")
	client, err := redis.Dial("tcp", host, redis.DialDatabase(12), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var daily WeekAdd

	i := 0
	for i < 11 {
		i++
		key := "weekly_trendingA_" + strconv.Itoa(i)
		player, _ := redis.String(client.Do("GET", key))
		switch i {
		case 1:
			daily.WeekAdd1 = player
			break
		case 2:
			daily.WeekAdd2 = player
			break
		case 3:
			daily.WeekAdd3 = player
			break
		case 4:
			daily.WeekAdd4 = player
			break
		case 5:
			daily.WeekAdd5 = player
			break
		case 6:
			daily.WeekAdd6 = player
			break
		case 7:
			daily.WeekAdd7 = player
			break
		case 8:
			daily.WeekAdd8 = player
			break
		case 9:
			daily.WeekAdd9 = player
			break
		case 10:
			daily.WeekAdd10 = player
			break
		}
	}
	return daily
}

// WeeklyDrop function that returns the top trending players past 5 days on sleeper API
func WeeklyDrop() WeekDrop {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")
	client, err := redis.Dial("tcp", host, redis.DialDatabase(12), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var daily WeekDrop

	i := 0
	for i < 11 {
		i++
		key := "weekly_trendingD_" + strconv.Itoa(i)
		player, _ := redis.String(client.Do("GET", key))
		switch i {
		case 1:
			daily.WeekDrop1 = player
			break
		case 2:
			daily.WeekDrop2 = player
			break
		case 3:
			daily.WeekDrop3 = player
			break
		case 4:
			daily.WeekDrop4 = player
			break
		case 5:
			daily.WeekDrop5 = player
			break
		case 6:
			daily.WeekDrop6 = player
			break
		case 7:
			daily.WeekDrop7 = player
			break
		case 8:
			daily.WeekDrop8 = player
			break
		case 9:
			daily.WeekDrop9 = player
			break
		case 10:
			daily.WeekDrop10 = player
			break
		}
	}
	return daily
}
