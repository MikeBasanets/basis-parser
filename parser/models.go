package parser

type Wardrobe struct {
	Pants     []Pants     `json:"pants"`
	Shirts    []Shirt     `json:"shirts"`
	Outerwear []Outerwear `json:"outerwear"`
}

type ClothingItem struct {
	ID          int64
	PageUrl     string
	ImageUrl    string
	Color       string
	Pattern     string
	Description string
	Brand       string
	Price       string
	Season      string
	Subcategory string
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
	FitType string
}

type Shirt struct {
	ClothingItem
	FitType        string
	LengthCm       int
	SleeveLengthCm int
	CollarOrCutout string
}
