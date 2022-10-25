package _2166_design_bitset

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	bs := Constructor(5)
	t.Log(bs.ToString())

	bs.Fix(3)
	t.Log(bs.ToString())

	bs.Fix(1)
	t.Log(bs.ToString())

	bs.Flip()
	t.Log(bs.ToString())

	t.Log(bs.All())

	bs.Unfix(0)
	t.Log(bs.ToString())

	bs.Flip()
	t.Log(bs.ToString())

	t.Log(bs.One())

	bs.Unfix(0)
	t.Log(bs.ToString())

	t.Log(bs.Count())
}

//["Bitset","flip","unfix","all","fix","fix","unfix","all","count","toString","toString","toString","unfix","flip","all","unfix","one","one","all","fix","unfix"]
//[[2],[],[1],[],[1],[1],[1],[],[],[],[],[],[0],[],[],[0],[],[],[],[0],[0]]
// [null,null,null,false,null,null,null,false,1,"10","10","10",null,null,true,null,true,true,false,null,null]
func TestConstructor2(t *testing.T) {
	bs := Constructor(2)
	bs.Flip()
	t.Log(bs.ToString())
	bs.Unfix(1)

	t.Log(bs.All())
	bs.Fix(1)
	bs.Fix(1)
	bs.Unfix(1)

	t.Log(bs.All())

	t.Log(bs.Count())

	t.Log(bs.ToString())
	t.Log(bs.ToString())
	t.Log(bs.ToString())

	bs.Unfix(0)
	bs.Flip()
	t.Log(bs.All())
}
