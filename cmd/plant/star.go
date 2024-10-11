package plant

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var StarCmd = &cobra.Command{
	Use:   "star",
	Short: "Creates a new star system in REI",
	Long:  `Creates a new directory that's a star type. A star type directory is a project in REI.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args[0]) == 0 {
			log.Fatal("You have to provide the star's name")
		} else {
			isWrongFormat, _ := regexp.MatchString(`(^[-\s]+.*)|(.*[-\s]+$)|(-(-)+)`, args[0])

			if isWrongFormat {
				log.Fatal(`The format is wrong. Your name can't be blank, start, or end with one or more "-" or spaces. Don't include consecutive "-" too.`)
			}
		}
		name := args[0]
		uuid := strings.ToLower(uuid.NewString())
		date := time.Now().Format("01-02-2006")

		// Format the name
		re, _ := regexp.Compile(`[^\w\s-]`)
		name = re.ReplaceAllString(name, "")

		re, _ = regexp.Compile(`\s+`)
		name = re.ReplaceAllString(name, "-")
		name = strings.ToLower(name)

		formatted_name := fmt.Sprintf("%s--star--%s--%s", name, date, uuid)

		// Create the folder
		err := os.Mkdir(formatted_name, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	},
}
