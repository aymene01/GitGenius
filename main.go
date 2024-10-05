package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/huh"
)

var (
	options = []string{}
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Choose what you want").
				Options(
					huh.NewOption("Commit name", "Commit name"),
					huh.NewOption("Branch name", "Branch name"),
					huh.NewOption("Pull Request name", "Pull Request name"),
				).
				Limit(4).
				Value(&options),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("opts: ", strings.Join(options, ", "))
}
