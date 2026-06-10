package io

import "github.com/IaK3lwin/ac/internal/domain/actions"

type Storage interface {
	GetActions(workspaceName string) (map[string]actions.Actions, error)
	GetWorkspaceList() []string
	UpdateWsCurrent(newCurrentWs string) error
}
