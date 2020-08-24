package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type smtpServer struct {
	host string
	port string
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

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// SendEmail function
func SendEmail(info Info) {
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
		return
	}
	fmt.Println("Email Sent!")
}
