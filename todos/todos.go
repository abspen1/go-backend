package todos

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

// Todos struct
type Todos struct {
	/*
			{
		    "UserId": 1,
		    "Id": 1,
		    "Title": "selected either the"
		    "Completed" false
		    }
	*/
	Title     string
	Completed string
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// GetTodos function pulls todo hash from redis database
func GetTodos() []Todos {
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
	// var todoSlice []Todos
	strMap, _ := redis.StringMap(client.Do("HGETALL", "Todos"))

	var todos Todos
	var todoSlice []Todos

	for key, item := range strMap {
		todos.Title = key
		todos.Completed = item
		todoSlice = append(todoSlice, todos)
	}
	return todoSlice
}

// AddTodo function
func AddTodo(todos Todos) bool {
	secret := goDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
	if err != nil {
		fmt.Println("Couldn't connect to Redis")
		return false
	}
	_, err = client.Do("AUTH", secret)
	if err != nil {
		fmt.Println("Invalid Redis authentification")
		return false
	}
	defer client.Close()
	// var todoSlice []Todos
	title := todos.Title
	completed := todos.Completed

	client.Do("HSET", "Todos", title, completed)

	return true
}
