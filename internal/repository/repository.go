package repository

import (
	"SpyCatAgency/internal/model"
	"context"
)

type CatRepository interface {
	Create(ctx context.Context, cat *model.Cat) error
	Update(ctx context.Context, cat *model.Cat) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.Cat, error)
	List(ctx context.Context) ([]model.Cat, error)
}

type MissionRepository interface {
	Create(ctx context.Context, mission *model.Mission) error
	Update(ctx context.Context, mission *model.Mission) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.Mission, error)
	List(ctx context.Context) ([]model.Mission, error)
	AssignCat(ctx context.Context, missionID, catID uint) error
}

type TargetRepository interface {
	Create(ctx context.Context, target *model.Target) error
	Update(ctx context.Context, target *model.Target) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.Target, error)
	ListByMissionID(ctx context.Context, missionID uint) ([]model.Target, error)
}
