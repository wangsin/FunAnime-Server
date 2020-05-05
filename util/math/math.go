package utilMath

import (
	"fmt"
	"math"
	"strconv"
)

func GetHumanFormatNumber(num int64) string {
	if num < 1000 {
		return strconv.FormatInt(num, 10)
	} else if num >= 1000 && num < 10000 {
		return fmt.Sprintf("%.1fk", math.Round(float64(num)/100.0) / 10)
	} else {
		return fmt.Sprintf("%.1fw", math.Round(float64(num)/1000.0) / 10)
	}
}
