package commands

type CommandInternal = func(args []string) error

var internalComands = map[string]CommandInternal{
	"switch": switchCommand,
}

func Resolve(nameCommand string) (bool, CommandInternal) {
	commandValue, ok := internalComands[nameCommand]
	return ok, commandValue
}
