package pointers

import "fmt"

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
