package usecase

import "golang-test-task/internal/repository"

type NumberUseCase interface {
	AddAndGetSorted(n int) ([]int, error)
}

type Service struct {
	repo repository.NumberRepository
}

func NewService(repo repository.NumberRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddAndGetSorted(n int) ([]int, error) {
	if err := s.repo.Save(n); err != nil {
		return nil, err
	}
	return s.repo.GetAllSorted()
}