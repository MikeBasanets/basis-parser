package main

import (
	"basis-parser/parser"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	setUpLogging()
	parser.ParseCategory("https://www.lamoda.by/c/517/clothes-muzhskie-bryuki/", parser.PantsType)
	duration := time.Since(start)
	fmt.Println(duration.Seconds(), " Seconds")
}

func setUpLogging() {
	file, err := os.OpenFile("logs/"+time.Now().Format(time.RFC3339)+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}
