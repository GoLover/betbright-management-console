package cmd

import (
	"betbright-management-console/apps/selection/adapter"
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type EventCommandLineHandler struct {
	su   domain.SelectionUseCase
	sa   adapter.SearchAdapter
	root *cobra.Command
}

func New(u domain.SelectionUseCase, sa adapter.SearchAdapter, cmd *cobra.Command) EventCommandLineHandler {
	return EventCommandLineHandler{su: u, sa: sa, root: cmd}
}

func (h *EventCommandLineHandler) Handle() {
	var cmd = &cobra.Command{
		Use:   "selection",
		Short: "selection related stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, SelectionOperator{sa: h.sa, su: h.su}))
		},
	}
	h.root.AddCommand(cmd)
}
