package workspace

import (
	"fmt"
	"slices"

	"github.com/IaK3lwin/ac/internal/domain/syserror"
)

func (ws *Workspace) Switch(workspaceName string) error {

	if !slices.Contains(ws.storage.GetWorkspaceList(), workspaceName) {
		message := fmt.Sprintf("Workspace called %s not found", workspaceName)
		return syserror.FactoryWsError(message, "worspace/switch")
	}

	//update current workspace field
	ws.currentWorkspace = workspaceName
	// storage to change current workspace cache
	actions, err := ws.storage.GetActions(workspaceName)

	if err != nil {
		return err
	}

	ws.actions = actions

	err = ws.storage.UpdateWsCurrent(ws.currentWorkspace)

	if err != nil {
		return syserror.FactoryWsError(err.Error(), "worspace/switch")
	}

	return nil
}
