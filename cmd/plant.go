package cmd

import (
	"log"
	"os"
	"regexp"

	"github.com/ryuuart/ceres/utils/plant"
	"github.com/spf13/cobra"
)

var plantType string

func init() {
	rootCmd.AddCommand(plantCmd)
	plantCmd.Flags().StringVarP(&plantType, "type", "t", "sphere", "Type of plant that'll be grown")
	plantCmd.MarkFlagRequired("type")
}

var plantCmd = &cobra.Command{
	Use:   "plant",
	Short: "Creates a new plant in the garden",
	Long:  `Creates a new formatted directory in REI`,
	Run: func(cmd *cobra.Command, args []string) {
		// Validate commands and arguments
		if !plant.IsValidPlantType(plantType) {
			log.Fatal(`You can only have a "sphere" or "star" plant type currently`)
		}

		if len(args) == 0 || len(args[0]) == 0 {
			log.Fatalf("You have to provide the %s's name", plantType)
		} else {
			isWrongFormat, _ := regexp.MatchString(`(^[-\s]+.*)|(.*[-\s]+$)|(-(-)+)`, args[0])

			if isWrongFormat {
				log.Fatal(`The format is wrong. Your name can't be blank, start, or end with one or more "-" or spaces. Don't include consecutive "-" too.`)
			}
		}

		// Create the folder
		formatted_name := plant.FormatPlantName(args[0], plantType)
		err := os.Mkdir(formatted_name, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	},
}
