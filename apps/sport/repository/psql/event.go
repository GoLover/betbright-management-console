package psql

import (
	"betbright-management-console/domain"
	"fmt"
)

func (s *SportRepository) getEventBySlug(slug string, complete bool) (domain.Event, error) {
	dao := &Event{}
	var err error
	if complete {
		err = errorTranslator(s.db.Where(&Event{Slug: slug}).Preload(`Markets`).Find(dao).Error)
	} else {
		err = errorTranslator(s.db.Where(&Event{Slug: slug}).First(dao).Error)
	}

	if err != nil {
		return domain.Event{}, err

	}
	return dao.ToDomain(), nil
}
func (s *SportRepository) CreateEvent(event domain.Event, sportSlug string) (domain.Event, error) {
	dao := &Event{}
	dao.FillFromDomain(event)
	sport, err := s.getSportBySlug(sportSlug, false)
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
	sport, err := s.getSportBySlug(sportSlug, false)
	if err != nil {
		return domain.Event{}, fmt.Errorf(`sport %w`, err)
	}
	dao.SportID = sport.Id

	updateResult := s.db.Model(&Event{Slug: eventSlug}).Updates(dao)
	if updateResult.RowsAffected == 0 {
		return domain.Event{}, fmt.Errorf(`%#v`, domain.ErrRepoRecordNotFound)
	}
	err = errorTranslator(updateResult.Error)
	if err != nil {
		return domain.Event{}, fmt.Errorf(`event %w`, err)
	}
	updatedEvent := dao.ToDomain()
	return updatedEvent, nil
}
func (s *SportRepository) DeactivateEvent(slug string) error {
	//TODO implement me
	panic("implement me")
}
