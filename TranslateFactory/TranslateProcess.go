package TranslateFactory

import (
	"GrisselChicasTest/Translate"
	"fmt"
)

type FactoryTranslator struct {
}

func NewFactoryTranslator() *FactoryTranslator {
	return &FactoryTranslator{}
}

func (f *FactoryTranslator) Factory(text string, destinationFormat string) (Itranslate.IConvertTypes, error) {

	if "TEXT" == text {
		if "BINARY" == destinationFormat {
			return ConvertTypes.NewConvertTypeTextToBinaryImplementation(), nil
		} else if "MORSE" == destinationFormat {
			return ConvertTypes.NewConvertTypeTextToMorseImplementation(), nil
		}
	}
	return nil, fmt.Errorf("not valid option")
}
