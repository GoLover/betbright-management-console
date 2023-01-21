package psql

import (
	"betbright-management-console/domain"
	"gorm.io/gorm/clause"
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

func (s *SportRepository) GetMarketsByEventId(eventId int) ([]domain.Market, error) {
	markets := make([]Market, 0)
	err := errorTranslator(s.db.Model(&Market{}).Where(&Market{EventID: eventId, IsActive: true}).Find(&markets).Error)
	if len(markets) == 0 {
		return []domain.Market{}, domain.ErrRepoRecordNotFound
	}
	if err != nil {
		return []domain.Market{}, err
	}
	result := make([]domain.Market, 0)
	for _, k := range markets {
		result = append(result, k.ToDomain())
	}
	return result, nil
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
		return domain.Market{}, domain.ErrRepoRecordNotFound
	}
	err := errorTranslator(updateResult.Error)
	if err != nil {
		return domain.Market{}, err
	}
	return dao.ToDomain(), nil
}
func (s *SportRepository) ChangeActivationMarket(marketId int, active bool) (domain.Market, error) {
	market := &Market{}
	updateResult := s.db.Model(&market).Clauses(clause.Returning{}).Where(map[string]interface{}{"id": marketId}).Updates(map[string]interface{}{"is_active": active})
	if updateResult.RowsAffected == 0 {
		return domain.Market{}, domain.ErrRepoRecordNotFound
	}
	return market.ToDomain(), errorTranslator(updateResult.Error)
}

func (s *SportRepository) GetMarketById(id int) (domain.Market, error) {
	var market Market
	err := errorTranslator(s.db.Model(&Market{}).Where(Market{Id: id}).First(&market).Error)
	if err != nil {
		return domain.Market{}, err
	}
	return market.ToDomain(), nil
}
