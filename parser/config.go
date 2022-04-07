package parser

type ParsingConfig struct {
	PantsConfig []struct {
		SubcategoryUrl string `json:"subcategoryUrl"`
		DefaultParams  Pants  `json:"defaultParams"`
	} `json:"pantsConfig"`
	ShirtsConfig []struct {
		SubcategoryUrl string `json:"subcategoryUrl"`
		DefaultParams  Shirt  `json:"defaultParams"`
	} `json:"shirtsConfig"`
	OuterwearConfig []struct {
		SubcategoryUrl string    `json:"subcategoryUrl"`
		DefaultParams  Outerwear `json:"defaultParams"`
	} `json:"outerwearConfig"`
}
