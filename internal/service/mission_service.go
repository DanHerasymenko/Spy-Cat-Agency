package service

import (
	"SpyCatAgency/internal/model"
	"SpyCatAgency/internal/repository"
	"context"
	"errors"
)

type MissionService struct {
	missionRepo repository.MissionRepository
	targetRepo  repository.TargetRepository
	catRepo     repository.CatRepository
}

func NewMissionService(
	missionRepo repository.MissionRepository,
	targetRepo repository.TargetRepository,
	catRepo repository.CatRepository,
) *MissionService {
	return &MissionService{
		missionRepo: missionRepo,
		targetRepo:  targetRepo,
		catRepo:     catRepo,
	}
}

func (s *MissionService) Create(ctx context.Context, create model.MissionCreate) (*model.Mission, error) {

	// Check if cat exists
	cat, err := s.catRepo.GetByID(ctx, create.CatID)
	if err != nil {
		return nil, err
	}

	mission := &model.Mission{
		CatID: create.CatID,
		Cat:   *cat,
	}

	if err := s.missionRepo.Create(ctx, mission); err != nil {
		return nil, err
	}

	// Create targets
	for _, targetCreate := range create.Targets {
		target := &model.Target{
			MissionID: mission.ID,
			Name:      targetCreate.Name,
			Country:   targetCreate.Country,
			Notes:     targetCreate.Notes,
		}
		if err := s.targetRepo.Create(ctx, target); err != nil {
			return nil, err
		}
		mission.Targets = append(mission.Targets, *target)
	}

	return mission, nil
}

func (s *MissionService) Update(ctx context.Context, id uint, update model.MissionUpdate) (*model.Mission, error) {

	mission, err := s.missionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	mission.Completed = update.Completed

	if err := s.missionRepo.Update(ctx, mission); err != nil {
		return nil, err
	}

	return mission, nil
}

func (s *MissionService) Delete(ctx context.Context, id uint) error {

	mission, err := s.missionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if mission.CatID != 0 {
		return errors.New("cannot delete mission that is assigned to a cat")
	}

	return s.missionRepo.Delete(ctx, id)
}

func (s *MissionService) GetByID(ctx context.Context, id uint) (*model.Mission, error) {
	return s.missionRepo.GetByID(ctx, id)
}

func (s *MissionService) List(ctx context.Context) ([]model.Mission, error) {
	return s.missionRepo.List(ctx)
}

func (s *MissionService) AssignCat(ctx context.Context, missionID, catID uint) error {
	return s.missionRepo.AssignCat(ctx, missionID, catID)
}

func (s *MissionService) AddTarget(
	ctx context.Context,
	missionID uint,
	targetCreate model.TargetCreate,
) (*model.Target, error) {
	mission, err := s.missionRepo.GetByID(ctx, missionID)
	if err != nil {
		return nil, err
	}

	if mission.Completed {
		return nil, errors.New("cannot add target to completed mission")
	}

	targets, err := s.targetRepo.ListByMissionID(ctx, missionID)
	if err != nil {
		return nil, err
	}

	if len(targets) >= 3 {
		return nil, errors.New("mission already has maximum number of targets")
	}

	target := &model.Target{
		MissionID: missionID,
		Name:      targetCreate.Name,
		Country:   targetCreate.Country,
		Notes:     targetCreate.Notes,
	}

	if err := s.targetRepo.Create(ctx, target); err != nil {
		return nil, err
	}

	return target, nil
}

func (s *MissionService) DeleteTarget(ctx context.Context, targetID uint) error {
	target, err := s.targetRepo.GetByID(ctx, targetID)
	if err != nil {
		return err
	}

	if target.Completed {
		return errors.New("cannot delete completed target")
	}

	return s.targetRepo.Delete(ctx, targetID)
}

func (s *MissionService) UpdateTarget(
	ctx context.Context,
	targetID uint,
	update model.TargetUpdate,
) (*model.Target, error) {
	target, err := s.targetRepo.GetByID(ctx, targetID)
	if err != nil {
		return nil, err
	}

	mission, err := s.missionRepo.GetByID(ctx, target.MissionID)
	if err != nil {
		return nil, err
	}

	if mission.Completed || target.Completed {
		return nil, errors.New("cannot update target in completed mission or completed target")
	}

	target.Notes = update.Notes
	target.Completed = update.Completed

	if err := s.targetRepo.Update(ctx, target); err != nil {
		return nil, err
	}

	return target, nil
}
