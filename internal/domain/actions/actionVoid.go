package actions

type ActionNil struct {
}

func (ac ActionNil) GetAlias() string {
	return ""

}

func (ac ActionNil) GetDescription() string {
	return ""
}

func (ac ActionNil) GetType() string {
	return ""
}
