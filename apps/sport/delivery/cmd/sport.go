package cmd

import (
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"fmt"
	"strconv"
)

type SportOperator struct {
	u domain.SportUseCase
}

func fillSportInteractive() domain.Sport {
	sport := domain.Sport{}
	pm := helper.PromptMessage{
		Msg:        "please enter sport name",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	sport.Name = helper.InputHandler(pm)
	pm = helper.PromptMessage{
		Msg:        "please enter sport display name",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	sport.DisplayName = helper.InputHandler(pm)
	pm = helper.PromptMessage{
		Msg:        "please enter order of sport in list",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	sport.Order, _ = strconv.Atoi(helper.InputHandler(pm))
	return sport
}
func (s SportOperator) Create(ctx context.Context) {
	sport := fillSportInteractive()
	sport, err := s.u.CreateSport(ctx, sport)
	if err != nil {
		fmt.Println(err)
	}
}

func (s SportOperator) Update(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter old sport slug",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	oldSportSlug := helper.InputHandler(pm)
	sport := fillSportInteractive()
	updateSport, err := s.u.UpdateSport(ctx, sport, oldSportSlug)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(`%#v`, updateSport)
}

func (s SportOperator) Deactivate(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter sport slug you want to deactivate",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	sportSlugForDelete := helper.InputHandler(pm)
	err := s.u.DeactivateSport(ctx, sportSlugForDelete)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`sport deactivated successfully`)
}
func (s SportOperator) Delete(ctx context.Context) {
	panic("imp")
}

func (s SportOperator) Search(ctx context.Context) {
	//TODO implement me
	panic("SS implement me")
}

func (s SportOperator) SearchAll(ctx context.Context) {
	s.Search(ctx)
}
