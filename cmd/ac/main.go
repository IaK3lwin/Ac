package main

import (
	"fmt"
	"os"

	"github.com/IaK3lwin/ac/internal/domain/commands"
	"github.com/IaK3lwin/ac/internal/domain/workspace"
	"github.com/IaK3lwin/ac/internal/infrastructure/storage"
)

func main() {
	args := os.Args[1:]

	if args == nil {
		return
	}

	ok, cmd := commands.Resolve(args[0])

	// validar se o input é um comando
	if ok {

		// se for comand executar ação do comando
		err := cmd(args[1:])

		if err != nil {
			fmt.Println("erro encontrado ao executar o metodo!\n ", err.Error())
		}
		return
	}

	storage := storage.YmlStorage{}

	// pegar o workspace atual
	ws, err := workspace.FactoryWorspace(storage)

	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	// verificar se o action existe

	action, err := ws.GetAction(args[0])

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	fmt.Println("action find: ", action)
	// executar o action
}
