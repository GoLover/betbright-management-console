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

func (s *SportRepository) getSportWithFilter(filter Sport, complete bool) (domain.Sport, error) {
	dao := &Sport{}
	var err error
	if complete {
		err = errorTranslator(s.db.Where(&filter).Preload(`Events`).Find(dao).Error)
	} else {
		err = errorTranslator(s.db.Where(&filter).First(dao).Error)
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
func (s *SportRepository) ChangeActivationSport(sportSlug string, active bool) error {
	fmt.Println(s.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&Sport{Slug: sportSlug}).Where(&Sport{Slug: sportSlug}).Updates(map[string]interface{}{"is_active": active})
	}))
	updateResult := s.db.Model(&Sport{Slug: sportSlug}).Where(&Sport{Slug: sportSlug}).Updates(map[string]interface{}{"is_active": active})
	if updateResult.RowsAffected == 0 {
		return fmt.Errorf(`sport %w`, domain.ErrRepoRecordNotFound)
	}
	return errorTranslator(updateResult.Error)
}

func (s *SportRepository) GetSportById(id int) (domain.Sport, error) {
	var sport Sport
	err := errorTranslator(s.db.Model(&Sport{}).Where(Sport{Id: id}).First(&sport).Error)
	if err != nil {
		return domain.Sport{}, err
	}
	return sport.ToDomain(), nil
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
