package _2_load_balancer

import "container/list"

type LoadBalancer struct {
	servers *list.List
	dict    map[int]*list.Element
	cur     int
}

func (lb *LoadBalancer) Add(serverID int) {
	lb.dict[serverID] = lb.servers.PushBack(serverID)
}

func New() *LoadBalancer {
	return &LoadBalancer{
		servers: list.New(),
		dict:    make(map[int]*list.Element),
		cur:     -1,
	}
}

func (lb *LoadBalancer) Remove(serverID int) {
	if lb.cur != serverID {
		lb.servers.Remove(lb.dict[serverID])
		delete(lb.dict, serverID)

		return
	}

	if lb.servers.Len() == 1 {
		lb.servers.Remove(lb.dict[serverID])
		delete(lb.dict, serverID)
		lb.cur = -1

		return
	}

	if next := lb.dict[serverID].Next(); next != nil {
		lb.cur = next.Value.(int)
	} else {
		lb.cur = lb.servers.Front().Value.(int)
	}

	lb.servers.Remove(lb.dict[serverID])
	delete(lb.dict, serverID)
}

// returns server id
func (lb *LoadBalancer) Pick() (serverID int) {
	res := lb.cur

	if res == -1 {
		return
	}

	if next := lb.dict[serverID].Next(); next != nil {
		lb.cur = next.Value.(int)
	} else {
		lb.cur = lb.servers.Front().Value.(int)
	}

	return res
}
