package cmd

import (
	"betbright-management-console/apps/sport/adapter"
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type SportCommandLineHandler struct {
	u    domain.SportUseCase
	sa   adapter.SearchAdapter
	root *cobra.Command
}

func New(u domain.SportUseCase, sa adapter.SearchAdapter, cmd *cobra.Command) SportCommandLineHandler {
	return SportCommandLineHandler{u: u, sa: sa, root: cmd}
}

func (h *SportCommandLineHandler) Handle() {
	var sportCmd = &cobra.Command{
		Use:   "sport",
		Short: "sport related stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, SportOperator{u: h.u, sa: h.sa}))
		},
	}
	h.root.AddCommand(sportCmd)
}
