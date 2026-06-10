package config

import (
	"fmt"
	"os/user"
)

var instanceConfig config

type config struct {
	pathToWorkspace        string
	pathTocurrentWorkspace string
	nameCurrentFileWs      string
	pathWs                 string
}

func Config() config {
	if (instanceConfig != config{}) {
		return instanceConfig
	}

	username, err := user.Current()

	if err != nil {
		return config{}
	}

	return config{
		pathToWorkspace:        fmt.Sprintf("%s/.cache/ac/workspace/", username.HomeDir),
		pathTocurrentWorkspace: fmt.Sprintf("%s/.cache/ac/current_ws", username.HomeDir),
		nameCurrentFileWs:      "current.txt",
		pathWs:                 "/home/ian-kelwin/.config/ac/workspaces",
	}
}
