package usecase

import (
	"testing"

	_auctionRepo "auctionservice/pkg/auction/repository"
)

func TestStartAuction(t *testing.T) {

	repo := _auctionRepo.NewMemoryRepository()
	usecase := New(repo, 250)

	usecase.StartAuction("test")
}
