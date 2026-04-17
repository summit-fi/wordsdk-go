package locale

import "strings"

func GetLocaleFromFileName(fileName string) string {
	var locale string
	// Assuming the file name format is "en_EU.ftl", we can split by '.' and take the first part
	parts := strings.Split(fileName, ".")
	if len(parts) > 0 {
		locale = parts[0]

		names := strings.Split(parts[0], "/")
		if len(names) > 0 {
			locale = names[len(names)-1]
		}

	}
	if !strings.Contains(locale, "_") {
		switch strings.ToLower(locale) {
		case "en":
			return "en_US"
		case "uk":
			return "uk_UA"
		case "sv":
			return "sv_SE"
		case "es":
			return "es_CO"
		}
	}
	return locale
}
