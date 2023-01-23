package usecase

import (
	"betbright-management-console/apps/sport/adapter"
	"betbright-management-console/domain"
	"context"
	"errors"
	"fmt"
	"github.com/gosimple/slug"
)

type SportUseCase struct {
	sa adapter.SearchAdapter
	r  domain.SportRepository
}

func (s *SportUseCase) Update(ctx context.Context) {
	sportId := ctx.Value(`sportId`).(int)
	_, err := s.r.GetEventsBySportId(sportId)
	if errors.Is(err, domain.ErrRepoRecordNotFound) {
		sport, err := s.r.GetSportById(sportId)
		if err != nil {
			fmt.Println(fmt.Errorf(`Update-GetSportId %w`, err))
			return
		}
		err = s.DeactivateSport(ctx, sport.Slug)
		if err != nil {
			fmt.Println(fmt.Errorf(`Update-DeactivateSport %w`, err))
			return
		}
		return
	}
	if err != nil {
		fmt.Println(fmt.Errorf(`Update-GetEventsBySportId %w`, err))
	}
}

func (s *SportUseCase) CreateSport(ctx context.Context, sport domain.Sport) (domain.Sport, error) {
	sport.Slug = slug.Make(sport.Name)
	sport.IsActive = true
	return s.r.CreateSport(sport)
}
func (s *SportUseCase) UpdateSport(ctx context.Context, sport domain.Sport, sportSlug string) (domain.Sport, error) {
	sport.Slug = slug.Make(sport.Name)
	return s.r.UpdateSport(sport, sportSlug)
}

func (s *SportUseCase) DeactivateSport(ctx context.Context, slug string) error {
	err := s.r.ChangeActivationSport(slug, false)
	if err != nil {
		return err
	}
	fmt.Println(`sport deactivated successfully`)
	return err
}
func (s *SportUseCase) ActivateSport(ctx context.Context, slug string) error {
	return s.r.ChangeActivationSport(slug, true)
}

func New(r domain.SportRepository, subjectsToObserve []domain.Observee) *SportUseCase {
	su := &SportUseCase{r: r}
	for _, k := range subjectsToObserve {
		k.Register(su)
	}
	return su
}
