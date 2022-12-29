package middleware

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCode() string{
	var stringCode, stringRandNumber string
	// Mudando a seed de geração com base no horario
	rand.Seed(time.Now().UnixNano())

	// Gerando 6 numeros
	for i := 0; i < 6; i++ {
		stringRandNumber = fmt.Sprintf("%d",rand.Intn(10))
		stringCode += stringRandNumber
	}
	return stringCode
}