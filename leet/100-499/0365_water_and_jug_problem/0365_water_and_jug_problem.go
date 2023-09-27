package _365_water_and_jug_problem

// bnsrg passed medium
// todo 2 there is a short and fast iterative solution without memorization - check it
func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	r := rec{
		jug1Capacity,
		jug2Capacity,
		targetCapacity,
		false,
		make(map[[2]int]bool),
	}

	r.do(0, 0)

	return r.success
}

type rec struct {
	jug1Cap, jug2Cap, target int
	success                  bool
	visitedState             map[[2]int]bool
}

func (r *rec) do(jug1, jug2 int) {
	key := [2]int{jug1, jug2}
	if r.success || r.visitedState[key] {
		return
	}

	r.visitedState[key] = true

	if jug1 == r.target || jug2 == r.target || jug1+jug2 == r.target {
		r.success = true

		return
	}

	if jug1 == 0 && jug2 == 0 {
		r.do(r.jug1Cap, 0)
		r.do(0, r.jug2Cap)

		return
	}

	if jug1 == 0 {
		r.do(r.jug1Cap, jug2)
		if jug2 > r.jug1Cap {
			r.do(r.jug1Cap, jug2-r.jug1Cap)
		} else {
			r.do(jug2, 0)
		}

		return
	}

	if jug2 == 0 {
		r.do(jug1, r.jug2Cap)
		if jug1 > r.jug2Cap {
			r.do(jug1-r.jug2Cap, r.jug2Cap)
		} else {
			r.do(0, jug1)
		}

		return
	}

	r.do(0, jug2)
	r.do(jug1, 0)

	cap1 := r.jug1Cap - jug1
	cap2 := r.jug2Cap - jug2

	if cap1 > jug2 {
		r.do(jug1+jug2, 0)
	} else {
		r.do(r.jug1Cap, jug2-cap1)
	}

	if cap2 > jug1 {
		r.do(0, jug1+jug2)
	} else {
		r.do(jug1-cap2, r.jug2Cap)
	}
}
