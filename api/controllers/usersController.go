package controllers

import(
  "fmt"
  "net/http"
  "net/mail"
  "DevCity/api/database"
  "DevCity/api/models"
  "DevCity/api/middleware"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
)

var user models.User

func RegisterUser(c *gin.Context){
  // Puxando dados do usuario da requisição
  if err := c.ShouldBindJSON(&user); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  //Validação de email
  validateEmail , _ := mail.ParseAddress(user.Email)
  if validateEmail == nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "invalid email",
    })
    return
  }
  //Criptografando senha
  sb, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
  if err != nil{ fmt.Println("Erro ao criptografar... ", err)}
  user.Password = string(sb)
  //Salvando dandos no banco
  database.DB.Create(&user)

  c.JSON(http.StatusOK, user)
}

func SearchUserByEmail(c *gin.Context){
  email := c.Params.ByName("email")
  database.DB.Where(&models.User{Email: email}).First(&user)
  if user.ID == 0{
    c.JSON(http.StatusNotFound, gin.H{
      "error":"User Not Found",
      "status":false,
    })
    return
  }
  c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context){
  var users []models.User
  database.DB.Find(&users)
  c.JSON(http.StatusOK, users)
}

func Authenticate(c *gin.Context){
  var userAuth models.User
  email := c.Param("email")
  password := c.Param("password")
  //Encontrando usuario por email
  database.DB.Where(&models.User{Email: email}).First(&userAuth)
  if userAuth.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{
      "error":"User Not Found",
      "status":false,
    })
    return
  }
  //Comparando senha do banco e senha informada
  err := bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(password))
  if err != nil{
    c.JSON(http.StatusNotFound, gin.H{
      "error":"Password don't macth",
      "status":false,
    })
    return
  }
  // Resposta caso autenticação bata
  c.JSON(http.StatusOK, gin.H{
    "status":true,
  })
}

func Test(c *gin.Context){
  err := middleware.SendWhitGomail("Ola victorrrrrrrrrrrr")
  if err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"erro":err})
    return
  }
  c.JSON(http.StatusOK, gin.H{"msg":"Email sended"})
}
