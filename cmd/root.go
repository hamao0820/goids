package cmd

import (
	"fmt"
	"goids/gui"
	"os"

	"github.com/spf13/cobra"
)

var width, height int
var n int
var speed float64

var rootCmd = &cobra.Command{
	Use:   "goids",
	Short: "gopher boids flocking algorithm animation",
	Long:  `gopher boids flocking algorithm animation.`,
	Run: func(cmd *cobra.Command, args []string) {
		if width < 0 || height < 0 {
			fmt.Println("width and height must be positive")
			os.Exit(1)
		}
		if n < 0 {
			fmt.Println("number of goids must be positive")
			os.Exit(1)
		}
		if speed < 0 {
			fmt.Println("max speed must be positive")
			os.Exit(1)
		}
		fmt.Printf("%d\n", n)
		gui.Run(width, height, n, speed)
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
	rootCmd.Flags().IntVarP(&width, "width", "w", 640, "width of the window")
	rootCmd.Flags().IntVar(&height, "height", 480, "height of the window")
	rootCmd.Flags().IntVarP(&n, "number", "n", 30, "number of gopher")
	rootCmd.Flags().Float64VarP(&speed, "speed", "s", 3, "max speed of the gopher")
}
