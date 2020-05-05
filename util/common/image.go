package common

import "fmt"

func BuildImageLink(img string) string {
	return fmt.Sprintf("http://192.168.127.130:4869/%s", img)
}
