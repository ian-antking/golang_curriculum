package character

import (
	"unicode"
	"errors"
)

func Shift(character rune, spaces uint) (newCharacter rune, err error){
	character = unicode.ToLower(character)
	if character > 122 || character < 92 {
		err = errors.New("character to shift must be a-z")
	}
	newCharacter = character + rune(spaces)
	if newCharacter > 122 {
		newCharacter -= 26
	}
	return
}