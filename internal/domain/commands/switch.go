package commands

import (
	"fmt"

	"github.com/IaK3lwin/ac/internal/domain/workspace"
	"github.com/IaK3lwin/ac/internal/infrastructure/storage"
)

func switchCommand(args []string) error {

	storage := storage.YmlStorage{}

	ws, err := workspace.FactoryWorspace(storage)

	fmt.Println("workspace current: ", ws)

	if err != nil {
		return err
	}

	fmt.Println("workspace: ", args[0])

	err = ws.Switch(args[0])
	storage.UpdateWsCurrent(args[0])

	if err != nil {
		fmt.Println("Erro found in switch command: ", err.Error())
		return err
	}

	return nil
}
