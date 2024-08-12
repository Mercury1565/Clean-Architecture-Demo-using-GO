package main

import (
	"Clean_Architecture/api/route"
	"Clean_Architecture/bootstrap"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	database := app.Mongo.Database(env.DBName)
	defer app.CloseMongoDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, *database, gin)
	gin.Run(env.ServerAddress)
}
