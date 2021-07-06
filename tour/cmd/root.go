package root

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Comman{}
	rootCmd.AddCommand(sqlCmd)
}

func Execute() error {
	return rootCmd.Execute()
}