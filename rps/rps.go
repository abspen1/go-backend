package rps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/badoux/checkmail"
	"github.com/gomodule/redigo/redis"
)

// User struct
type User struct {
	Username string
	Wins     int
	Losses   int
}

// GetRPS func to show the info of this endpoint
func GetRPS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rock Paper Scissors game save endpoint, nothing to see here!")
}

// PostRPS func to save the data
func PostRPS(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var rpsUser User
	_ = json.Unmarshal(info, &rpsUser)

	rpsUser = SaveData(rpsUser)

	json.NewEncoder(w).Encode(rpsUser)
}

// GetRPSLogin func to display what the endpoint is for
func GetRPSLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rock Paper Scissors game login endpoint, nothing to see here!")
}

// PostRPSLogin func to display RPS stats for user
func PostRPSLogin(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var rpsUser User
	_ = json.Unmarshal(info, &rpsUser)

	err := checkmail.ValidateFormat(rpsUser.Username)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Format Error")
		return
	}
	err = checkmail.ValidateHost(rpsUser.Username)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		fmt.Fprintf(w, "Error")
		return
	}

	rpsUser = GetData(rpsUser)

	json.NewEncoder(w).Encode(rpsUser)
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
