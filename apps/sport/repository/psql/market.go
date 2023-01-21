package psql

import (
	"betbright-management-console/domain"
	"fmt"
)

func (s *SportRepository) CreateMarket(market domain.Market, eventSlug string) (domain.Market, error) {
	dao := &Market{}
	dao.FillFromDomain(market)
	event, err := s.getEventWithFilter(Event{Slug: eventSlug, IsActive: true}, false)
	if err != nil {
		return domain.Market{}, err
	}
	dao.EventID = event.Id
	err = errorTranslator(s.db.Create(dao).Error)
	if err != nil {
		return domain.Market{}, err
	}
	return dao.ToDomain(), nil
}

func (s *SportRepository) UpdateMarket(market domain.Market, marketId int, eventSlug string) (domain.Market, error) {
	dao := &Market{}
	dao.FillFromDomain(market)
	if eventSlug != `` {
		event, err := s.getEventWithFilter(Event{Slug: eventSlug, IsActive: true}, false)
		if err != nil {
			return domain.Market{}, err
		}
		dao.EventID = event.Id
	}
	updateResult := s.db.Where(&Market{Id: marketId}).Updates(dao)
	if updateResult.RowsAffected == 0 {
		return domain.Market{}, fmt.Errorf(`market %#v`, domain.ErrRepoRecordNotFound)
	}
	err := errorTranslator(updateResult.Error)
	if err != nil {
		return domain.Market{}, err
	}
	return dao.ToDomain(), nil
}
func (s *SportRepository) ChangeActivationMarket(marketId int, active bool) error {
	updateResult := s.db.Model(&Market{}).Where(&Market{Id: marketId}).Updates(map[string]interface{}{"is_active": active})
	if updateResult.RowsAffected == 0 {
		return fmt.Errorf(`market %w`, domain.ErrRepoRecordNotFound)
	}
	return errorTranslator(updateResult.Error)
}
