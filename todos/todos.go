package todos

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/abspen1/restful-go/auth"
	"github.com/gomodule/redigo/redis"
)

// Todos struct
type Todos struct {
	/*
			{
		    "Title": "Example",
		    "Completed": "false"
		    }
	*/
	Title     string
	Completed bool
}

// FullTodo struct
type FullTodo struct {
	Title     string
	Completed bool
	Id        int
}

// GetTodos function pulls todo hash from redis database
func GetTodos() []FullTodo {
	secret := auth.GoDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var unencoded *FullTodo

	// var todoSlice []Todos
	todoSlice, err := redis.Strings(client.Do("SMEMBERS", "Todos"))
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(todoSlice)

	len, err := redis.Int(client.Do("SCARD", "Todos"))
	if err != nil {
		fmt.Println(err)
	}

	i := 0

	var fullTodo FullTodo
	var todoSliceUnencoded []FullTodo

	for i < len {
		json.Unmarshal([]byte(todoSlice[i]), &unencoded)
		fullTodo.Title = unencoded.Title
		fullTodo.Completed = unencoded.Completed
		fullTodo.Id = unencoded.Id
		todoSliceUnencoded = append(todoSliceUnencoded, fullTodo)
		i++
	}

	return todoSliceUnencoded
}

// AddTodo function
func AddTodo(todos Todos) bool {
	secret := auth.GoDotEnvVariable("REDIS")

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
	client.Do("INCR", "todo-id")

	var fullTodo FullTodo

	fullTodo.Title = todos.Title
	fullTodo.Completed = todos.Completed
	id, err := redis.Int(client.Do("GET", "todo-id"))
	fullTodo.Id = id
	if err != nil {
		return false
	}

	todoEn, err := json.Marshal(fullTodo)
	if err != nil {
		return false
	}

	client.Do("SADD", "Todos", todoEn)

	return true
}

// RmTodo function
func RmTodo(fullTodo FullTodo) bool {
	secret := auth.GoDotEnvVariable("REDIS")

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

	todo := FullTodo{
		Title:     fullTodo.Title,
		Completed: fullTodo.Completed,
		Id:        fullTodo.Id,
	}

	todoEn, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Marshal error")
		return false
	}
	client.Do("SREM", "Todos", todoEn)
	client.Do("SADD", "Todos-stashed", todoEn)

	return true
}
