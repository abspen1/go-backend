package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/abspen1/restful-go/players/trending"

	"github.com/abspen1/restful-go/todos"

	"github.com/gomodule/redigo/redis"
)

type smtpServer struct {
	host string
	port string
}

//Sender struct
type Sender struct {
	User     string
	Password string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// Info struct
type Info struct {
	Name    string
	Email   string
	Message string
}

// BotsFFL struct
type BotsFFL struct {
	Name  string
	Email string
	State string
}

// Birthday struct
type Birthday struct {
	Name  string
	Email string
}

// SendEmail function
func SendEmail(info Info) bool {
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL-PASS")
	// Receiver Email address.
	to := []string{
		from,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	strMessage := fmt.Sprintf("Name: %s  Email: %s Message: %s", info.Name, info.Email, info.Message)
	msg := "From: " + from + "\n" + "Subject: webapp\n\n" + strMessage
	// fmt.Println(strMessage)
	message := []byte(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending Email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Email Sent!")
	return true
}

// SaveBotsInfo function
func SaveBotsInfo(info BotsFFL) bool {
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL-PASS")
	// Receiver Email address.
	to := []string{
		from,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	strMessage := fmt.Sprintf("Name: %s  Email: %s State: %s", info.Name, info.Email, info.State)
	msg := "From: " + from + "\n" + "Subject: BotsFFL\n\n" + strMessage
	// fmt.Println(strMessage)
	message := []byte(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending Email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Email Sent!")
	return true
}

// SendBdayEmail function
func SendBdayEmail(info Birthday) bool {
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL-PASS")
	// Receiver Email address.
	to := []string{
		from, info.Email,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	strMessage := fmt.Sprintf("Happy birthday %s! Enjoy your day!\n\n\n\nBest,\nAustin's Automation", info.Name)
	msg := "From: " + from + "\n" +
		"To: " + info.Email + "\n" +
		"Subject: Happy Birthday\n\n" + strMessage
	// fmt.Println(strMessage)
	message := []byte(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending Email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Email Sent!")
	return true
}

func testTodos() {
	secret := os.Getenv("REDIS_PASS")
	host := os.Getenv("REDIS_HOST_PORT")

	client, err := redis.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	client.Do("DEL", "Todos")

	// first := todos.Todos{
	// 	"Write essay for CSE 301",
	// 	false,
	// }

	// if todos.AddTodo(first) {
	// 	fmt.Println("It worked")
	// }
	// todos.RmTodo(first)

	fmt.Println(todos.GetTodos())
}

func main() {
	// var roster players.Roster
	// roster = players.GetMidwestTeamNames()
	// fmt.Println(roster)
	// roster = players.GetNortheastTeamNames()
	// fmt.Println(roster)
	// roster = players.GetSoutheastTeamNames()
	// fmt.Println(roster)
	// roster = players.GetWestTeamNames()
	// fmt.Println(roster)
	fmt.Println(trending.DailyAdd())
	fmt.Println(trending.WeeklyAdd())
}
