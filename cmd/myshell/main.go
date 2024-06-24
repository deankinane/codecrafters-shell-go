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
	com := strings.ToLower(args[0])

	switch com {
	case CommandExit:
		err = command_exit(args)
	case CommandEcho:
		command_echo(args)
		return nil
	case CommandType:
		err = command_type(args)
	default:
		err = errors.New(com + ": command not found")
	}

	if err != nil {
		return err
	}

	return nil
}

func command_exit(args []string) error {
	if len(args) == 1 {
		return errors.New("no error code provided")
	}
	code := args[1]
	int_code, err := strconv.Atoi(code)
	if err != nil {
		return errors.New("invalid exit code")
	}

	os.Exit(int_code)
	return nil
}

func command_echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func command_type(args []string) error {
	if len(args) == 1 {
		return errors.New("no command provided")
	}
	com := args[1]
	t, ok := CommandTypes[com]
	if !ok {
		fmt.Println(com + ": not found")
		return nil
	}

	switch t {
	case TypeBuiltin:
		fmt.Println(com, "is a shell builtin")
	}

	return nil
}

const (
	CommandExit = "exit"
	CommandEcho = "echo"
	CommandType = "type"
)

const (
	TypeBuiltin = "builtin"
)

var CommandTypes = map[string]string{
	CommandExit: TypeBuiltin,
	CommandEcho: TypeBuiltin,
	CommandType: TypeBuiltin,
}
