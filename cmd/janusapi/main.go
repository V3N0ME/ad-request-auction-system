package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	database "janusapi/infrastructure/database"
	models "janusapi/pkg/models"

	templateHandler "janusapi/pkg/template/delivery/http"
	_templateRepo "janusapi/pkg/template/repository"
	_templateUsecase "janusapi/pkg/template/usecase"
)

var (
	templateRepo *_templateRepo.Repository
)

var (
	templateUseCase *_templateUsecase.Usecase
)

var (
	masterDB *sql.DB
)

func initInfrastructure() {

	var err error

	masterDB, err = database.GetMysqlConnection(&models.MysqlConfig{
		Host:     "127.0.0.1",
		User:     "root",
		Password: "root",
		Database: "test_master",
		Port:     "8889",
	})

	if err != nil {
		panic(err)
	}
}

func initRepositories() {
	templateRepo = _templateRepo.NewMysqlRepo(masterDB)
}

func initUseCases() {
	templateUseCase = _templateUsecase.New(templateRepo)
}

func initHandlers(router gin.IRouter) {
	templateHandler.InitHandler(router, templateUseCase)
}

func main() {
	initInfrastructure()
	initRepositories()
	initUseCases()

	r := gin.Default()
	initHandlers(r)
	r.Run()
}
