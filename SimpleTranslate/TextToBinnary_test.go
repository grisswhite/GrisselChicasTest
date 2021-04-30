package SimpleTranslate

import (
	"testing"
)

func TestTranslateTextToBinaryImpl_Translate(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Text-Binary",
			args: args{input: "Cadena de texto que se desea traducir"},
			want: "01000011 01100001 01100100 01100101 01101110 01100001 00100000 01100100 01100101 00100000 01110100 01100101 01111000 01110100 01101111 00100000 01110001 01110101 01100101 00100000 01110011 01100101 00100000 01100100 01100101 01110011 01100101 01100001 00100000 01110100 01110010 01100001 01100100 01110101 01100011 01101001 01110010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TranslateTextToBinaryImpl{}
			if got := c.Translate(tt.args.input); got != tt.want {
				t.Errorf("ConvertByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
