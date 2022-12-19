package routes

import(
  "DevCity/api/controllers"
  "github.com/gin-gonic/gin"
)

func HandlerFunc(){
  app := gin.Default()
  app.POST("/user", controllers.Register)
  app.Run()
}
