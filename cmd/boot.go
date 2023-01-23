package cmd

import (
	eventCmd "betbright-management-console/apps/event/delivery/cmd"
	eventUseCase "betbright-management-console/apps/event/usecase"
	marketCmd "betbright-management-console/apps/market/delivery/cmd"
	marketUseCase "betbright-management-console/apps/market/usecase"
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
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Boot() {
	db, err := gorm.Open(postgres.Open("host=localhost user=betbright password=password dbname=betbright port=5432 sslmode=disable TimeZone=Asia/Tehran"), &gorm.Config{})
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
		HTTPS:    true,
		Addr:     "localhost:9200",
		User:     "elastic",
		Password: "Dic*f23qSC8ayxU1N5jW",
	})
	pgxConfig, err := pgx.ParseConnectionString("host=localhost user=betbright password=password dbname=betbright port=5432 sslmode=disable TimeZone=Asia/Tehran")
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
	sportHandler := sportCmd.New(sportUsecase, cmd)
	eventHandler := eventCmd.New(eventUsecase, cmd)
	marketHandler := marketCmd.New(marketUsecase, cmd)
	selectionHandler := selectionCmd.New(selectionUsecase, cmd)
	sportHandler.Handle()
	eventHandler.Handle()
	marketHandler.Handle()
	selectionHandler.Handle()
	for {
		_ = cmd.Execute()
	}
}
