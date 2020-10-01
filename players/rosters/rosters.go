package rosters

import (
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

// Team struct
type Team struct {
	Team1  []Roster
	Team2  []Roster
	Team3  []Roster
	Team4  []Roster
	Team5  []Roster
	Team6  []Roster
	Team7  []Roster
	Team8  []Roster
	Team9  []Roster
	Team10 []Roster
	Team11 []Roster
	Team12 []Roster
}

// Roster struct
type Roster struct {
	Player1    string `json:"player1"`
	Position1  string `json:"position1"`
	Player2    string `json:"player2"`
	Position2  string `json:"position2"`
	Player3    string `json:"player3"`
	Position3  string `json:"position3"`
	Player4    string `json:"player4"`
	Position4  string `json:"position4"`
	Player5    string `json:"player5"`
	Position5  string `json:"position5"`
	Player6    string `json:"player6"`
	Position6  string `json:"position6"`
	Player7    string `json:"player7"`
	Position7  string `json:"position7"`
	Player8    string `json:"player8"`
	Position8  string `json:"position8"`
	Player9    string `json:"player9"`
	Position9  string `json:"position9"`
	Player10   string `json:"player10"`
	Position10 string `json:"position10"`
	Player11   string `json:"player11"`
	Position11 string `json:"position11"`
	Player12   string `json:"player12"`
	Position12 string `json:"position12"`
	Player13   string `json:"player13"`
	Position13 string `json:"position13"`
	Player14   string `json:"player14"`
	Position14 string `json:"position14"`
	Player15   string `json:"player15"`
	Position15 string `json:"position15"`
	Player16   string `json:"player16"`
	Position16 string `json:"position16"`
	Player17   string `json:"player17"`
	Position17 string `json:"position17"`
	Player18   string `json:"player18"`
	Position18 string `json:"position18"`
	Player19   string `json:"player19"`
	Position19 string `json:"position19"`
	Player20   string `json:"player20"`
	Position20 string `json:"position20"`
	Player21   string `json:"player21"`
	Position21 string `json:"position21"`
	Player22   string `json:"player22"`
	Position22 string `json:"position22"`
	Player23   string `json:"player23"`
	Position23 string `json:"position23"`
	Player24   string `json:"player24"`
	Position24 string `json:"position24"`
	Player25   string `json:"player25"`
	Position25 string `json:"position25"`
	Player26   string `json:"player26"`
	Position26 string `json:"position26"`
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// GetMidwestTeamRosters function
func GetMidwestTeamRosters() Team {
	secret := goDotEnvVariable("REDIS")
	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(10), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var Rosters Roster
	var Teams Team

	i := 0
	for i < 12 {
		i++
		var RosterList []Roster
		key := "mw_team_" + strconv.Itoa(i)
		name, _ := redis.String(client.Do("GET", key))
		hash, _ := redis.StringMap(client.Do("HGETALL", name))
		j := 0
		for key, value := range hash {
			j++
			switch j {
			case 1:
				Rosters.Player1 = key
				Rosters.Position1 = value
				break
			case 2:
				Rosters.Player2 = key
				Rosters.Position2 = value
				break
			case 3:
				Rosters.Player3 = key
				Rosters.Position3 = value
				break
			case 4:
				Rosters.Player4 = key
				Rosters.Position4 = value
				break
			case 5:
				Rosters.Player5 = key
				Rosters.Position5 = value
				break
			case 6:
				Rosters.Player6 = key
				Rosters.Position6 = value
				break
			case 7:
				Rosters.Player7 = key
				Rosters.Position7 = value
				break
			case 8:
				Rosters.Player8 = key
				Rosters.Position8 = value
				break
			case 9:
				Rosters.Player9 = key
				Rosters.Position9 = value
				break
			case 10:
				Rosters.Player10 = key
				Rosters.Position10 = value
				break
			case 11:
				Rosters.Player11 = key
				Rosters.Position11 = value
				break
			case 12:
				Rosters.Player12 = key
				Rosters.Position12 = value
				break
			case 13:
				Rosters.Player13 = key
				Rosters.Position13 = value
				break
			case 14:
				Rosters.Player14 = key
				Rosters.Position14 = value
				break
			case 15:
				Rosters.Player15 = key
				Rosters.Position15 = value
				break
			case 16:
				Rosters.Player16 = key
				Rosters.Position16 = value
				break
			case 17:
				Rosters.Player17 = key
				Rosters.Position17 = value
				break
			case 18:
				Rosters.Player18 = key
				Rosters.Position18 = value
				break
			case 19:
				Rosters.Player19 = key
				Rosters.Position19 = value
				break
			case 20:
				Rosters.Player20 = key
				Rosters.Position20 = value
				break
			case 21:
				Rosters.Player21 = key
				Rosters.Position21 = value
				break
			case 22:
				Rosters.Player22 = key
				Rosters.Position22 = value
				break
			case 23:
				Rosters.Player23 = key
				Rosters.Position23 = value
				break
			case 24:
				Rosters.Player24 = key
				Rosters.Position24 = value
				break
			case 25:
				Rosters.Player25 = key
				Rosters.Position25 = value
				break
			case 26:
				Rosters.Player26 = key
				Rosters.Position26 = value
				break
			}
		}
		RosterList = append(RosterList, Rosters)
		Teams.setTeamNames(client, i, RosterList)
	}
	return Teams
}

// GetNortheastTeamRosters function
func GetNortheastTeamRosters() Team {
	secret := goDotEnvVariable("REDIS")
	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(10), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var Rosters Roster
	var Teams Team

	i := 0
	for i < 12 {
		i++
		var RosterList []Roster
		key := "ne_team_" + strconv.Itoa(i)
		name, _ := redis.String(client.Do("GET", key))
		hash, _ := redis.StringMap(client.Do("HGETALL", name))
		j := 0
		for key, value := range hash {
			j++
			switch j {
			case 1:
				Rosters.Player1 = key
				Rosters.Position1 = value
				break
			case 2:
				Rosters.Player2 = key
				Rosters.Position2 = value
				break
			case 3:
				Rosters.Player3 = key
				Rosters.Position3 = value
				break
			case 4:
				Rosters.Player4 = key
				Rosters.Position4 = value
				break
			case 5:
				Rosters.Player5 = key
				Rosters.Position5 = value
				break
			case 6:
				Rosters.Player6 = key
				Rosters.Position6 = value
				break
			case 7:
				Rosters.Player7 = key
				Rosters.Position7 = value
				break
			case 8:
				Rosters.Player8 = key
				Rosters.Position8 = value
				break
			case 9:
				Rosters.Player9 = key
				Rosters.Position9 = value
				break
			case 10:
				Rosters.Player10 = key
				Rosters.Position10 = value
				break
			case 11:
				Rosters.Player11 = key
				Rosters.Position11 = value
				break
			case 12:
				Rosters.Player12 = key
				Rosters.Position12 = value
				break
			case 13:
				Rosters.Player13 = key
				Rosters.Position13 = value
				break
			case 14:
				Rosters.Player14 = key
				Rosters.Position14 = value
				break
			case 15:
				Rosters.Player15 = key
				Rosters.Position15 = value
				break
			case 16:
				Rosters.Player16 = key
				Rosters.Position16 = value
				break
			case 17:
				Rosters.Player17 = key
				Rosters.Position17 = value
				break
			case 18:
				Rosters.Player18 = key
				Rosters.Position18 = value
				break
			case 19:
				Rosters.Player19 = key
				Rosters.Position19 = value
				break
			case 20:
				Rosters.Player20 = key
				Rosters.Position20 = value
				break
			case 21:
				Rosters.Player21 = key
				Rosters.Position21 = value
				break
			case 22:
				Rosters.Player22 = key
				Rosters.Position22 = value
				break
			case 23:
				Rosters.Player23 = key
				Rosters.Position23 = value
				break
			case 24:
				Rosters.Player24 = key
				Rosters.Position24 = value
				break
			case 25:
				Rosters.Player25 = key
				Rosters.Position25 = value
				break
			case 26:
				Rosters.Player26 = key
				Rosters.Position26 = value
				break
			}
		}
		RosterList = append(RosterList, Rosters)
		Teams.setTeamNames(client, i, RosterList)
	}
	return Teams
}

// GetWestTeamRosters function
func GetWestTeamRosters() Team {
	secret := goDotEnvVariable("REDIS")
	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(10), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var Rosters Roster
	var Teams Team

	i := 0
	for i < 12 {
		i++
		var RosterList []Roster
		key := "w_team_" + strconv.Itoa(i)
		name, _ := redis.String(client.Do("GET", key))
		hash, _ := redis.StringMap(client.Do("HGETALL", name))
		j := 0
		for key, value := range hash {
			j++
			switch j {
			case 1:
				Rosters.Player1 = key
				Rosters.Position1 = value
				break
			case 2:
				Rosters.Player2 = key
				Rosters.Position2 = value
				break
			case 3:
				Rosters.Player3 = key
				Rosters.Position3 = value
				break
			case 4:
				Rosters.Player4 = key
				Rosters.Position4 = value
				break
			case 5:
				Rosters.Player5 = key
				Rosters.Position5 = value
				break
			case 6:
				Rosters.Player6 = key
				Rosters.Position6 = value
				break
			case 7:
				Rosters.Player7 = key
				Rosters.Position7 = value
				break
			case 8:
				Rosters.Player8 = key
				Rosters.Position8 = value
				break
			case 9:
				Rosters.Player9 = key
				Rosters.Position9 = value
				break
			case 10:
				Rosters.Player10 = key
				Rosters.Position10 = value
				break
			case 11:
				Rosters.Player11 = key
				Rosters.Position11 = value
				break
			case 12:
				Rosters.Player12 = key
				Rosters.Position12 = value
				break
			case 13:
				Rosters.Player13 = key
				Rosters.Position13 = value
				break
			case 14:
				Rosters.Player14 = key
				Rosters.Position14 = value
				break
			case 15:
				Rosters.Player15 = key
				Rosters.Position15 = value
				break
			case 16:
				Rosters.Player16 = key
				Rosters.Position16 = value
				break
			case 17:
				Rosters.Player17 = key
				Rosters.Position17 = value
				break
			case 18:
				Rosters.Player18 = key
				Rosters.Position18 = value
				break
			case 19:
				Rosters.Player19 = key
				Rosters.Position19 = value
				break
			case 20:
				Rosters.Player20 = key
				Rosters.Position20 = value
				break
			case 21:
				Rosters.Player21 = key
				Rosters.Position21 = value
				break
			case 22:
				Rosters.Player22 = key
				Rosters.Position22 = value
				break
			case 23:
				Rosters.Player23 = key
				Rosters.Position23 = value
				break
			case 24:
				Rosters.Player24 = key
				Rosters.Position24 = value
				break
			case 25:
				Rosters.Player25 = key
				Rosters.Position25 = value
				break
			case 26:
				Rosters.Player26 = key
				Rosters.Position26 = value
				break
			}
		}
		RosterList = append(RosterList, Rosters)
		Teams.setTeamNames(client, i, RosterList)
	}
	return Teams
}

// GetSoutheastTeamRosters function
func GetSoutheastTeamRosters() Team {
	secret := goDotEnvVariable("REDIS")
	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(10), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var Rosters Roster
	var Teams Team

	i := 0
	for i < 12 {
		i++
		var RosterList []Roster
		key := "se_team_" + strconv.Itoa(i)
		name, _ := redis.String(client.Do("GET", key))
		hash, _ := redis.StringMap(client.Do("HGETALL", name))
		j := 0
		for key, value := range hash {
			j++
			switch j {
			case 1:
				Rosters.Player1 = key
				Rosters.Position1 = value
				break
			case 2:
				Rosters.Player2 = key
				Rosters.Position2 = value
				break
			case 3:
				Rosters.Player3 = key
				Rosters.Position3 = value
				break
			case 4:
				Rosters.Player4 = key
				Rosters.Position4 = value
				break
			case 5:
				Rosters.Player5 = key
				Rosters.Position5 = value
				break
			case 6:
				Rosters.Player6 = key
				Rosters.Position6 = value
				break
			case 7:
				Rosters.Player7 = key
				Rosters.Position7 = value
				break
			case 8:
				Rosters.Player8 = key
				Rosters.Position8 = value
				break
			case 9:
				Rosters.Player9 = key
				Rosters.Position9 = value
				break
			case 10:
				Rosters.Player10 = key
				Rosters.Position10 = value
				break
			case 11:
				Rosters.Player11 = key
				Rosters.Position11 = value
				break
			case 12:
				Rosters.Player12 = key
				Rosters.Position12 = value
				break
			case 13:
				Rosters.Player13 = key
				Rosters.Position13 = value
				break
			case 14:
				Rosters.Player14 = key
				Rosters.Position14 = value
				break
			case 15:
				Rosters.Player15 = key
				Rosters.Position15 = value
				break
			case 16:
				Rosters.Player16 = key
				Rosters.Position16 = value
				break
			case 17:
				Rosters.Player17 = key
				Rosters.Position17 = value
				break
			case 18:
				Rosters.Player18 = key
				Rosters.Position18 = value
				break
			case 19:
				Rosters.Player19 = key
				Rosters.Position19 = value
				break
			case 20:
				Rosters.Player20 = key
				Rosters.Position20 = value
				break
			case 21:
				Rosters.Player21 = key
				Rosters.Position21 = value
				break
			case 22:
				Rosters.Player22 = key
				Rosters.Position22 = value
				break
			case 23:
				Rosters.Player23 = key
				Rosters.Position23 = value
				break
			case 24:
				Rosters.Player24 = key
				Rosters.Position24 = value
				break
			case 25:
				Rosters.Player25 = key
				Rosters.Position25 = value
				break
			case 26:
				Rosters.Player26 = key
				Rosters.Position26 = value
				break
			}
		}
		RosterList = append(RosterList, Rosters)
		Teams.setTeamNames(client, i, RosterList)
	}
	return Teams
}

func (teamNames *Team) setTeamNames(client redis.Conn, n int, name []Roster) {
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
