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

// ["Robot","step","getPos","getDir","step","getPos","getDir","step","step","getPos","getDir","step","step","step","getPos","getDir","step","step","step","getPos","getDir","step"]
//[[8,2],[17],[],[],[21],[],[],[22],[34],[],[],[1],[46],[35],[],[],[44],[14],[31],[],[],[50]]

// got:  [null,null,[1,0],"East",null,[6,0],"East",null,null,[1,1],"West",null,null,null,[0,0],"East",null,null,null,[6,1],"West",null]
// want: [null,null,[1,0],"East",null,[6,0],"East",null,null,[1,1],"West",null,null,null,[0,0],"South",null,null,null,[6,1],"West",null]
