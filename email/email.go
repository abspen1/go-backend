package email

import (
	"fmt"
	"net/smtp"

	"github.com/abspen1/restful-go/auth"
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

// BotsFFL struct
type BotsFFL struct {
	Name  string
	Email string
	State string
}

// Birthday struct
type Birthday struct {
	Name          string
	Email         string
	JokeSetup     string
	JokePunchLine string
	Auth          string
}

// SendEmail function
func SendEmail(info Info) bool {
	// Sender data.
	from := auth.GoDotEnvVariable("EMAIL")
	password := auth.GoDotEnvVariable("EMAIL-PASS")
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
	from := auth.GoDotEnvVariable("EMAIL")
	password := auth.GoDotEnvVariable("EMAIL-PASS")
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
func SendBdayEmail(info Birthday) string {
	// Sender data.
	from := auth.GoDotEnvVariable("EMAIL")
	password := auth.GoDotEnvVariable("EMAIL-PASS")
	pass := auth.GoDotEnvVariable("BACK-END-AUTH")
	if info.Auth != pass {
		return "Auth Err"
	}
	// Receiver Email address.
	to := []string{
		from, info.Email,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	strMessage := fmt.Sprintf("Happy birthday %s! Here is a joke to get your day started right!\n\n\n%s\n\n%s\n\n\n\n\nBest Regards,\n\nAustin", info.Name, info.JokeSetup, info.JokePunchLine)
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
		return "Email Err"
	}
	fmt.Println("Email Sent!")
	return "Success"
}
