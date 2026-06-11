package main

import (
	"fmt"
	"os"

	"github.com/IaK3lwin/ac/internal/domain/commands"
	configsh "github.com/IaK3lwin/ac/internal/domain/config_sh"
	"github.com/IaK3lwin/ac/internal/domain/workspace"
	"github.com/IaK3lwin/ac/internal/infrastructure/executor"
	"github.com/IaK3lwin/ac/internal/infrastructure/storage"
)

func main() {
	args := os.Args[1:]

	if args == nil {
		return
	}

	if args[0] == "config" {
		fmt.Print(commands.Config("zsh"))
		return
	}

	if args[0] == "config-sh" {
		err := configsh.Install()
		if err != nil {
			fmt.Println("Erro ao configurar o sh", err.Error())
		}
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

	// executar o action

	err = executor.Execute(action)

	if err != nil {
		fmt.Println("[Error executing action]\n ", err.Error())
	}
}
