package repository

import models "auctionservice/pkg/models"

//MemoryRepository is the instance of auction's in memory repository
type MemoryRepository struct {
	bidders map[string]models.Bidder
}

//NewMemoryRepository returns a new instance of auction's memory repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		bidders: make(map[string]models.Bidder),
	}
}
