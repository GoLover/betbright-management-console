package cmd

import (
	eventCmd "betbright-management-console/apps/event/delivery/cmd"
	eventUseCase "betbright-management-console/apps/event/usecase"
	marketCmd "betbright-management-console/apps/market/delivery/cmd"
	marketUseCase "betbright-management-console/apps/market/usecase"
	searchCmd "betbright-management-console/apps/search/delivery/cmd"
	"betbright-management-console/apps/search/delivery/wal"
	"betbright-management-console/apps/search/repository/elasticsearch"
	searchUseCase "betbright-management-console/apps/search/usecase"
	selectionCmd "betbright-management-console/apps/selection/delivery/cmd"
	selectionUseCase "betbright-management-console/apps/selection/usecase"
	sportCmd "betbright-management-console/apps/sport/delivery/cmd"
	"betbright-management-console/apps/sport/repository/psql"
	sportUseCase "betbright-management-console/apps/sport/usecase"
	"betbright-management-console/domain"
	"betbright-management-console/infra/elastic"
	"encoding/json"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type ElasticSearchConfig struct {
	HTTPS    bool
	Addr     string
	User     string
	Password string
}
type Config struct {
	PostgresStrConn string
	ElasticSearch   ElasticSearchConfig
}

func parseConfig() Config {
	configData, err := os.ReadFile(`config.json`)
	if err != nil {
		panic(err)
	}
	c := Config{}
	err = json.Unmarshal(configData, &c)
	if err != nil {
		panic(err)
	}
	return c
}
func Boot() {
	config := parseConfig()
	db, err := gorm.Open(postgres.Open(config.PostgresStrConn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	cmd := &cobra.Command{Use: "betbright",
		Short: "betbright interactive command line"}
	repo := psql.New(db)
	repo.Migrate()
	var _ domain.Operator = sportCmd.SportOperator{}
	var _ domain.Operator = marketCmd.MarketOperator{}
	var _ domain.Operator = eventCmd.EventOperator{}
	var _ domain.Operator = selectionCmd.SelectionOperator{}

	selectionUsecase := selectionUseCase.New(repo)
	eventUsecase := eventUseCase.New(repo, []domain.Observee{})
	sportUsecase := sportUseCase.New(repo, []domain.Observee{eventUsecase})
	marketUsecase := marketUseCase.New(repo, []domain.Observee{selectionUsecase})
	eventUsecase.BindObserveLately([]domain.Observee{marketUsecase})

	ec := elastic.NewClient(&elastic.ClientConfig{
		HTTPS:    config.ElasticSearch.HTTPS,
		Addr:     config.ElasticSearch.Addr,
		User:     config.ElasticSearch.User,
		Password: config.ElasticSearch.Password,
	})
	pgxConfig, err := pgx.ParseConnectionString(config.PostgresStrConn)
	if err != nil {
		panic(err)
	}
	pgxConn, err := pgx.Connect(pgxConfig)
	if err != nil {
		panic(err)
	}
	pgxReplicationConn, err := pgx.ReplicationConnect(pgxConfig)
	if err != nil {
		panic(err)
	}
	searchUsecase := searchUseCase.New(elasticsearch.New(ec))
	logger := logrus.New()
	loggerEntry := logrus.NewEntry(logger)
	searchHandler := wal.New(searchUsecase)
	go searchHandler.BindListener(loggerEntry, pgxConn, pgxReplicationConn)
	sportHandler := sportCmd.New(sportUsecase, searchUsecase, cmd)
	eventHandler := eventCmd.New(eventUsecase, searchUsecase, cmd)
	marketHandler := marketCmd.New(marketUsecase, searchUsecase, cmd)
	selectionHandler := selectionCmd.New(selectionUsecase, searchUsecase, cmd)
	searchCmd := searchCmd.New(searchUsecase, cmd)
	sportHandler.Handle()
	eventHandler.Handle()
	marketHandler.Handle()
	selectionHandler.Handle()
	searchCmd.Handle()
	for {
		_ = cmd.Execute()
	}
}
