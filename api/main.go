package main

import(
  "fmt"
  "DevCity/api/routes"
  "DevCity/api/database"
)

func main() {
  fmt.Println("Iniciando servidor...")
  database.ConnectDB()
  routes.HandlerFunc()
}
