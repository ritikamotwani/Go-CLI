package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func readAndWriteFile(directory string, sourceFile string, fileName string) {
	// read
	content, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		panic(err)
	}
	// write
	filePath := directory + "/" + fileName

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create(directory + "/" + fileName)
	} else {
		fmt.Println(fileName, ": File already exists in the home directory")
	}
	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		panic(err)
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project of the given language that prints Hello World",
	Long:  `Create a project of the given language that prints Hello World. 💫 The language is passed as a string in the command.`,
	Run: func(cmd *cobra.Command, args []string) {
		var language = args[0]
		homeDir, _ := os.UserHomeDir()

		switch language {
		case "Python":
			readAndWriteFile(homeDir, "hello_world.py", "hello_world.py")

		case "C++":
			readAndWriteFile(homeDir, "hello-world.txt", "hello-world.cpp")

		case "NodeJS":
			appDir := "/my-app-hello-world"

			nodeJSApp := homeDir + appDir
			if _, err := os.Stat(nodeJSApp); os.IsNotExist(err) {
				os.Mkdir(nodeJSApp, 0700)
			} else {
				fmt.Println("my-app-hello-world named directory already exists")
				panic(err)
			}

			cmd := exec.Command("npm", "init", "-y")
			cmd.Dir = nodeJSApp
			out, err := cmd.Output()
			if err != nil {
				fmt.Print(err)
			}

			cmd1 := exec.Command("npm", "install", "express", "--save")
			cmd1.Dir = nodeJSApp
			out, err = cmd1.Output()
			if err != nil {
				fmt.Print(err)
			}
			fmt.Println(string(out), cmd)

			readAndWriteFile(nodeJSApp, "hello-world.js", "index.js")

		default:
			fmt.Println("Uh oh! We don't support this language yet. 🥕")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
