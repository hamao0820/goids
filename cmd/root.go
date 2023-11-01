package cmd

import (
	"goids/gui"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goids",
	Short: "gopher boids flocking algorithm animation",
	Long:  `gopher boids flocking algorithm animation.`,
	Run: func(cmd *cobra.Command, args []string) {
		gui.Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
