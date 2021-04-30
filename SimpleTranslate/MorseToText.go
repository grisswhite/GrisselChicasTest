package SimpleTranslate

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
	var listC []string
	listC = strings.Split(text, " ")
	for _, c := range listC {
		resultText += getTextChar(c)
	}

	fmt.Println(resultText)
	return resultText

}

func getTextChar(c string) string {
	val := TranslateRepository.GetMorseToText(c)
	return val
}
