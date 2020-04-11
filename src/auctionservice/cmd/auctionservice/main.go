package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	errors "auctionservice/pkg/errors"

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

func initRepositories() {
	auctionRepo = _auctionRepo.NewMemoryRepository()
}

func initUseCases() {
	bidDelay := os.Getenv("BID_DELAY")
	bidDelayInt, err := strconv.Atoi(bidDelay)
	if err != nil {
		panic(errors.DelayBidError)
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
	r.Run()
}
