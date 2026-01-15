package service

import (
	"strings"
)

var textMorseCode = map[string]string{
	"А": ".-", "Б": "-...", "В": ".--", "Г": "--.", "Д": "-..", "Е": ".",
	"Ж": "...-", "З": "--..", "И": "..", "Й": ".---", "К": "-.-", "Л": ".-..",
	"М": "--", "Н": "-.", "О": "---", "П": ".--.", "Р": ".-.", "С": "...",
	"Т": "-", "У": "..-", "Ф": "..-.", "Х": "....", "Ц": "-.-.", "Ч": "---.",
	"Ш": "----", "Щ": "--.-", "Ъ": ".--.-.", "Ы": "-.--", "Ь": "-..-", "Э": "..-..",
	"Ю": "..--", "Я": ".-.-", "1": ".----", "2": "..---", "3": "...--", "4": "....-",
	"5": ".....", "6": "-....", "7": "--...", "8": "---..", "9": "----.", "0": "-----", " ": "/",
}

var morseCodeText = func() map[string]string {
	m := make(map[string]string)
	for key, value := range textMorseCode {
		m[value] = key
	}
	return m
}()

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
		parts := strings.Split(input, " ")
		var decoded []string
		for _, code := range parts {
			if code == "" {
				continue
			}
			if char, ok := morseCodeText[code]; ok {
				decoded = append(decoded, char)
			} else {
				return ""
			}
		}
		return strings.Join(decoded, "")
	}
	// кодируем текст в Морзе
	input = strings.ToUpper(input)
	var encoded []string
	for _, r := range input {
		if code, ok := textMorseCode[string(r)]; ok {
			encoded = append(encoded, code)
		} else {
			return ""
		}
	}
	return strings.Join(encoded, " ")
}
