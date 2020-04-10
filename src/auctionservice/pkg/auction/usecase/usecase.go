package usecase

import (
	auction "auctionservice/pkg/auction"
	models "auctionservice/pkg/models"
)

//Usecase is the usecase of auction
type Usecase struct {
	repo auction.Repository
}

//New returns a new instance of auction's usecase
func New(repo auction.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

//StartAuction starts the bidding round and returns the auction response
func (u *Usecase) StartAuction(auctionID string) (models.AuctionResponse, error) {
	return models.AuctionResponse{
		BidderID:    "123",
		MaxBidValue: 100,
	}, nil
}
