package SimpleTranslate

import (
	"reflect"
	"testing"
)

func TestConvertTypeTextToMorseImplementation_ConvertByType(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Translating text to morse",
			args: args{input: "Cadena de texto que se desea traducir"},
			want: "-.-. .- -.. . -. .-    -.. .    - . -..- - ---    --.- .    ... .    -.. . ... . .-    - .-. .- -.. -.-. .. .-.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &TranslateTextToMorseImpl{}
			if got := t.translate(tt.args.input); got != tt.want {
				t.Errorf("ConvertByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
