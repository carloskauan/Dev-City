package routes

import(
  "DevCity/api/controllers"
  "github.com/gin-gonic/gin"
)

func HandlerFunc(){
  app := gin.Default()

  //Users Routes
  app.GET("/user", controllers.GetAllUsers)
  app.POST("/user", controllers.RegisterUser)
  app.GET("/user/:email", controllers.SeachUserByEmail)

  app.Run()
}
