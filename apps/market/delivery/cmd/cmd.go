package cmd

import (
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type EventCommandLineHandler struct {
	mu   domain.MarketUseCase
	root *cobra.Command
}

func New(u domain.MarketUseCase, cmd *cobra.Command) EventCommandLineHandler {
	return EventCommandLineHandler{mu: u, root: cmd}
}

func (h *EventCommandLineHandler) Handle() {
	var cmd = &cobra.Command{
		Use:   "market",
		Short: "market related stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, MarketOperator{h.mu}))
		},
	}
	h.root.AddCommand(cmd)
}
