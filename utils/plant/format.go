package plant

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

func IsValidPlantType(plantType string) bool {
	plantType = strings.ToLower(plantType)
	switch plantType {
	case "sphere":
		fallthrough
	case "star":
		return true
	default:
		return false
	}
}

func FormatPlantName(name string, plantType string) string {
	plantType = strings.ToLower(plantType)
	uuid := strings.ToLower(uuid.NewString())
	date := time.Now().Format("01-02-2006")

	// Format the name
	re, _ := regexp.Compile(`[^\w\s-]`)
	name = re.ReplaceAllString(name, "")

	re, _ = regexp.Compile(`\s+`)
	name = re.ReplaceAllString(name, "-")
	name = strings.ToLower(name)

	return fmt.Sprintf("%s--%s--%s--%s", name, plantType, date, uuid)
}
