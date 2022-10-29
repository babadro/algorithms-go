package _2069_walking_robot_simulation_ii

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	r := Constructor(6, 3)
	r.Step(2)
	r.Step(2)

	t.Log(r.GetPos())
	t.Log(r.GetDir())

	r.Step(2)
	r.Step(1)
	r.Step(4)

	t.Log(r.GetPos())
	t.Log(r.GetDir())
}
