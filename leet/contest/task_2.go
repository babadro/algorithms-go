package contest

const off = 0
const on = 1
const blue = 2

func numTimesAllBlue(light []int) int {
	m, counter := 0, 0
	for idx, bulbIdx := range light {
		i := idx + 1
		if bulbIdx > m {
			m = bulbIdx
		}
		if m == i {
			counter++
		}
	}
	return counter
}

func numTimesAllBlueNaive(light []int) int {
	blueCounter := 0
	bulbStates := make([]int, len(light))
	for _, no := range light {
		idx := no - 1
		bulbStates[idx] = on
		increment := true
		allPreviousIsOn := true
		for k := 0; k < len(light); k++ {
			if allPreviousIsOn && bulbStates[k] != off {
				bulbStates[k] = blue
			} else {
				allPreviousIsOn = false
			}
			if bulbStates[k] != off && bulbStates[k] == on {
				increment = false
			}
		}
		if increment {
			blueCounter++
		}
	}
	return blueCounter
}

/*
	allPreviousIsOn := true
			for k := 0; k < idx; k++ {
				if bulbStates[k] == 0 {
					allPreviousIsOn = false
					break
				}
			}
			if allPreviousIsOn {
				bulbStates[idx] = blue
			} else {
				bulbStates[idx] = on
			}
		}
		incrementCounter := true
		for j := 0; j < len(light); j++ {
			if bulbStates[j] != off  && bulbStates[j] == on {
				incrementCounter = false
				break
			}
		}
		if incrementCounter {
			blueCounter++
		}
*/
