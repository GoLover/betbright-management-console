package psql

import (
	"betbright-management-console/domain"
	"fmt"
	"gorm.io/gorm/clause"
)

func (s *SportRepository) getEventWithFilter(event Event, complete bool) (domain.Event, error) {
	dao := &Event{}
	var err error
	if complete {
		err = errorTranslator(s.db.Where(&event).Preload(`Markets`).Find(dao).Error)
	} else {
		err = errorTranslator(s.db.Where(&event).First(dao).Error)
	}

	if err != nil {
		return domain.Event{}, err

	}
	return dao.ToDomain(), nil
}
func (s *SportRepository) CreateEvent(event domain.Event, sportSlug string) (domain.Event, error) {
	dao := &Event{}
	dao.FillFromDomain(event)
	sport, err := s.getSportWithFilter(Sport{Slug: sportSlug, IsActive: true}, false)
	if err != nil {
		return domain.Event{}, err
	}
	dao.SportID = sport.Id
	err = errorTranslator(s.db.Create(dao).Error)
	if err != nil {
		return domain.Event{}, err
	}
	return dao.ToDomain(), nil
}
func (s *SportRepository) UpdateEvent(event domain.Event, eventSlug, sportSlug string) (domain.Event, error) {
	dao := &Event{}
	dao.FillFromDomain(event)
	if sportSlug != `` {
		sport, err := s.getSportWithFilter(Sport{Slug: sportSlug, IsActive: true}, false)
		if err != nil {
			return domain.Event{}, fmt.Errorf(`sport %w`, err)
		}
		dao.SportID = sport.Id
	}
	updateResult := s.db.Model(&Event{Slug: eventSlug}).Updates(dao)
	if updateResult.RowsAffected == 0 {
		return domain.Event{}, fmt.Errorf(`%#v`, domain.ErrRepoRecordNotFound)
	}
	err := errorTranslator(updateResult.Error)
	if err != nil {
		return domain.Event{}, fmt.Errorf(`event %w`, err)
	}
	updatedEvent := dao.ToDomain()
	return updatedEvent, nil
}
func (s *SportRepository) ChangeActivationEvent(eventSlug string, active bool) (domain.Event, error) {
	event := &Event{}
	updateResult := s.db.Model(event).Clauses(clause.Returning{}).Where(&Event{Slug: eventSlug}).Updates(map[string]interface{}{"is_active": active})
	if updateResult.RowsAffected == 0 {
		return domain.Event{}, fmt.Errorf(`event %w`, domain.ErrRepoRecordNotFound)
	}
	return event.ToDomain(), errorTranslator(updateResult.Error)
}

func (s *SportRepository) GetEventsBySportId(sportId int) ([]domain.Event, error) {
	var events []Event
	err := errorTranslator(s.db.Model(&Event{}).Where(map[string]interface{}{"sport_id": sportId}).Find(&events).Error)
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, domain.ErrRepoRecordNotFound
	}
	result := make([]domain.Event, 0)
	for _, k := range events {
		result = append(result, k.ToDomain())
	}
	return result, nil
}

func (s *SportRepository) GetEventById(id int) (domain.Event, error) {
	var event Event
	err := errorTranslator(s.db.Model(&Event{}).Where(Event{Id: id}).First(&event).Error)
	if err != nil {
		return domain.Event{}, err
	}
	return event.ToDomain(), nil
}
