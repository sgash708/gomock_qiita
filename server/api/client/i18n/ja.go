package i18n

type LocaleMessage struct {
	Name    string `yaml:"name"`
	Message string `yaml:"message"`
}

var LocaleMessages = []*LocaleMessage{
	{Name: "required", Message: "{0}はゼロ、または空にできません"},
	{Name: "exists", Message: "{0}は必須項目です"},
	{Name: "url", Message: "{0}はURL形式にしてください"},
}
