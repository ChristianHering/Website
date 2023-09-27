package main
/*
import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

//Email holds all the information used to send an email
type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

//Validate an email to make sure all values are valid
func (e Email) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.From, validation.Required, is.Email),
		validation.Field(&e.To, validation.Required, is.Email),
		validation.Field(&e.Subject, validation.Required),
		validation.Field(&e.Body, validation.Required),
	)
}

//SendMail Sends out an email with the given information
func SendMail(e *Email) error {
	m := gomail.NewMessage() //Init email message, and define options

	m.SetHeader("Reply-To", e.From)
	m.SetHeader("From", Config.EmailConfig.Username)
	m.SetHeader("To", e.To)
	m.SetHeader("Subject", e.Subject)
	m.SetBody("text/html", e.Body)

	server := gomail.NewDialer(Config.EmailConfig.Host, Config.EmailConfig.Port, Config.EmailConfig.Username, Config.EmailConfig.Password)

	err := server.DialAndSend(m) //Send the email
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
*/