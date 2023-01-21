package helper

import (
	"betbright-management-console/domain"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
)

func OperationHandler(cmd *cobra.Command, args []string) {
	oHandler := cmd.Context().Value(`area`).(domain.Operator)
	if cmd.Use != "search" {
		availableOperations := []string{"create", "update", "search", "delete", "deactivate"}
		pm := PromptMessage{
			Msg:        "what do you want to do?",
			ErrMsg:     "Err!",
			Selectable: availableOperations,
		}
		var operation string
		if len(args) > 0 {
			for _, k := range availableOperations {
				if args[0] == k {
					operation = k
				}
			}
		}
		if operation == "" {
			operation = SelectHandler(pm)
		}
		switch operation {
		case "create":
			oHandler.Create(cmd.Context())
		case "update":
			oHandler.Update(cmd.Context())
		case "delete":
			oHandler.Delete(cmd.Context())
		case "deactivate":
			oHandler.Deactivate(cmd.Context())
		case "search":
			oHandler.Search(cmd.Context())
		default:
			panic("what is your choice?")
		}
		return
	}
	oHandler.SearchAll(cmd.Context())
}

type PromptMessage struct {
	Msg        string
	ErrMsg     string
	Selectable []string
}

func InputHandler(pc PromptMessage) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Msg,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func SelectHandler(pc PromptMessage) string {
	items := pc.Selectable
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: pc.Msg,
			Items: items,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
