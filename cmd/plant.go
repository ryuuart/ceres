package cmd

import (
	"github.com/ryuuart/ceres/cmd/plant"
	"github.com/spf13/cobra"
)

var plantCmd = &cobra.Command{
	Use:   "plant",
	Short: "Creates a new plant in the garden",
	Long:  `Creates a new formatted directory in REI`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(plantCmd)
	plantCmd.AddCommand(plant.StarCmd)
}
