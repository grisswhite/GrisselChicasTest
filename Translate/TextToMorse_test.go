package Translate

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
			c := &ConvertTypeTextToMorseImplementation{}
			if got := c.ConvertByType(tt.args.input); got != tt.want {
				t.Errorf("ConvertByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConvertTypeTextToMorseImplementation(t *testing.T) {
	tests := []struct {
		name string
		want *ConvertTypeTextToMorseImplementation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConvertTypeTextToMorseImplementation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConvertTypeTextToMorseImplementation() = %v, want %v", got, tt.want)
			}
		})
	}
}
