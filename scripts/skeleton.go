package scripts

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/SuperDulli/advent-of-code/util"
)

//go:embed templates/*.go
var fs embed.FS

func CreateSkeleton(day, year int) {
	if day > 25 || day <= 0 {
		log.Fatalf("invalid day value, must be 1 through 25, got %v", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	templates, err := template.ParseFS(fs, "templates/*.go")
	if err != nil {
		log.Fatalf("parsing templates directory: %s", err)
	}

	mainFilename := filepath.Join(util.DirectoryName(), fmt.Sprintf("%d/day%02d/main.go", year, day))
	testFilename := filepath.Join(util.DirectoryName(), fmt.Sprintf("%d/day%02d/main_test.go", year, day))
	fmt.Printf("main file: %s\n", mainFilename)

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	ensureNotOverwriting(mainFilename)
	ensureNotOverwriting(testFilename)

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatalf("creating main.go file: %v", err)
	}

	testFile, err := os.Create(testFilename)
	if err != nil {
		log.Fatalf("creating main_test.go file: %v", err)
	}

	templates.ExecuteTemplate(mainFile, "main.go", nil)
	templates.ExecuteTemplate(testFile, "main_test.go", nil)
	fmt.Printf("templates made for %d-day%d\n", year, day)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("File already exists: %s", filename)
	}
}
