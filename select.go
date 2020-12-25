package gotoant

type antSelect struct {
	options []AntSelectOption
}

type AntSelectOption struct {
	Title string      `json:"title"`
	Value interface{} `json:"value"`
}

func NewSelect() *antSelect {
	return &antSelect{}
}

func (a *antSelect) SetOption(title string, value interface{}) {
	a.options = append(a.options, AntSelectOption{
		Title: title,
		Value: value,
	})
}

func (a *antSelect) GetOptions() []AntSelectOption {
	return a.options
}
