package cmd

import (
	"fmt"
	"os"

	"github.com/shunsukehamada/goids/gui"

	"github.com/spf13/cobra"
)

var width, height int
var n int
var speed float64
var force float64
var fullScreen bool
var sight float64

var rootCmd = &cobra.Command{
	Use:     "goids",
	Version: "1.0.0",
	Short:   "Run a Boids Flocking animation with a Gopher in a GUI.",
	Long: `Run a Boids Flocking animation with a Gopher in a GUI.

This CLI application allows you to run a captivating Boids Flocking animation featuring our beloved Gopher in a graphical user interface (GUI).
The animation simulates the collective behavior of Gopher-like creatures, following the Boids Flocking algorithm.
You can experience the mesmerizing movement patterns of Gophers as they interact with one another, all while being displayed in a visually appealing graphical interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		if width <= 0 || height <= 0 {
			fmt.Println("width and height must be positive")
			os.Exit(1)
		}
		if n <= 0 {
			fmt.Println("number of gopher must be positive")
			os.Exit(1)
		}
		if speed < 0 {
			fmt.Println("max speed must be non negative")
			os.Exit(1)
		}
		if force < 0 {
			fmt.Println("max force must be non negative")
			os.Exit(1)
		}
		if sight < 0 {
			fmt.Println("sight must be non negative")
			os.Exit(1)
		}

		gui.Run(width, height, n, speed, force, sight, fullScreen)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&width, "width", "w", 640, "width of the window")
	rootCmd.Flags().IntVar(&height, "height", 480, "height of the window")
	rootCmd.Flags().IntVarP(&n, "number", "n", 30, "number of gopher")
	rootCmd.Flags().Float64VarP(&speed, "speed", "s", 3, "max speed of the gopher")
	rootCmd.Flags().Float64VarP(&force, "force", "f", 2, "max force of the gopher")
	rootCmd.Flags().Float64Var(&sight, "sight", 100, "sight of the gopher")
	rootCmd.Flags().BoolVar(&fullScreen, "full", false, "full screen mode")
}
