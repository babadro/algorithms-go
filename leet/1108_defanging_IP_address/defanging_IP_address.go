package _1108_defanging_IP_address

// passed. easy.
func defangIPaddre(address string) string {
	res := make([]byte, 0, len(address))

	for _, v := range []byte(address) {
		if v == '.' {
			res = append(res, "[.]"...)
		} else {
			res = append(res, v)
		}
	}

	return string(res)
}
