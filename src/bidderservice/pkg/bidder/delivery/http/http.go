package http

import (
	"github.com/gin-gonic/gin"

	helpers "bidderservice/helpers"
	bidder "bidderservice/pkg/bidder"
)

//Handler is the instance od the auction's handler
type Handler struct {
	usecase bidder.Usecase
}

//InitHandler initialises template's http handler
func InitHandler(router gin.IRouter, uc bidder.Usecase) {
	handler := &Handler{uc}
	route := router.Group("/bid")
	route.GET("/", handler.getBid)
}

func (h *Handler) getBid(c *gin.Context) {
	bidderResponse := h.usecase.MakeBid()
	helpers.Respond(c, 200, "Success", map[string]interface{}{
		"bidder_id": bidderResponse.BidderID,
		"bid_value": bidderResponse.BidValue,
	})
}
