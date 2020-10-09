package _380_insert_delete_getrandom_o1

import "testing"

func TestRandomizedSet(t *testing.T) {
	set := Constructor()
	set.Insert(0)
	set.Insert(1)
	set.Insert(2)

	set.Remove(2)
	t.Log(set.Insert(2))
	set.Remove(1)
	set.Remove(0)
	//t.Log(set.Insert(3))
	//t.Log(set.Insert(3))
	//t.Log(set.Remove(3))
	//t.Log(set.Remove(3))
	//
	//for i := 0; i < 20; i++ {
	//	set.Insert(i)
	//}
	//
	//for i := 0; i < 20; i++ {
	//	t.Log(set.GetRandom())
	//}
	//
	//t.Log("start removing\n")
	//
	//for i := 10; i < 20; i++ {
	//	set.Remove(i)
	//}
	//
	//for i := 0; i < 20; i++ {
	//	t.Log(set.GetRandom())
	//}

}
