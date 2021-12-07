package unit

import (
	"adventOfCode2021/day3/part2/helper"
	"log"
	"strconv"
)

type Unit struct {
	Name     string
	BitValue string
}

func (obj *Unit) CalcDecimal() int {
	val, err := strconv.ParseInt((*obj).BitValue, 2, 64)
	if err != nil {
		log.Fatalf("Parse error %s", err)
	}

	return int(val)
}

func (obj *Unit) CalcValues(
	values []string,
	comparator func(zero []string, one []string, index int) []string,
) {
	localValues := values
	for i := 0; i < len(values[0]); i++ {
		if len(localValues) == 1 {
			break
		}

		zero, one := helper.SplitByBitOccurrencesAtIndex(localValues, i)
		localValues = comparator(zero, one, i)
	}

	(*obj).BitValue = localValues[0]
}
