package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/bolinette/bolinette-cli/generator"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var questions = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "First name your app: "},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "app",
		Prompt: &survey.Select{
			Message: "What kind of app do you want to build ?",
			Options: []string{"A simple bolinette API"},
		},
	},
	{
		Name: "database",
		Prompt: &survey.Select{
			Message: "What type of database do you want ?",
			Options: []string{
				"SQLITE",
				"MySql",
				"MariaDB",
				"PostgreSQL",
			},
		},
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a Bolinette application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initBolinetteAndVenv()
		getBolinetteVersion()
		getCliVersion()
		welcomeMessage :=
			`
 ▄▄▄▄    ▒█████   ██▓     ██▓ ███▄    █ ▓█████▄▄▄█████▓▄▄▄█████▓▓█████ 
 ▓█████▄ ▒██▒  ██▒▓██▒    ▓██▒ ██ ▀█   █ ▓█   ▀▓  ██▒ ▓▒▓  ██▒ ▓▒▓█   ▀ 
 ▒██▒ ▄██▒██░  ██▒▒██░    ▒██▒▓██  ▀█ ██▒▒███  ▒ ▓██░ ▒░▒ ▓██░ ▒░▒███   
 ▒██░█▀  ▒██   ██░▒██░    ░██░▓██▒  ▐▌██▒▒▓█  ▄░ ▓██▓ ░ ░ ▓██▓ ░ ▒▓█  ▄ 
 ░▓█  ▀█▓░ ████▓▒░░██████▒░██░▒██░   ▓██░░▒████▒ ▒██▒ ░   ▒██▒ ░ ░▒████▒
 ░▒▓███▀▒░ ▒░▒░▒░ ░ ▒░▓  ░░▓  ░ ▒░   ▒ ▒ ░░ ▒░ ░ ▒ ░░     ▒ ░░   ░░ ▒░ ░
 ▒░▒   ░   ░ ▒ ▒░ ░ ░ ▒  ░ ▒ ░░ ░░   ░ ▒░ ░ ░  ░   ░        ░     ░ ░  ░
  ░    ░ ░ ░ ░ ▒    ░ ░    ▒ ░   ░   ░ ░    ░    ░        ░         ░   
  ░          ░ ░      ░  ░ ░           ░    ░  ░                    ░  ░
	   ░                                                                
																		
																		
																		
																		
																		
																		
																		
																		
																		
																		
		`
		fmt.Printf("%s\nBolinette version: %s, CLI version: %s\n", welcomeMessage, bolinettVersion, cliVersion)
		answers := struct {
			Name     string
			AppType  string `survey:"app"`
			Database string
		}{}
		err := survey.Ask(questions, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if strings.Contains(answers.AppType, "A simple bolinette API") {
			_, err := os.Stat(answers.Name)
			if os.IsNotExist(err) {
				generator.GenerateHeadlessBolinetteApi(answers.Name, answers.Database)
			} else {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("Your already have a folder %s in your current directory.", answers.Name))
			}
		} else {
			fmt.Println("Error processing response")
			fmt.Println("Exiting...")
		}
		os.Exit(1)
	},
}

func init() {
	blntCmd.AddCommand(initCmd)
}

func initBolinetteAndVenv() {
	fmt.Print("Initialising a virtualenv and installing Bolinette...")
	switch runtime.GOOS {
	case "windows":
		fmt.Println("To be implemented")
	default:
		cmd := exec.Command("bash", "-c", "python -m venv venv")
		_, err := cmd.CombinedOutput()
		if err != nil {
			cmd = exec.Command("bash", "-c", "python3 -m venv venv")
			_, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Cannot find a version of Python installed.")
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
	cmd := exec.Command("bash", "-c", "source venv/bin/activate && pip install pip --upgrade && pip install bolinette")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("done")
}
