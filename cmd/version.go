package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var bolinettVersion string
var cliVersion string

func init() {
	blntCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Bolinette and Bolinette cli",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Bolinette version: %s, CLI version: %s\n", bolinettVersion, cliVersion)
	},
}

func getBolinetteVersion() string {
	var cmd = &exec.Cmd{}
	switch runtime.GOOS {
	case "windows":
		fmt.Println("To be implemented")
	default:
		cmd = exec.Command("bash", "-c", "source venv/bin/activate && pip freeze")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("jnkh")
	freeze := string(out)
	re := regexp.MustCompile(`Bolinette==[\d+.]+`)
	freeze = re.FindString(freeze)
	split := strings.Split(freeze, "==")
	if len(split) < 2 {
		fmt.Fprintln(os.Stderr, "Error when fetching the version of bolinette")
		os.Exit(1)
	}
	bolinettVersion = split[1]
	return bolinettVersion
}

func getCliVersion() string {
	cliVersion = "0.0.1"
	return cliVersion
}
