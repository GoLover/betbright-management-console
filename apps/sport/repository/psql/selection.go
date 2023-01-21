package psql

import "betbright-management-console/domain"

func (s *SportRepository) CreateSelection(selection domain.Selection, marketId, eventId int) (domain.Selection, error) {
	dao := &Selection{}
	dao.FillFromDomain(selection)
	dao.MarketID = marketId
	dao.EventID = eventId
	err := errorTranslator(s.db.Create(dao).Error)
	if err != nil {
		return domain.Selection{}, err
	}
	return dao.ToDomain(), nil
}

func (s *SportRepository) DeactivateSelection(id int) error {
	//TODO implement me
	panic("implement me")
}
