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

func cowlourful (messages []string) string {

	if len(messages) == 0{
		return cowArt
	}
	
	maxLen := 0
	for _, msg := range messages{
		if len([]rune(msg))>maxLen{
			maxLen = len([]rune(msg))
		}
	}

	var bubble strings.Builder

	//top of bubble
	bubble.WriteString(" "+strings.Repeat("-", maxLen+2)+"\n")
	
	for _, msg := range messages{
		lineRunes := []rune(msg)
		padding := maxLen - len(lineRunes)

		bubble.WriteString("< ")

		for i, char:= range lineRunes{
			bubble.WriteString(getColour(i,len(lineRunes)))
			bubble.WriteRune(char)
		}

		bubble.WriteString("\x1b[0m")
		bubble.WriteString(strings.Repeat(" ",padding)+" >\n")
	}

	//bottom of bubble
	bubble.WriteString(" "+strings.Repeat("-",maxLen+2)+"\n")

	return bubble.String() + cowArt

}

