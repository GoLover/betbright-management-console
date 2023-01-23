package cmd

import (
	"betbright-management-console/apps/market/adapter"
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type EventCommandLineHandler struct {
	mu   domain.MarketUseCase
	sa   adapter.SearchAdapter
	root *cobra.Command
}

func New(u domain.MarketUseCase, sa adapter.SearchAdapter, cmd *cobra.Command) EventCommandLineHandler {
	return EventCommandLineHandler{mu: u, sa: sa, root: cmd}
}

func (h *EventCommandLineHandler) Handle() {
	var cmd = &cobra.Command{
		Use:   "market",
		Short: "market related stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, MarketOperator{sa: h.sa, mu: h.mu}))
		},
	}
	h.root.AddCommand(cmd)
}
