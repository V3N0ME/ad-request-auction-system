package usecase

import (
	"testing"
	"time"
)

func TestMakeBid(t *testing.T) {

	bidderID := "123"
	bidderPort := "3000" 
	bidDelay := 200

	bidderUsecase := New(bidderID, bidderPort, bidDelay)

	endTime := time.Now().Add(  time.Duration(bidDelay) * time.Millisecond )

	resp := bidderUsecase.MakeBid()
	if resp.BidderID != "123" {
		t.Errorf("Incorrect bidder id")
	}

	if(time.Now().Before(endTime)) {
		t.Errorf("Bid was not delayed")
	}
}