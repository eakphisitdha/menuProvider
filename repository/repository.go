package repository

import (
	"database/sql"
	"menuProvider/model"
)

type IRepository interface {
	Find() ([]*model.MenuItem, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{db: db}
}

func (r *Repository) Find() ([]*model.MenuItem, error) {

	// query
	rows, err := r.db.Query("SELECT * FROM menu")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menuItems []*model.MenuItem

	// loop from query
	for rows.Next() {
		var menuItem model.MenuItem
		err := rows.Scan(
			&menuItem.ID,
			&menuItem.ParentID,
			&menuItem.Title,
			&menuItem.Name,
			&menuItem.Route,
			&menuItem.Icon,
			&menuItem.IsChildren,
		)
		if err != nil {
			return nil, err
		}
		menuItems = append(menuItems, &menuItem)
	}
	return menuItems, nil
}
