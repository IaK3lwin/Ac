package executor

import "github.com/IaK3lwin/ac/internal/domain/actions"

type executeMethod func(actions.Actions) error

var executorMethodList = map[string]executeMethod{
	"dir": executorDir,
}

func Execute(currentAct actions.Actions) error {
	err := executorMethodList[currentAct.GetType()](currentAct)

	if err != nil {
		return err
	}

	return nil

}
