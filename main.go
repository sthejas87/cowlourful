package main

import (
	"bufio"
	"os"
	"fmt"
	"math"
	"golang.org/x/term"
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
	
	var messages []string

	reader := bufio.NewScanner(os.Stdin) // to read input from Stdin

	for reader.Scan(){
		messages = append(messages, reader.Text())
	}	
	
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err!=nil || width<4 { //min required for bubble
		width = 80
	}

	finalOutput := cowlourful(messages, width)
	fmt.Println(finalOutput)
}
