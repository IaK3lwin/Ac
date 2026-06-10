package workspace

import (
	"fmt"
	"slices"

	"github.com/IaK3lwin/ac/internal/domain/actions"
	"github.com/IaK3lwin/ac/internal/domain/syserror"
)

func (ws Workspace) GetAction(alias string) (actions.Actions, error) {
	var aliasExisted []string

	for _, action := range ws.actions {

		aliasAction := action.GetAlias()

		aliasExisted = append(aliasExisted, aliasAction)
	}

	fmt.Println("array slias: ", aliasExisted)

	if slices.Contains(aliasExisted, alias) {
		return ws.actions[alias], nil
	}

	return actions.ActionNil{}, syserror.WsActionsNotFound(alias, "", "domain/workspace/getAction()")
}
