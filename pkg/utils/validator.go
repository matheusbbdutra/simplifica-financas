package utils

import (
    "github.com/go-playground/locales/pt_BR"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    ptbr_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)


type Validate struct {
	validate  *validator.Validate
	translator ut.Translator
}

func NewValidator() *Validate {
	validate := validator.New()
    uni := ut.New(pt_BR.New())
    trans, _ := uni.GetTranslator("pt_BR")
    ptbr_translations.RegisterDefaultTranslations(validate, trans)

	return &Validate{
		validate:   validate,
        translator: trans,
	}
}

func (v *Validate) TranslateError(err error) map[string]string {
    errors := make(map[string]string)
    if validationErrs, ok := err.(validator.ValidationErrors); ok {
        for _, fieldErr := range validationErrs {
            errors[fieldErr.Field()] = fieldErr.Translate(v.translator)
        }
    }
    return errors
}

func (v *Validate) ValidateStruct(i interface{}) error {
	return v.validate.Struct(i)
}

