package cmd

import (
    "bufio"
    "fmt"
    "os"
		"os/exec"
    "strings"

    "github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
    Use: "new [name]",
    Run: func(cmd *cobra.Command, args []string) {
        var name string
        if len(args) > 0 {
            name = args[0]
        } else {
						//if name isn't provied take user input for name
            reader := bufio.NewReader(os.Stdin)
            fmt.Print("Enter app name: ")
            name, _ = reader.ReadString('\n')

						// remove the newline character
            name = strings.TrimSpace(name)
        }
				//create dir in current dir with name of [name] as the parent dir for backend and frontend
				//within the parent dir, init a git repo
				//create a folder named frontend for next.js app
				//create a folder called backend for gin app
				
				// Create directories
        os.MkdirAll(name+"/frontend", os.ModePerm)
        os.MkdirAll(name+"/backend", os.ModePerm)

        // Initialize git repo in the new directory
        os.Chdir(name)
        gitCmd := exec.Command("git", "init")
        err := gitCmd.Run()
        if err != nil {
            fmt.Println("Failed to initialize git repository:", err)
            return
        }

				//init nextjs app w/ default settings 
				os.Chdir("./frontend")

				//Check if create-next-app is installed
				checkNextCmd := exec.Command("npx", "create-next-app", "--version")
				if err := checkNextCmd.Run(); err != nil {
						//If not installed, install it
						installCmd := exec.Command("npm", "install", "-g", "create-next-app")
						if err := installCmd.Run(); err != nil {
								fmt.Println("Failed to install create-next-app:", err)
								return
						}
				}

				initNextjs := exec.Command("npx", "create-next-app", ".", "--typescript", "--tailwind", "--eslint", "--app", "--use-npm")
        err = initNextjs.Run()
				if err != nil {
					fmt.Println("Failed to initialize git repository:", err)
            return
				}

        // Change back to the original directory
        os.Chdir("..")

				//init gin app
				os.Chdir("./backend")

				//check is go is installed
				checkGoCmd := exec.Command("go", "version")
				if err := checkGoCmd.Run(); err != nil {
					fmt.Println("Go isn't installed: ", err)
					return
				}

				//init go mod
				initGoMod := exec.Command("go", "mod", "init", name)
				err = initGoMod.Run()
				if err != nil {
					fmt.Println("failed to init go mod: ", err)
				}
				// Create main.go file
				file, err := os.Create("main.go")
				if err != nil {
					fmt.Println("Failed to create main.go file:", err)
					return
				}
				defer file.Close()

				// Write initial content to main.go
				//below code is an exmaple. to be written.
				initialContent := `package main

				import (
					"fmt"
				)

				func main() {
					fmt.Println("Hello, world!")
				}
				`
				_, err = file.WriteString(initialContent)
				if err != nil {
					fmt.Println("Failed to write initial content to main.go:", err)
					return
				}
    },
}