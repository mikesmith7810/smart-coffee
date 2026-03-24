package repository

import (
	"database/sql"
	"errors"
	"smart-coffee/domain"
)

var ErrNotFound = errors.New("coffee not found")

type CoffeeRepository struct {
	db *sql.DB
}

func NewCoffeeRepository(db *sql.DB) *CoffeeRepository {
	return &CoffeeRepository{db: db}
}

func (r *CoffeeRepository) FindByID(id string) (domain.Coffee, error) {
	var coffee domain.Coffee
	err := r.db.QueryRow(
		"SELECT id, name, calories FROM coffees WHERE id = ?", id,
	).Scan(&coffee.Id, &coffee.Name, &coffee.Calories)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Coffee{}, ErrNotFound
	}
	return coffee, err
}

func (r *CoffeeRepository) Upsert(coffee domain.Coffee) error {
	_, err := r.db.Exec(
		`INSERT INTO coffees (id, name, calories) VALUES (?, ?, ?)
		 ON DUPLICATE KEY UPDATE name = VALUES(name), calories = VALUES(calories)`,
		coffee.Id, coffee.Name, coffee.Calories,
	)
	return err
}
