package repository

import (
	"SpyCatAgency/internal/model"
	"SpyCatAgency/internal/repository"
	"context"
	"database/sql"
	"fmt"
)

type TargetRepository struct {
	db *sql.DB
}

func NewTargetRepository(db *sql.DB) repository.TargetRepository {
	return &TargetRepository{db: db}
}

func (r *TargetRepository) Create(ctx context.Context, target *model.Target) error {
	query := `
		INSERT INTO targets (name, country, notes, mission_id, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(
		ctx, query,
		target.Name,
		target.Country,
		target.Notes,
		target.MissionID,
		target.Completed,
	).Scan(&target.ID, &target.CreatedAt, &target.UpdatedAt)
}

func (r *TargetRepository) Update(ctx context.Context, target *model.Target) error {
	query := `
		UPDATE targets
		SET name = $1, mission_id = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING updated_at`

	return r.db.QueryRowContext(
		ctx, query,
		target.Name,
		target.MissionID,
		target.ID,
	).Scan(&target.UpdatedAt)
}

func (r *TargetRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM targets WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *TargetRepository) GetByID(ctx context.Context, id uint) (*model.Target, error) {
	query := `
		SELECT id, name, mission_id, created_at, updated_at
		FROM targets
		WHERE id = $1`

	target := &model.Target{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&target.ID,
		&target.Name,
		&target.MissionID,
		&target.CreatedAt,
		&target.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("target not found")
	}
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (r *TargetRepository) ListByMissionID(ctx context.Context, missionID uint) ([]model.Target, error) {
	query := `
		SELECT id, name, mission_id, created_at, updated_at
		FROM targets
		WHERE mission_id = $1
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query, missionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var targets []model.Target
	for rows.Next() {
		var target model.Target
		if err := rows.Scan(
			&target.ID,
			&target.Name,
			&target.MissionID,
			&target.CreatedAt,
			&target.UpdatedAt,
		); err != nil {
			return nil, err
		}
		targets = append(targets, target)
	}
	return targets, rows.Err()
}
