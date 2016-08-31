package email

import (
	"fmt"
	"net/mail"
	"net/smtp"
)

//Mail email account info
type Mail struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

//Send Send Email
func (m *Mail) Send(msg *Message) error {
	msg.From = mail.Address{
		Name:    m.Name,
		Address: m.Account,
	}
	addr := fmt.Sprintf("%s:%s", m.Host, m.Port)
	auth := smtp.PlainAuth("", m.Account, m.Password, m.Host)

	return smtp.SendMail(addr, auth, m.Account, msg.Tolist(), msg.Bytes())
}
