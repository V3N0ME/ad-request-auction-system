package models

//Bidder is the model of a bidder
type Bidder struct {
	BidderID string `json:"bidder_id" binding:"required"`
	Port     string `json:"port" binding:"required"`
}

//BidderResponse is the http response from the bidder
type BidderResponse struct {
	BidderID string  `json:"bidder_id"`
	BidValue float64 `json:"bid_value"`
}
