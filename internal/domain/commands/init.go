package commands

import (
	"fmt"
	"os"
)

func Config(shell string) string {
	exe, _ := os.Executable()

	return fmt.Sprintf(`ac() {
	local result

	result="$("%s" "$@")"

	if [[ "$result" == __AC_CD__:* ]]; then
		cd "${result#__AC_CD__:}"
		return
	fi

	echo "$result"
}
`, exe)

}
