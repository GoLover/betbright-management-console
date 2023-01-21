package cmd

import (
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type EventCommandLineHandler struct {
	su   domain.SelectionUseCase
	root *cobra.Command
}

func New(u domain.SelectionUseCase, cmd *cobra.Command) EventCommandLineHandler {
	return EventCommandLineHandler{su: u, root: cmd}
}

func (h *EventCommandLineHandler) Handle() {
	var cmd = &cobra.Command{
		Use:   "selection",
		Short: "selection related stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, SelectionOperator{h.su}))
		},
	}
	h.root.AddCommand(cmd)
}
