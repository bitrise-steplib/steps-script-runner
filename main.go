package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-tools/go-steputils/input"
)

func main() {
	//Validate inputs
	if err := input.ValidateIfPathExists(os.Getenv("file_path")); err != nil {
		fmt.Printf("file_path: %s", err.Error())
		os.Exit(1)
	}
	if err := input.ValidateWithOptions(os.Getenv("runner_bin"), "bash", "ruby", "go"); err != nil {
		fmt.Printf("runner_bin: %s", err.Error())
		os.Exit(1)
	}

	exitCode, err := runScript(os.Getenv("runner_bin"), os.Getenv("file_path"))

	if err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(exitCode)
}

func runScript(runnerBin string, filePath string) (int, error) {
	paramList := []string{}

	if runnerBin == "go" {
		paramList = append(paramList, "run")
	}

	paramList = append(paramList, filePath)

	script := command.NewWithStandardOuts(runnerBin, paramList...)

	exitCode, err := script.RunAndReturnExitCode()

	if err != nil {
		errorString := fmt.Sprintf("Error: %s, command: %#v", err.Error(), script.GetCmd())

		if exitCode == 0 {
			return 1, errors.New(errorString)
		}

		return exitCode, errors.New(errorString)
	}

	return exitCode, nil
}
