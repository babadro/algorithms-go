package _1736_latest_time_by_replacing_hidden_digits

//https://leetcode.com/contest/weekly-contest-225
//passed. tptl
func maximumTime(time string) string {
	arr := [5]byte{'2', '3', ':', '5', '9'}

	res := []byte(time)
	for i := range res {
		if res[i] == '?' {
			if i == 1 && (res[0] == '0' || res[0] == '1') {
				res[i] = '9'
			} else {
				res[i] = arr[i]
			}

			if i == 0 && (res[i+1] > '3' && res[i+1] <= '9') {
				res[0] = '1'
				continue
			}
		}
	}

	return string(res)
}
