package main

import(
  "fmt"
  "DevCity/api/routes"
  "DevCity/api/database"
  "DevCity/api/middleware"
)

func main() {
  fmt.Println(middleware.GenerateCode())
  fmt.Println("Iniciando servidor...")
  database.ConnectDB()
  routes.HandlerFunc()
}
