package service

import (
	"errors"
	"github.com/EgorMizerov/portfolio/pkg/models"
	"github.com/EgorMizerov/portfolio/pkg/repository"
)

type WorkService struct {
	repo repository.Work
}

func NewWorkService(repo repository.Work) *WorkService {
	return &WorkService{repo: repo}
}

func (s *WorkService) Create(work models.Work) (int, error) {
	id, err := s.repo.Create(work)
	if err != nil {
		return 0, errors.New("Ошибка при работе с базой данных")
	}

	return id, nil
}

func (s *WorkService) GetAll() ([]models.Work, error) {
	return s.repo.GetAll()
}

func (s *WorkService) Update(id int, work models.Work) error {
	return s.repo.Update(id, work)
}

func (s *WorkService) Delete(id int) error {
	return s.repo.Delete(id)
}
