package database

import(
  "log"
  "DevCity/api/models"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

var (
  DB *gorm.DB
  err error
)

func ConnectDB(){
  prepare := "root:1597@tcp(127.0.0.1:3306)/devcity?charset=utf8mb4&parseTime=True&loc=Local"
  DB, err = gorm.Open(mysql.Open(prepare))
  if err != nil{
    log.Panic("Erro ao se connectar com o banco ",err)
  }
  DB.AutoMigrate(&models.User{})
  DB.AutoMigrate(&models.Codes{})
}
