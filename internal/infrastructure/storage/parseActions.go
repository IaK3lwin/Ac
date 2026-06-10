package storage

import (
	"github.com/IaK3lwin/ac/internal/domain/actions"
	"gopkg.in/yaml.v3"
)

type parseCommand func(data []byte) (actions.Actions, error)

var parseAction = map[string]parseCommand{
	"command": func(data []byte) (actions.Actions, error) {
		var commandType actions.ActCommand
		err := yaml.Unmarshal(data, &commandType)
		if err != nil {
			return actions.ActCommand{}, err
		}

		return commandType, nil
	},
	"dir": func(data []byte) (actions.Actions, error) {
		var dirType actions.ActDir
		err := yaml.Unmarshal(data, &dirType)

		if err != nil {
			return actions.ActionNil{}, err
		}

		return dirType, nil
	},

	"session": func(data []byte) (actions.Actions, error) {
		var sessionType actions.ActSession

		err := yaml.Unmarshal(data, &sessionType)

		if err != nil {
			return actions.ActionNil{}, err
		}

		return sessionType, nil
	},
}
