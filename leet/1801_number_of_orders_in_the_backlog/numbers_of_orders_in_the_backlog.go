package _1801_number_of_orders_in_the_backlog

import "container/heap"

const (
	buy  = 0
	sell = 1
	d    = 1_000_000_007
)

// todo 1
func getNumberOfBacklogOrders(orders [][]int) int {
	buyOr, sellOr := &minHeap{}, &minHeap{}

	for _, batch := range orders {
		price, amount, typ := batch[0], batch[1], batch[2]
		h1, h2 := buyOr, sellOr
		if typ == sell {
			h1, h2 = sellOr, buyOr
		}

		if h2.Len() > 0 {
			minPrice := (*h2)[h2.Len()-1][0]
			if (typ == buy && minPrice <= price) || (typ == sell && minPrice >= price) {
				for h2.Len() > 0 && amount > 0 {
					last := h2.Len() - 1
					itemPrice := (*h2)[last][0]
					if (typ == buy && itemPrice > price) || (typ == sell && itemPrice < price) {
						break
					}

					itemAmount := (*h2)[last][1]
					if itemAmount <= amount {
						heap.Pop(h2)
						amount -= itemAmount

						continue
					}

					(*h2)[last][1] -= amount
					amount = 0
				}
			}
		}

		if amount > 0 {
			heap.Push(h1, []int{price, amount})
		}
	}

	sum := 0
	for _, batch := range *buyOr {
		sum += batch[1]
		sum %= d
	}

	for _, batch := range *sellOr {
		sum += batch[1]
		sum %= d
	}

	return sum
}

type minHeap [][]int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]

	return res
}
