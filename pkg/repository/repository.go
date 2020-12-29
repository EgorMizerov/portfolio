package repository

import (
	"github.com/EgorMizerov/portfolio/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Work interface {
	Create(work models.Work) (int, error)
	GetAll() ([]models.Work, error)
	Update(id int, work models.Work) error
	Delete(id int) error
}

type Repository struct {
	Work
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Work: NewWorkPostgres(db),
	}
}
