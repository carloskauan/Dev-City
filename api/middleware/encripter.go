package middleware

import( 
"fmt"
"golang.org/x/crypto/bcrypt"
)

func Encripter(data string)string{
	sb, err := bcrypt.GenerateFromPassword([]byte(data), 10)
  if err != nil{ fmt.Println("Erro ao criptografar... ", err)}
  return string(sb)
}