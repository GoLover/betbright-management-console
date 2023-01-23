package cmd

import (
	"betbright-management-console/apps/selection/adapter"
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

type SelectionOperator struct {
	su domain.SelectionUseCase
	sa adapter.SearchAdapter
}

func fillSelectionInteractive() domain.Selection {
	selection := domain.Selection{}
	pm := helper.PromptMessage{
		Msg:        "please enter selection name",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	selection.Name = helper.InputHandler(pm)
	pm = helper.PromptMessage{
		Msg:        "please enter selection price",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	var err error
	selection.Price, err = decimal.NewFromString(helper.InputHandler(pm))
	if err != nil {
		panic(err)
	}
	pm = helper.PromptMessage{
		Msg:        "please enter event id for selection",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	selection.SelectedEvent.Id, _ = strconv.Atoi(helper.InputHandler(pm))
	pm = helper.PromptMessage{
		Msg:        "please enter market id for selection",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	selection.SelectedMarket.Id, _ = strconv.Atoi(helper.InputHandler(pm))
	outcomes := []string{"Unsettled", "Void", "Lose", "Place", "Win"}
	pm = helper.PromptMessage{
		Msg:        "please select outcome",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: outcomes,
	}
	outcome := helper.SelectHandler(pm)
	for i, k := range outcomes {
		if k == outcome {
			selection.Outcome = domain.OutcomeState(i) + 1
			break
		}
	}
	return selection
}
func (s SelectionOperator) Create(ctx context.Context) {
	selection := fillSelectionInteractive()
	_, err := s.su.CreateSelection(ctx, selection, selection.SelectedMarket.Id, selection.SelectedEvent.Id)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (s SelectionOperator) Update(ctx context.Context) {
	selection := fillSelectionInteractive()
	pm := helper.PromptMessage{
		Msg:        "please enter selection id for update",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	selectionId, _ := strconv.Atoi(helper.InputHandler(pm))
	_, err := s.su.UpdateSelection(ctx, selection, selectionId)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (s SelectionOperator) Delete(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
func (s SelectionOperator) Deactivate(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter selection id to deactivate",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	selectionId, _ := strconv.Atoi(helper.InputHandler(pm))
	err := s.su.DeactivateSelection(ctx, selectionId)
	if err != nil {
		fmt.Println(err)
	}
}
func (s SelectionOperator) Activate(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter selection id to activate",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	selectionId, _ := strconv.Atoi(helper.InputHandler(pm))
	err := s.su.ActivateSelection(ctx, selectionId)
	if err != nil {
		fmt.Println(err)
	}
}

func (s SelectionOperator) Search(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "enter your query: ",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	searchQuery := helper.InputHandler(pm)
	result, err := s.sa.Search(ctx, `events`, searchQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`%#v`, result)
}

func (s SelectionOperator) SearchAll(ctx context.Context) {
	s.Search(ctx)
}
