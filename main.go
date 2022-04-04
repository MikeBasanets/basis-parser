package main

import (
	"fmt"
	"time"
	"basis-parser/parser"
)

func main() {
	start := time.Now()
	parser.ParseCategory("https://www.lamoda.by/c/517/clothes-muzhskie-bryuki/", parser.PantsType)
	duration := time.Since(start)
	fmt.Println(duration.Seconds(), " Seconds")
}
