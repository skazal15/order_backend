package utils

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	Dto                 interface{}
	Validator           *validator.Validate
	UniversalTranslator ut.Translator
}

func (d *StructValidator) Validate() []string {
	err := d.Validator.Struct(d.Dto)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errTrans := validationErrors.Translate(d.UniversalTranslator)
		var errStr []string
		for _, value := range errTrans {
			errStr = append(errStr, value)
		}
		return errStr
	}
	return nil
}
