package TranslateAbsFactory

import "GrisselChicasTest/SimpleTranslate"

type TransLateFachada struct {
	morse SimpleTranslate.TranslateTextToMorseImpl
	text  SimpleTranslate.TranslateBinaryToTextImpl
}

func NewTransLateFachada() *TransLateFachada {
	return &TransLateFachada{
		morse: SimpleTranslate.NewTranslateMorseToTextImpl(),
		text:  SimpleTranslate.NewTranslateBinaryToTextImpl(),
	}
}

func (t TransLateFachada) Translate(text string) string {
	text = t.text.Translate(text)
	text = t.morse.Translate(text)
}
