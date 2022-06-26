package _2_load_balancer

import "math/rand"

type LoadBalancer2 struct {
	servers []int
	dict    map[int]int
}

func New2() *LoadBalancer2 {
	return &LoadBalancer2{dict: make(map[int]int)}
}

func (lb *LoadBalancer2) Remove2(serverID int) {
	idx := lb.dict[serverID]
	last := len(lb.servers) - 1
	lb.servers[idx], lb.servers[last] = lb.servers[last], lb.servers[idx]
	lb.servers = lb.servers[:last]

	delete(lb.dict, serverID)
}

func (lb *LoadBalancer2) Add(serverID int) {
	lb.dict[serverID] = len(lb.servers)
	lb.servers = append(lb.servers, serverID)
}

func (lb *LoadBalancer2) Pick() int {
	if len(lb.servers) == 0 {
		return -1
	}

	return rand.Intn(len(lb.servers))
}
