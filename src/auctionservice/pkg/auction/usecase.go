package auction

import models "auctionservice/pkg/models"

//Usecase is the interface of auction's interface
type Usecase interface {
	StartAuction(string) (models.AuctionResponse, error)
}
