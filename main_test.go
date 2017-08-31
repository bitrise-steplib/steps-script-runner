package main

import (
	"os"
	"testing"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/stretchr/testify/require"
)

var workDir = os.Getenv("BITRISE_SOURCE_DIR")

func Test_runScriptBash(t *testing.T) {
	require.NoError(t, pathutil.EnsureDirExist("_tmp"))

	t.Log("Successful execution")
	{
		require.NoError(t, fileutil.WriteStringToFile("_tmp/test.sh", "echo 'This is a Bash script'"))

		exitCode, err := runScript("bash", "_tmp/test.sh", workDir)
		require.NoError(t, err)
		require.Equal(t, 0, exitCode)
	}

	t.Log("Exit with code 222")
	{
		require.NoError(t, fileutil.WriteStringToFile("_tmp/test_failing.sh", "exit 222"))

		exitCode, err := runScript("bash", "_tmp/test_failing.sh", workDir)
		require.Equal(t, 222, exitCode)
		require.Error(t, err)
	}
}

func Test_runScriptRuby(t *testing.T) {
	require.NoError(t, pathutil.EnsureDirExist("_tmp"))

	t.Log("Successful Ruby execution")
	{
		require.NoError(t, fileutil.WriteStringToFile("_tmp/test.rb", "puts 'This is a Ruby script'"))

		exitCode, err := runScript("ruby", "_tmp/test.rb", workDir)
		require.NoError(t, err)
		require.Equal(t, 0, exitCode)
	}

	t.Log("Check if working_dir is set properly")
	{
		require.NoError(t, fileutil.WriteStringToFile("_tmp/test_workdir.rb", "puts IO.read(\".gitignore\")"))

		exitCode, err := runScript("ruby", "_tmp/test_workdir.rb", workDir)
		require.NoError(t, err)
		require.Equal(t, 0, exitCode)
	}
}

func Test_runScriptGo(t *testing.T) {
	require.NoError(t, pathutil.EnsureDirExist("_tmp"))
	goScript := `package main
		
import (
	"fmt"
)

func main() {
	fmt.Println("This is a Go script")
}
	`
	require.NoError(t, fileutil.WriteStringToFile("_tmp/test.go", goScript))

	exitCode, err := runScript("go", "_tmp/test.go", workDir)
	require.NoError(t, err)
	require.Equal(t, 0, exitCode)
}
