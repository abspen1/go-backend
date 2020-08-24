package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/abspen1/restful-go/email"

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

// SetString function
func SetString(proj Project) {
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

	projEn, _ := json.Marshal(proj)
	// fmt.Println(proj)
	client.Do("SADD", "projects", projEn)
	fmt.Println("Added to database")
}

//RmString func
func RmString(proj RmProject) bool {
	secret := goDotEnvVariable("REDIS")
	pass := goDotEnvVariable("PASSWORD-TEST")

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
	fmt.Println(proj2)

	projEn, _ := json.Marshal(proj2)
	client.Do("SREM", "projects", projEn)
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

	client.Do("DEL", "projects")

	project1, _ := redis.Strings(client.Do("SMEMBERS", "projects"))
	fmt.Println(project1)

	len, _ := redis.Int(client.Do("SCARD", "projects"))
	fmt.Println(len)

	i := 0
	if len > 0 {
		for i < len {
			json.Unmarshal([]byte(project1[i]), &unencoded)
			fmt.Println(unencoded.Language)
			fmt.Println(unencoded.Description)
			i++
		}
	}
}

func main() {
	// 	var proj Project
	// 	proj.Language = "PYTHON"
	// 	proj.Description = "Something test"
	// fullStr = proj.Language + ""

	// SetString(proj)
	// proj := RmProject{
	// 	Language:    "PYTHON",
	// 	Description: "Testing This Out",
	// 	Password:    "Secure97",
	// }
	// proj2 := Project{
	// 	Language:    "PYTHON",
	// 	Description: "Testing This Out",
	// }
	// RmString(proj)
	// SetString(proj2)
	// GetProjects()
	var info email.Info

	info.Name = "Austin"
	info.Email = "abspencer2097@gmail.com"
	info.Message = "Hello World"
	email.SendEmail(info)
}
