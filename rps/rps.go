package rps

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// User struct
type User struct {
	Username string
	Wins     int
	Losses   int
}

// SaveData function
func SaveData(user User) User {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!", response)
	defer client.Close()

	check, _ := redis.String(client.Do("HGET", user.Username, "wins"))
	if check == "" {
		client.Do("HSET", user.Username, "wins", user.Wins)
		client.Do("HSET", user.Username, "losses", user.Losses)
		return user
	}

	client.Do("HSET", user.Username, "wins", user.Wins)
	client.Do("HSET", user.Username, "losses", user.Losses)

	return user
}

// GetData function
func GetData(user User) User {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!", response)
	defer client.Close()

	check, _ := redis.String(client.Do("HGET", user.Username, "wins"))
	if check == "" {
		client.Do("HSET", user.Username, "wins", user.Wins)
		client.Do("HSET", user.Username, "losses", user.Losses)
		return user
	}
	hash, _ := redis.StringMap(client.Do("HGETALL", user.Username))

	for key, value := range hash {
		if key == "wins" {
			wins, _ := strconv.Atoi(value)
			user.Wins = user.Wins + wins
		} else if key == "losses" {
			losses, _ := strconv.Atoi(value)
			user.Losses = user.Losses + losses
		}
	}
	client.Do("HSET", user.Username, "wins", user.Wins)
	client.Do("HSET", user.Username, "losses", user.Losses)

	return user
}
