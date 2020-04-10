package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	errors "bidderservice/pkg/errors"

	bidderHandler "bidderservice/pkg/bidder/delivery/http"
	_bidderUsecase "bidderservice/pkg/bidder/usecase"

	"github.com/google/uuid"
)

var (
	bidderUseCase *_bidderUsecase.Usecase
)

func initUseCases() {

	id, err := uuid.NewUUID()
	if err != nil {
		panic(errors.BidderIDGenerationError)
	}
	fileID := id.String()

	bidDelay := os.Getenv("BID_DELAY")
	bidDelayInt, err := strconv.Atoi(bidDelay)
	if err != nil {
		panic(errors.DelayBidError)
	}

	bidderUseCase = _bidderUsecase.New(fileID, os.Getenv("PORT"), bidDelayInt)
	bidderUseCase.Register()
}

func initHandlers(router gin.IRouter) {
	bidderHandler.InitHandler(router, bidderUseCase)
}

func main() {
	initUseCases()
	r := gin.Default()
	initHandlers(r)
	r.Run()
}
