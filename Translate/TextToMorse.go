package Translate

import (
	"GrisselChicasTest/TranslateRepository"
	"fmt"
	"strings"
)

type TranslateMorseToTextImpl struct {
}

func NewTranslateMorseToTextImpl() *TranslateMorseToTextImpl {
	return &TranslateMorseToTextImpl{}
}

func (c *TranslateMorseToTextImpl) Translate(text string) string {
	var resultText string
	upperText := strings.ToUpper(text)
	var listC []string
	listC = strings.Split(upperText, "")
	for _, c := range listC {
		resultText += getMorseChar(c)
	}

	fmt.Println(resultText)
	return resultText

}

func getMorseChar(c string) string {
	val := TranslateRepository.GetTextToMorse()[c]
	return val
}
