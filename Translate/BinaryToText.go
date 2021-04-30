package Translate

type TranslateBinaryToTextImpl struct {
}

func NewTranslateBinaryToTextImpl() *TranslateBinaryToTextImpl {
	return &TranslateBinaryToTextImpl{}
}

func (c *TranslateBinaryToTextImpl) Translate(textToTranslate string) string {
	var text string
	text = "Text to Translate"
	return text

}
