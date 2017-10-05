package main

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-utils/fileutil"

	"github.com/kballard/go-shellquote"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/bitrise-tools/go-steputils/input"
)

var isDebug bool

func main() {
	//Validate inputs
	isDebug = os.Getenv("is_debug") == "yes"

	filePath := os.Getenv("file_path")
	if absPath, err := pathutil.AbsPath(filePath); err != nil {
		fmt.Printf("* file_path: %s", err.Error())
		os.Exit(1)
	} else {
		filePath = absPath
	}
	if err := input.ValidateIfPathExists(filePath); err != nil {
		fmt.Printf("* file_path: %s", err.Error())
		os.Exit(1)
	}
	debugLogf("* file_path: %s", filePath)

	runner := os.Getenv("runner")
	if err := input.ValidateIfNotEmpty(runner); err != nil {
		fmt.Printf("* runner: %s", err.Error())
		os.Exit(1)
	}
	workingDir := os.Getenv("working_dir")
	if err := input.ValidateIfPathExists(workingDir); err != nil {
		fmt.Printf("* working_dir: %s", err.Error())
		os.Exit(1)
	}
	debugLogf("* working_dir: %s", workingDir)

	fileContent, err := fileutil.ReadStringFromFile(filePath)
	if err != nil {
		fmt.Printf("* error: %s", err.Error())
		os.Exit(1)
	}
	debugLogf("* script")
	debugLogf("%s", fileContent)
	debugLogf("* end of script")

	exitCode, err := runScript(runner, filePath, workingDir)
	if err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(exitCode)
}

func runScript(runner string, filePath string, workingDir string) (int, error) {
	var paramList []string
	var binary string

	if splitRunner, err := shellquote.Split(runner); err == nil {
		binary = splitRunner[0]
		paramList = splitRunner[1:]
	} else {
		return 1, fmt.Errorf("Error: %s", err.Error())
	}

	paramList = append(paramList, filePath)

	script := command.NewWithStandardOuts(binary, paramList...).SetStdin(os.Stdin).SetDir(workingDir)

	exitCode, err := script.RunAndReturnExitCode()
	if err != nil {
		retError := fmt.Errorf("error: %s, command: %#v", err.Error(), script.GetCmd())

		if exitCode == 0 {
			return 1, retError
		}

		return exitCode, retError
	}

	return exitCode, nil
}

func debugLogf(format string, a ...interface{}) {
	if isDebug {
		fmt.Println(fmt.Sprintf(format, a...))
	}
}
