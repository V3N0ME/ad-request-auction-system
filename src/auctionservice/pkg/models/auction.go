package models

//Auction is the model of an auction
type Auction struct {
	AuctionID string `json:"auction_id"`
}

//AuctionResponse is the model of the reponse from the auction
type AuctionResponse struct {
	BidderID    string  `json:"bidder_id"`
	MaxBidValue float64 `json:"price"`
}
