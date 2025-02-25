package validation

import (
	"errors"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	goplayground "github.com/go-playground/validator/v10"
	jatranslation "github.com/go-playground/validator/v10/translations/ja"
	"devport/adapter/validator"
)

type goPlayground struct {
	validator *goplayground.Validate
	translate ut.Translator
	err       error
	msg       []string
}

func NewGoPlayground() (validator.Validator, error) {
	var (
		lang             = ja.New()
		uni              = ut.New(lang, lang)
		translate, found = uni.GetTranslator("ja")
	)

	if !found {
		return nil, errors.New("translator not found")
	}

	v := goplayground.New()
	if err := jatranslation.RegisterDefaultTranslations(v, translate); err != nil {
		return nil, err
	}

	return &goPlayground{
		validator: v,
		translate: translate,
	}, nil
}

func (g *goPlayground) Validate(i interface{}) error {
	if len(g.msg) > 0 {
		g.msg = nil
	}

	g.err = g.validator.Struct(i)
	if g.err != nil {
		return g.err
	}

	return nil
}

func (g *goPlayground) Messages() []string {
	if g.err != nil {
		for _, err := range g.err.(goplayground.ValidationErrors) {
			g.msg = append(g.msg, err.Translate(g.translate))
		}
	}

	return g.msg
}
