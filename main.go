package main

import (
	"fmt"
	"github.com/mstch/gotrack/provider"
	"github.com/mstch/gotrack/utils/language"
)

const (
	token = "TuJRyLm1pYNrM+p+9rNLd4/ZeIpYpAD4Abma3ot2g0llimorjYNfF338D4grlAWy"
)

func main() {
	lang, _ := language.Get()
	request := provider.NewLaPosteRequestBuilder().AddToken(token).
		AddPackageId("LU680211095FR").
		AddLanguage(lang).
		Build()

	fmt.Println(request)
}
