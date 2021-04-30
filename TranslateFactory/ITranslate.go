package TranslateFactory

import "GrisselChicasTest/Translate"

type ITranslate interface {
	Translate(textoATraducir string, formatoOrigen string, formatoDestino string) (Translate.ISimpleTranslate, error)
}
