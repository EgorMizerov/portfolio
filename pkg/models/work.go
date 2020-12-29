package models

type Work struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Date        string `json:"date_up" db:"date_up"`
	Tag         string `json:"tag" db:"tag"`
	Img         string `json:"img" db:"img"`
	Url         string `json:"url" db:"url"`
}
