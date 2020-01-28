package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	ValidateCodeLength = 6
)

func GenValidateCode() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < ValidateCodeLength; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
