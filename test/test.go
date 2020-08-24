package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

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

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

//RmString func
func RmString(proj RmProject) bool {
	secret := goDotEnvVariable("REDIS")
	pass := goDotEnvVariable("PASSWORD")

	if pass != proj.Password {
		fmt.Println("Incorrect Password")
		return false
	}

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

	projEn, _ := json.Marshal(proj2)
	client.Do("LREM", "projects", 0, projEn)
	fmt.Println("Removed from database")
	return true
}

//GetProjects function
func GetProjects() {
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

	var unencoded *Project

	project1, _ := redis.Strings(client.Do("LRANGE", "projects", 0, -1))

	len, _ := redis.Int(client.Do("LLEN", "projects"))

	i := 0

	var s []string

	for i < len {
		json.Unmarshal([]byte(project1[i]), &unencoded)
		s = append(s, unencoded.Language)
		s = append(s, unencoded.Description)
		i++
	}
}

func main() {
	// proj := RmProject{
	// 	Language:    "PYTHON",
	// 	Description: "Machine learning",
	// 	Password:    "Secure97",
	// }
	// RmString(proj)
	GetProjects()
}
