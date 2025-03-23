package services

import (
	"context"
	"go-book/pkg/models"
	"go-book/pkg/repositories"
)

type TopicService struct {
	repo *repositories.TopicRepository
}

func NewTopicService(repo *repositories.TopicRepository) *TopicService {
	return &TopicService{repo: repo}
}

func (s *TopicService) CreateTopic(ctx context.Context, topic models.Topic) error {
	return s.repo.InsertTopic(ctx, topic)
}

func (s *TopicService) GetTopics(ctx context.Context) ([]models.Topic, error) {
	return s.repo.GetTopics(ctx)
}

func (s *TopicService) DeleteTopic(ctx context.Context, topic models.Topic) error {
	return s.repo.DeleteTopic(ctx, topic)
}
