package main

import (
	"fmt"
	"github.com/mstch/gotrack/provider"
)

const (
	token = "TuJRyLm1pYNrM+p+9rNLd4/ZeIpYpAD4Abma3ot2g0llimorjYNfF338D4grlAWy"
)

func main() {
	request := provider.LaPosteRequest{}
	request.AddToken(token)
	request.AddPackageId("LU680211095FR")
	request.AddLanguage("fr_FR")
	request.Build()
	fmt.Println(request)
}
