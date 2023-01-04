package main

import (
	"DevCity/api/database"
	"DevCity/api/middleware"
	"DevCity/api/routes"
	"fmt"
)

func main() {
	fmt.Println(middleware.GenerateCode())
	fmt.Println("Iniciando servidor...")
	database.ConnectDB()
	routes.HandlerFunc()
}
