package psql

import "betbright-management-console/domain"

func (s *SportRepository) CreateMarket(market domain.Market, eventSlug string) (domain.Market, error) {
	dao := &Market{}
	dao.FillFromDomain(market)
	event, err := s.getEventBySlug(eventSlug, false)
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

func (s *SportRepository) DeactivateMarket(id int) error {
	//TODO implement me
	panic("implement me")
}
