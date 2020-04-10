package repository

import "database/sql"

//Repository is the instance of template's repository
type Repository struct {
	db *sql.DB
}

//NewMysqlRepo returns an instance of template's repository
func NewMysqlRepo(db *sql.DB) *Repository {
	return &Repository{db}
}
