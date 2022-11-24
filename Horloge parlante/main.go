package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Nous sommes le %d %s, il est %dh%d", time.Now().Day(), getMonthName(int(time.Now().Month())), time.Now().Hour(), time.Now().Minute())
}

func getMonthName(month int) string {
	switch month {
	case 1:
		return "Janvier"
	case 2:
		return "Février"
	case 3:
		return "Mars"
	case 4:
		return "Avril"
	case 5:
		return "Mai"
	case 6:
		return "Juin"
	case 7:
		return "Juillet"
	case 8:
		return "Août"
	case 9:
		return "Septembre"
	case 10:
		return "Octobre"
	case 11:
		return "Novembre"
	case 12:
		return "Décembre"
	default:
		return "Inconnu"
	}
}
