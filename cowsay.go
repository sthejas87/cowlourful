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

func wrapText(text string, width int) []string {
	var lines []string
	words := strings.Fields(text)
	if len(words)==0{
		return []string{""}
	}

	currLine := words[0]
	for i:=1;i<len(words);i++{
		word := words[i]
		if len(currLine)+1+len(word)<=width{
		currLine += " " + word
		}else{
			lines = append(lines, currLine)
			currLine = word
		}
	}
	lines = append(lines, currLine)
	return lines
}

func cowlourful (messages []string, width int) string {

	if len(messages) == 0{
		return cowArt
	}
	
	var wrappedLines []string
	for _, msg := range messages{
		//subtract 4 for '< ' and ' >'
		wrappedLines = append(wrappedLines, wrapText(msg, width-4)...)
	}
	var bubble strings.Builder

	//top of bubble
	bubble.WriteString(" "+strings.Repeat("-", width-2)+"\n")
	
	for _, line := range wrappedLines{
		
		padding := width-4-len([]rune(line))
		bubble.WriteString("< ")

		for j, char:= range []rune(line){
			bubble.WriteString(getColour(j, len([]rune(line))))
			bubble.WriteRune(char)
		}

		bubble.WriteString("\x1b[0m")
		bubble.WriteString(strings.Repeat(" ",padding)+" >\n")
	}

	//bottom of bubble
	bubble.WriteString(" "+strings.Repeat("-",width-2)+"\n")

	return bubble.String() + cowArt

}

