package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
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

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// SendEmail function
func SendEmail(info Info) bool {
	// Sender data.
	from := goDotEnvVariable("EMAIL")
	password := goDotEnvVariable("EMAIL-PASS")
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
	from := goDotEnvVariable("EMAIL")
	password := goDotEnvVariable("EMAIL-PASS")
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
	from := goDotEnvVariable("EMAIL")
	password := goDotEnvVariable("EMAIL-PASS")
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

// GetTwitterData function
func GetTwitterData() {
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
	// client.Do("HDEL", "Todos", "try to add something")
	hash, _ := redis.StringMap(client.Do("HGETALL", "Todos"))
	fmt.Println(hash)
}

func main() {
	GetTwitterData()
}
