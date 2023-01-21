package usecase

import (
	"betbright-management-console/domain"
	"context"
	"github.com/gosimple/slug"
)

type SportUseCase struct {
	r domain.SportRepository
}

func (s SportUseCase) CreateSport(ctx context.Context, sport domain.Sport) (domain.Sport, error) {
	sport.Slug = slug.Make(sport.Name)
	sport.IsActive = true
	return s.r.CreateSport(sport)
}
func (s SportUseCase) UpdateSport(ctx context.Context, sport domain.Sport, sportSlug string) (domain.Sport, error) {
	sport.Slug = slug.Make(sport.Name)
	return s.r.UpdateSport(sport, sportSlug)
}

func (s SportUseCase) DeactivateSport(ctx context.Context, slug string) error {
	return s.r.ChangeActivationSport(slug, false)
}
func (s SportUseCase) ActivateSport(ctx context.Context, slug string) error {
	return s.r.ChangeActivationSport(slug, true)
}

func New(r domain.SportRepository) *SportUseCase {
	return &SportUseCase{r: r}
}
