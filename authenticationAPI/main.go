package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/microservices/authenticationAPI/apiEngine"
	"github.com/humbertovnavarro/microservices/authenticationAPI/config"
	"github.com/humbertovnavarro/microservices/authenticationAPI/context"
	_ "github.com/lib/pq"
)

func main() {
	appConfig := config.LoadConfig()
	context.DB = startStorage(&appConfig.Database)
	context.API = startApi(&appConfig.Api)
}

func startApi(config *config.ApiConfig) *gin.Engine {
	app := apiEngine.New()
	err := app.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.Fatal(err.Error())
	}
	return app
}

func startStorage(config *config.DatabaseConfig) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if (err) != nil {
		log.Fatal(err.Error())
		db.Close()
	}
	return db
}
