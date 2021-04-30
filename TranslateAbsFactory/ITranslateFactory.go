package TranslateAbsFactory

type ITranslate interface {
	Translate(textoATraducir string, formatoOrigen string, formatoDestino string)
}
