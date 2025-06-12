package repository

import (
	"SpyCatAgency/internal/model"
	"SpyCatAgency/internal/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type CatRepository struct {
	db *sql.DB
}

func NewCatRepository(db *sql.DB) repository.CatRepository {
	return &CatRepository{db: db}
}

func (r *CatRepository) Create(ctx context.Context, cat *model.Cat) error {
	query := `
		INSERT INTO cats (name, years_experience, breed, salary, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(
		ctx, query,
		cat.Name,
		cat.YearsExperience,
		cat.Breed,
		cat.Salary,
	).Scan(&cat.ID, &cat.CreatedAt, &cat.UpdatedAt)
}

func (r *CatRepository) Update(ctx context.Context, cat *model.Cat) error {
	query := `
		UPDATE cats
		SET salary = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING updated_at`

	return r.db.QueryRowContext(ctx, query, cat.Salary, cat.ID).Scan(&cat.UpdatedAt)
}

func (r *CatRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM cats WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *CatRepository) GetByID(ctx context.Context, id uint) (*model.Cat, error) {
	query := `
		SELECT id, name, years_experience, breed, salary, created_at, updated_at
		FROM cats
		WHERE id = $1`

	cat := &model.Cat{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&cat.ID,
		&cat.Name,
		&cat.YearsExperience,
		&cat.Breed,
		&cat.Salary,
		&cat.CreatedAt,
		&cat.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("cat not found")
	}
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (r *CatRepository) List(ctx context.Context) ([]model.Cat, error) {
	query := `
		SELECT id, name, years_experience, breed, salary, created_at, updated_at
		FROM cats
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []model.Cat
	for rows.Next() {
		var cat model.Cat
		if err := rows.Scan(
			&cat.ID,
			&cat.Name,
			&cat.YearsExperience,
			&cat.Breed,
			&cat.Salary,
			&cat.CreatedAt,
			&cat.UpdatedAt,
		); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	return cats, rows.Err()
}
