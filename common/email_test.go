package common

import (
	"gopkg.in/gomail.v2"
	"testing"
)

func TestSendEmail(t *testing.T) {
	m := gomail.NewMessage()
	m.SetHeader("From", "1029841651@qq.com")
	m.SetHeader("To", "jackanszu@gmail.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello 123456")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 465, "1029841651@qq.com", "test")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		t.Fatal(err)
	}
}
