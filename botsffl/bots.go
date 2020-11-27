package botsffl

import (
	"log"
	"os"

	"github.com/abspen1/restful-go/email"

	"github.com/gomodule/redigo/redis"
)

// Leaders struct
type Leaders struct {
	//Midwest
	StandingsMw1  string
	StandingsMw2  string
	StandingsMw3  string
	StandingsMw4  string
	StandingsMw5  string
	StandingsMw6  string
	StandingsMw7  string
	StandingsMw8  string
	StandingsMw9  string
	StandingsMw10 string
	StandingsMw11 string
	StandingsMw12 string
	PointsMw1     string
	PointsMw2     string
	PointsMw3     string
	PointsMw4     string
	PointsMw5     string
	PointsMw6     string
	PointsMw7     string
	PointsMw8     string
	PointsMw9     string
	PointsMw10    string
	PointsMw11    string
	PointsMw12    string
	//Northeast
	StandingsNe1  string
	StandingsNe2  string
	StandingsNe3  string
	StandingsNe4  string
	StandingsNe5  string
	StandingsNe6  string
	StandingsNe7  string
	StandingsNe8  string
	StandingsNe9  string
	StandingsNe10 string
	StandingsNe11 string
	StandingsNe12 string
	PointsNe1     string
	PointsNe2     string
	PointsNe3     string
	PointsNe4     string
	PointsNe5     string
	PointsNe6     string
	PointsNe7     string
	PointsNe8     string
	PointsNe9     string
	PointsNe10    string
	PointsNe11    string
	PointsNe12    string
	// Southeast
	StandingsSe1  string
	StandingsSe2  string
	StandingsSe3  string
	StandingsSe4  string
	StandingsSe5  string
	StandingsSe6  string
	StandingsSe7  string
	StandingsSe8  string
	StandingsSe9  string
	StandingsSe10 string
	StandingsSe11 string
	StandingsSe12 string
	PointsSe1     string
	PointsSe2     string
	PointsSe3     string
	PointsSe4     string
	PointsSe5     string
	PointsSe6     string
	PointsSe7     string
	PointsSe8     string
	PointsSe9     string
	PointsSe10    string
	PointsSe11    string
	PointsSe12    string
	// West
	StandingsW1  string
	StandingsW2  string
	StandingsW3  string
	StandingsW4  string
	StandingsW5  string
	StandingsW6  string
	StandingsW7  string
	StandingsW8  string
	StandingsW9  string
	StandingsW10 string
	StandingsW11 string
	StandingsW12 string
	PointsW1     string
	PointsW2     string
	PointsW3     string
	PointsW4     string
	PointsW5     string
	PointsW6     string
	PointsW7     string
	PointsW8     string
	PointsW9     string
	PointsW10    string
	PointsW11    string
	PointsW12    string
}

// SetLeaders function
func SetLeaders() Leaders {
	var status Leaders
	secret := os.Getenv("REDIS_PASS")

	client, err := redis.Dial("tcp", os.Getenv("REDIS_HOST_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	status.setStandingsMw(client)
	status.setPointsMw(client)
	status.setStandingsNe(client)
	status.setPointsNe(client)
	status.setStandingsSe(client)
	status.setPointsSe(client)
	status.setStandingsW(client)
	status.setPointsW(client)

	return status
}

// Midwest
func (status *Leaders) setStandingsMw(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.StandingsMw1, _ = redis.String(client.Do("GET", "mw_standings_1"))
		case 2:
			status.StandingsMw2, _ = redis.String(client.Do("GET", "mw_standings_2"))
		case 3:
			status.StandingsMw3, _ = redis.String(client.Do("GET", "mw_standings_3"))
		case 4:
			status.StandingsMw4, _ = redis.String(client.Do("GET", "mw_standings_4"))
		case 5:
			status.StandingsMw5, _ = redis.String(client.Do("GET", "mw_standings_5"))
		case 6:
			status.StandingsMw6, _ = redis.String(client.Do("GET", "mw_standings_6"))
		case 7:
			status.StandingsMw7, _ = redis.String(client.Do("GET", "mw_standings_7"))
		case 8:
			status.StandingsMw8, _ = redis.String(client.Do("GET", "mw_standings_8"))
		case 9:
			status.StandingsMw9, _ = redis.String(client.Do("GET", "mw_standings_9"))
		case 10:
			status.StandingsMw10, _ = redis.String(client.Do("GET", "mw_standings_10"))
		case 11:
			status.StandingsMw11, _ = redis.String(client.Do("GET", "mw_standings_11"))
		case 12:
			status.StandingsMw12, _ = redis.String(client.Do("GET", "mw_standings_12"))
		}
		i++
	}
}

func (status *Leaders) setPointsMw(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.PointsMw1, _ = redis.String(client.Do("GET", "mw_points_1"))
		case 2:
			status.PointsMw2, _ = redis.String(client.Do("GET", "mw_points_2"))
		case 3:
			status.PointsMw3, _ = redis.String(client.Do("GET", "mw_points_3"))
		case 4:
			status.PointsMw4, _ = redis.String(client.Do("GET", "mw_points_4"))
		case 5:
			status.PointsMw5, _ = redis.String(client.Do("GET", "mw_points_5"))
		case 6:
			status.PointsMw6, _ = redis.String(client.Do("GET", "mw_points_6"))
		case 7:
			status.PointsMw7, _ = redis.String(client.Do("GET", "mw_points_7"))
		case 8:
			status.PointsMw8, _ = redis.String(client.Do("GET", "mw_points_8"))
		case 9:
			status.PointsMw9, _ = redis.String(client.Do("GET", "mw_points_9"))
		case 10:
			status.PointsMw10, _ = redis.String(client.Do("GET", "mw_points_10"))
		case 11:
			status.PointsMw11, _ = redis.String(client.Do("GET", "mw_points_11"))
		case 12:
			status.PointsMw12, _ = redis.String(client.Do("GET", "mw_points_12"))
		}
		i++
	}
}

// Northeast
func (status *Leaders) setStandingsNe(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.StandingsNe1, _ = redis.String(client.Do("GET", "ne_standings_1"))
		case 2:
			status.StandingsNe2, _ = redis.String(client.Do("GET", "ne_standings_2"))
		case 3:
			status.StandingsNe3, _ = redis.String(client.Do("GET", "ne_standings_3"))
		case 4:
			status.StandingsNe4, _ = redis.String(client.Do("GET", "ne_standings_4"))
		case 5:
			status.StandingsNe5, _ = redis.String(client.Do("GET", "ne_standings_5"))
		case 6:
			status.StandingsNe6, _ = redis.String(client.Do("GET", "ne_standings_6"))
		case 7:
			status.StandingsNe7, _ = redis.String(client.Do("GET", "ne_standings_7"))
		case 8:
			status.StandingsNe8, _ = redis.String(client.Do("GET", "ne_standings_8"))
		case 9:
			status.StandingsNe9, _ = redis.String(client.Do("GET", "ne_standings_9"))
		case 10:
			status.StandingsNe10, _ = redis.String(client.Do("GET", "ne_standings_10"))
		case 11:
			status.StandingsNe11, _ = redis.String(client.Do("GET", "ne_standings_11"))
		case 12:
			status.StandingsNe12, _ = redis.String(client.Do("GET", "ne_standings_12"))
		}
		i++
	}
}

func (status *Leaders) setPointsNe(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.PointsNe1, _ = redis.String(client.Do("GET", "ne_points_1"))
		case 2:
			status.PointsNe2, _ = redis.String(client.Do("GET", "ne_points_2"))
		case 3:
			status.PointsNe3, _ = redis.String(client.Do("GET", "ne_points_3"))
		case 4:
			status.PointsNe4, _ = redis.String(client.Do("GET", "ne_points_4"))
		case 5:
			status.PointsNe5, _ = redis.String(client.Do("GET", "ne_points_5"))
		case 6:
			status.PointsNe6, _ = redis.String(client.Do("GET", "ne_points_6"))
		case 7:
			status.PointsNe7, _ = redis.String(client.Do("GET", "ne_points_7"))
		case 8:
			status.PointsNe8, _ = redis.String(client.Do("GET", "ne_points_8"))
		case 9:
			status.PointsNe9, _ = redis.String(client.Do("GET", "ne_points_9"))
		case 10:
			status.PointsNe10, _ = redis.String(client.Do("GET", "ne_points_10"))
		case 11:
			status.PointsNe11, _ = redis.String(client.Do("GET", "ne_points_11"))
		case 12:
			status.PointsNe12, _ = redis.String(client.Do("GET", "ne_points_12"))
		}
		i++
	}
}

// Southeast
func (status *Leaders) setStandingsSe(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.StandingsSe1, _ = redis.String(client.Do("GET", "se_standings_1"))
		case 2:
			status.StandingsSe2, _ = redis.String(client.Do("GET", "se_standings_2"))
		case 3:
			status.StandingsSe3, _ = redis.String(client.Do("GET", "se_standings_3"))
		case 4:
			status.StandingsSe4, _ = redis.String(client.Do("GET", "se_standings_4"))
		case 5:
			status.StandingsSe5, _ = redis.String(client.Do("GET", "se_standings_5"))
		case 6:
			status.StandingsSe6, _ = redis.String(client.Do("GET", "se_standings_6"))
		case 7:
			status.StandingsSe7, _ = redis.String(client.Do("GET", "se_standings_7"))
		case 8:
			status.StandingsSe8, _ = redis.String(client.Do("GET", "se_standings_8"))
		case 9:
			status.StandingsSe9, _ = redis.String(client.Do("GET", "se_standings_9"))
		case 10:
			status.StandingsSe10, _ = redis.String(client.Do("GET", "se_standings_10"))
		case 11:
			status.StandingsSe11, _ = redis.String(client.Do("GET", "se_standings_11"))
		case 12:
			status.StandingsSe12, _ = redis.String(client.Do("GET", "se_standings_12"))
		}
		i++
	}
}

func (status *Leaders) setPointsSe(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.PointsSe1, _ = redis.String(client.Do("GET", "se_points_1"))
		case 2:
			status.PointsSe2, _ = redis.String(client.Do("GET", "se_points_2"))
		case 3:
			status.PointsSe3, _ = redis.String(client.Do("GET", "se_points_3"))
		case 4:
			status.PointsSe4, _ = redis.String(client.Do("GET", "se_points_4"))
		case 5:
			status.PointsSe5, _ = redis.String(client.Do("GET", "se_points_5"))
		case 6:
			status.PointsSe6, _ = redis.String(client.Do("GET", "se_points_6"))
		case 7:
			status.PointsSe7, _ = redis.String(client.Do("GET", "se_points_7"))
		case 8:
			status.PointsSe8, _ = redis.String(client.Do("GET", "se_points_8"))
		case 9:
			status.PointsSe9, _ = redis.String(client.Do("GET", "se_points_9"))
		case 10:
			status.PointsSe10, _ = redis.String(client.Do("GET", "se_points_10"))
		case 11:
			status.PointsSe11, _ = redis.String(client.Do("GET", "se_points_11"))
		case 12:
			status.PointsSe12, _ = redis.String(client.Do("GET", "se_points_12"))
		}
		i++
	}
}

// West
func (status *Leaders) setStandingsW(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.StandingsW1, _ = redis.String(client.Do("GET", "w_standings_1"))
		case 2:
			status.StandingsW2, _ = redis.String(client.Do("GET", "w_standings_2"))
		case 3:
			status.StandingsW3, _ = redis.String(client.Do("GET", "w_standings_3"))
		case 4:
			status.StandingsW4, _ = redis.String(client.Do("GET", "w_standings_4"))
		case 5:
			status.StandingsW5, _ = redis.String(client.Do("GET", "w_standings_5"))
		case 6:
			status.StandingsW6, _ = redis.String(client.Do("GET", "w_standings_6"))
		case 7:
			status.StandingsW7, _ = redis.String(client.Do("GET", "w_standings_7"))
		case 8:
			status.StandingsW8, _ = redis.String(client.Do("GET", "w_standings_8"))
		case 9:
			status.StandingsW9, _ = redis.String(client.Do("GET", "w_standings_9"))
		case 10:
			status.StandingsW10, _ = redis.String(client.Do("GET", "w_standings_10"))
		case 11:
			status.StandingsW11, _ = redis.String(client.Do("GET", "w_standings_11"))
		case 12:
			status.StandingsW12, _ = redis.String(client.Do("GET", "w_standings_12"))
		}
		i++
	}
}

func (status *Leaders) setPointsW(client redis.Conn) {
	i := 1
	for i < 13 {
		switch i {
		case 1:
			status.PointsW1, _ = redis.String(client.Do("GET", "w_points_1"))
		case 2:
			status.PointsW2, _ = redis.String(client.Do("GET", "w_points_2"))
		case 3:
			status.PointsW3, _ = redis.String(client.Do("GET", "w_points_3"))
		case 4:
			status.PointsW4, _ = redis.String(client.Do("GET", "w_points_4"))
		case 5:
			status.PointsW5, _ = redis.String(client.Do("GET", "w_points_5"))
		case 6:
			status.PointsW6, _ = redis.String(client.Do("GET", "w_points_6"))
		case 7:
			status.PointsW7, _ = redis.String(client.Do("GET", "w_points_7"))
		case 8:
			status.PointsW8, _ = redis.String(client.Do("GET", "w_points_8"))
		case 9:
			status.PointsW9, _ = redis.String(client.Do("GET", "w_points_9"))
		case 10:
			status.PointsW10, _ = redis.String(client.Do("GET", "w_points_10"))
		case 11:
			status.PointsW11, _ = redis.String(client.Do("GET", "w_points_11"))
		case 12:
			status.PointsW12, _ = redis.String(client.Do("GET", "w_points_12"))
		}
		i++
	}
}

// SaveProspect function
func SaveProspect(info email.BotsFFL) {
	secret := os.Getenv("REDIS_PASS")

	client, err := redis.Dial("tcp", os.Getenv("REDIS_HOST_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	client.Do("HSET", "BotsProspectsState", info.Email, info.State)
	client.Do("HSET", "BotsProspectsName", info.Email, info.Name)
}
