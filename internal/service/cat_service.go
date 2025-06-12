package service

import (
	"SpyCatAgency/internal/client"
	"SpyCatAgency/internal/model"
	"SpyCatAgency/internal/repository"
	"context"
	"errors"
)

type CatService struct {
	repo   repository.CatRepository
	catAPI *client.CatAPI
}

func NewCatService(repo repository.CatRepository, catAPI *client.CatAPI) *CatService {
	return &CatService{
		repo:   repo,
		catAPI: catAPI,
	}
}

func (s *CatService) Create(ctx context.Context, catCreate model.CatCreate) (*model.Cat, error) {

	// Validate breed using CatAPI
	valid, err := s.catAPI.ValidateBreed(ctx, catCreate.Breed)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("invalid cat breed")
	}

	cat := &model.Cat{
		Name:            catCreate.Name,
		YearsExperience: catCreate.YearsExperience,
		Breed:           catCreate.Breed,
		Salary:          catCreate.Salary,
	}

	if err = s.repo.Create(ctx, cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func (s *CatService) Update(ctx context.Context, id uint, update model.CatUpdate) (*model.Cat, error) {

	cat, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	cat.Salary = update.Salary

	if err = s.repo.Update(ctx, cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func (s *CatService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *CatService) GetByID(ctx context.Context, id uint) (*model.Cat, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CatService) List(ctx context.Context) ([]model.Cat, error) {
	return s.repo.List(ctx)
}
