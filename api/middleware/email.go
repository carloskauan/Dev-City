package middleware

import (
  "crypto/tls"
  gomail "gopkg.in/gomail.v2"
)

func SendWhitGomail(destin string, code string)error{
  m := gomail.NewMessage()
  m.SetHeader("From", "dev.city@hotmail.com")
  m.SetHeader("To", destin)
  m.SetHeader("Subject", "Verificação de email")
  m.SetBody("text/html", "Seu codigo de verificação é:\n"+code)
  d := gomail.NewDialer("smtp.office365.com", 587, "dev.city@hotmail.com", "15971597Ck_.")
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
  err := d.DialAndSend(m)
  return err
}
