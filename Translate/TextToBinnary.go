package Translate

import "fmt"

type TranslateTextToBinaryImpl struct {
}

func NewTranslateTextToBinaryImpl() *TranslateTextToBinaryImpl {
	return &TranslateTextToBinaryImpl{}
}

func (c *TranslateTextToBinaryImpl) Translate(textToTranslate string) string {
	var text string
	for _, i := range textToTranslate {
		text += fmt.Sprintf("%.8b", i)
	}
	return text

}
