package psql

import (
	"betbright-management-console/domain"
	"github.com/shopspring/decimal"
)

type Sport struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	DisplayName string
	Slug        string `gorm:"uniqueIndex"`
	OrderInList int    `gorm:"uniqueIndex"`
	IsActive    bool
	Events      []Event
}

func (s *Sport) FillFromDomain(sport domain.Sport) {
	s.Id = sport.Id
	s.Name = sport.Name
	s.OrderInList = sport.Order
	s.Slug = sport.Slug
	s.IsActive = sport.IsActive
	s.DisplayName = sport.DisplayName
}
func (s *Sport) ToDomain() domain.Sport {
	sport := domain.Sport{
		Id:          s.Id,
		Name:        s.Name,
		DisplayName: s.DisplayName,
		Slug:        s.Slug,
		Order:       s.OrderInList,
		IsActive:    s.IsActive,
		Events:      nil,
	}
	return sport
}

type Event struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	EType    int
	Status   int
	Slug     string `gorm:"uniqueIndex"`
	IsActive bool
	SportID  int
	Markets  []Market
}

func (e *Event) FillFromDomain(event domain.Event) {
	e.Id = event.Id
	e.Name = event.Name
	e.EType = int(event.EType)
	e.Status = int(event.Status)
	e.Slug = event.Slug
	e.IsActive = event.IsActive
	e.Markets = make([]Market, 0)
	for _, k := range event.Markets {
		e.Markets = append(e.Markets, Market{
			Id:          k.Id,
			Name:        k.Name,
			DisplayName: k.DisplayName,
			OrderInList: k.Order,
			IsActive:    k.IsActive,
			Schema:      k.Schema,
			Columns:     k.Columns,
			EventID:     e.Id,
		})
	}
}

func (e *Event) ToDomain() domain.Event {
	event := domain.Event{
		Id:       e.Id,
		Name:     e.Name,
		EType:    domain.EventType(e.EType),
		Status:   domain.EventStatus(e.Status),
		Slug:     e.Slug,
		IsActive: e.IsActive,
	}
	event.Markets = make([]domain.Market, 0)
	for _, k := range e.Markets {
		event.Markets = append(event.Markets, domain.Market{
			Id:          k.Id,
			Name:        k.Name,
			DisplayName: k.DisplayName,
			Order:       k.OrderInList,
			IsActive:    k.IsActive,
			Schema:      k.Schema,
			Columns:     k.Columns,
		})
	}
	return event
}

type Market struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	DisplayName string
	OrderInList int `gorm:"uniqueIndex"`
	IsActive    bool
	Schema      int
	Columns     int
	EventID     int
}

func (e *Market) FillFromDomain(market domain.Market) {
	e.Id = market.Id
	e.Name = market.Name
	e.DisplayName = market.DisplayName
	e.OrderInList = market.Order
	e.IsActive = market.IsActive
	e.Schema = market.Schema
	e.Columns = market.Columns
	e.EventID = market.EventId
}

func (e *Market) ToDomain() domain.Market {
	market := domain.Market{
		Id:          e.Id,
		Name:        e.Name,
		DisplayName: e.DisplayName,
		Order:       e.OrderInList,
		IsActive:    e.IsActive,
		Schema:      e.Schema,
		Columns:     e.Columns,
		EventId:     e.EventID,
	}
	return market
}

type Selection struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	EventID  int
	Event    Event
	MarketID int
	Market   Market
	Name     string
	Price    decimal.Decimal `gorm:"type:decimal(10,2)"`
	IsActive bool
	Outcome  int
}

func (s *Selection) FillFromDomain(selection domain.Selection) {
	s.EventID = selection.SelectedEvent.Id
	s.MarketID = selection.SelectedMarket.Id
	s.Id = selection.Id
	s.Outcome = int(selection.Outcome)
	s.Price = selection.Price
	s.Name = selection.Name
	s.IsActive = selection.IsActive
}

func (s *Selection) ToDomain() domain.Selection {
	selection := domain.Selection{
		Id:             s.Id,
		Name:           s.Name,
		SelectedEvent:  s.Event.ToDomain(),
		SelectedMarket: s.Market.ToDomain(),
		Price:          s.Price,
		IsActive:       s.IsActive,
		Outcome:        domain.OutcomeState(s.Outcome),
	}
	return selection
}
