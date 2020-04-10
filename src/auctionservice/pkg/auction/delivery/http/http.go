package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	helpers "auctionservice/helpers"
	auction "auctionservice/pkg/auction"
	models "auctionservice/pkg/models"
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
	route.POST("/bidder", handler.registerBidder)
	route.GET("/bidder", handler.getAllBidders)
}

func (h *Handler) create(c *gin.Context) {
	auctionResponse := h.usecase.StartAuction("123")
	if auctionResponse.BidderID == "" {
		helpers.Respond(c, 404, "No Bidder Online", nil)
	} else {
		helpers.Respond(c, 200, "Auction Completed", map[string]interface{}{
			"bidder_id":     auctionResponse.BidderID,
			"max_bid_value": auctionResponse.MaxBidValue,
		})
	}
}

func (h *Handler) getAllBidders(c *gin.Context) {
	bidders := h.usecase.GetAllBidders()
	if len(bidders) == 0 {
		helpers.Respond(c, 404, "No bidder registered", nil)
	} else {
		helpers.Respond(c, 200, "Success", map[string]interface{}{
			"bidders": bidders,
		})
	}
}

func (h *Handler) registerBidder(c *gin.Context) {

	var bidder models.Bidder

	if err := c.ShouldBindWith(&bidder, binding.JSON); err != nil {
		helpers.Respond(c, 400, err.Error(), nil)
		return
	}

	h.usecase.RegisterBidder(bidder)
	helpers.Respond(c, 200, "Bidder Registered", nil)
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
