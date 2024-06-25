package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"slices"
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
	case CommandPwd:
		err = command_pwd(args)
	case CommandCd:
		err = command_cd(args)
	default:
		err = command_run(args)
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
		path, ok := command_type_check_path(com)
		if !ok {
			fmt.Println(com + ": not found")
			return nil
		} else {
			fmt.Println(com, "is", path)
			return nil
		}
	}

	switch t {
	case TypeBuiltin:
		fmt.Println(com, "is a shell builtin")
	}

	return nil
}

func command_type_check_path(path string) (string, bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, p := range paths {
		files, _ := os.ReadDir(p)
		idx := slices.IndexFunc(files, func(f fs.DirEntry) bool { return f.Name() == path })
		if idx > -1 {
			return p + "/" + path, true
		}
	}

	return "", false
}

func command_run(args []string) error {
	path := args[0]
	_, ok := command_type_check_path(path)
	if !ok {
		return errors.New(path + ": command not found")
	}
	cmd := exec.Command(path, strings.Join(args[1:], " "))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	return err
}

func command_pwd(args []string) error {
	if len(args) > 1 {
		return errors.New("invalid arguments for pwd command")
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(cwd)
	return nil
}

func command_cd(args []string) error {
	if len(args) > 2 {
		return errors.New("too many arguments")
	}

	if len(args) == 1 {
		return nil
	}

	path := args[1]

	err := os.Chdir(path)
	if err != nil {
		return errors.New(fmt.Sprint("cd: ", path+": ", "No such file or directory"))
	}
	return nil
}

const (
	CommandExit = "exit"
	CommandEcho = "echo"
	CommandType = "type"
	CommandPwd  = "pwd"
	CommandCd   = "cd"
)

const (
	TypeBuiltin = "builtin"
)

var CommandTypes = map[string]string{
	CommandExit: TypeBuiltin,
	CommandEcho: TypeBuiltin,
	CommandType: TypeBuiltin,
	CommandPwd:  TypeBuiltin,
	CommandCd:   TypeBuiltin,
}
