package psql

import (
	"betbright-management-console/domain"
	"fmt"
	"gorm.io/gorm/clause"
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
func (s *SportRepository) ChangeActivationSelection(selectionId int, active bool) (domain.Selection, error) {
	selection := &Selection{}
	updateResult := s.db.Model(selection).Clauses(clause.Returning{}).Where(&Selection{Id: selectionId}).Updates(map[string]interface{}{"is_active": active})
	if updateResult.RowsAffected == 0 {
		return domain.Selection{}, fmt.Errorf(`selection %w`, domain.ErrRepoRecordNotFound)
	}
	return selection.ToDomain(), errorTranslator(updateResult.Error)
}

func (s *SportRepository) GetSelectionByMarketId(marketId int) ([]domain.Selection, error) {
	var selection []Selection
	err := errorTranslator(s.db.Model(&Selection{}).Where(map[string]interface{}{"market_id": marketId}).Find(&selection).Error)
	if err != nil {
		return nil, err
	}
	if len(selection) == 0 {
		return nil, domain.ErrRepoRecordNotFound
	}
	result := make([]domain.Selection, 0)
	for _, k := range selection {
		result = append(result, k.ToDomain())
	}
	return result, nil
}

func (s *SportRepository) GetSelectionById(id int) (domain.Selection, error) {
	var selection Selection
	err := errorTranslator(s.db.Model(&Selection{}).Where(Selection{Id: id}).First(&selection).Error)
	if err != nil {
		return domain.Selection{}, err
	}
	return selection.ToDomain(), nil
}
