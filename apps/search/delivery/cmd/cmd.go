package cmd

import (
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type SearchCommandLineHandler struct {
	su   domain.SearchUsecase
	root *cobra.Command
}

func New(u domain.SearchUsecase, cmd *cobra.Command) SearchCommandLineHandler {
	return SearchCommandLineHandler{su: u, root: cmd}
}

func (h *SearchCommandLineHandler) Handle() {
	var cmd = &cobra.Command{
		Use:   "search",
		Short: "search in all stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, SearchCommandLineHandler{su: h.su}))
		},
	}
	h.root.AddCommand(cmd)
}
