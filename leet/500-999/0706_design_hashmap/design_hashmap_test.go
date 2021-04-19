package _706_design_hashmap

import (
	"reflect"
	"testing"
)

func TestHashMap(t *testing.T) {
	h := Constructor()
	h.Put(1, 1)
	h.Put(2, 2)
	t.Log(h.Get(1))
	t.Log(h.Get(3))
	h.Put(2, 1)
	t.Log(h.Get(2))
	h.Remove(2)
	t.Log(h.Get(2))
}

func TestHashMap2(t *testing.T) {
	h := Constructor()
	t.Log(h.Get(1))
	t.Log(h.Get(1))

	h.Put(1, 1)
	h.Put(1, 1)
	h.Put(1, 1)
	h.Put(1, 1)
	h.Put(2, 2)

	t.Log(h.Get(1))
	t.Log(h.Get(3))

	h.Put(2, 1)

	t.Log(h.Get(2))

	h.Remove(2)

	t.Log(h.Get(2))
}

//func Call(commands []string, args [][]int, expected [][]int)

func TestHashMap3(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		want     [][]int
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
