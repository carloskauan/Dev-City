package routes

import(
  "DevCity/api/controllers"
  "github.com/gin-gonic/gin"
)

func HandlerFunc(){
  app := gin.Default()

  //Users Routes
  app.GET("/user", controllers.GetAllUsers)
  app.GET("/user/:email", controllers.SearchUserByEmail)
  app.POST("/user", controllers.RegisterUser)
  app.GET("/user/:email/:password", controllers.Authenticate)
  // Verify
  app.POST("/user/verify/:email", controllers.SendCodeEmailAndSave)
  app.GET("/user/verify/:email/:code", controllers.ComapareCode)
  app.DELETE("/user/verify/:id", controllers.DeleteCode)

  app.Run()
}
