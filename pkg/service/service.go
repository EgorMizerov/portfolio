package service

import (
	"github.com/EgorMizerov/portfolio/pkg/models"
	"github.com/EgorMizerov/portfolio/pkg/repository"
)

type Work interface {
	Create(work models.Work) (int, error)
	GetAll() ([]models.Work, error)
	Update(id int, work models.Work) error
	Delete(id int) error
}

type Service struct {
	Work
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Work: NewWorkService(rep),
	}
}
