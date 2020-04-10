package models

//Bidder is the model of a bidder
type Bidder struct {
	BidderID      string `json:"bidder_id"`
	Port          string `json:"port"`
	ResponseDelay int    `json:"responseDelay"`
}
