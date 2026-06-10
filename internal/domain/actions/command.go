package actions

type ActCommand struct {
	Alias       string `yaml:"alias"`
	Command     string `yaml:"command"`
	Description string `yaml:"description"`
	TypeAct     string `yaml:"type"`
}

func (ac ActCommand) GetAlias() string {
	return ac.Alias
}

func (ac ActCommand) GetDescription() string {
	return ac.Description
}

func (ac ActCommand) GetType() string {
	return ac.TypeAct
}
