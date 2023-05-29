package main

import (
	"fmt"
	"github.com/mstch/gotrack/provider"
)

func main() {
	request := provider.NewRequest("laposte", "LU680211095FR")

	fmt.Println(request)
}
