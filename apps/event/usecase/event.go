package usecase

import (
	"betbright-management-console/domain"
	"context"
	"errors"
	"fmt"
	"github.com/gosimple/slug"
)

type EventUseCase struct {
	observers         []domain.Observer
	subjectsToObserve []domain.Observee
	r                 domain.SportRepository
}

func (s *EventUseCase) Update(ctx context.Context) {
	eventId := ctx.Value(`eventId`).(int)
	_, err := s.r.GetMarketsByEventId(eventId)
	if errors.Is(err, domain.ErrRepoRecordNotFound) {
		event, err := s.r.GetEventById(eventId)
		if err != nil {
			fmt.Println(fmt.Errorf(`GetEventById-UpdateSignal %w`, err))
			return
		}
		err = s.DeactivateEvent(ctx, event.Slug)
		if err != nil {
			fmt.Println(fmt.Errorf(`DeactivateEvent-UpdateSignal %w`, err))
			return
		}
		return
	}
	if err != nil {
		fmt.Println(fmt.Errorf(`Event-UpdateSignal %w`, err))
	}

}

func (s *EventUseCase) Register(observer domain.Observer) {
	if s.observers == nil {
		s.observers = make([]domain.Observer, 0)
	}
	s.observers = append(s.observers, observer)
}

func (s *EventUseCase) Notify(ctx context.Context) {
	for _, k := range s.observers {
		k.Update(ctx)
	}
}

func (s *EventUseCase) CreateEvent(ctx context.Context, event domain.Event, sportSlug string) (domain.Event, error) {
	event.Slug = slug.Make(event.Name)
	event.IsActive = true
	_, err := s.r.CreateEvent(event, sportSlug)
	if err != nil {
		if errors.Is(err, domain.ErrRepoRecordNotFound) {
			return domain.Event{}, domain.ErrUseCaseEnteredSportSlugNotFound
		}
	}
	return domain.Event{}, err
}
func (s *EventUseCase) UpdateEvent(ctx context.Context, event domain.Event, currentEventSlug, newSportSlug string) (domain.Event, error) {
	event.Slug = slug.Make(event.Name)
	return s.r.UpdateEvent(event, currentEventSlug, newSportSlug)
}

func (s *EventUseCase) DeactivateEvent(ctx context.Context, slug string) error {
	event, err := s.r.ChangeActivationEvent(slug, false)
	if err != nil {
		return err
	}
	s.Notify(context.WithValue(ctx, `sportId`, event.SportId))
	fmt.Println(`event deactivated successfully`)
	return err
}

func (s *EventUseCase) ActivateEvent(ctx context.Context, slug string) error {
	_, err := s.r.ChangeActivationEvent(slug, true)
	return err
}

func (s *EventUseCase) DeleteEvent(ctx context.Context, slug string) error {
	return s.r.DeleteEvent(slug)
}
func (s *EventUseCase) BindObserveLately(subjectsToObserve []domain.Observee) {
	for _, k := range subjectsToObserve {
		k.Register(s)
	}
}
func New(r domain.SportRepository, subjectsToObserve []domain.Observee) *EventUseCase {
	eu := &EventUseCase{r: r}
	for _, k := range subjectsToObserve {
		k.Register(eu)
	}
	return &EventUseCase{r: r}
}
