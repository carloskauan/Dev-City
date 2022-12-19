package controllers

import(
  "net/http"
  "DevCity/api/database"
  "DevCity/api/models"
  "github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
  var user models.User
  if err := c.ShouldBindJSON(&user); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  database.DB.Create(&user)
  c.JSON(http.StatusOK, user)
}
