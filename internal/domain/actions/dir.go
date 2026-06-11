package actions

type ActDir struct {
	Alias       string `yaml:"alias"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Dir         string `yaml:"dir"`
}

func (ad ActDir) GetAlias() string {
	return ad.Alias
}

func (ad ActDir) GetDescription() string {
	return ad.Description
}

func (ad ActDir) GetType() string {
	return ad.Type
}

func (ad ActDir) GetDir() string {
	return ad.Dir
}
