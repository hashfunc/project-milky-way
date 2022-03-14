package main

import (
	"log"

	"github.com/hashfunc/project-milky-way/internal/setup"
)

func main() {
	files, err := setup.GetCsvFiles()

	if err != nil {
		log.Fatal(err)
	}

	stars, err := setup.GetStarsFromFiles(files)

	if err != nil {
		log.Fatal(err)
	}

	if err := setup.Database(stars); err != nil {
		log.Fatal(err)
	}
}
