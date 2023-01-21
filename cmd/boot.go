package cmd

import (
	eventCmd "betbright-management-console/apps/event/delivery/cmd"
	eventUseCase "betbright-management-console/apps/event/usecase"
	marketCmd "betbright-management-console/apps/market/delivery/cmd"
	marketUseCase "betbright-management-console/apps/market/usecase"
	selectionCmd "betbright-management-console/apps/selection/delivery/cmd"
	selectionUseCase "betbright-management-console/apps/selection/usecase"
	sportCmd "betbright-management-console/apps/sport/delivery/cmd"
	"betbright-management-console/apps/sport/repository/psql"
	sportUseCase "betbright-management-console/apps/sport/usecase"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Boot() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	cmd := &cobra.Command{Use: "betbright",
		Short: "betbright interactive command line"}
	repo := psql.New(db)
	repo.Migrate()
	sportUsecase := sportUseCase.New(repo)
	selectionUsecase := selectionUseCase.New(repo)
	marketUsecase := marketUseCase.New(repo)
	eventUsecase := eventUseCase.New(repo)
	sportHandler := sportCmd.New(sportUsecase, cmd)
	eventHandler := eventCmd.New(eventUsecase, cmd)
	marketHandler := marketCmd.New(marketUsecase, cmd)
	selectionHandler := selectionCmd.New(selectionUsecase, cmd)
	sportHandler.Handle()
	eventHandler.Handle()
	marketHandler.Handle()
	selectionHandler.Handle()
	cmd.Execute()
}