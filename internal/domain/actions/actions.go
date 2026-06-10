package actions

type Actions interface {
	GetAlias() string
	GetDescription() string
	GetType() string
}

type ActionsBaseType struct {
	Type string `yaml:"type"`
}
