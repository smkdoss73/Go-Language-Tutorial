package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buff := bytes.NewBufferString("life is an empty dream so enjoy every second im doing like this")
	for {
		line, err := buff.ReadString('m')
		// println("Masterline", line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("eof: ", line)
				break
			}
			fmt.Println("error: ", err)
			break
		}
		fmt.Println("Line: ", line)
	}
}
