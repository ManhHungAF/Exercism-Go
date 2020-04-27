package clock

import (
	"fmt"
)

// Clock type definition
type Clock struct {
	hours int
	mins  int
}

// New docs
func New(h, m int) Clock {
	hour, min := calcExactTime(h, m, 0)
	return Clock{
		hours: hour,
		mins:  min,
	}
}

// Add docs
func (c Clock) Add(a int) Clock {
	hour, min := calcExactTime(c.hours, c.mins, a)
	return Clock{
		hours: hour,
		mins:  min,
	}
}

// Subtract docs
func (c Clock) Subtract(a int) Clock {
	hour, min := calcExactTime(c.hours, c.mins, -a)
	return Clock{
		hours: hour,
		mins:  min,
	}
}

// String : Method to return the string representation of Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.mins)
}

func calcExactTime(h, m, a int) (exactHour, exactMin int) {
	totalMin := h*60 + m + a
	// if totalMin < 0 -> roll over some day until 0 <= totalMin <= 1440 (mins in one day)
	if totalMin < 0 {
		totalMin += ((1440 - totalMin) / 1440) * 1440
	}
	exactMin = totalMin % 60
	exactHour = (totalMin / 60) % 24

	return exactHour, exactMin
}
