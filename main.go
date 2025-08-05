package main

import (
	"bufio"
	"os"
	"io"
	"fmt"
	"math"
)

func getColour (index, total int) string {
	
	//change spread and intensity of colours
	freq := 0.3
	phase := 4.0

	//calculate RGB values by shifting sine waves from [-1,1] to [0,255] for full colour range
	r := math.Sin(freq*float64(index)+phase)*127.5 + 127.5
	g := math.Sin(freq*float64(index)+phase+2)*127.5 + 127.5
	b := math.Sin(freq*float64(index)+phase+4)*127.5 + 127.5

	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", int(r), int(g), int(b))
}

func main(){
	
	//check whether piped or not
	info,_ := os.Stdin.Stat()
	if info.Mode()&os.ModeCharDevice != 0{
		fmt.Println("The command is intended to work with pipes")
		os.Exit(1)
	}
	
	RESET := "\x1b[0m"
	
	var output[] rune

	reader := bufio.NewReader(os.Stdin) // to read input from Stdin
	for{
		input, _, err := reader.ReadRune() // read Rune
		if err!= nil && err == io.EOF {
		break
		}// break when EOF reached
		output = append(output, input)
	}

	totalLength := len(output)
	
	for i, char := range output{
		colour := getColour(i, totalLength)
		fmt.Printf("%s%c", colour, char)
	}	

	fmt.Print(RESET)
}
