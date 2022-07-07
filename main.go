package main

import "fmt"

type animalsList interface {
	foodWeigh() float64
	howWeigh() float64
	fmt.Stringer
}

func main() {
	monthlyFood := .0
	animals := []animalsList{
		cat{weight: 6.2},
		cat{weight: 3.1},
		cat{weight: 5.7},
		cat{weight: 5.5},
		cat{weight: 4.4},
		dog{weight: 36.4},
		dog{weight: 42},
		dog{weight: 54},
		cow{weight: 679.3},
		cow{weight: 730.7},
		cow{weight: 690.1},
		cow{weight: 709.6},
	}
	for _, ani := range animals {
		fmt.Printf("%v весит: %vкг, ест: %vкг корма в месяц\n", animalsList(ani), ani.howWeigh(), ani.foodWeigh())
		monthlyFood += ani.foodWeigh()
	}
	fmt.Printf("Всего: %vкг корма в месяц", monthlyFood)
}
