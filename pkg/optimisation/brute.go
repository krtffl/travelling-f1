package optimisation

import (
	"fmt"

	"github.com/krtffl/travelling-f1/pkg/distance"
	"github.com/krtffl/travelling-f1/pkg/ergast"
)

func BruteForce() {
	season := ergast.GetSeason(2023)
	numOfRaces := len(season.Races)

	comb := make([][]float64, numOfRaces)
	for i := range comb {
		comb[i] = make([]float64, numOfRaces)
	}

	writes := 0
	diagonal := 0

	for i := range comb {
		for j := range comb[i] {
			if i == j {
				comb[i][j] = 0
				diagonal++
				continue
			}

			if j > i {
				continue
			}

			d := distance.DistanceBetweenTwoRaces(season.Races[i], season.Races[j])
			comb[i][j] = d
			comb[j][i] = d
			writes++

		}
	}

	fmt.Println(comb)
	fmt.Println("writes", writes)
	fmt.Println("diagonal", diagonal)
	fmt.Println("number of races", numOfRaces)
	fmt.Println("total elements", numOfRaces*numOfRaces)
	fmt.Println("total minus diag", numOfRaces*numOfRaces-diagonal)
	fmt.Println("total writes", (numOfRaces*numOfRaces-diagonal)/2)
}
