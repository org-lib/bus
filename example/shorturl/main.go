package main

import (
	"fmt"
	ShortUrlGenerator "github.com/org-lib/bus/short-url"
)

func main() {
	result, err := ShortUrlGenerator.Transform("hippo.baozun.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
