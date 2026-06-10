package actions

type ActSession struct {
	Alias       string   `yaml:"alias"`
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Sessions    []string `yaml:"sessions"`
}

func (as ActSession) GetAlias() string {
	return as.Alias
}

func (as ActSession) GetDescription() string {
	return as.Description
}

func (as ActSession) GetType() string {
	return as.Type
}
