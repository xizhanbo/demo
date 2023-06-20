package main

import (
	"fmt"
	"time"
)

type Schedule struct {
	Start time.Time // 开始时间
	End   time.Time // 结束时间
}

func between(a, c Schedule) bool {
	astart := a.Start.Unix()
	aend := a.End.Unix()
	cstart := c.Start.Unix()
	cend := c.End.Unix()
	fmt.Println(astart, aend, cstart, cend)
	if astart >= cstart && astart <= cend {
		return true
	}
	if aend >= cstart && aend <= cend {
		return true
	}
	return false
}

func can(a Schedule) bool {
	for _, s := range schedule {
		if between(a, s) {
			return false
		}
	}
	return true
}

var schedule = []Schedule{
	{
		Start: time.Date(2023, 6, 20, 14, 0, 0, 0, time.Local),
		End:   time.Date(2023, 6, 20, 15, 30, 0, 0, time.Local),
	},
}

func main() {
	insert := Schedule{
		Start: time.Date(2023, 6, 20, 13, 35, 0, 0, time.Local),
		End:   time.Date(2023, 6, 20, 13, 55, 0, 0, time.Local),
	}
	fmt.Println(can(insert))
}
