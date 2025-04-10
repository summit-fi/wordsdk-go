package source

type Source interface {
	LoadAllStatic(checksumIn string) (result []Object, checksumOut string, err error)
	LoadAllDynamic(dynamicKey string, checksumIn string) (result []Object, checkSumOut string, err error)
	LoadOneDynamic(accessKey, lang, key string) (string, error)
	SaveDynamic(accessKey string, data []Object) error
	Save(data []Object) error
}

type Object struct {
	LocaleCode string `json:"localeCode"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}
