package cmd

import (
	"github.com/sorahashiroi/pomodomu/pkg"
	"github.com/spf13/cobra"
)

var focusDuration int
var breakDuration int
var longBreakDuration int
var cycles int

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new Pomodoro session",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.RunPomodoro(focusDuration, breakDuration, longBreakDuration, cycles)
	},
}

func init() {
	// フラグの設定
	startCmd.Flags().IntVarP(&focusDuration, "focus", "f", 25, "Focus duration in minutes")
	startCmd.Flags().IntVarP(&breakDuration, "break", "b", 5, "Short break duration in minutes")
	startCmd.Flags().IntVarP(&longBreakDuration, "long-break", "l", 15, "Long break duration in minutes")
	startCmd.Flags().IntVarP(&cycles, "cycles", "c", 4, "Number of focus/break cycles")
}
