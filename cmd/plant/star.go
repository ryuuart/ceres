package plant

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var StarCmd = &cobra.Command{
	Use:   "star",
	Short: "Creates a new star system in REI",
	Long:  `Creates a new directory that's a star type. A star type directory is a project in REI.`,
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		uuid := strings.ToLower(uuid.NewString())
		date := time.Now().Format("01-02-2006")

		// get the name
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("What's the name of the star?").Value(&name),
			),
		)

		err := form.Run()
		if err != nil {
			log.Fatal(err)
		} else {
			// Format the name
			re, _ := regexp.Compile(`[^\w\s-]`)
			name = re.ReplaceAllString(name, "")

			re, _ = regexp.Compile(`\s+`)
			name = re.ReplaceAllString(name, "-")
			name = strings.ToLower(name)

			formatted_name := fmt.Sprintf("%s--star--%s--%s", name, date, uuid)
			err = os.Mkdir(formatted_name, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}
