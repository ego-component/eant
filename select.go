package gotoant

type AntSelect struct {
	options []AntSelectOption
}

type AntSelectOption struct {
	Title string      `json:"title"`
	Value interface{} `json:"value"`
}

func NewSelect() *AntSelect {
	return &AntSelect{
		options: make([]AntSelectOption, 0),
	}
}

func (a *AntSelect) SetOption(title string, value interface{}) {
	a.options = append(a.options, AntSelectOption{
		Title: title,
		Value: value,
	})
}

func (a *AntSelect) GetOptions() []AntSelectOption {
	return a.options
}
