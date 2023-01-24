package cmd

import (
	"betbright-management-console/apps/market/adapter"
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"fmt"
	"strconv"
)

type MarketOperator struct {
	sa adapter.SearchAdapter
	mu domain.MarketUseCase
}

func fillMarketInteractive() domain.Market {
	market := domain.Market{}
	pm := helper.PromptMessage{
		Msg:        "please enter market name",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	market.Name = helper.InputHandler(pm)
	pm = helper.PromptMessage{
		Msg:        "please enter market display name",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	market.DisplayName = helper.InputHandler(pm)
	pm = helper.PromptMessage{
		Msg:        "please enter market order in list",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	market.Order, _ = strconv.Atoi(helper.InputHandler(pm))
	pm = helper.PromptMessage{
		Msg:        "please enter market columns number",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	market.Columns, _ = strconv.Atoi(helper.InputHandler(pm))
	pm = helper.PromptMessage{
		Msg:        "please enter market schema number",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	market.Schema, _ = strconv.Atoi(helper.InputHandler(pm))
	return market
}
func (s MarketOperator) Create(ctx context.Context) {
	market := fillMarketInteractive()
	pm := helper.PromptMessage{
		Msg:        "please enter event slug name for market",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	eventSlug := helper.InputHandler(pm)
	market, err := s.mu.CreateMarket(ctx, market, eventSlug)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(`%#v`, market)

}

func (s MarketOperator) Update(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter id of market for update",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	marketId, _ := strconv.Atoi(helper.InputHandler(pm))
	market := fillMarketInteractive()
	pm = helper.PromptMessage{
		Msg:        "please enter event slug name for market",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	eventSlug := helper.InputHandler(pm)
	market, err := s.mu.UpdateMarket(ctx, market, marketId, eventSlug)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(`%#v`, market)
}

func (s MarketOperator) Deactivate(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter id of market to deactivate",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	marketId, _ := strconv.Atoi(helper.InputHandler(pm))
	err := s.mu.DeactivateMarket(ctx, marketId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(`market deactivated successfully`)
}

func (s MarketOperator) Activate(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter id of market to activate",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	marketId, _ := strconv.Atoi(helper.InputHandler(pm))
	err := s.mu.ActivateMarket(ctx, marketId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(`market activated successfully`)
}

func (s MarketOperator) Delete(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter market id you want to delete",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	marketIdForDelete, err := strconv.Atoi(helper.InputHandler(pm))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = s.mu.DeleteMarket(ctx, marketIdForDelete)
	if err != nil {
		fmt.Println(err)
	}
}

func (s MarketOperator) Search(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "enter your query: ",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	searchQuery := helper.InputHandler(pm)
	result, err := s.sa.Search(ctx, `markets`, searchQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	result.PrettyPrint()
}

func (s MarketOperator) SearchAll(ctx context.Context) {
	s.Search(ctx)
}
