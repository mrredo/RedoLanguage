package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "redolang",
	Short: "RedoLanguage - a simple CLI to build, run RedoLanguage code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Do you want to do something or else")
	},
}
var build = &cobra.Command{
	Use: "build",
	//Aliases: []string{""},
	Short: "parses RedoLanguage code to a json file in AST",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("build command", args)
	},
}
var run = &cobra.Command{
	Use: "run",
	//Aliases: []string{""},
	Short: "Runs redolanguage code",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("run command", args)
	},
}

func Execute() {
	rootCmd.AddCommand(build, run)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an err while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
