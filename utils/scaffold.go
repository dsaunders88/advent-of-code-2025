package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

// TODO: get current day from CLI flag

const (
	TEMPLATE_MAIN string = "days/templates/main.go"
	TEMPLATE_TEST string = "days/templates/main_test.go"
	DAY           string = "01"
)

type Day struct {
	Name string
}

// Create a directory and associated files for each day. Go templates are located in `days/templates` and use template variables for the `text/template` package to parse. Also adds to blank text files for inputs.
func ScaffoldDayTemplates() error {
	templates := []string{TEMPLATE_MAIN, TEMPLATE_TEST}
	currentDay := Day{DAY}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// get/create new day folder
	// get path to days folder
	dirPath := path.Join("days", currentDay.Name)
	err = os.Mkdir(path.Join(wd, dirPath), 0755)
	if err != nil && !os.IsExist(err) {
		return err
	} else if os.IsExist(err) {
		fmt.Printf("directory already exists: %s\n", err)
	}

	// create go templates
	for _, t := range templates {
		// get template data
		data, err := os.ReadFile(path.Join(wd, t))
		if err != nil {
			return err
		}

		// process template
		templ, err := template.New("day").Parse(string(data))
		if err != nil {
			return err
		}

		// create/write file
		pathParts := strings.Split(t, "/")
		filename := pathParts[len(pathParts)-1]
		filepath := path.Join(wd, dirPath, filename)

		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
		if err != nil {
			if os.IsExist(err) {
				return fmt.Errorf("file %s already exists", filepath)
			} else {
				return err
			}
		}

		defer file.Close()

		// execute template with writer (`file`) and data (`currentDay`)
		err = templ.Execute(file, currentDay)
		if err != nil {
			return err
		}
		fmt.Printf("main template file written to %s\n", filepath)
	}

	// create empty text files
	txt, err := os.Create(path.Join(wd, dirPath, "input.txt"))
	if err != nil {
		return err
	}
	defer txt.Close()
	txt.Write(nil)

	txt, err = os.Create(path.Join(wd, dirPath, "input-test.txt"))
	if err != nil {
		return err
	}
	defer txt.Close()
	txt.Write(nil)

	return nil
}
