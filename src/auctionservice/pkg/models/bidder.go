package models

//Bidder is the model of a bidder
type Bidder struct {
	BidderID      string `json:"bidder_id"`
	Port          string `json:"port"`
	ResponseDelay int    `json:"responseDelay"`
}

//BidderResponse is the http response from the bidder
type BidderResponse struct {
	BidderID string  `json:"bidder_id"`
	BidValue float64 `json:"bidder_value"`
}
