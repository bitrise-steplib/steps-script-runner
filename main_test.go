package main

import (
	"testing"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/stretchr/testify/require"
)

func Test_runScriptBash(t *testing.T) {
	require.NoError(t, pathutil.EnsureDirExist("_tmp"))

	t.Log("Successful execution")
	{
		require.NoError(t, fileutil.WriteStringToFile("_tmp/test.sh", "echo 'This is a Bash script'"))

		exitCode, err := runScript("bash", "_tmp/test.sh")
		require.Equal(t, 0, exitCode)
		require.Equal(t, nil, err)
	}

	t.Log("Exit with code 222")
	{
		require.NoError(t, fileutil.WriteStringToFile("_tmp/test_failing.sh", "exit 222"))

		exitCode, err := runScript("bash", "_tmp/test_failing.sh")
		require.Equal(t, 222, exitCode)
		require.Error(t, err)
	}
}

func Test_runScriptRuby(t *testing.T) {
	require.NoError(t, pathutil.EnsureDirExist("_tmp"))
	require.NoError(t, fileutil.WriteStringToFile("_tmp/test.rb", "puts 'This is a Ruby script'"))

	exitCode, err := runScript("ruby", "_tmp/test.rb")
	require.Equal(t, 0, exitCode)
	require.Equal(t, nil, err)
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

	exitCode, err := runScript("go", "_tmp/test.go")
	require.Equal(t, 0, exitCode)
	require.Equal(t, nil, err)
}
