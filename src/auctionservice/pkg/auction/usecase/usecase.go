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

//RegisterBidder registers a new bidder
func (u *Usecase) RegisterBidder(bidder models.Bidder) {
	u.repo.RegisterBidder(bidder)
}

//GetAllBidders returns all registered bidders
func (u *Usecase) GetAllBidders() []models.Bidder {
	allBidders := make([]models.Bidder, 0)
	bidders := u.repo.GetAllBidders()

	for _, bidder := range bidders {
		allBidders = append(allBidders, bidder)
	}

	return allBidders
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

	httpRes := make(chan httpResponse, len(bidders))

	for _, b := range bidders {
		go func(bidder models.Bidder) {

			type Request struct {
				URL     string
				Method  string
				Payload []byte
				Headers map[string]string
			}

			body, statusCode, err := u.req.MakeRequest(request.Request{
				URL:    "http://127.0.0.1:3000/bid/",
				Method: "GET",
			})

			httpRes <- httpResponse{
				body:       body,
				statusCode: statusCode,
				err:        err,
			}

		}(b)
	}

	var maxBidValue float64
	var maxBidderID string

	type apiResponse struct {
		Data models.BidderResponse `json:"data"`
	}

	for range bidders {

		response := <-httpRes

		if response.err != nil {
			continue
		}
		if response.statusCode != 200 {
			continue
		}

		var apiResp apiResponse
		json.Unmarshal([]byte(response.body), &apiResp)

		bidderResponse := apiResp.Data

		if bidderResponse.BidValue > maxBidValue {
			maxBidderID = bidderResponse.BidderID
			maxBidValue = bidderResponse.BidValue
		}
	}

	close(httpRes)

	return models.AuctionResponse{
		BidderID:    maxBidderID,
		MaxBidValue: maxBidValue,
	}
}
