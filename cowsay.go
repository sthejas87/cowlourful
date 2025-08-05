package main

import (
	"strings"
)

const cowArt = `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`

func cowlourful (message string) string {
	var bubble strings.Builder
	msgRunes := []rune(message)
	msgLen := len(msgRunes)
	//top of bubble
	bubble.WriteString(" "+strings.Repeat("-", msgLen+2)+"\n")
	
	bubble.WriteString("< ")

	for i, char:= range msgRunes{
		bubble.WriteString(getColour(i,msgLen))
		bubble.WriteRune(char)
	}

	bubble.WriteString("\x1b[0m")
	bubble.WriteString(" >\n")
	//bottom of bubble
	bubble.WriteString(" "+strings.Repeat("-",msgLen+2)+"\n")

	return bubble.String() + cowArt

}

