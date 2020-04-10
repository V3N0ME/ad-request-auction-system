package usecase

import (
	auction "auctionservice/pkg/auction"
	models "auctionservice/pkg/models"
	request "auctionservice/pkg/request"
	"encoding/json"
	"time"
)

//Usecase is the usecase of auction
type Usecase struct {
	repo auction.Repository
	req  *request.CustomHTTP
}

//New returns a new instance of auction's usecase
func New(repo auction.Repository, bidderTimeout int) *Usecase {

	req := request.New(request.Config{
		Timeout:            time.Duration(bidderTimeout) * time.Microsecond,
		MaxOpenConnections: 500,
	})

	return &Usecase{
		repo: repo,
		req:  req,
	}
}

//StartAuction starts the bidding round and returns the auction response
func (u *Usecase) StartAuction(auctionID string) models.AuctionResponse {

	bidders := u.repo.GetAllBidders()
	auctionResponse := u.getAuctionResult(bidders)

	return auctionResponse
}

func (u *Usecase) getAuctionResult(bidders map[string]models.Bidder) models.AuctionResponse {

	type httpResponse struct {
		body       string
		statusCode int
		err        error
	}

	httpRes := make(chan httpResponse)

	for _, b := range bidders {
		go func(bidder models.Bidder) {

			body, statusCode, err := u.req.MakeRequest(request.Request{})
			httpRes <- httpResponse{
				body:       body,
				statusCode: statusCode,
				err:        err,
			}

		}(b)
	}

	var maxBidValue float64
	var maxBidderID string

	for range bidders {
		response := <-httpRes
		if response.err != nil {
			continue
		}
		if response.statusCode != 200 {
			continue
		}

		var bidderResponse models.BidderResponse
		json.Unmarshal([]byte(response.body), &bidderResponse)

		if bidderResponse.BidValue > maxBidValue {
			maxBidderID = bidderResponse.BidderID
			maxBidValue = bidderResponse.BidValue
		}
	}

	return models.AuctionResponse{
		BidderID:    maxBidderID,
		MaxBidValue: maxBidValue,
	}
}
