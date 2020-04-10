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

//GetAllBidders returns all the registered bidders
func (m *MemoryRepository) GetAllBidders() map[string]models.Bidder {
	return m.bidders
}
