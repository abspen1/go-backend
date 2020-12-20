package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/abspen1/restful-go/twitter/tweet"

	"github.com/abspen1/restful-go/players/trending"

	"github.com/abspen1/restful-go/players/rosters"

	"github.com/abspen1/restful-go/players"
	"github.com/abspen1/restful-go/todos"

	"github.com/abspen1/restful-go/botsffl"

	"github.com/abspen1/restful-go/twitter"

	"github.com/abspen1/restful-go/email"

	"github.com/abspen1/restful-go/projects"
	"github.com/abspen1/restful-go/rps"
	"github.com/badoux/checkmail"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func getProjects(w http.ResponseWriter, r *http.Request) {
	s := projects.GetString()

	json.NewEncoder(w).Encode(s)
}

func postProjects(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var proj projects.Project
	_ = json.Unmarshal(info, &proj)
	projects.SetString(proj)
	fmt.Fprintf(w, string(info))

	// fmt.Fprintf(w, "Test POST endpoint worked!")
}

func postRmprojects(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var proj projects.RmProject
	_ = json.Unmarshal(info, &proj)

	if projects.CheckPass(proj) {
		if projects.RmString(proj) {
			fmt.Fprintf(w, "POST remove worked!")
		} else {
			fmt.Fprintf(w, "Error")
		}
	} else {
		fmt.Fprintf(w, "Err")
	}

	// fmt.Fprintf(w, "Test POST endpoint worked!")
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info email.Info
	_ = json.Unmarshal(body, &info)

	err := checkmail.ValidateFormat(info.Email)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Format Error")
		return
	}
	err = checkmail.ValidateHost(info.Email)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		fmt.Fprintf(w, "Error")
		return
	}

	if email.SendEmail(info) {
		fmt.Fprintf(w, "Email sent successfully")
		return
	}
	fmt.Fprintf(w, "Email not sent")
}

func postRPS(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var rpsUser rps.User
	_ = json.Unmarshal(info, &rpsUser)

	rpsUser = rps.SaveData(rpsUser)

	json.NewEncoder(w).Encode(rpsUser)
}

func postRPSLogin(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var rpsUser rps.User
	_ = json.Unmarshal(info, &rpsUser)

	err := checkmail.ValidateFormat(rpsUser.Username)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Format Error")
		return
	}
	err = checkmail.ValidateHost(rpsUser.Username)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		fmt.Fprintf(w, "Error")
		return
	}

	rpsUser = rps.GetData(rpsUser)

	json.NewEncoder(w).Encode(rpsUser)
}

func getRPSLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rock Paper Scissors game login endpoint, nothing to see here!")
}

func getRPS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rock Paper Scissors game save endpoint, nothing to see here!")
}

func getEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Send email endpoint, nothing to see here!")
}

func getRmprojects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove projects endpoint, nothing to see here!")
}

func getTweet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<body style="text-align:center;">
	<h1>Go twitter bot post tweet endpoint, nothing to see here!<h1>
	<img src="https://www.logo.wine/a/logo/Go_(programming_language)/Go_(programming_language)-Logo.wine.svg" alt="Go Logo">
	</body>`)
}

func getTwitterData(w http.ResponseWriter, r *http.Request) {
	data := twitter.GetTwitterData()

	json.NewEncoder(w).Encode(data)
}

func getBotsFFL(w http.ResponseWriter, r *http.Request) {
	leaders := botsffl.SetLeaders()

	json.NewEncoder(w).Encode(leaders)
}

func getMwTeams(w http.ResponseWriter, r *http.Request) {
	var roster players.Roster
	roster = players.GetMidwestTeamNames()
	json.NewEncoder(w).Encode(roster)
}

func getWTeams(w http.ResponseWriter, r *http.Request) {
	var roster players.Roster
	roster = players.GetWestTeamNames()
	json.NewEncoder(w).Encode(roster)
}

func getNeTeams(w http.ResponseWriter, r *http.Request) {
	var roster players.Roster
	roster = players.GetNortheastTeamNames()
	json.NewEncoder(w).Encode(roster)
}

func getSeTeams(w http.ResponseWriter, r *http.Request) {
	var roster players.Roster
	roster = players.GetSoutheastTeamNames()
	json.NewEncoder(w).Encode(roster)
}

func getMwRosters(w http.ResponseWriter, r *http.Request) {
	var roster rosters.Team
	roster = rosters.GetMidwestTeamRosters()
	json.NewEncoder(w).Encode(roster)
}

func getNeRosters(w http.ResponseWriter, r *http.Request) {
	var roster rosters.Team
	roster = rosters.GetNortheastTeamRosters()
	json.NewEncoder(w).Encode(roster)
}

func getWRosters(w http.ResponseWriter, r *http.Request) {
	var roster rosters.Team
	roster = rosters.GetWestTeamRosters()
	json.NewEncoder(w).Encode(roster)
}

func getSeRosters(w http.ResponseWriter, r *http.Request) {
	var roster rosters.Team
	roster = rosters.GetSoutheastTeamRosters()
	json.NewEncoder(w).Encode(roster)
}

func getDailyTrendAdd(w http.ResponseWriter, r *http.Request) {
	var daily trending.DayAdd
	daily = trending.DailyAdd()
	json.NewEncoder(w).Encode(daily)
}

func getDailyTrendDrop(w http.ResponseWriter, r *http.Request) {
	var daily trending.DayDrop
	daily = trending.DailyDrop()
	json.NewEncoder(w).Encode(daily)
}

func getWeeklyTrendAdd(w http.ResponseWriter, r *http.Request) {
	var weekly trending.WeekAdd
	weekly = trending.WeeklyAdd()
	json.NewEncoder(w).Encode(weekly)
}

func getWeeklyTrendDrop(w http.ResponseWriter, r *http.Request) {
	var weekly trending.WeekDrop
	weekly = trending.WeeklyDrop()
	json.NewEncoder(w).Encode(weekly)
}

func postBotsFFL(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info email.BotsFFL
	_ = json.Unmarshal(body, &info)

	err := checkmail.ValidateFormat(info.Email)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Format Error")
		return
	}

	err = checkmail.ValidateHost(info.Email)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		fmt.Fprintf(w, "Error")
		return
	}

	if email.SaveBotsInfo(info) {
		botsffl.SaveProspect(info)
		fmt.Fprintf(w, "Email sent successfully")
		return
	}
	fmt.Fprintf(w, "Email not sent")
}

func getBdayEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Austin's bday emailer post address :)")
}

func postBdayEmail(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info email.Birthday
	_ = json.Unmarshal(body, &info)
	resp := email.SendBdayEmail(info)
	if resp == "Success" {
		fmt.Fprintf(w, "Email sent successfully")
		return
	}

	if resp == "Auth Err" {
		fmt.Fprintf(w, "Invalid Authentification")
		return
	}

	fmt.Fprintf(w, "Email not sent")
}

func postTweet(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var content tweet.Content
	_ = json.Unmarshal(body, &content)
	if content.Auth != os.Getenv("SECRET") {
		fmt.Fprintf(w, "Invalid Authentification")
		return
	}
	resp := tweet.Tweet(content)
	if resp == true {
		fmt.Fprintf(w, "Tweet sent successfully")
	} else {
		fmt.Fprintf(w, "Error in postTweet")
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos := todos.GetTodos()

	json.NewEncoder(w).Encode(todos)
}

func postTodos(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info todos.Todos
	_ = json.Unmarshal(body, &info)
	if todos.AddTodo(info) {
		fmt.Fprintf(w, "Added todo successfully")
	}

	fmt.Fprintf(w, "Todo not added due to an error")
}
func rmTodos(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info todos.FullTodo
	_ = json.Unmarshal(body, &info)
	if todos.RmTodo(info) {
		fmt.Fprintf(w, "Removed todo successfully")
	}

	fmt.Fprintf(w, "Todo wasn't removed due to an error")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<body style="text-align:center;">
	<h1>Austin's back-end API</h1>
	<h3>Random</h3>
	<p><a href="/austinapi/projects">Projects</a></p>
	<p><a href="/austinapi/rmprojects">Remove Projects</a></p>
	<p><a href="/austinapi/rps/login">Rock Paper Scissors Login</a></p>
	<p><a href="/austinapi/rps/">Rock Paper Scissors Save</a></p>
	<p><a href="/austinapi/bdayemail">Bday Emailer</a></p>
	<p><a href="/austinapi/email">Contact Page Emailer</a></p>
	<p><a href="/austinapi/todos">Todos</a></p>
	<p><a href="/austinapi/tendie-intern">Twitter Data</a></p>
	<p><a href="/austinapi/go-tweet">Go Twitter Bot</a></p>
	<h3>Battle of the States Flag Football League</h3>
	<p><a href="/austinapi/botsffl">Standings</a></p>
	<p><a href="/austinapi/botsffl/teams/midwest">Midwest Teams</a></p>
	<p><a href="/austinapi/botsffl/teams/midwest/roster">Midwest Teams' Rosters</a></p>
	<p><a href="/austinapi/botsffl/teams/west">West Teams</a></p>
	<p><a href="/austinapi/botsffl/teams/west/roster">West Teams' Rosters</a></p>
	<p><a href="/austinapi/botsffl/teams/northeast">Northeast Teams</a></p>
	<p><a href="/austinapi/botsffl/teams/northeast/roster">Northeast Teams' Rosters</a></p>
	<p><a href="/austinapi/botsffl/teams/southeast">Southeast Teams</a></p>
	<p><a href="/austinapi/botsffl/teams/southeast/roster">Southeast Teams' Rosters</a></p>
	<p><a href="/austinapi/botsffl/trending/daily/add">sleeper's trending added players(24hrs)</a></p>
	<p><a href="/austinapi/botsffl/trending/daily/drop">sleeper's trending dropped players(24hrs)</a></p>
	<p><a href="/austinapi/botsffl/trending/weekly/add">sleeper's trending added players(5days)</a></p>
	<p><a href="/austinapi/botsffl/trending/weekly/drop">sleeper's trending added players(5days)</a></p>
	</body>`)
}

func handleRequests() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://abspen1.github.io", "https://austinspencer.works"},
		AllowCredentials: true,
		Debug:            false,
	})
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/austinapi/", homePage)
	myRouter.HandleFunc("/austinapi/go-tweet", postTweet).Methods("POST")
	myRouter.HandleFunc("/austinapi/go-tweet", getTweet)
	myRouter.HandleFunc("/austinapi/projects", getProjects).Methods("GET")
	myRouter.HandleFunc("/austinapi/projects", postProjects).Methods("POST")
	myRouter.HandleFunc("/austinapi/rmprojects", postRmprojects).Methods("POST")
	myRouter.HandleFunc("/austinapi/rmprojects", getRmprojects)
	myRouter.HandleFunc("/austinapi/email", sendEmail).Methods("POST")
	myRouter.HandleFunc("/austinapi/email", getEmail)
	myRouter.HandleFunc("/austinapi/rps/login", postRPSLogin).Methods("POST")
	myRouter.HandleFunc("/austinapi/rps", postRPS).Methods("POST")
	myRouter.HandleFunc("/austinapi/rps/login", getRPSLogin)
	myRouter.HandleFunc("/austinapi/rps", getRPS)
	myRouter.HandleFunc("/austinapi/tendie-intern", getTwitterData).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl", getBotsFFL).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl", postBotsFFL).Methods("POST")
	myRouter.HandleFunc("/austinapi/bdayemail", postBdayEmail).Methods("POST")
	myRouter.HandleFunc("/austinapi/bdayemail", getBdayEmail)
	myRouter.HandleFunc("/austinapi/todos", getTodos).Methods("GET")
	myRouter.HandleFunc("/austinapi/todos", postTodos).Methods("POST")
	myRouter.HandleFunc("/austinapi/todos/rm", rmTodos).Methods("POST")
	myRouter.HandleFunc("/austinapi/botsffl/teams/midwest", getMwTeams).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/west", getWTeams).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/northeast", getNeTeams).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/southeast", getSeTeams).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/midwest/roster", getMwRosters).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/west/roster", getWRosters).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/northeast/roster", getNeRosters).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/southeast/roster", getSeRosters).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/trending/daily/add", getDailyTrendAdd).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/trending/daily/drop", getDailyTrendDrop).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/trending/weekly/add", getWeeklyTrendAdd).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/trending/weekly/drop", getWeeklyTrendDrop).Methods("GET")

	handler := c.Handler(myRouter)
	log.Fatal(http.ListenAndServe(":8558", handler))
}

func main() {
	handleRequests()
}
