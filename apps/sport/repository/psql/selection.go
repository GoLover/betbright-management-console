package psql

import (
	"betbright-management-console/domain"
	"fmt"
)

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
func (s *SportRepository) UpdateSelection(selection domain.Selection, selectionId, marketId, eventId int) (domain.Selection, error) {
	dao := &Selection{}
	dao.FillFromDomain(selection)
	dao.MarketID = marketId
	dao.EventID = eventId
	updateResult := s.db.Where(&Selection{Id: selectionId}).Updates(dao)
	if updateResult.RowsAffected == 0 {
		return domain.Selection{}, fmt.Errorf(`selection %w`, domain.ErrRepoRecordNotFound)
	}
	err := errorTranslator(updateResult.Error)
	if err != nil {
		return domain.Selection{}, err
	}
	return dao.ToDomain(), nil
}
func (s *SportRepository) ChangeActivationSelection(selectionId int, active bool) error {
	updateResult := s.db.Model(&Selection{}).Where(&Selection{Id: selectionId}).Updates(map[string]interface{}{"is_active": active})
	if updateResult.RowsAffected == 0 {
		return fmt.Errorf(`selection %w`, domain.ErrRepoRecordNotFound)
	}
	return errorTranslator(updateResult.Error)
}
