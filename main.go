package main

import (
	"fmt"
	"os"
	"github.com/mstch/gotrack/provider"
	"github.com/mstch/gotrack/utils/commandline"
)

func main() {
	var args commandline.Args
	args.Parse(os.Args[1:])

	request := provider.NewRequest(args.GetTracker(), args.GetPackageId())
	fmt.Println(request)
	result := request.GetResult()

	for line, _ := range result {
		fmt.Println(line)
	}
}
