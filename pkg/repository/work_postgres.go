package repository

import (
	"fmt"
	"github.com/EgorMizerov/portfolio/pkg/models"
	"github.com/jmoiron/sqlx"
)

type WorkPostgres struct {
	db *sqlx.DB
}

func NewWorkPostgres(db *sqlx.DB) *WorkPostgres {
	return &WorkPostgres{db: db}
}

func (r *WorkPostgres) Create(work models.Work) (int, error) {
	fmt.Println(work)
	var id int
	query := fmt.Sprintf("INSERT INTO works (title, description, date_up, tag, img, url) values ($1, $2, $3, $4, $5, $6)")
	row := r.db.QueryRow(query, work.Title, work.Description, work.Date, work.Tag, work.Img, work.Url)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return id, nil
}

func (r *WorkPostgres) GetAll() ([]models.Work, error) {
	var lists []models.Work
	query := fmt.Sprintf("SELECT * FROM %s", works)
	err := r.db.Select(&lists, query)

	return lists, err
}
func (r *WorkPostgres) Update(id int, work models.Work) error {
	//return r.d.Update(id, work)
	return nil
}

func (r *WorkPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", works)
	_, err := r.db.Exec(query, id)

	return err
}
