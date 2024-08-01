package goi18n

import (
	"fmt"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestTranslate(t *testing.T) {
	//新建语言包
	bundle := i18n.NewBundle(language.English)

	loc := i18n.NewLocalizer(bundle, language.English.String())
	messages := &i18n.Message{
		ID:          "Emails",
		Description: "The number of unread emails a user has",
		One:         "{{.Name}} has {{.Count}} email.",
		Other:       "{{.Name}} has {{.Count}} emails.",
	}

	messageCount := 2
	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: messages,
		TemplateData: map[string]interface{}{
			"Name":  "Theo",
			"Count": messageCount,
		},
		PluralCount: messageCount,
	})

	fmt.Println(translation)

}
