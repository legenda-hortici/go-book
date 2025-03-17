package services

import (
	"go-book/pkg/models"
	"go-book/pkg/repositories"
)

type TopicService struct {
	repo *repositories.TopicRepository
}

func NewTopicService(repo *repositories.TopicRepository) *TopicService {
	return &TopicService{repo: repo}
}

func (s *TopicService) CreateTopic(topic models.Topic) error {
	return s.repo.InsertTopic(topic)
}

func (s *TopicService) GetTopics() ([]models.Topic, error) {
	return s.repo.GetTopics()
}

func (s *TopicService) DeleteTopic(topic models.Topic) error {
	return s.repo.DeleteTopic(topic)
}
