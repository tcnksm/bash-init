package main

import (
	"fmt"
	"os"
	"text/template"
)

// Source has besic information for generate artifact.
type Source struct {
	Name     string
	Template template.Template
}

func (s Source) generate(definition application) error {
	wr, err := os.Create(s.Name)
	if err != nil {
		return err
	}

	defer wr.Close()
	return s.Template.Execute(wr, definition)
}

// safeRemove remove file and if file is not exit, return false
func (s Source) safeRemove(force bool) (bool, error) {

	if _, err := os.Stat(s.Name); err == nil && force {
		err = os.Remove(s.Name)
		return true, err
	}

	if _, err := os.Stat(s.Name); err == nil {
		fmt.Fprintf(os.Stderr, "%s is already exists, overwrite it? [Y/n]: ", s.Name)
		var ans string
		_, err := fmt.Scanf("%s", &ans)

		if err != nil {
			return false, err
		}

		if ans == "Y" {
			err = os.Remove(s.Name)
			return true, err
		} else {
			return false, err
		}
	}

	return true, nil
}
