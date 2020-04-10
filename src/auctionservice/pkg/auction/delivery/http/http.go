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
	router.GET("/list", handler.listEndpoints)
	route := router.Group("/auction")
	route.POST("/", handler.create)
}

func (h *Handler) create(c *gin.Context) {
	auctionResponse := h.usecase.StartAuction("123")
	if auctionResponse.BidderID == "" {
		helpers.Respond(c, 404, "No bidder online", nil)
	} else {
		helpers.Respond(c, 200, "Auction Completed", map[string]interface{}{
			"bidder_id":     auctionResponse.BidderID,
			"max_bid_value": auctionResponse.MaxBidValue,
		})
	}
}

func (h *Handler) listEndpoints(c *gin.Context) {

	type Endpoint struct {
		Path    string   `json:"path"`
		Methods []string `json:"methods"`
	}

	helpers.Respond(c, 200, "Success", map[string]interface{}{
		"endpoints": []Endpoint{
			Endpoint{
				Path:    "/auction",
				Methods: []string{"POST"},
			},
			Endpoint{
				Path:    "/bidder",
				Methods: []string{"POST"},
			},
			Endpoint{
				Path:    "/list",
				Methods: []string{"GET"},
			},
		},
	},
	)
}
