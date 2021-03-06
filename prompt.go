package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

func createprompt(label string, choices []language) language {
	template := &promptui.SelectTemplates{
		Label:    "{{ . | cyan }}",
		Active:   "> {{ .Name | yellow }}",
		Inactive: "{{ .Name | white }}",
		Selected: "{{ .Name | green }}",
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     choices,
		Templates: template,
		HideHelp:  true,
	}

	i, _, _ := prompt.Run()

	return choices[i]
}

func addlanguageprompt() {
	scanner := bufio.NewScanner(os.Stdin)
	// Prompt user asking which category they would like to add their language to
	devType := createprompt("Choose category to add your language to:", []language{{Name: "Front End Web", Url: ""}, {Name: "Back End Web", Url: ""}})

	// Prompt and store for the NAME of their language
	time.Sleep(1 * time.Second)
	fmt.Print("Name of language: ")
	time.Sleep(1 * time.Second)
	scanner.Scan()
	languagename := scanner.Text()
	// Prompt and store for the URL to the documentation
	fmt.Print("URL to the language documentation: ")
	time.Sleep(1 * time.Second)
	scanner.Scan()
	languageurl := scanner.Text()

	// Create new language object
	new := &language{
		Name:     string(languagename),
		Url:      string(languageurl),
		Category: devType.Name,
	}
	newJson, _ := json.Marshal(new)
	filename := "languages/" + languagename + ".json"

	// Write new language to text file
	if _, existsErr := os.Stat("languages/"); errors.Is(existsErr, os.ErrNotExist) {
		os.Mkdir("languages", os.ModePerm)
	}
	if _, existsErr := os.Stat(filename); errors.Is(existsErr, os.ErrNotExist) {
		ioutil.WriteFile(filename, newJson, 0644)
	}
}
