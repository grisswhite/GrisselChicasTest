package main

import "GrisselChicasTest/TranslateFactory"

func main() {
	fachada := TranslateFactory.Translate()
	fachada.Encode("Kennia")
}
