package main

import (
	"fmt"
	"github.com/mstch/gotrack/provider"
)

const (
	token = "TuJRyLm1pYNrM+p+9rNLd4/ZeIpYpAD4Abma3ot2g0llimorjYNfF338D4grlAWy"
)

func main() {
	request := provider.NewLaPosteRequestBuilder().AddToken(token).
		AddPackageId("LU680211095FR").
		AddLanguage("fr_FR").
		Build()

	fmt.Println(request)
}
