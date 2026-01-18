package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

func Converter(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}

	// определяем: это Морзе или текст
	isMorse := true
	for _, sign := range input {
		if sign != '.' && sign != '-' && sign != ' ' && sign != '/' {
			isMorse = false
			break
		}
	}

	if isMorse {
		// декодируем Морзе в текст
		return morse.DefaultConverter.ToText(input)
	}
	// кодируем текст в Морзе
	return morse.DefaultConverter.ToMorse(input)
}
