package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
)

var questions = []*survey.Question{
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
				"SQLITE file",
				"MySql",
				"PostgreSQL",
				"None",
			},
		},
	},
}

func init() {
	blntCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a Bolinette application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
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
		// add the cli version and bolinette version
		// check current directory and if bolinette is installed
		fmt.Println(welcomeMessage)
		answers := struct {
			AppType  string `survey:"app"`
			Database string
		}{}
		err := survey.Ask(questions, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}
