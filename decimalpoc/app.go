package decimalpoc

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func Start() error {
	test, err := decimal.NewFromString("12.40000000006")
	if err != nil {
		return err
	}
	fmt.Println(test.RoundCeil(2))

	x := 12.3456
	fmt.Println(math.Floor(x*100) / 100) // 12.34 (round down)
	fmt.Println(math.Round(x*100) / 100) // 12.35 (round to nearest)
	fmt.Println(math.Ceil(x))
	return nil
}
