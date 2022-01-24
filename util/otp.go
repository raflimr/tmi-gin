package util

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	timeout        = 20
	ch             = make(chan bool)
	otp     string = GenerateRandomNumber()
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
}

func GenerateRandomNumber() string {
	rand.Seed(time.Now().UnixMilli())
	var out strings.Builder
	for i := 0; i < 4; i++ {
		out.WriteString(strconv.Itoa(rand.Intn(9)))
	}

	return out.String()
}

func OTPEmail(email string) (*string, error) {

	if otp == "" {
		otp = GenerateRandomNumber()
	} else {
		otp = GenerateRandomNumber()
	}

	// Reveal Configs Vars
	emailData := os.Getenv("EMAIL_OTP")
	if emailData == "" {
		emailData = "your-email"
	}

	passwordData := os.Getenv("PASSWORD_OTP")
	if passwordData == "" {
		passwordData = "your-password"
	}

	// Sender data.
	from := emailData
	password := passwordData

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(otp)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Email Sent Successfully!")
	go timer(timeout, ch)
	go watcher(timeout, ch)

	return &otp, nil
}
