package players

import (
	"log"
	"strconv"

	"github.com/abspen1/restful-go/auth"
	"github.com/gomodule/redigo/redis"
)

// Roster struct
type Roster struct {
	Team1  string
	Team2  string
	Team3  string
	Team4  string
	Team5  string
	Team6  string
	Team7  string
	Team8  string
	Team9  string
	Team10 string
	Team11 string
	Team12 string
}

// GetMidwestTeamNames function
func GetMidwestTeamNames() Roster {
	secret := auth.GoDotEnvVariable("REDIS")
	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(10), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var roster Roster
	i := 0
	for i < 12 {
		i++
		key := "mw_team_" + strconv.Itoa(i)
		name, _ := redis.String(client.Do("GET", key))
		roster.setTeamNames(client, i, name)
	}
	return roster
}

// GetWestTeamNames function
func GetWestTeamNames() Roster {
	secret := auth.GoDotEnvVariable("REDIS")
	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(10), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var roster Roster
	i := 0
	for i < 12 {
		i++
		key := "w_team_" + strconv.Itoa(i)
		name, _ := redis.String(client.Do("GET", key))
		roster.setTeamNames(client, i, name)
	}
	return roster
}

// GetNortheastTeamNames function
func GetNortheastTeamNames() Roster {
	secret := auth.GoDotEnvVariable("REDIS")
	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(10), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var roster Roster
	i := 0
	for i < 12 {
		i++
		key := "ne_team_" + strconv.Itoa(i)
		name, _ := redis.String(client.Do("GET", key))
		roster.setTeamNames(client, i, name)
	}
	return roster
}

// GetSoutheastTeamNames function
func GetSoutheastTeamNames() Roster {
	secret := auth.GoDotEnvVariable("REDIS")
	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(10), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var roster Roster
	i := 0
	for i < 12 {
		i++
		key := "se_team_" + strconv.Itoa(i)
		name, _ := redis.String(client.Do("GET", key))
		roster.setTeamNames(client, i, name)
	}
	return roster
}

func (teamNames *Roster) setTeamNames(client redis.Conn, n int, name string) {
	if n == 1 {
		teamNames.Team1 = name
	} else if n == 2 {
		teamNames.Team2 = name
	} else if n == 3 {
		teamNames.Team3 = name
	} else if n == 4 {
		teamNames.Team4 = name
	} else if n == 5 {
		teamNames.Team5 = name
	} else if n == 6 {
		teamNames.Team6 = name
	} else if n == 7 {
		teamNames.Team7 = name
	} else if n == 8 {
		teamNames.Team8 = name
	} else if n == 9 {
		teamNames.Team9 = name
	} else if n == 10 {
		teamNames.Team10 = name
	} else if n == 11 {
		teamNames.Team11 = name
	} else if n == 12 {
		teamNames.Team12 = name
	}
}
