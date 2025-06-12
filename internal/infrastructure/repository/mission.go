package repository

import (
	"SpyCatAgency/internal/model"
	"SpyCatAgency/internal/repository"
	"context"
	"database/sql"
	"fmt"
)

type MissionRepository struct {
	db *sql.DB
}

func NewMissionRepository(db *sql.DB) repository.MissionRepository {
	return &MissionRepository{db: db}
}

func (r *MissionRepository) Create(ctx context.Context, mission *model.Mission) error {
	query := `
		INSERT INTO missions (name, cat_id, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(
		ctx, query,
		mission.Name,
		mission.CatID,
	).Scan(&mission.ID, &mission.CreatedAt, &mission.UpdatedAt)
}

func (r *MissionRepository) Update(ctx context.Context, mission *model.Mission) error {
	query := `
		UPDATE missions
		SET name = $1, cat_id = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING updated_at`

	return r.db.QueryRowContext(
		ctx, query,
		mission.Name,
		mission.CatID,
		mission.ID,
	).Scan(&mission.UpdatedAt)
}

func (r *MissionRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM missions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *MissionRepository) GetByID(ctx context.Context, id uint) (*model.Mission, error) {
	query := `
		SELECT id, name, cat_id, created_at, updated_at
		FROM missions
		WHERE id = $1`

	mission := &model.Mission{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&mission.ID,
		&mission.Name,
		&mission.CatID,
		&mission.CreatedAt,
		&mission.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("mission not found")
	}
	if err != nil {
		return nil, err
	}
	return mission, nil
}

func (r *MissionRepository) List(ctx context.Context) ([]model.Mission, error) {
	query := `
		SELECT id, name, cat_id, created_at, updated_at
		FROM missions
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []model.Mission
	for rows.Next() {
		var mission model.Mission
		if err := rows.Scan(
			&mission.ID,
			&mission.Name,
			&mission.CatID,
			&mission.CreatedAt,
			&mission.UpdatedAt,
		); err != nil {
			return nil, err
		}
		missions = append(missions, mission)
	}
	return missions, rows.Err()
}

func (r *MissionRepository) AssignCat(ctx context.Context, missionID, catID uint) error {
	query := `
		UPDATE missions
		SET cat_id = $1, updated_at = NOW()
		WHERE id = $2`

	res, err := r.db.ExecContext(ctx, query, catID, missionID)
	if err != nil {
		return fmt.Errorf("failed to assign cat to mission: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("mission with id %d not found", missionID)
	}

	return nil
}
