package cmd

import (
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type SportCommandLineHandler struct {
	u    domain.SportUseCase
	root *cobra.Command
}

func New(u domain.SportUseCase, cmd *cobra.Command) SportCommandLineHandler {
	return SportCommandLineHandler{u: u, root: cmd}
}

func (h *SportCommandLineHandler) Handle() {
	var sportCmd = &cobra.Command{
		Use:   "sport",
		Short: "sport related stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, SportOperator{u: h.u}))
		},
	}
	h.root.AddCommand(sportCmd)
}
