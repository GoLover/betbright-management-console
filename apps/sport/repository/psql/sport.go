package psql

import (
	"betbright-management-console/domain"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type SportRepository struct {
	db *gorm.DB
}

func New(database *gorm.DB) *SportRepository {
	return &SportRepository{db: database}
}
func (s *SportRepository) Migrate() {
	s.db.AutoMigrate(Sport{})
	s.db.AutoMigrate(Event{})
	s.db.AutoMigrate(Market{})
	s.db.AutoMigrate(Selection{})
}
func (s *SportRepository) CreateSport(sport domain.Sport) (domain.Sport, error) {
	dao := &Sport{}
	dao.FillFromDomain(sport)
	err := errorTranslator(s.db.Create(dao).Error)
	if err != nil {
		return domain.Sport{}, err
	}
	return dao.ToDomain(), nil
}

func (s *SportRepository) getSportBySlug(slug string, complete bool) (domain.Sport, error) {
	dao := &Sport{}
	var err error
	if complete {
		err = errorTranslator(s.db.Where(&Sport{Slug: slug}).Preload(`Events`).Find(dao).Error)
	} else {
		err = errorTranslator(s.db.Where(&Sport{Slug: slug}).First(dao).Error)
	}
	if err != nil {
		return domain.Sport{}, err

	}
	return dao.ToDomain(), nil
}

func (s *SportRepository) UpdateSport(sport domain.Sport, sportSlug string) (domain.Sport, error) {
	dao := &Sport{}
	dao.FillFromDomain(sport)
	updateResult := s.db.Where(&Sport{Slug: sportSlug}).Updates(dao)
	if updateResult.RowsAffected == 0 {
		return domain.Sport{}, fmt.Errorf(`sport %w`, domain.ErrRepoRecordNotFound)
	}
	err := errorTranslator(updateResult.Error)
	if err != nil {
		return domain.Sport{}, err
	}
	return dao.ToDomain(), nil
}

func (s *SportRepository) DeactivateSport(slug string) error {
	//TODO implement me
	panic("implement me")
}

func errorTranslator(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), `UNIQUE constraint failed`) {
		return domain.ErrRepoRecordAlreadyExist
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrRepoRecordNotFound
	}
	return err
}
