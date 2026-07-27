/*
Copyright © 2026 Tom Zaspel tom@zaspel.it
*/
package rootCmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "go-git-pull",
	Version: version,
	Short:   "Simple tool to frequently update one or more git repositories",
	Long:    `This simple tool let you update one or more git repositories. Furtheremore you can set an interval to update the defined repositories frequently.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		basePath, err := cmd.Flags().GetString("basePath")
		if err != nil {
			return err
		}
		fmt.Printf("BasePath is :%s\n", basePath)
		return nil
	}}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("help", "h", false, "Help message for for this tool")

	rootCmd.Flags().StringP("basePath", "b", ".", "root directory for searching Git repositories")

	rootCmd.Flags().BoolP("recursive", "r", false, "If specified this tool search recursively.")

	rootCmd.Flags().IntP("internalMin", "i", 10, "The interval of updating the selected Git repositories.")
}
