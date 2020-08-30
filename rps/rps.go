package rps

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

// User struct
type User struct {
	Username string
	Wins     int
	Losses   int
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// SaveData function
func SaveData(user User) User {
	secret := goDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
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
			fmt.Println("wins:", wins)
			user.Wins = user.Wins + wins
		} else if key == "losses" {
			losses, _ := strconv.Atoi(value)
			fmt.Println("losses:", losses)
			user.Losses = user.Losses + losses
		}
	}
	client.Do("HSET", user.Username, "wins", user.Wins)
	client.Do("HSET", user.Username, "losses", user.Losses)

	return user
}

// GetData function
func GetData(user User) User {
	secret := goDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
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
