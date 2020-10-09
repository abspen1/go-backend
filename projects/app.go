package projects

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/abspen1/restful-go/auth"
	"github.com/gomodule/redigo/redis"
)

/*
redis.Bool() – converts a single reply to a bool
redis.Bytes() – converts a single reply to a byte slice ([]byte)
redis.Float64() – converts a single reply to a float64
redis.Int() – converts a single reply to a int
redis.String() – converts a single reply to a string
redis.Values() – converts an array reply to an slice of individual replies
redis.Strings() – converts an array reply to an slice of strings ([]string)
redis.ByteSlices() – converts an array reply to an slice of byte slices ([][]byte)
redis.StringMap() – converts an array of strings (alternating key, value) into a map[string]string. Useful for HGETALL etc
*/

//Project struct
type Project struct {
	Language    string
	Description string
}

//RmProject struct
type RmProject struct {
	Language    string
	Description string
	Password    string
}

// GetString function
func GetString() []string {
	secret := auth.GoDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!", response)

	var unencoded *Project

	project1, _ := redis.Strings(client.Do("SMEMBERS", "projects"))
	fmt.Println(project1)

	len, _ := redis.Int(client.Do("SCARD", "projects"))

	i := 0

	var s []string

	for i < len {
		json.Unmarshal([]byte(project1[i]), &unencoded)
		s = append(s, unencoded.Language)
		s = append(s, unencoded.Description)
		i++
	}
	return (s)
}

// SetString function
func SetString(proj Project) {
	secret := auth.GoDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!", response)

	projEn, _ := json.Marshal(proj)
	client.Do("SADD", "projects", projEn)
	fmt.Println("Added to database")
}

//RmString func
func RmString(proj RmProject) bool {
	secret := auth.GoDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!", response)

	proj2 := Project{
		Language:    proj.Language,
		Description: proj.Description,
	}
	fmt.Println(proj2)

	projEn, _ := json.Marshal(proj2)
	client.Do("SREM", "projects", projEn)
	fmt.Println("Removed from database")
	return true
}

// CheckPass function
func CheckPass(proj RmProject) bool {
	pass := auth.GoDotEnvVariable("PASSWORD")

	if pass != proj.Password {
		fmt.Println("Incorrect Password")
		return false
	}
	return true
}
