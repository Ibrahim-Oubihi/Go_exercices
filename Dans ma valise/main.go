package main

import (
	"fmt"
)

func main() {
	phrase := "Dans ma valise, il y a…"
	mot := []string{}
	temp := ""

	for temp != "fin" {
		fmt.Printf("\nEntrez un mot :")
		fmt.Scan(&temp)
		mot = append(mot, temp)
		fmt.Printf(phrase)
		showPhrase(mot)
		if showPhrase(mot []string) < 4{
			
		} 
	}

}

func showPhrase(mot []string) {
	for i := 0; i < len(mot); i++ {
		fmt.Printf(" %s ", mot[i])
	}
}
