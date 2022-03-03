package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"strings"
)

func main() {
	action := flag.String("action", "get", "Type of action when running the program.")
	flag.Parse()

	switch strings.ToLower(*action) {
	case "get":
		devType := createprompt("Choose Category:", []language{{Name: "Front End Web", Url: "", Category: "Front End Web"}, {Name: "Back End Web", Url: "", Category: "Back End Web"}})
		stored := getStoredLanguages()
		temp := []language{}
		for i := range stored {
			if stored[i].Category == devType.Category {
				temp = append(temp, stored[i])
			}
		}
		lang := createprompt("Choose Language:", temp)
		openbrowser(lang.Url)
	case "add":
		addlanguageprompt()
	}
}

func getStoredLanguages() []language {
	folder, _ := ioutil.ReadDir("languages/")
	var languges []language
	for _, item := range folder {
		var lang language
		data, _ := ioutil.ReadFile("languages/" + item.Name())
		json.Unmarshal(data, &lang)
		languges = append(languges, lang)
	}
	return languges
}
