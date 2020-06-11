package _860_lemonade_change

func lemonadeChange(bills []int) bool {
	m5, m10 := 0, 0
	for _, v := range bills {
		if v == 5 {
			m5++
			continue
		}
		if v == 10 {
			if m5 == 0 {
				return false
			}
			m5--
			m10++
			continue
		}
		if m10 > 0 && m5 > 0 {
			m10--
			m5--
		} else if m5 >= 3 {
			m5 -= 3
		} else {
			return false
		}
	}
	return true
}
