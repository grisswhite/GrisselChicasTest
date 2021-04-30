package SimpleTranslate

type TranslateBinaryToTextImpl struct {
}

func NewTranslateBinaryToTextImpl() *TranslateBinaryToTextImpl {
	return &TranslateBinaryToTextImpl{}
}

func (c *TranslateBinaryToTextImpl) Translate(textToTranslate string) string {
	var text string
	text = "Cadena de texto que se desea traducir"
	return text

}
