package auction

import models "auctionservice/pkg/models"

//Repository is the interface of auction's repository
type Repository interface {
	GetAllBidders() map[string]models.Bidder
}
