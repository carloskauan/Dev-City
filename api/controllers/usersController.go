package controllers

import(
  "net/http"
  "DevCity/api/database"
  "DevCity/api/models"
  "github.com/gin-gonic/gin"
)

var user models.User

func RegisterUser(c *gin.Context){
  if err := c.ShouldBindJSON(&user); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  database.DB.Create(&user)
  c.JSON(http.StatusOK, user)
}

func SeachUserByEmail(c *gin.Context){
  email := c.Params.ByName("email")
  database.DB.Where(&models.User{Email: email}).First(&user)
  c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context){
  var users []models.User
  database.DB.Find(&users)
  c.JSON(http.StatusOK, users)
}
