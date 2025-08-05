package main

import (
	"bufio"
	"os"
	"io"
	"fmt"
)

func main(){

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
