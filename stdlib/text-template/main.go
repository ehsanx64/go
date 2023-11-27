package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Gender  string
	CanHunt bool
	Skills  []string
}

const personsTpl = "persons.tmpl"

// Define a shortcut for quick panics ;)
func abort(err error) {
	if err != nil {
		panic(err)
	}
}

// Define a function so data can be moved outta code's face
func getPersons() []Person {
	return []Person{
		{
			Name:    "Adam",
			Age:     30,
			Gender:  "Male",
			CanHunt: true,
			Skills:  []string{"hunting", "foraging", "building", "loving"},
		},
		{
			Name:    "Eve",
			Age:     30,
			Gender:  "Female",
			CanHunt: false,
			Skills:  []string{"cooking", "foraging", "healing", "loving"},
		},
		{
			Name:    "John",
			Age:     3,
			Gender:  "Male",
			CanHunt: false,
			Skills:  []string{"playing", "loving"},
		},
		{
			Name:    "Alan",
			Age:     1,
			Gender:  "Male",
			CanHunt: false,
			Skills:  []string{},
		},
	}
}

func main() {
	funcMap := template.FuncMap{
		"getFormattedSkills": func(skills []string) string {
			var out string
			lastIndex := len(skills) - 1

			if lastIndex < 0 {
				out = "No skills to be listed"
				return out
			}

			for i, val := range skills {
				if lastIndex != i {
					out += val + ", "
				} else {
					out += val
				}
			}

			return out
		},
		"getMainSkill": func(skills []string) string {
			if len(skills) == 0 {
				return "No skill at all!"
			}

			return skills[0]
		},
		"getLastSkill": func(skills []string) string {
			if len(skills) == 0 {
				return "No skill at all!"
			}

			return skills[len(skills)-1]
		},
		"addTwo": func(num int) int {
			return num + 2
		},
		"getNameLen": func(name string) string {
			return fmt.Sprintf("%s is %2d characters long.", name, len(name))
		},
		"dec": func(i int) int {
			if i == 0 {
				return 0
			}
			return i - 1
		},
	}

	// Load the template
	tpl, err := template.New(personsTpl).Funcs(funcMap).ParseFiles(personsTpl)
	abort(err)

	// Get the persons data
	persons := getPersons()

	// Feed the data into the template engine
	err = tpl.Execute(os.Stdout, persons)
	abort(err)
}
