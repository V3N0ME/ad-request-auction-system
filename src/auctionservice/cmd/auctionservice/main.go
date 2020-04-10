package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	auctionHandler "auctionservice/pkg/auction/delivery/http"
	_auctionRepo "auctionservice/pkg/auction/repository"
	_auctionUsecase "auctionservice/pkg/auction/usecase"
)

var (
	auctionRepo *_auctionRepo.MemoryRepository
)

var (
	auctionUseCase *_auctionUsecase.Usecase
)

var (
	masterDB *sql.DB
)

func initRepositories() {
	auctionRepo = _auctionRepo.NewMemoryRepository()
}

func initUseCases() {
	auctionUseCase = _auctionUsecase.New(auctionRepo, 250)
}

func initHandlers(router gin.IRouter) {
	auctionHandler.InitHandler(router, auctionUseCase)
}

func main() {
	initRepositories()
	initUseCases()

	r := gin.Default()
	initHandlers(r)
	r.Run()
}
