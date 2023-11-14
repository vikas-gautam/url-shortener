package internal

import (
	"auth-service/models"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	mail "github.com/xhit/go-simple-mail/v2"
)

var app models.AppConfig

func ListenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			SendMsg(msg)
		}
	}()
}

func SendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "mailhog"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		logrus.Error(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	err = email.Send(client)
	if err != nil {
		logrus.Error(err)
	}

	log.Println("Email has been sent!")
}
