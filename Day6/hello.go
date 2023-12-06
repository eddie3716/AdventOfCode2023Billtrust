package main

import "fmt"

type race struct {
	Time     int
	Distance int
}

func distanceTravelled(actualRaceTime int, timeButtonPressed int) int {
	acc := 1 //1 ml/sec^2

	velocity := timeButtonPressed * acc

	distance := velocity * (actualRaceTime - timeButtonPressed)

	return distance
}

func main() {
	/*
		Time:      7  15   30
		Distance:  9  40  200
	*/
	//races := []race{race{7, 9}, race{15, 40}, race{30, 200}}
	//races := []race{race{71530, 940200}}

	/*
		Time:        40     82     84     92
		Distance:   233   1011   1110   1487
	*/
	//races := []race{race{40, 233}, race{82, 1011}, race{84, 1110}, race{92, 1487}}
	races := []race{race{40828492, 233101111101487}}

	raceHypotheticals := [][]int{}

	marginOfError := 1
	for raceNumber, race := range races {
		raceHypotheticals = append(raceHypotheticals, []int{})
		for timeButtonPressed := 0; timeButtonPressed <= race.Time; timeButtonPressed++ {
			distance := distanceTravelled(race.Time, timeButtonPressed)
			raceHypotheticals[raceNumber] = append(raceHypotheticals[raceNumber], distance)
		}
		numOfBetterRaces := 0
		for _, result := range raceHypotheticals[raceNumber] {
			if result > race.Distance {
				numOfBetterRaces++
			}
		}
		fmt.Println(numOfBetterRaces)
		marginOfError *= numOfBetterRaces
	}

	fmt.Println(marginOfError)
	//fmt.Println(raceHypotheticals)
}
