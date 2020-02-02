package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type menuitem struct {
	Name   string
	Method func() bool
}

func main() {
	menuitems := []menuitem{
		{"Addition upp 10", additionUppTill10},
		{"Addition upp 20", func() bool {
			fmt.Println("Addition upp till 20")
			return true
		}},
		{"Subtraktion upp 20", func() bool {
			fmt.Println("Subtraktion upp till 10")
			return true
		}},
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
