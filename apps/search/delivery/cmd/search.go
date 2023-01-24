package cmd

import (
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"fmt"
)

type SearchOperator struct {
	u domain.SearchUsecase
}

func (s SearchOperator) Create(ctx context.Context) {
	s.SearchAll(ctx)
}

func (s SearchOperator) Update(ctx context.Context) {
	s.SearchAll(ctx)
}

func (s SearchOperator) Delete(ctx context.Context) {
	s.SearchAll(ctx)
}

func (s SearchOperator) Search(ctx context.Context) {
	s.SearchAll(ctx)
}

func (s SearchOperator) SearchAll(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter query phrase:",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	searchQuery := helper.InputHandler(pm)
	result, err := s.u.Search(ctx, `*`, searchQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	result.PrettyPrint()
}
