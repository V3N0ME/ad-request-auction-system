package main

import (
	"database/sql"
	"os"
	"strconv"

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
	bidDelay := os.Getenv("BID_DELAY")
	bidDelayInt, err := strconv.Atoi(bidDelay)
	if err != nil {
		panic(err)
	}
	auctionUseCase = _auctionUsecase.New(auctionRepo, bidDelayInt)
}

func initHandlers(router gin.IRouter) {
	auctionHandler.InitHandler(router, auctionUseCase)
}

func main() {
	initRepositories()
	initUseCases()

	r := gin.Default()
	initHandlers(r)
	r.Run("0.0.0.0:8080")
}
