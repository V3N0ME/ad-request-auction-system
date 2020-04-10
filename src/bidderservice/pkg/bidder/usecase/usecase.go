package usecase

import (
	models "bidderservice/pkg/models"
	request "bidderservice/pkg/request"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"time"
)

//Usecase is the usecase of auction
type Usecase struct {
	req      *request.CustomHTTP
	bidderID string
	port     string
}

const retryDelay = 5

const minBid = 10
const maxBid = 10000

//New returns a new instance of auction's usecase
func New(bidderID, port string, bidderTimeout int) *Usecase {

	req := request.New(request.Config{
		Timeout:            time.Duration(bidderTimeout) * time.Microsecond,
		MaxOpenConnections: 500,
	})

	return &Usecase{
		req:      req,
		bidderID: bidderID,
		port:     port,
	}
}

//MakeBid returns the bid value
func (u *Usecase) MakeBid() models.BidderResponse {

	bidValue := minBid + rand.Float64()*(maxBid-minBid)

	return models.BidderResponse{
		BidderID: u.bidderID,
		BidValue: math.Floor(bidValue*100) / 100, //rounds off to 2 decimal places
	}
}

//Register registers the bidder with the auctioner
func (u *Usecase) Register() {

	log.Println("Registering Bidder...")

	values := map[string]string{"bidder_id": u.bidderID, "port": u.port}
	jsonValue, _ := json.Marshal(values)

	_, statusCode, err := u.req.MakeRequest(request.Request{
		URL:     "http://127.0.0.1:8080/auction/bidder",
		Method:  "POST",
		Payload: jsonValue,
	})

	if err != nil {
		time.Sleep(time.Duration(retryDelay) * time.Second)
		u.Register()
	}

	if statusCode != 200 {
		time.Sleep(time.Duration(retryDelay) * time.Second)
		u.Register()
	}

	log.Println("Bidder Registered")
}
