package cmd

import (
	"betbright-management-console/apps/event/adapter"
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type EventCommandLineHandler struct {
	eu   domain.EventUseCase
	sa   adapter.SearchAdapter
	root *cobra.Command
}

func New(u domain.EventUseCase, sa adapter.SearchAdapter, cmd *cobra.Command) EventCommandLineHandler {
	return EventCommandLineHandler{eu: u, sa: sa, root: cmd}
}

func (h *EventCommandLineHandler) Handle() {
	var cmd = &cobra.Command{
		Use:   "event",
		Short: "event related stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, EventOperator{sa: h.sa, eu: h.eu}))
		},
	}
	h.root.AddCommand(cmd)
}
