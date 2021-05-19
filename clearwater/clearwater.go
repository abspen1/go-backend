package clearwater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

//Trip struct
type Trip struct {
	Name      string
	Email     string
	StartDate string
	EndDate   string
	Password  string
	Color     string
}

// GetTrips func to get the current trips
func GetTrips(w http.ResponseWriter, r *http.Request) {
	s := GetString()

	json.NewEncoder(w).Encode(s)
}

// PostTrips func to add trips to database
func PostTrips(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var trip Trip
	_ = json.Unmarshal(info, &trip)
	SetString(trip)
	fmt.Fprintf(w, "Success")
}

// PostRemoveTrips func to remove trips from database
func PostRemoveTrips(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var trip Trip
	_ = json.Unmarshal(info, &trip)

	if RmString(trip) {
		fmt.Fprintf(w, "POST remove worked!")
	} else {
		fmt.Fprintf(w, "Error")
	}

}

// GetString function
func GetString() []string {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host, redis.DialDatabase(4), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}

	// var unencoded *Trip

	trips1, _ := redis.Strings(client.Do("SMEMBERS", "clearwater-trips"))
	fmt.Println(trips1)

	// len, _ := redis.Int(client.Do("SCARD", "clearwater-trips"))

	// i := 0

	// var s []string

	// for i < len {
	// 	json.Unmarshal([]byte(trips1[i]), &unencoded)
	// 	s = append(s, unencoded.StartDate)
	// 	s = append(s, unencoded.EndDate)
	// 	s = append(s, unencoded.Email)
	// 	s = append(s, unencoded.Name)
	// 	s = append(s, unencoded.Color)
	// 	s = append(s, unencoded.Password)
	// 	i++
	// }
	return (trips1)
}

// SetString function
func SetString(trip Trip) {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host, redis.DialDatabase(4), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}

	projEn, _ := json.Marshal(trip)
	client.Do("SADD", "clearwater-trips", projEn)
	fmt.Println("Added a new trip to clearwater DB")
}

//RmString func
func RmString(trip Trip) bool {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host, redis.DialDatabase(4), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}

	proj2 := Trip{
		StartDate: trip.StartDate,
		EndDate:   trip.EndDate,
		Email:     trip.Email,
		Name:      trip.Name,
		Color:     trip.Color,
		Password:  trip.Password,
	}

	fmt.Println(proj2)

	projEn, _ := json.Marshal(proj2)
	client.Do("SREM", "clearwater-trips", projEn)
	fmt.Println("Removed trip from clearwater DB")
	return true
}
