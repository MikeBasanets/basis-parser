package db

import "time"

type ClothingItem struct {
	PageUrl     string
	ImageUrl    string
	Color       string
	Pattern     string
	Description string
	Brand       string
	Price       int
	Season      string
	Subcategory string `json:"subcategory"`
	LastUpdated time.Time
}

type Outerwear struct {
	ClothingItem
	HoodType              string
	LengthCm              int
	SleeveLengthCm        int
	InsulationComposition string
}

type Pants struct {
	ClothingItem
	FitType      string
	LegOpeningCm int
}

type Shirt struct {
	ClothingItem
	FitType        string
	LengthCm       int
	SleeveLengthCm int
	CollarOrCutout string `json:"collarOrCutout"`
}
