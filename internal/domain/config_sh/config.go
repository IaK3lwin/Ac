package configsh

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Install() error {
	shell := filepath.Base(os.Getenv("SHELL"))

	configFile, err := getShellConfig(shell)
	if err != nil {
		return err
	}

	executable, err := os.Executable()
	if err != nil {
		return err
	}

	hook := fmt.Sprintf(
		`eval "$(%s config %s)"`,
		executable,
		shell,
	)

	content, err := os.ReadFile(configFile)
	if err == nil {
		if strings.Contains(string(content), hook) {
			fmt.Println("AC já está configurado.")
			return nil
		}
	}

	file, err := os.OpenFile(
		configFile,
		os.O_APPEND|os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(
		fmt.Sprintf(
			"\n# AC\n%s\n",
			hook,
		),
	)
	if err != nil {
		return err
	}

	fmt.Printf(
		"AC configurado com sucesso em %s\n\n"+
			"Reabra o terminal ou execute:\n"+
			"source %s\n",
		configFile,
		configFile,
	)

	return nil
}

func getShellConfig(shell string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	switch shell {
	case "zsh":
		return filepath.Join(home, ".zshrc"), nil

	case "bash":
		return filepath.Join(home, ".bashrc"), nil

	default:
		return "", fmt.Errorf(
			"shell não suportado: %s",
			shell,
		)
	}
}
