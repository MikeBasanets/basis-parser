package main

import (
	"basis-parser/parser"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	setUpLogging()
	configStr, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	var config parser.ParsingConfig
	err = json.Unmarshal([]byte(configStr), &config)
	if err != nil {
		log.Fatalln(err)
	}
	for i := range config.OuterwearConfig {
		parser.ParseOuterwearSubcategory(config.OuterwearConfig[i].SubcategoryUrl, config.OuterwearConfig[i].DefaultParams)
	}
	for i := range config.ShirtsConfig {
		parser.ParseShirtSubcategory(config.ShirtsConfig[i].SubcategoryUrl, config.ShirtsConfig[i].DefaultParams)
	}
	for i := range config.PantsConfig {
		parser.ParsePantsSubcategory(config.PantsConfig[i].SubcategoryUrl, config.PantsConfig[i].DefaultParams)
	}
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
