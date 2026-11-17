package vo

import (
	"strings"
	"unicode"
)

type CleanStringOptions struct {
	MaxLength int
	Lowercase bool
	Uppercase bool
}

type CleanStringVO struct {
	value string
}

func NewCleanStringVO(
	value string,
	options ...CleanStringOptions,
) CleanStringVO {
	config := CleanStringOptions{
		MaxLength: 255,
	}

	if len(options) > 0 {
		config = options[0]

		if config.MaxLength <= 0 {
			config.MaxLength = 255
		}
	}

	cleaned := strings.Map(func(r rune) rune {
		switch {
		case unicode.IsLetter(r), unicode.IsMark(r), unicode.IsNumber(r):
			return r
		case unicode.IsSpace(r):
			return ' '
		case strings.ContainsRune(".,!?;:-()'\"", r):
			return r
		default:
			return -1
		}
	}, value)

	// Remove espaços nas extremidades e agrupa espaços repetidos.
	cleaned = strings.Join(strings.Fields(cleaned), " ")

	// Lowercase tem prioridade se ambas estiverem habilitadas.
	if config.Lowercase {
		cleaned = strings.ToLower(cleaned)
	} else if config.Uppercase {
		cleaned = strings.ToUpper(cleaned)
	}

	characters := []rune(cleaned)
	if len(characters) > config.MaxLength {
		cleaned = string(characters[:config.MaxLength])
	}

	return CleanStringVO{
		value: strings.TrimSpace(cleaned),
	}
}

func (v CleanStringVO) Value() string {
	return v.value
}
