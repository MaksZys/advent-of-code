package helper

func SplitByBitOccurrencesAtIndex(values []string, index int) ([]string, []string) {
	zero := make([]string, 0)
	one := make([]string, 0)

	if (len(values[0]) - 1) < index {
		panic("To big index")
	}

	for _, el := range values {
		if el[index] == '0' {
			zero = append(zero, el)
		} else {
			one = append(one, el)
		}
	}

	return zero, one
}

func OxygenComparator(zero []string, one []string, index int) []string {
	zeroLen := len(zero)
	oneLen := len(one)

	if zeroLen <= oneLen {
		return one
	}

	return zero
}

func Co2ScrubberComparator(zero []string, one []string, index int) []string {
	zeroLen := len(zero)
	oneLen := len(one)

	if zeroLen <= oneLen {
		return zero
	}

	return one
}
