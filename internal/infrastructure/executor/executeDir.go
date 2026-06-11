package executor

import (
	"fmt"

	"github.com/IaK3lwin/ac/internal/domain/actions"
)

func executorDir(ac actions.Actions) error {

	dir := ac.(actions.ActDir).GetDir()

	fmt.Printf("__AC_CD__:%s\n", dir)

	return nil
}
