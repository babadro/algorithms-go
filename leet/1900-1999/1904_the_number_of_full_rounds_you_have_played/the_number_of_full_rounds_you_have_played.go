package _1904_the_number_of_full_rounds_you_have_played

import "strconv"

// math. passed. easy. best solution. #time
func numberOfRounds2(startTime string, finishTime string) int {
	startH, _ := strconv.Atoi(startTime[:2])
	startM, _ := strconv.Atoi(startTime[3:])
	finishH, _ := strconv.Atoi(finishTime[:2])
	finishM, _ := strconv.Atoi(finishTime[3:])

	start := startH*60 + startM
	finish := finishH*60 + finishM
	if finish < start {
		finish += 1440
	}

	if remainder := finish % 15; remainder != 0 {
		finish -= remainder
	}

	return (finish - start) / 15
}

// simulation
func numberOfRounds(startTime string, finishTime string) int {
	startH, _ := strconv.Atoi(startTime[:2])
	startM, _ := strconv.Atoi(startTime[3:])
	finishH, _ := strconv.Atoi(finishTime[:2])
	finishM, _ := strconv.Atoi(finishTime[3:])

	start := startH*60 + startM
	finish := finishH*60 + finishM
	if finish < start {
		finish += 1440
	}

	flag := false
	counter := 0

	for t := start; t <= finish; t++ {
		if t%15 == 0 {
			if flag {
				counter++
			} else {
				flag = true
			}
		}
	}

	return counter
}
