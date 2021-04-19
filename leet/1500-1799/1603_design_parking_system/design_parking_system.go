package _1603_design_parking_system

// tptl. easy. passed
type ParkingSystem struct {
	slots [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{slots: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	idx := carType - 1
	if this.slots[idx] == 0 {
		return false
	}

	this.slots[idx]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
