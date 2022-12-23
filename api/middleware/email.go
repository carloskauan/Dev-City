package middleware

import (
  "crypto/tls"
  gomail "gopkg.in/gomail.v2"
)

func SendWhitGomail(body string)error{
  m := gomail.NewMessage()
  m.SetHeader("From", "dev.city@hotmail.com")
  m.SetHeader("To", "vcladara@gmail.com")
  m.SetHeader("Subject", "Teste de envio")
  m.SetBody("text/html", body)
  d := gomail.NewDialer("smtp.office365.com", 587, "dev.city@hotmail.com", "15971597Ck_.")
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
  err := d.DialAndSend(m)
  return err
}
