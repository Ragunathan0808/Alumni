package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kabi175/alumini-app-backend/controllers"
	"github.com/kabi175/alumini-app-backend/repo"
	service "github.com/kabi175/alumini-app-backend/services"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getGinEngine() *gin.Engine {
	return gin.Default()
}

func getGormPgDB() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	url := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, dbname, password)
	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func buildContainer() *dig.Container {
	container := dig.New()

	err := container.Provide(getGinEngine)

	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(getGormPgDB)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(repo.NewDefaultUserRepo)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(service.NewJwtTokenService)

	if err != nil {

		log.Fatal(err)
	}

	err = container.Provide(service.NewDefaultAuthService)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(controllers.NewAuthMiddleware)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(service.NewDefaultUserService)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(controllers.AddGinAuthController)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(controllers.AddGinUserController)
	if err != nil {
		log.Fatal(err)
	}

	return container
}
