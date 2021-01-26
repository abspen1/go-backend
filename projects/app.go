package projects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

// GetProjects func to get the current projects
func GetProjects(w http.ResponseWriter, r *http.Request) {
	s := GetString()

	json.NewEncoder(w).Encode(s)
}

// GetRmProjects func to show what this endpoint displays
func GetRmProjects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove projects endpoint, nothing to see here!")
}

// PostProjects func to add projects to database
func PostProjects(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var proj Project
	_ = json.Unmarshal(info, &proj)
	SetString(proj)
	fmt.Fprintf(w, string(info))

	// fmt.Fprintf(w, "Test POST endpoint worked!")
}

// PostRmProjects func to remove projects from database
func PostRmProjects(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var proj RmProject
	_ = json.Unmarshal(info, &proj)

	if CheckPass(proj) {
		if RmString(proj) {
			fmt.Fprintf(w, "POST remove worked!")
		} else {
			fmt.Fprintf(w, "Error")
		}
	} else {
		fmt.Fprintf(w, "Err")
	}

	// fmt.Fprintf(w, "Test POST endpoint worked!")
}

// GetString function
func GetString() []string {
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

	projEn, _ := json.Marshal(proj)
	client.Do("SADD", "projects", projEn)
	fmt.Println("Added to database")
}

//RmString func
func RmString(proj RmProject) bool {
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
	pass := os.Getenv("PASSWORD")

	if pass != proj.Password {
		fmt.Println("Incorrect Password")
		return false
	}
	return true
}
