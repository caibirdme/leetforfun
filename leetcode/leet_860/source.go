package leet_860

func lemonadeChange(bills []int) bool {
	count5 := 0
	count10 := 0
	for _, b := range bills {
		if b == 5 {
			count5++
		} else if b == 10 {
			count5--
			count10++
		} else {
			if count10 > 0 {
				count10--
				count5--
			} else {
				count5 -= 3
			}
		}
		if count5 < 0 {
			return false
		}
	}
	return true
}
