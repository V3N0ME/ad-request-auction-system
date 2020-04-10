package http

import (
	"github.com/gin-gonic/gin"

	helpers "auctionservice/helpers"
	auction "auctionservice/pkg/auction"
)

//Handler is the instance od the auction's handler
type Handler struct {
	usecase auction.Usecase
}

//InitHandler initialises template's http handler
func InitHandler(router gin.IRouter, uc auction.Usecase) {
	handler := &Handler{uc}
	route := router.Group("/auction")
	route.POST("/", handler.create)
}

func (h *Handler) create(c *gin.Context) {
	h.usecase.StartAuction("123")
	helpers.Respond(c, 200, "Report Created", nil)
}
