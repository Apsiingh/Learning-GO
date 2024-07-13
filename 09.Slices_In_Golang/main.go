package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in Go lang")

	trpCustomName := []string{"Nirmal_da", "Pramit_da", "Bhaskar_da"}

	trpCustomName = append(trpCustomName, "Somnath_da", "Avi_da", "Disha", "Sneha", "Akhand")

	fmt.Println(trpCustomName)

	frontName := append(trpCustomName[:3])
	fmt.Println("FrontEnd team name", frontName)

	endName := append(trpCustomName[3:])
	fmt.Println("Backend team name", endName)

	mainName := append(trpCustomName[3:4])
	fmt.Println("Manager of the custom Team", mainName)

	highScores := make([]int, 5)

	highScores[0] = 879
	highScores[1] = 258
	highScores[2] = 147
	highScores[3] = 159
	highScores[4] = 899
	// highScores[5] = 999

	highScores = append(highScores, 5464, 4578, 4587)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

}
