package vo

import (
	"kakapo/internal/shared"
	"regexp"
	"strings"
)

var slugPattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

type SlugVO struct {
	value string
}

func NewSlugVO(value string) (SlugVO, error) {
	value = strings.ToLower(strings.TrimSpace(value))

	if value == "" {
		return SlugVO{}, shared.ValidationError(shared.ErrorInput{
			Message: "Slug inválido",
			Fields: []shared.FieldError{
				{
					Field:   "slug",
					Message: "Slug é obrigatório",
				},
			},
		})
	}

	if len(value) > 100 {
		return SlugVO{}, shared.ValidationError(shared.ErrorInput{
			Message: "Slug inválido",
			Fields: []shared.FieldError{
				{
					Field:   "slug",
					Message: "Slug deve ter no máximo 100 caracteres",
				},
			},
		})
	}

	if !slugPattern.MatchString(value) {
		return SlugVO{}, shared.ValidationError(shared.ErrorInput{
			Message: "Slug inválido",
			Fields: []shared.FieldError{
				{
					Field:   "slug",
					Message: "Slug deve conter apenas letras sem acento, números e hífens entre palavras",
				},
			},
		})
	}

	return SlugVO{value: value}, nil
}

func (s SlugVO) Value() string {
	return s.value
}
