package services

import (
	"go-book/pkg/models"
	"go-book/pkg/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlockService struct {
	repo *repositories.BlockRepository
}

func NewBlockService(repo *repositories.BlockRepository) *BlockService {
	return &BlockService{repo: repo}
}

func (b *BlockService) AddBlock(block models.Block) error {
	return b.repo.InsertBlock(block)
}

func (s *BlockService) GetBlocks(id primitive.ObjectID) ([]models.Block, error) {
	return s.repo.GetBlocks(id)
}

func (s *BlockService) DeleteAllBlocks(id primitive.ObjectID) error {
	return s.repo.DeleteAllBlocks(id)
}

func (s *BlockService) DeleteBlock(block models.Block) error {
	return s.repo.DeleteBlock(block.ID)
}
