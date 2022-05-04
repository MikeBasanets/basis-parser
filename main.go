package main

import (
	"basis-parser/db"
	"basis-parser/parser"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	setUpLogging()
	godotenv.Load(".env")
	db.Connect()
	startOfParsing := time.Now()
	parseAndSaveToDb()
	err := db.RemoveClothingUpdatedBefore(startOfParsing)
	if err != nil {
		log.Println(err)
	}
	db.Disconnect()
}

func setUpLogging() {
	file, err := os.OpenFile("logs/"+time.Now().Format(time.RFC3339)+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func parseAndSaveToDb() {
	config := loadConfig()
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go parseOuterwearSubcategories(config, wg)
	go parseShirtSubcategories(config, wg)
	go parsePantsSubcategories(config, wg)
	wg.Wait()
}

func parseOuterwearSubcategories(config parser.ParsingConfig, wg *sync.WaitGroup) {
	for i := range config.OuterwearConfig {
		parser.ParseOuterwearSubcategory(config.OuterwearConfig[i].SubcategoryUrl, config.OuterwearConfig[i].DefaultParams, func(result db.Outerwear) {
			err := db.UpsertOuterwear(result)
			if err != nil {
				log.Println(err)
			}
		})
	}
	wg.Done()
}

func parseShirtSubcategories(config parser.ParsingConfig, wg *sync.WaitGroup) {
	for i := range config.ShirtsConfig {
		parser.ParseShirtSubcategory(config.ShirtsConfig[i].SubcategoryUrl, config.ShirtsConfig[i].DefaultParams, func(result db.Shirt) {
			err := db.UpsertShirt(result)
			if err != nil {
				log.Println(err)
			}
		})
	}
	wg.Done()
}

func parsePantsSubcategories(config parser.ParsingConfig, wg *sync.WaitGroup) {
	for i := range config.PantsConfig {
		parser.ParsePantsSubcategory(config.PantsConfig[i].SubcategoryUrl, config.PantsConfig[i].DefaultParams, func(result db.Pants) {
			err := db.UpsertPants(result)
			if err != nil {
				log.Println(err)
			}
		})
	}
	wg.Done()
}

func loadConfig() parser.ParsingConfig {
	configStr, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	var config parser.ParsingConfig
	err = json.Unmarshal([]byte(configStr), &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}
