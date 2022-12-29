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
  user.Password = middleware.Encripter(user.Password)

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

func SendCodeEmailAndSave(c *gin.Context){
  email := c.Param("email")

  //Gerando e criptografando codigo
	code := middleware.GenerateCode()
  hashCode := string(middleware.Encripter(code))
	
	// Salvando codigo e email no banco
	dataCode := models.Codes{EmailUser: email, Code: hashCode}
	database.DB.Create(&dataCode)

	//Enviando Email
	err := middleware.SendWhitGomail(email, code)
  if err != nil{
    c.JSON(http.StatusOK, gin.H{ 
      "error":"Erro ao enviar codigo de verificação", 
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{ 
    "status":"Emeil e codigo de verificação enviado",
  })
}

func ComapareCode(c *gin.Context){
  var codeRegister models.Codes
  email := c.Param("email")
  code := c.Param("code")

  // Query do codigo do banco pelo email
  database.DB.Where(&models.Codes{EmailUser: email}).First(&codeRegister)
  if codeRegister.ID == 0{ 
    c.JSON(http.StatusBadRequest, gin.H{"error":"Email sem registro de envio de codigo"})
    return
  }
  fmt.Println("\n",codeRegister.Code,"\n",codeRegister.Code, middleware.Encripter(code),"\n", codeRegister.Code == middleware.Encripter(code))
  // Comparação de codigo recebido pelo ending point e codigo registrado no banco
  err := bcrypt.CompareHashAndPassword([]byte(codeRegister.Code), []byte(code))
  if err != nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "equal":false,
      "status":"Codigo informado diferente do gerado",
      "ID":codeRegister.ID,
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "equal":true,
    "status":"Codigo infomado igual ao gerado",
    "ID":codeRegister.ID,
  })
}

func DeleteCode(c *gin.Context){
  var code models.Codes
  database.DB.Delete(&code, c.Param("id"))
  database.DB.First(&code, c.Param("id"))
  if code.ID != 0{
    c.JSON(http.StatusOK, gin.H{
      "status":"Erro ao apagar codigo de verificação do banco",
      "deleted":false,
    })
    return
  } 
  c.JSON(http.StatusOK, gin.H{
    "status":"Codigo de verificação do banco apagado com sucesso",
    "deleted":true,
  })
}