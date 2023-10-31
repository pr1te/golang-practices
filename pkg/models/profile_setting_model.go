package models

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

type Language string

var (
	TH Language = Language(display.Self.Name(language.Thai))
)

type ProfileSetting struct {
	Model
	Language Language `json:"language"`
}

func (lang *Language) Scan(value interface{}) error {
	if value == nil {
		*lang = ""

		return nil
	}

	str, ok := value.(string)

	if !ok {
		return fmt.Errorf("invalud lang. got %s", value)
	}

	matcher := language.NewMatcher([]language.Tag{
		language.Thai,
		language.AmericanEnglish,
	})

	tag := language.MustParse(str)
	_, _, confidence := matcher.Match(tag)

	if confidence == language.No {
		return fmt.Errorf("unsupported language. got '%s'", str)
	}

	return nil
}
