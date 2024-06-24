package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		in, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		in = strings.Trim(in, "\r\n ")

		if len(in) == 0 {
			continue
		}

		args := strings.Fields(strings.Trim(in, "\r\n "))
		err := handle_args(args)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func handle_args(args []string) error {
	var err error
	for i := 0; i < len(args); i++ {
		com := strings.ToLower(args[i])

		switch com {
		case CommandExit:
			i += 1
			err = command_exit(args[i])
		default:
			err = errors.New(com + ": command not found")
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func command_exit(code string) error {
	int_code, err := strconv.Atoi(code)
	if err != nil {
		return errors.New("invalid exit code")
	}

	os.Exit(int_code)
	return nil
}

const (
	CommandExit = "exit"
)
