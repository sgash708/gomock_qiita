package i18n

import (
	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/locales/ja_JP"
	ut "github.com/go-playground/universal-translator"
)

type i18nClient struct {
	DefaultTranslator ut.Translator
	Translators       []ut.Translator
}

type I18nClientInterface interface {
	T(key string) string
	EmbedT(key, field string, ps ...string) string
}

func NewI18nClient() (I18nClientInterface, error) {
	trans := genTranslator()
	i18n := &i18nClient{}

	if ja, found := trans.GetTranslator("ja_JP"); found {
		if err := i18n.loadJaMessages(ja); err != nil {
			return nil, err
		}
		i18n.DefaultTranslator = ja
		i18n.Translators = append(i18n.Translators, ja)
	}
	return I18nClientInterface(i18n), nil
}

func genTranslator() *ut.UniversalTranslator {
	japanese := ja_JP.New()
	english := en_US.New()
	return ut.New(japanese, japanese, english)
}

func (c *i18nClient) loadJaMessages(ja ut.Translator) error {
	for _, msg := range LocaleMessages {
		if err := ja.Add(msg.Name, msg.Message, false); err != nil {
			return err
		}
	}
	return nil
}

func (c *i18nClient) T(key string) string {
	t, _ := c.DefaultTranslator.T(key)
	if len(t) == 0 {
		t = key
	}
	return t
}

func (c *i18nClient) EmbedT(key, field string, ps ...string) string {
	t, _ := c.DefaultTranslator.T(field)
	if len(t) == 0 {
		t = field
	}

	params := []string{t}
	params = append(params, ps...)
	t, _ = c.DefaultTranslator.T(key, params...)
	return t
}
