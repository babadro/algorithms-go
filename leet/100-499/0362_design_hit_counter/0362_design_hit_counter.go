package _362_design_hit_counter

// #bnsrg passed
type hit struct {
	timestamp, count int
}

type HitCounter struct {
	hits []hit
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (this *HitCounter) Hit(timestamp int) {
	if len(this.hits) == 0 {
		this.hits = []hit{{timestamp: timestamp, count: 1}}

		return
	}

	last := len(this.hits) - 1

	if this.hits[last].timestamp == timestamp {
		this.hits[last].count++

		return
	}

	this.hits = append(this.hits, hit{timestamp: timestamp, count: 1})
}

func (this *HitCounter) GetHits(timestamp int) int {
	var res int

	for i := len(this.hits) - 1; i >= 0; i-- {
		t := this.hits[i].timestamp

		if t > timestamp {
			continue
		}

		if t <= timestamp-300 {
			return res
		}

		res += this.hits[i].count
	}

	return res
}
