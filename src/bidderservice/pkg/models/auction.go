package models

//AuctionResponse is the model of the reponse from the auction
type AuctionResponse struct {
	BidderID    string  `json:"bidder_id"`
	MaxBidValue float64 `json:"price"`
}
