package rosters

import (
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
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
	Player1    string `json:"player1,omitempty"`
	Position1  string `json:"position1,omitempty"`
	Player2    string `json:"player2,omitempty"`
	Position2  string `json:"position2,omitempty"`
	Player3    string `json:"player3,omitempty"`
	Position3  string `json:"position3,omitempty"`
	Player4    string `json:"player4,omitempty"`
	Position4  string `json:"position4,omitempty"`
	Player5    string `json:"player5,omitempty"`
	Position5  string `json:"position5,omitempty"`
	Player6    string `json:"player6,omitempty"`
	Position6  string `json:"position6,omitempty"`
	Player7    string `json:"player7,omitempty"`
	Position7  string `json:"position7,omitempty"`
	Player8    string `json:"player8,omitempty"`
	Position8  string `json:"position8,omitempty"`
	Player9    string `json:"player9,omitempty"`
	Position9  string `json:"position9,omitempty"`
	Player10   string `json:"player10,omitempty"`
	Position10 string `json:"position10,omitempty"`
	Player11   string `json:"player11,omitempty"`
	Position11 string `json:"position11,omitempty"`
	Player12   string `json:"player12,omitempty"`
	Position12 string `json:"position12,omitempty"`
	Player13   string `json:"player13,omitempty"`
	Position13 string `json:"position13,omitempty"`
	Player14   string `json:"player14,omitempty"`
	Position14 string `json:"position14,omitempty"`
	Player15   string `json:"player15,omitempty"`
	Position15 string `json:"position15,omitempty"`
	Player16   string `json:"player16,omitempty"`
	Position16 string `json:"position16,omitempty"`
	Player17   string `json:"player17,omitempty"`
	Position17 string `json:"position17,omitempty"`
	Player18   string `json:"player18,omitempty"`
	Position18 string `json:"position18,omitempty"`
	Player19   string `json:"player19,omitempty"`
	Position19 string `json:"position19,omitempty"`
	Player20   string `json:"player20,omitempty"`
	Position20 string `json:"position20,omitempty"`
	Player21   string `json:"player21,omitempty"`
	Position21 string `json:"position21,omitempty"`
	Player22   string `json:"player22,omitempty"`
	Position22 string `json:"position22,omitempty"`
	Player23   string `json:"player23,omitempty"`
	Position23 string `json:"position23,omitempty"`
	Player24   string `json:"player24,omitempty"`
	Position24 string `json:"position24,omitempty"`
	Player25   string `json:"player25,omitempty"`
	Position25 string `json:"position25,omitempty"`
	Player26   string `json:"player26,omitempty"`
	Position26 string `json:"position26,omitempty"`
}

// GetMidwestTeamRosters function
func GetMidwestTeamRosters() Team {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host, redis.DialDatabase(10), redis.DialPassword(secret))
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
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")
	client, err := redis.Dial("tcp", host, redis.DialDatabase(10), redis.DialPassword(secret))
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
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")
	client, err := redis.Dial("tcp", host, redis.DialDatabase(10), redis.DialPassword(secret))
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
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host, redis.DialDatabase(10), redis.DialPassword(secret))
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
