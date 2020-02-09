package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/manifoldco/promptui"
)

type menuitem struct {
	Name   string
	Method func() bool
}

func main() {
	rand.Seed(time.Now().UnixNano())
	menuitems := []menuitem{
		{"Addition upp 10", func() bool { return additionUppTill(10, 7) }},
		{"Addition upp 20", func() bool { return additionUppTill(20, 7) }},
		{"Addition upp 40", func() bool { return additionUppTill(40, 7) }},
		{"Subtraktion upp 10", func() bool { return subtraktionUppTill(10, 7) }},
		{"Subtraktion upp 20", func() bool { return subtraktionUppTill(20, 7) }},
		{"10-kompisar", func() bool { return tioKompisar() }},
		{"Avsluta", func() bool {
			fmt.Println("Hejdå")
			return false
		}},
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Name }}?",
		Active:   "> {{ .Name | green }} <", // a rose: \U0001F337
		Inactive: "  {{ .Name | white }}",
		Selected: "{{ .Name | red | cyan }}",
		// 		Details: `
		// --------- Pepper ----------
		// {{ "Name:" | faint }}	{{ .Name }}
		// {{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
		// {{ "Peppers:" | faint }}	{{ .Peppers }}`,
	}
	prompt := promptui.Select{
		Label: "Välj svårighet",
		// Items: []string{"Addition upp 20",
		// 	"Subtraktion upp till 17",
		// 	"10-kompisar"},
		Items:     menuitems,
		Templates: templates,
		Size:      len(menuitems),
	}

	for runagain := true; runagain; {
		i, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		//fmt.Printf("You choose %q %d\n", result, i)
		prompt.CursorPos = i
		//fmt.Println(menuitems[i].Name)
		runagain = menuitems[i].Method()
	}

}
