package source

type Source interface {
	LoadAll(checksumIn string) (result []Object, checksumOut string, err error)
	Save(data []Object) error
}

type Object struct {
	LocaleCode string      `json:"localeCode"`
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
}
