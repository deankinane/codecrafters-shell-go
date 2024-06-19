package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		in, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. ", err.Error())
		}

		in = strings.Trim(in, "\r\n ")
		fmt.Fprint(os.Stdout, in+": command not found\r\n")
	}
}
