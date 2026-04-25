package main

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func noForbiddenWords(fl validator.FieldLevel) bool {
	text, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	forbidden := []string{"вайб", "рофл", "кринж"}
	lowerText := strings.ToLower(text)

	for _, word := range forbidden {
		// Ищем слово как отдельную лексему
		if strings.Contains(lowerText, word) {
			// Проверяем, что это действительно слово, а не часть другого слова
			idx := strings.Index(lowerText, word)
			if idx == 0 || !isLetter(rune(lowerText[idx-1])) {
				if idx+len(word) == len(lowerText) || !isLetter(rune(lowerText[idx+len(word)])) {
					return false
				}
			}
		}
	}
	return true
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'а' && r <= 'я')
}
