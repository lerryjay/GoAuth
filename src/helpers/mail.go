package helpers

import (
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(recipient, subject, message string) bool {
	smtp_server := os.Getenv("SMTP_SERVER")
	smtp_port, err := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 64)
	if err != nil {
		return false
	}
	smtp_username := os.Getenv("SMTP_USERNAME")
	smtp_password := os.Getenv("SMTP_PASSWORD")

	msg := gomail.NewMessage()
	msg.SetHeader("From", "<paste your gmail account here>")
	msg.SetHeader("To", recipient)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", message)
	// msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer(smtp_server, int(smtp_port), smtp_username, smtp_password)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

	return true
}
