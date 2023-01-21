package cmd

import (
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"github.com/spf13/cobra"
)

type EventCommandLineHandler struct {
	eu   domain.EventUseCase
	root *cobra.Command
}

func New(u domain.EventUseCase, cmd *cobra.Command) EventCommandLineHandler {
	return EventCommandLineHandler{eu: u, root: cmd}
}

func (h *EventCommandLineHandler) Handle() {
	var cmd = &cobra.Command{
		Use:   "event",
		Short: "event related stuffs",
		Run:   helper.OperationHandler,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetContext(context.WithValue(cmd.Context(), `area`, EventOperator{h.eu}))
		},
	}
	h.root.AddCommand(cmd)
}
