package auction

import models "bidderservice/pkg/models"

//Usecase is the interface of bidder's interface
type Usecase interface {
	MakeBid() models.BidderResponse
	Register()
}
