package main

import (
	"bufio"
	"os"
	"io"
	"fmt"
)

func main(){
	
	//check whether piped or not
	info,_ := os.Stdin.Stat()
	if info.Mode()&os.ModeCharDevice != 0{
		fmt.Println("The command is intended to work with pipes")
		os.Exit(1)
	}

	var output[] rune

	reader := bufio.NewReader(os.Stdin) // to read input from Stdin
	for{
		input, _, err := reader.ReadRune() // read Rune
		if err!= nil && err == io.EOF {
		break
		}// break when EOF reached
		output = append(output, input)
	}

	fmt.Println(string(output))
}
