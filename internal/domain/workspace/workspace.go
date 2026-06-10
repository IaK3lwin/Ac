package workspace

import (
	"github.com/IaK3lwin/ac/internal/domain/actions"
	"github.com/IaK3lwin/ac/internal/domain/io"
)

type Workspace struct {
	currentWorkspace string
	actions          map[string]actions.Actions
	storage          io.Storage
}

var instance *Workspace

func FactoryWorspace(storage io.Storage) (*Workspace, error) {
	if instance != nil {
		return instance, nil
	}

	instance = &Workspace{}

	currentWs, err := GetNameWsCurrent()

	if err != nil {
		return nil, err
	}

	instance.currentWorkspace = currentWs

	instance.storage = storage
	actions, err := storage.GetActions(instance.currentWorkspace)

	if err != nil {
		return instance, err
	}

	instance.actions = actions

	return instance, nil
}
