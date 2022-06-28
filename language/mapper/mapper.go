package mapper

import (
	"github.com/comasz/GoLang-Learning"
)
  
func Greets(s string) string {
	info := getlang.FromString(s)
	switch info.LanguageCode() {
	case "en":
		return "Hello"
	case "de":
		return "Guten Tag"
	case "fr":
		return "Bonjour"
	default:
		return "I dont know your language yet"
	}
}
