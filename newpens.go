package main

import "fmt"

func main() {

	penCase := []Pen{}

	penCase = append(penCase, Pen{Fullness: 80})
	penCase = append(penCase, Pen{Fullness: 0})
	penCase = append(penCase, Pen{Fullness: 20})

	if caseHasNotEmptyPen(penCase) {
		fmt.Println("В пенале есть пишущая ручка.")
	} else {
		fmt.Println("В пенале нет пишущих ручек.")
	}
}

func caseHasNotEmptyPen(penCase []Pen) bool {
	for i := 0; i < (len(penCase)); i++ {
		pen := penCase[i]
		if pen.Fullness > 0 {
			return true
		}
	}
	return false
}

type Pen struct {
	Fullness int
}

func (pen Pen) IsEmpty() bool {
	return pen.Fullness == 0
}
