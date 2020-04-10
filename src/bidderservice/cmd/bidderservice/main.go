package main

import (
	"github.com/gin-gonic/gin"

	bidderHandler "bidderservice/pkg/bidder/delivery/http"
	_bidderUsecase "bidderservice/pkg/bidder/usecase"
)

var (
	bidderUseCase *_bidderUsecase.Usecase
)

func initUseCases() {
	bidderUseCase = _bidderUsecase.New(250)
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
