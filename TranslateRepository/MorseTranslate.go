package TranslateRepository

import "encoding/json"

func GetTextToMorse() map[string]string {

	var morseMap map[string]string
	val := `{"A": "°- ","B": "-°°° ","C": "-°-° ","D": "-°° ","E": "° ","F": "°°-° ","G": "--° ","H": "°°°° ","J": "°--- ","K": "-°- ","L": "°-°° ", "M": "-- ","N": "-° ","O": "--- ","P": "°--° ","Q": "--°- ","R": "°-° ","S": "°°° ","T": "- ","U ": "°°- ","V": "°°°- ","W": "°-- ","X": "-°°- ","Y": "-°-- ","Z": "--°° ", " ":"   " }`
	json.Unmarshal([]byte(val), &morseMap)
	return morseMap
}

func GetMorseToText() map[string]string {

	var morseMap map[string]string
	val := `{"°- ": "A","-°°° ": "B","-°-° ": "C","-°° ": "D","° ": "E","°°-° ": "F","--° ": "G","°°°° ": "H","°--- ": "J","-°- ": "K","°-°° ": "L","-- ": "M","-° ": "N","--- ": "O","°--° ": "P","--°- ": "Q","°-° ": "R","°°° ": "S","- ": "T","°°- ": "U","°°°- ": "V","°-- ": "W","-°°- ": "X","-°-- ": "Y","--°° ": "Z","   ": " " }`
	json.Unmarshal([]byte(val), &morseMap)
	return morseMap
}
