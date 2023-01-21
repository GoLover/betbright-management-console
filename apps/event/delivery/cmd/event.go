package cmd

import (
	"betbright-management-console/domain"
	"betbright-management-console/domain/helper"
	"context"
	"fmt"
)

type EventOperator struct {
	eu domain.EventUseCase
}

func fillEventInteractively() (domain.Event, string) {
	event := domain.Event{}
	pm := helper.PromptMessage{
		Msg:        "please enter event name",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	event.Name = helper.InputHandler(pm)
	eventStatuses := []string{"PrePlay", "InPlay", "Ended"}
	pm = helper.PromptMessage{
		Msg:        "please choose event status",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: eventStatuses,
	}
	selectedStatus := helper.SelectHandler(pm)
	for i, k := range eventStatuses {
		if k == selectedStatus {
			event.Status = domain.EventStatus(i) + 1
			break
		}
	}

	eventTypes := []string{"PrePlay", "InPlay"}
	pm = helper.PromptMessage{
		Msg:        "please choose event type",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: eventTypes,
	}
	selectedType := helper.SelectHandler(pm)
	for i, k := range eventTypes {
		if k == selectedType {
			event.EType = domain.EventType(i) + 1
			break
		}
	}

	pm = helper.PromptMessage{
		Msg:        "please enter sport slug for add event to it",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	sportSlug := helper.InputHandler(pm)

	return event, sportSlug
}
func (s EventOperator) Create(ctx context.Context) {
	event, sportSlug := fillEventInteractively()

	event, err := s.eu.CreateEvent(ctx, event, sportSlug)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (s EventOperator) Update(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter event slug for update",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	eventSlugForUpdate := helper.InputHandler(pm)
	event, sportSlug := fillEventInteractively()
	updateEvent, err := s.eu.UpdateEvent(ctx, event, eventSlugForUpdate, sportSlug)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(`%#v`, updateEvent)

}

func (s EventOperator) Delete(ctx context.Context) {
	pm := helper.PromptMessage{
		Msg:        "please enter event slug for update",
		ErrMsg:     domain.ErrDeliveryIncorrectInput.Error(),
		Selectable: nil,
	}
	eventSlugForDelete := helper.InputHandler(pm)
	s.eu.DeleteEvent(ctx, eventSlugForDelete)
}

func (s EventOperator) Deactivate(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s EventOperator) Search(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s EventOperator) SearchAll(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
