package main

import (
	"fmt"
	"os"

	"github.com/kballard/go-shellquote"

	"github.com/bitrise-io/go-utils/colorstring"
	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/bitrise-tools/go-steputils/input"
)

func main() {
	//Validate inputs
	filePath := os.Getenv("file_path")
	if absPath, err := pathutil.AbsPath(filePath); err != nil {
		fmt.Printf("file_path: %s", err.Error())
		os.Exit(1)
	} else {
		filePath = absPath
	}
	if err := input.ValidateIfPathExists(filePath); err != nil {
		fmt.Printf("file_path: %s", err.Error())
		os.Exit(1)
	}
	runner := os.Getenv("runner")
	if err := input.ValidateIfNotEmpty(runner); err != nil {
		fmt.Printf("runner: %s", err.Error())
		os.Exit(1)
	}
	workingDir := os.Getenv("working_dir")
	if err := input.ValidateIfPathExists(workingDir); err != nil {
		fmt.Printf("working_dir: %s", err.Error())
		os.Exit(1)
	}

	exitCode, err := runScript(runner, filePath, workingDir)
	if err != nil {
		prettyError := fmt.Sprintf(
			`Your script returned an error. Check the output above for the root cause.
Additional details:
Script: \t%s
Working directory: \t%s
Exit code: \t%s
Error: \t%s
`,
			colorstring.Cyan(filePath),
			colorstring.Cyan(workingDir),
			colorstring.Red(fmt.Sprintf("%d", exitCode)),
			colorstring.Red(err.Error()),
		)
		fmt.Println("-------------------") // Separate streamed stdout/stderr from the step's error message
		fmt.Println(prettyError)
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
		return 1, fmt.Errorf("shell quote split error: %s", err.Error())
	}

	paramList = append(paramList, filePath)

	script := command.NewWithStandardOuts(binary, paramList...).SetDir(workingDir)

	return script.RunAndReturnExitCode()
}
