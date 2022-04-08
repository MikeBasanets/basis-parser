package parser

import "basis-parser/db"

type ParsingConfig struct {
	PantsConfig []struct {
		SubcategoryUrl string   `json:"subcategoryUrl"`
		DefaultParams  db.Pants `json:"defaultParams"`
	} `json:"pantsConfig"`
	ShirtsConfig []struct {
		SubcategoryUrl string   `json:"subcategoryUrl"`
		DefaultParams  db.Shirt `json:"defaultParams"`
	} `json:"shirtsConfig"`
	OuterwearConfig []struct {
		SubcategoryUrl string       `json:"subcategoryUrl"`
		DefaultParams  db.Outerwear `json:"defaultParams"`
	} `json:"outerwearConfig"`
}
