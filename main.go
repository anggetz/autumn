package main

import (
	"autumn/controllers"
	"autumn/models"
	"autumn/pkg/autumncore/generic"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	var configPath = flag.String("config", ".env", "please input config path")

	flag.Parse()

	godotenv.Load(*configPath)

	route := gin.New()

	fmt.Println("Server listening on port " + os.Getenv("APP_PORT"))

	groupApi := route.Group("v1")
	{
		generic.Resources[models.User]("user", groupApi, controllers.NewUserController())
	}

	route.Run(":" + os.Getenv("APP_PORT"))
}
