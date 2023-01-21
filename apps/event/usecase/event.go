package usecase

import (
	"betbright-management-console/domain"
	"context"
	"errors"
	"github.com/gosimple/slug"
)

type EventUseCase struct {
	observers         []domain.Observer
	subjectsToObserve []domain.Observee
	r                 domain.SportRepository
}

func (s *EventUseCase) Update() {
	//TODO implement me
	panic("implement me")
}

func (s *EventUseCase) Register(observer domain.Observer) {
	if s.observers == nil {
		s.observers = make([]domain.Observer, 0)
	}
	s.observers = append(s.observers, observer)
}

func (s *EventUseCase) Notify() {
	for _, k := range s.observers {
		k.Update()
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
	return s.r.ChangeActivationEvent(slug, false)
}

func (s *EventUseCase) SyncActiveStatus(ctx context.Context) {

}

func (s *EventUseCase) ActivateEvent(ctx context.Context, slug string) error {
	return s.r.ChangeActivationEvent(slug, true)
}

func (s *EventUseCase) DeleteEvent(ctx context.Context, slug string) error {
	//TODO implement me
	panic("implement me")
}

func New(r domain.SportRepository, subjectsToObserve []domain.Observee) *EventUseCase {
	eu := &EventUseCase{r: r}
	for _, k := range subjectsToObserve {
		k.Register(eu)
	}
	return &EventUseCase{r: r}
}
