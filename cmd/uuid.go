package cmd

import (
	"fmt"

	"github.com/google/uuid"
)

func PrintUUID() {
	uuid_string := uuid.NewString()
	fmt.Printf("potato potato? %s\n", uuid_string)
	fmt.Printf("length: %d\n", len(uuid_string))
}
