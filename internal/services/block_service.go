package services

import (
	"context"
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

func (b *BlockService) AddBlock(ctx context.Context, block models.Block) error {
	return b.repo.InsertBlock(ctx, block)
}

func (s *BlockService) GetBlocks(ctx context.Context, id primitive.ObjectID) ([]models.Block, error) {
	return s.repo.GetBlocks(ctx, id)
}

func (s *BlockService) DeleteAllBlocks(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.DeleteAllBlocks(ctx, id)
}

func (s *BlockService) DeleteBlock(ctx context.Context, block models.Block) error {
	return s.repo.DeleteBlock(ctx, block.ID)
}

func (s *BlockService) UpdateBlock(ctx context.Context, block models.Block) error {
	return s.repo.UpdateBlock(ctx, block.ID, block.Content)
}
