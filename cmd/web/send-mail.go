package main

import (
	"fmt"
	"github.com/ivanpopov/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MainChan
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)

	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%v", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}

		mailTemplate := string(data)
		msgToSent := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSent)
	}

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("email sent!")
	}
}
