package _380_insert_delete_getrandom_o1

import "math/rand"

// 88% 8%
type RandomizedSet struct {
	arr        []int
	valueToIDx map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{valueToIDx: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.valueToIDx[val]
	if ok {
		return false
	}

	this.valueToIDx[val] = len(this.arr)
	this.arr = append(this.arr, val)

	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.valueToIDx[val]
	if !ok {
		return false
	}

	delete(this.valueToIDx, val)
	lastIDx := len(this.arr) - 1
	lastValue := this.arr[lastIDx]

	if len(this.arr) != 1 && idx != lastIDx {
		this.arr[idx] = lastValue
		this.valueToIDx[lastValue] = idx
	}

	this.arr = this.arr[:lastIDx]

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
