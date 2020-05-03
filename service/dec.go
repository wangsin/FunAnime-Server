package main

import (
	"fmt"
	"github.com/mervick/aes-everywhere/go/aes256"
)

func main() {
	fmt.Println(aes256.Decrypt("U2FsdGVkX1940s8Xv5vqEXdscZuCLNjzV1oF8FwaJsRT6c33W0Zefd/BCnp6djpd/7fwQm5Pus3oSTu4vhLBmg==", "123456"))
}
